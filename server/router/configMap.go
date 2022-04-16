package router

import (
	v1 "gin-vue-admin/api/v1/k8s"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitConfigMapRouter(Router *gin.RouterGroup) {
	ConfigMapRouter := Router.Group("cm").Use(middleware.OperationRecord())
	{
		ConfigMapRouter.POST("createConfigMap", v1.CreateConfigMap)   // 新建ConfigMap
		ConfigMapRouter.DELETE("deleteConfigMap", v1.DeleteConfigMap) // 删除ConfigMap
		ConfigMapRouter.PUT("updateConfigMap", v1.UpdateConfigMap)    // 更新ConfigMap
		ConfigMapRouter.GET("findConfigMap", v1.FindConfigMap)        // 根据ID获取ConfigMap
		ConfigMapRouter.GET("getConfigMapList", v1.GetConfigMapList)  // 获取ConfigMap列表
	}
}
