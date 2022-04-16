// 自动生成模板BuildTask
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type BuildTask struct {
	global.GVA_MODEL
	ProjectId   int    `json:"projectId" form:"projectId" gorm:"column:project_id;comment:所属工程的ID;type:int;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;"`
	Deploy      int    `json:"deploy" form:"deploy" gorm:"column:deploy;comment:部署环境;type:int;"`
	BuildWay    int    `json:"buildWay" form:"buildWay" gorm:"column:build_way;comment:构建方式;type:int;"`
	State       int    `json:"state" form:"state" gorm:"column:state;comment:状态;type:int;"`
	BuildLog    string `json:"buildLog" form:"buildLog" gorm:"column:build_log;comment:构建日志;type:text;"`
	TestLog     string `json:"testLog" form:"testLog" gorm:"column:test_log;comment:测试日志;type:text;"`
}

func (BuildTask) TableName() string {
	return "cicd_task"
}
