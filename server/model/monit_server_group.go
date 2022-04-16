package model

import "gin-vue-admin/global"

type ServerGroup struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"comment:组名"`
	Description string `json:"description" gorm:"comment:说明"`
	Servers     string `json:"servers" gorm:"comment:服务器ID列表"`
}
