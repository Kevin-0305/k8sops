package k8s

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service/k8s"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Service
// @Summary 创建Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "创建Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/createService [post]
func CreateService(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	if err := k8s.CreateService(global.GVA_K8SCLIENTS[0], req.Namespace, req.ConfigData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Service
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deleteService [delete]
func DeleteService(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindJSON(&req)
	if err := k8s.DeleteService(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Service
// @Summary 更新Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "更新Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /svc/updateService [put]
func UpdateService(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	if err := k8s.UpdateService(global.GVA_K8SCLIENTS[0], req.Namespace, req.ConfigData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Service
// @Summary 用id查询Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "用id查询Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /svc/findService [get]
func FindService(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindQuery(&req)
	if reData, err := k8s.GetService(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reData": reData}, c)
	}
}

// @Tags Service
// @Summary 分页获取Service列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ServiceSearch true "分页获取Service列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/getServiceList [get]
func GetServiceList(c *gin.Context) {
	var req request.K8sRequestSearch
	_ = c.ShouldBindQuery(&req)
	fmt.Println(req)
	if list, err := k8s.ListService(global.GVA_K8SCLIENTS[0], req.Namespace); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		total := len(list.Items)
		start := (req.PageInfo.Page - 1) * req.PageInfo.PageSize
		end := start + req.PageInfo.PageSize
		if end > total {
			end = total
		}
		response.OkWithDetailed(response.PageResult{
			List:     list.Items[start:end],
			Total:    int64(total),
			Page:     req.PageInfo.Page,
			PageSize: req.PageInfo.PageSize,
		}, "获取成功", c)
	}
}
