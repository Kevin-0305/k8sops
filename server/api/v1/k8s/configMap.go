package k8s

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service/k8s"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
)

// @Tags ConfigMap
// @Summary 创建ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "创建ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cm/createConfigMap [post]
func CreateConfigMap(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	var cm *v1.ConfigMap
	err = json.Unmarshal([]byte(req.ConfigData), &cm)
	if err != nil {
		global.GVA_LOG.Error("json数据错误", zap.Any("err", err))
	}
	if err := k8s.CreateConfigMap(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name, cm.Data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ConfigMap
// @Summary 删除ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "删除ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cm/deleteConfigMap [delete]
func DeleteConfigMap(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindJSON(&req)
	if err := k8s.DeleteConfigMap(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ConfigMap
// @Summary 更新ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "更新ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cm/updateConfigMap [put]
func UpdateConfigMap(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	var cm *v1.ConfigMap
	err = json.Unmarshal([]byte(req.ConfigData), &cm)
	if err != nil {
		global.GVA_LOG.Error("json数据错误", zap.Any("err", err))
	}
	if err := k8s.UpdateConfigMap(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name, cm.Data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ConfigMap
// @Summary 用id查询ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "用id查询ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cm/findConfigMap [get]
func FindConfigMap(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindQuery(&req)
	if reData, err := k8s.GetConfigMap(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reData": reData}, c)
	}
}

// @Tags ConfigMap
// @Summary 分页获取ConfigMap列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ConfigMapSearch true "分页获取ConfigMap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cm/getConfigMapList [get]
func GetConfigMapList(c *gin.Context) {
	var req request.K8sRequestSearch
	_ = c.ShouldBindQuery(&req)
	fmt.Println(req)
	if list, err := k8s.ListConfigMap(global.GVA_K8SCLIENTS[0], req.Namespace); err != nil {
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
