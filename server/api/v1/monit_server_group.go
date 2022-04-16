package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags MonitorServer
// @Summary 拉取所有服务器分组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitorServer/getServerList/{page}/{pageSize} [get]
func GetServerGroupList(c *gin.Context) {

	err, list := service.GetServerGroupList()
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

// @Tags ServerGroup
// @Summary 添加服务器分组
// @Produce  application/json
// @Param data body model.MonitorServer true "分组名,描述,服务器列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /serverGroup/CreateServerGroup [post]
func CreateServerGroup(c *gin.Context) {
	var R model.ServerGroup
	err := c.ShouldBindJSON(&R)
	if err != nil {
		log.Println("err", err)
	}
	if err, serverGroupReturn := service.CreateServerGroup(R); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.ServerGroupResponse{ServerGroup: serverGroupReturn}, "添加成功", c)
	}
}

// @Tags ServerGroup
// @Summary 修改服务器分组
// @Produce  application/json
// @Param data body model.MonitorServer true "分组名,描述,服务器列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /serverGroup/CreateServerGroup [post]
func PutServerGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var R model.ServerGroup
	err := c.ShouldBindJSON(&R)
	if err != nil {
		log.Println("err", err)
	}
	if err, serverGroupReturn := service.SetServerGroupInfo(id, R); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.ServerGroupResponse{ServerGroup: serverGroupReturn}, "修改成功", c)
	}

}
