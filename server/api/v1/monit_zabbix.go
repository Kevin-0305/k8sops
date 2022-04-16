package v1

import (
	"encoding/json"
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"

	"github.com/gin-gonic/gin"
)

// @Tags MonitorServer
// @Summary 添加新的服务器
// @Produce  application/json
// @Param data body model.Service true "用id查询"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /monitorServer/GetZabbixStat [get]
func GetZabbixStat(c *gin.Context) {
	serverId := c.Param("id")
	data := global.GVA_REDIS.HGet("zabbixStat", serverId)
	var statList []map[string]string
	err := json.Unmarshal([]byte(data.Val()), &statList)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     statList,
			Total:    1,
			Page:     1,
			PageSize: 1,
		}, "获取成功", c)
	}
}
