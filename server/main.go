package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/service"
	"os"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {

	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	global.GVA_PWD, _ = os.Getwd()    //获取当前的工作目录
	initialize.Redis()                // 初始化redis
	initialize.Timer()
	initialize.InitK8sClient()
	// go service.SetServerStatToRedis() //对列表中的服务器进行监控，数据放入redis
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	go service.SetZabbixStatToReddis()
	core.RunWindowsServer()
}
