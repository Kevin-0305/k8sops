// 自动生成模板SharedStorage
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type SharedStorage struct {
      global.GVA_MODEL
      StorageName  string `json:"storageName" form:"storageName" gorm:"column:storage_name;comment:存储名;type:varchar(20);size:20;"`
      ServerIP  string `json:"serverIP" form:"serverIP" gorm:"column:server_ip;comment:主机ip;type:varchar(20);size:20;"`
      ServerID  int `json:"serverID" form:"serverID" gorm:"column:server_id;comment:主机id;type:int;"`
      Volume  float64 `json:"volume" form:"volume" gorm:"column:volume;comment:容量;type:float;"`
      UseVolume  float64 `json:"useVolume" form:"useVolume" gorm:"column:use_volume;comment:使用容量;type:float;"`
      Path  string `json:"path" form:"path" gorm:"column:path;comment:路径;type:varchar(200);size:200;"`
      User  string `json:"user" form:"user" gorm:"column:user;comment:账号;type:varchar(100);size:100;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;type:varchar(200);size:200;"`
      Protocol  int `json:"protocol" form:"protocol" gorm:"column:protocol;comment:;type:int;"`
}


func (SharedStorage) TableName() string {
  return "shared_storage"
}

