package request

type CreateServer struct {
	PubIpAddress string `json:"pubIpAddress"  `
	PveIpAddress string `json:"pveIpAddress"`
	IpV6Address  string `json:"ipV6Address"`
	User         string `json:"user"`
	Password     string `json:"password"`
	SshPort      int64  `json:"sshPort" gorm:"default:22"`
	HostName     string `json:"hostName"`
	GroupId      int64  `json:"groupId"`
	KeyFile      string `json:"keyFile"`
	State        int64  `json:"state"`
	ServerType   int64  `json:"serverType"`
	System       int64  `json:"system" gorm:"default:1"`
}

type UpdateServer struct {
	PubIpAddress string `json:"pubIpAddress"  `
	PveIpAddress string `json:"pveIpAddress"`
	IpV6Address  string `json:"ipV6Address"`
	User         string `json:"user"`
	Password     string `json:"-"`
	SshPort      int64  `json:"sshPort" gorm:"default:22"`
	HostName     string `json:"hostName"`
	GroupId      int64  `json:"groupId"`
	KeyFile      string `json:"keyFile"`
	State        int64  `json:"state"`
	ServerType   int64  `json:"serverType"`
	System       int64  `json:"system" gorm:"default:1"`
}
