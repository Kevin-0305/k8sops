package service

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"log"
	"time"

	zabbixgo "github.com/canghai908/zabbix-go"
)

func SetZabbixStatToReddis() {
	configMap := config.InitServersConfig()
	apiMap := make(map[string]*zabbixgo.API)
	for _, v := range configMap.Servers {
		url := fmt.Sprintf("http://%s/api_jsonrpc.php", v.Host)
		api := zabbixgo.NewAPI(url)
		api.Login(v.User, v.Password)
		apiMap[v.ID] = api
	}
	for _, v := range configMap.Servers {
		itemStat, err := ItemStat(apiMap[v.ID], configMap.Items, configMap.Group)
		if err != nil {
			log.Fatal(err)
		}
		global.GVA_REDIS.HSet("zabbixStat", v.ID, itemStat)

	}
	time.Sleep(30 * 60 * time.Second)

}

func ItemStat(api *zabbixgo.API, statConfigMap map[string]string, groupName string) (string, error) {
	statMap := make(map[string]zabbixgo.Items)
	for k, v := range statConfigMap {
		statMap[k] = ItemStatGet(api, v, groupName)
	}
	resultMap := make(map[string]map[string]interface{})
	for _, v := range HostsGetByGroupName(api, groupName) {
		resultMap[v] = make(map[string]interface{})
		resultMap[v]["hostId"] = v
	}
	for k, _ := range statConfigMap {
		for _, v2 := range statMap[k] {
			_, ok := resultMap[v2.HostId]
			if ok {
				resultMap[v2.HostId][k] = v2.Lastvalue
			}
		}
	}
	var resultList []interface{}
	for _, v := range resultMap {
		resultList = append(resultList, v)
	}
	data, err := json.Marshal(resultList)
	if err != nil {
		return "", err
	}
	return string(data), nil

}

func ItemStatGet(api *zabbixgo.API, key string, groupName string) zabbixgo.Items {
	params := map[string]interface{}{
		"output":      []string{"hostid", "name", "lastvalue"}, //需求数据，监控项的name 和最新的值
		"search":      map[string]string{"key_": key},          //监控项
		"selectHosts": []string{groupName}}                     //群组名称
	ret, err := api.ItemsGet(params)
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func HostsGetByGroupName(api *zabbixgo.API, groupName string) []string {
	params := map[string]interface{}{
		"output":      []string{"extend"},  //需求数据，监控项的name 和最新的值
		"selectHosts": []string{groupName}} //群组名称
	ret, err := api.HostsGet(params)
	if err != nil {
		log.Fatal(err)
	}
	var hostIDs []string
	for _, v := range ret {
		hostIDs = append(hostIDs, v.HostId)
	}
	return hostIDs
}
