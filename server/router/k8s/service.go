package k8s

import (
	v1 "gin-vue-admin/api/v1/k8s"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitServiceRouter(Router *gin.RouterGroup) {
	ServiceRouter := Router.Group("svc").Use(middleware.OperationRecord())
	{
		ServiceRouter.POST("createService", v1.CreateService)   // 新建Service
		ServiceRouter.DELETE("deleteService", v1.DeleteService) // 删除Service
		ServiceRouter.PUT("updateService", v1.UpdateService)    // 更新Service
		ServiceRouter.GET("findService", v1.FindService)        // 根据ID获取Service
		ServiceRouter.GET("getServiceList", v1.GetServiceList)  // 获取Service列表
	}
}
