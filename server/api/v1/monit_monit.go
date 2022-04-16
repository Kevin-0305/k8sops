package v1

import (
	"fmt"

	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Test() {
	fmt.Println("hello word")
}

func GetServerStatList(c *gin.Context) {
	//result, err:=GetServerStatList()
	//jsons := `{"Uptime":5159327550000000,"Hostname":"10-7-133-31","Load1":"0.34","Load5":"0.17","Load10":"0.14","RunningProcs":"4","TotalProcs":"189","MemTotal":884633600,"MemFree":66011136,"MemBuffers":0,"MemCached":482484224,"SwapTotal":536866816,"SwapFree":529256448,"FSInfos":[{"MountPoint":"/","Used":4821450752,"Free":16641851392}],"NetIntf":{"docker0":{"IPv4":"172.17.0.1/16","IPv6":"fe80::42:d8ff:fe4a:1cd2/64","Rx":0,"Tx":446},"eth0":{"IPv4":"10.7.133.31/16","IPv6":"fe80::5054:ff:feea:e39c/64","Rx":3482571675,"Tx":3528558916},"lo":{"IPv4":"127.0.0.1/8","IPv6":"::1/128","Rx":0,"Tx":0},"vethd141c63":{"IPv4":"","IPv6":"fe80::9c55:81ff:fe19:17b/64","Rx":0,"Tx":1102}},"CPU":{"User":10.526316,"Nice":0,"System":21.052631,"Idle":68.42105,"Iowait":0,"Irq":0,"SoftIrq":0,"Steal":0,"Guest":0}}`
	if list, err := service.GetServerStatList(); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    1,
			Page:     1,
			PageSize: 1,
		}, "获取成功", c)
	}
}

// func getStats(client *redis.Client, i int) {
// 	t := time.NewTicker(time.Second)
// 	t1 := time.NewTicker(time.Second * 5)
// 	for {
// 		select {
// 		case <-t.C:
// 			pong, _ := client.Ping().Result()
// 			fmt.Println(pong)
// 		case <-t1.C:
// 			statsJosn, _ := client.Get(strconv.Itoa(i)).Result()
// 			fmt.Println(strconv.Itoa(i))
// 			fmt.Println(statsJosn)
// 		}
// 	}

// }

// func statServers() {
// 	client := until.Connect()
// 	fmt.Println("hello world")
// 	//set("name", "x", 100)
// 	err := client.Set("name", "x", time.Duration(100)*time.Second).Err()
// 	if err != nil {
// 		fmt.Printf("set score failed, err:%v\n", err)
// 		return
// 	}
// 	//s, ok := get("name")
// 	r, err := client.Get("name").Result()
// 	if err != nil {
// 		fmt.Println("err", err)
// 	}
// 	fmt.Println(r)
// 	for i := 1; i < 4; i++ {
// 		go rtop.GetState(client, strconv.Itoa(i)) //向redis中塞入key为i的数据
// 		go getStats(client, i)                    //向redis中获取key为i的数据,并打印

// 	}
// 	time.Sleep(100 * time.Second)

// }
