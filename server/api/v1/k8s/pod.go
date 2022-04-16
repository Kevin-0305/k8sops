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

// @Tags PersistentVolumeClaim
// @Summary 创建PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "创建PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/createPersistentVolumeClaim [post]
func CreatePersistentVolumeClaim(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	if err := k8s.CreatePersistentVolumeClaim(global.GVA_K8SCLIENTS[0], req.Namespace, req.ConfigData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PersistentVolumeClaim
// @Summary 删除PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "删除PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePersistentVolumeClaim [delete]
func DeletePersistentVolumeClaim(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindJSON(&req)
	if err := k8s.DeletePersistentVolumeClaim(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PersistentVolumeClaim
// @Summary 更新PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "更新PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /svc/updatePersistentVolumeClaim [put]
func UpdatePersistentVolumeClaim(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	if err := k8s.UpdatePersistentVolumeClaim(global.GVA_K8SCLIENTS[0], req.Namespace, req.ConfigData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PersistentVolumeClaim
// @Summary 用id查询PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "用id查询PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /svc/findPersistentVolumeClaim [get]
func FindPersistentVolumeClaim(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindQuery(&req)
	if reData, err := k8s.GetPersistentVolumeClaim(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reData": reData}, c)
	}
}

// @Tags PersistentVolumeClaim
// @Summary 分页获取PersistentVolumeClaim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PersistentVolumeClaimSearch true "分页获取PersistentVolumeClaim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/getPersistentVolumeClaimList [get]
func GetPersistentVolumeClaimList(c *gin.Context) {
	var req request.K8sRequestSearch
	_ = c.ShouldBindQuery(&req)
	fmt.Println(req)
	if list, err := k8s.ListPersistentVolumeClaim(global.GVA_K8SCLIENTS[0], req.Namespace); err != nil {
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
