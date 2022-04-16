package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags MonitorServer
// @Summary 添加新的服务器
// @Produce  application/json
// @Param data body model.MonitorServer true "公网IP，私网IP，IPV6，ssh端口，主机名，分组id，状态服务器类型，系统"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /monitorServer/createServer [post]
func CreateServer(c *gin.Context) {
	var R request.CreateServer
	err := c.ShouldBindJSON(&R)
	if err != nil {
		fmt.Println("err", err)
	}

	//password := utils.RSA_Encrypt([]byte(R.Password), path.Join(global.GVA_CONFIG.Rsa.Dir, "rsa_public_key.pem"))
	AesPassword := utils.AESCBCEncrypt(global.GVA_CONFIG.Aes, R.Password)
	server := &model.MonitorServer{
		PubIpAddress: R.PubIpAddress,
		PveIpAddress: R.PveIpAddress,
		IpV6Address:  R.IpV6Address,
		SshPort:      R.SshPort,
		User:         R.User,
		Password:     AesPassword,
		HostName:     R.HostName,
		GroupId:      R.GroupId,
		KeyFile:      R.KeyFile,
		State:        R.State,
		ServerType:   R.ServerType,
		System:       R.System,
	}
	err, serverReturn := service.CreateServer(*server)
	if err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithDetailed(response.MonitorServerResponse{Server: serverReturn}, "添加失败", c)
	} else {
		response.OkWithDetailed(response.MonitorServerResponse{Server: serverReturn}, "添加成功", c)
	}
}

// @Tags MonitorServer
// @Summary 分页获取服务器列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param   page     path    int     true        "page"
// @Param   pageSize     path    int     true        "pageSize"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitorServer/getServerList/{page}/{pageSize} [get]
func GetServerList(c *gin.Context) {
	// var pageInfo request.PageInfo
	// _ = c.ShouldBindJSON(&pageInfo)
	page, _ := strconv.Atoi(c.Param("page"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	pageInfo := request.PageInfo{
		Page:     page,
		PageSize: pageSize,
	}
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetServerList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)

	}

}

// @Tags MonitorServer
// @Summary 设置服务器
// @Produce  application/json
// @Param id path int true "id"
// @Param data body model.MonitorServer true "公网IP，私网IP，IPV6，ssh端口，主机名，分组id，状态服务器类型，系统"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /monitorServer/putServerInfo/{id}/ [put]
func PutServerInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var R request.CreateServer
	err := c.ShouldBindJSON(&R)
	if err != nil {
		fmt.Println("err", err)
	}
	server := &model.MonitorServer{
		PubIpAddress: R.PubIpAddress,
		PveIpAddress: R.PveIpAddress,
		IpV6Address:  R.IpV6Address,
		User:         R.User,
		SshPort:      R.SshPort,
		HostName:     R.HostName,
		GroupId:      R.GroupId,
		KeyFile:      R.KeyFile,
		State:        R.State,
		ServerType:   R.ServerType,
		System:       R.System,
	}
	err, serverReturn := service.SetServerInfo(id, *server)
	if err != nil {
		global.GVA_LOG.Error("配置失败", zap.Any("err", err))
		response.FailWithDetailed(response.MonitorServerResponse{Server: serverReturn}, "配置失败", c)
	} else {
		response.OkWithDetailed(response.MonitorServerResponse{Server: serverReturn}, "配置成功", c)
	}
}

// @Tags MonitorServer
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitorServer/getServerInfo/{id}/ [get]
func GetServerInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, serverReturn := service.FindServerById(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithDetailed(response.MonitorServerResponse{Server: serverReturn}, "获取失败", c)
	} else {
		response.OkWithDetailed(response.MonitorServerResponse{Server: serverReturn}, "获取成功", c)
	}
}

// @Tags MonitorServer
// @Summary 删除服务器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monitorServer/deleteServer/{id}/ [delete]
func DeleteServer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.DeleteServer(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Any("err", err))
		response.FailWithDetailed(response.MonitorServerResponse{}, "删除失败", c)
	} else {
		response.OkWithDetailed(response.MonitorServerResponse{}, "删除成功", c)
	}
}
