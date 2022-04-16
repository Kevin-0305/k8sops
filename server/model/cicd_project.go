// 自动生成模板Project
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Project struct {
	global.GVA_MODEL
	Name              string `json:"name" form:"name" gorm:"column:name;comment:项目名"`
	Description       string `json:"description" form:"description" gorm:"column:description;comment:说明;type:text;"`
	Language          int    `json:"language" form:"language" gorm:"column:language;comment:语言;type:int;"`
	WareHouse         string `json:"wareHouse" form:"wareHouse" gorm:"column:ware_house;comment:仓库地址"`
	WareHouseAccout   string `json:"wareHouseAccout" form:"wareHouseAccout" gorm:"column:ware_house_accout;comment:仓库账号"`
	WareHousePassword string `json:"wareHousePassword" form:"wareHousePassword" gorm:"column:ware_house_password;comment:仓库密码"`
	BuildScript       string `json:"buildScript" form:"buildScript" gorm:"column:build_script;comment:构建脚本;type:text;"`
	TestScript        string `json:"testScript" form:"testScript" gorm:"column:test_script;comment:测试脚本;type:text;"`
	Servers           string `json:"servers" form:"servers" gorm:"column:servers;comment:需要部署的服务器组ID"`
	DeploymentDir     string `json:"deploymentDir" form:"deploymentDir" gorm:"column:deployment_dir;comment:需要部署的服务器工程位置"`
	RunScript         string `json:"runScript" form:"runScript" gorm:"column:run_script;comment:生产环境中运行的脚本;type:text;"`
}
