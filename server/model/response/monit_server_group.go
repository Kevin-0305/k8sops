package response

import (
	"gin-vue-admin/model"
)

type ServerGroupResponse struct {
	ServerGroup model.ServerGroup `json:"serverGroup"`
}
