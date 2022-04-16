package k8s

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service/k8s"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Namespace
// @Summary 创建Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sRequest true "创建Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ns/createNamespace [post]
func CreateNamespace(c *gin.Context) {
	var req request.K8sRequest
	_ = c.ShouldBindJSON(&req)
	if err := k8s.CreateNamespace(global.GVA_K8SCLIENTS[0], req.Name); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Namespace
// @Summary 分页获取Namespace列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sRequestSearch true "分页获取ConfigMap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ns/GetNamespaceList [get]
func GetNamespaceList(c *gin.Context) {
	var req request.K8sRequestSearch
	_ = c.ShouldBindQuery(&req)
	if list, err := k8s.ListNamespaces(global.GVA_K8SCLIENTS[0]); err != nil {
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
