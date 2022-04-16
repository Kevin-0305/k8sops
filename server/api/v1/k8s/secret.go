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

type SecretData struct {
	v1.Secret
	Data map[string]string `json:"data"`
}

// @Tags Secret
// @Summary 创建Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "创建Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sct/createSecret [post]
func CreateSecret(c *gin.Context) {
	var req request.K8sSecretRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	if err := k8s.CreateSecret(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name, req.Data); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags Secret
// @Summary 删除Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "删除Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sct/deleteSecret [delete]
func DeleteSecret(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindJSON(&req)
	if err := k8s.DeleteSecret(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Secret
// @Summary 更新Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "更新Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sct/updateSecret [put]
func UpdateSecret(c *gin.Context) {
	var req request.K8sRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("数据错误", zap.Any("err", err))
	}
	var secretData map[string]string
	err = json.Unmarshal([]byte(req.ConfigData), &secretData)
	if err != nil {
		global.GVA_LOG.Error("json数据错误", zap.Any("err", err))
	}
	if err := k8s.UpdateConfigMap(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name, secretData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Secret
// @Summary 用id查询Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "用id查询Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sct/findSecret [get]
func FindSecret(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindQuery(&req)
	if sctData, err := k8s.GetSecret(global.GVA_K8SCLIENTS[0], req.Namespace, req.Name); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		secretMap := make(map[string]string)
		for k, v := range sctData.Data {
			secretMap[k] = string(v)
		}
		reData := SecretData{*sctData, secretMap}
		fmt.Println(reData)
		response.OkWithData(gin.H{"reData": reData}, c)
	}
}

// @Tags Secret
// @Summary 分页获取Secret列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SecretSearch true "分页获取Secret列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sct/getSecretList [get]
func GetSecretList(c *gin.Context) {
	var req request.K8sRequestSearch
	_ = c.ShouldBindQuery(&req)
	if list, err := k8s.ListSecret(global.GVA_K8SCLIENTS[0], req.Namespace); err != nil {
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
