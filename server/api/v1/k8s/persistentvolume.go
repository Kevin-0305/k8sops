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

// @Tags PersistentVolume
// @Summary 创建PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "创建PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/createPersistentVolume [post]
func CreatePersistentVolume(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	if err := k8s.CreatePersistentVolume(global.GVA_K8SCLIENTS[0], req.ConfigData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PersistentVolume
// @Summary 删除PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "删除PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePersistentVolume [delete]
func DeletePersistentVolume(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindJSON(&req)
	if err := k8s.DeletePersistentVolume(global.GVA_K8SCLIENTS[0], req.Name); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PersistentVolume
// @Summary 更新PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "更新PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /svc/updatePersistentVolume [put]
func UpdatePersistentVolume(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	if err := k8s.UpdatePersistentVolume(global.GVA_K8SCLIENTS[0], req.ConfigData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PersistentVolume
// @Summary 用id查询PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "用id查询PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /svc/findPersistentVolume [get]
func FindPersistentVolume(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindQuery(&req)
	if reData, err := k8s.GetPersistentVolume(global.GVA_K8SCLIENTS[0], req.Name); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reData": reData}, c)
	}
}

// @Tags PersistentVolume
// @Summary 分页获取PersistentVolume列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PersistentVolumeSearch true "分页获取PersistentVolume列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/getPersistentVolumeList [get]
func GetPersistentVolumeList(c *gin.Context) {
	var req request.K8sRequestSearch
	_ = c.ShouldBindQuery(&req)
	fmt.Println(req)
	if list, err := k8s.ListPersistentVolume(global.GVA_K8SCLIENTS[0]); err != nil {
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
