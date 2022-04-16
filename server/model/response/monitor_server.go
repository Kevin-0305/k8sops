package response

import (
	"gin-vue-admin/model"
)

type MonitorServerResponse struct {
	Server model.MonitorServer `json:"server"`
}
