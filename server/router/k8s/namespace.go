package k8s

import (
	v1 "gin-vue-admin/api/v1/k8s"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitNamespaceRouter(Router *gin.RouterGroup) {
	NamespaceRouter := Router.Group("ns").Use(middleware.OperationRecord())
	{
		NamespaceRouter.POST("createNamespaces", v1.CreateNamespace)  // 新建ConfigMap
		NamespaceRouter.GET("getNamespacesList", v1.GetNamespaceList) // 获取ConfigMap列表
	}
}
