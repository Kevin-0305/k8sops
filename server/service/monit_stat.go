package service

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	"gin-vue-admin/utils/sshStat"
	"log"
	"sync"
	"time"
)

func SetServerStatToRedis() {
	var wg sync.WaitGroup
	err, serverList := GetAllServerList()
	if err != nil {
		log.Println(err)
		return
	}
	ticker := time.NewTicker(15 * time.Second)

	for {
		<-ticker.C
		var hostStatsMap sync.Map
		for i := 0; i < len(serverList); i++ {
			wg.Add(1)
			server := serverList[i]
			password := utils.AESCBCDecrypter(global.GVA_CONFIG.Aes, server.Password)
			go sshStat.GetStats(server.PubIpAddress, server.User, password, int(server.SshPort), &hostStatsMap, &wg)
			// go rtop.GetState("103.218.243.41", "root", "wangkai@2020", 22, "", global.GVA_REDIS, "103.218.243.41") //向redis中塞入key为i的数据
			//向redis中获取key为i的数据,并打印
		}
		wg.Wait()
		//var hostMap map[string]string
		hostMap := make(map[string]string)
		hostStatsMap.Range(func(key, value interface{}) bool {
			hostMap[key.(string)] = value.(string)
			//fmt.Printf("%s,%s", key, value)
			return true
		})
		data, err := json.Marshal(hostMap)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
		global.GVA_REDIS.Set("hostStats", string(data), 30*time.Second)
	}
}

func GetServerStatList() (List interface{}, err error) {

	// jsons := `{"Uptime":5159327550000000,"Hostname":"10-7-133-31","Load1":"0.34","Load5":"0.17","Load10":"0.14","RunningProcs":"4","TotalProcs":"189","MemTotal":884633600,"MemFree":66011136,"MemBuffers":0,"MemCached":482484224,"SwapTotal":536866816,"SwapFree":529256448,"FSInfos":[{"MountPoint":"/","Used":4821450752,"Free":16641851392}],"NetIntf":{"docker0":{"IPv4":"172.17.0.1/16","IPv6":"fe80::42:d8ff:fe4a:1cd2/64","Rx":0,"Tx":446},"eth0":{"IPv4":"10.7.133.31/16","IPv6":"fe80::5054:ff:feea:e39c/64","Rx":3482571675,"Tx":3528558916},"lo":{"IPv4":"127.0.0.1/8","IPv6":"::1/128","Rx":0,"Tx":0},"vethd141c63":{"IPv4":"","IPv6":"fe80::9c55:81ff:fe19:17b/64","Rx":0,"Tx":1102}},"CPU":{"User":10.526316,"Nice":0,"System":21.052631,"Idle":68.42105,"Iowait":0,"Irq":0,"SoftIrq":0,"Steal":0,"Guest":0}}`
	// err, serverList := GetAllServerList()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	var MonitStatList []response.MonitStatResponse
	var hostStatsMap map[string]string
	hostStatJson, err := global.GVA_REDIS.Get("hostStats").Result()
	err = json.Unmarshal([]byte(hostStatJson), &hostStatsMap)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for k, v := range hostStatsMap {
		var stats sshStat.Stats
		err = json.Unmarshal([]byte(v), &stats)
		if err != nil {
			log.Println(err.Error())
			return
		}
		MemUse := stats.MemTotal - stats.MemFree
		Ram := float64(MemUse) / float64(stats.MemTotal)
		//Disk := (serverStat.DiskTotal-serverStat.DiskFree)/serverStat.DiskTotal
		fmt.Println("serverStat", stats)
		DiskTotal := stats.FSInfos[0].Used + stats.FSInfos[0].Free
		MonitStatResponse := response.MonitStatResponse{
			HostName:  stats.Hostname,
			IP:        k,
			CPU:       stats.Load5,
			Ram:       Ram,
			RamTotal:  stats.MemTotal,
			RamUse:    MemUse,
			Disk:      0.2,
			DiskTotal: DiskTotal,
			DiskUse:   stats.FSInfos[0].Used,
			Flow:      stats.NetIntf["eth0"].Rx,
		}
		MonitStatList = append(MonitStatList, MonitStatResponse)
	}
	return MonitStatList, nil

}
