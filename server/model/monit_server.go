package model

import "gin-vue-admin/global"

type MonitorServer struct {
	global.GVA_MODEL
	PubIpAddress string `json:"pubIpAddress" gorm:"comment:公网服务器ip地址"`
	PveIpAddress string `json:"pveIpAddress" gorm:"comment:私网服务器ip地址"`
	IpV6Address  string `json:"ipV6Address" gorm:"comment:公网ipv6"`
	User         string `json:"user" gorm:"comment:登录用户"`
	Password     string `json:"-" gorm:"comment:登录密码"`
	SshPort      int64  `json:"sshPort" gorm:"default:22;comment:ssh端口"`
	HostName     string `json:"hostName" gorm :"主机描述名"`
	GroupId      int64  `json:"groupId" gorm:"主机分组ID"`
	KeyFile      string `json:"keyFile" gorm:"comment:秘钥文件"`
	State        int64  `json:"state" gorm:"comment:状态"`
	ServerType   int64  `json:"serverType" gorm:"comment:服务器类型"`
	System       int64  `json:"system" gorm:"comment:系统"`
}
