package k8s

import (
	v1 "gin-vue-admin/api/v1/k8s"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitPodRouter(Router *gin.RouterGroup) {
	PodRouter := Router.Group("svc").Use(middleware.OperationRecord())
	{
		PodRouter.POST("createPod", v1.CreatePod)   // 新建Pod
		PodRouter.DELETE("deletePod", v1.DeletePod) // 删除Pod
		PodRouter.PUT("updatePod", v1.UpdatePod)    // 更新Pod
		PodRouter.GET("findPod", v1.FindPod)        // 根据ID获取Pod
		PodRouter.GET("getPodList", v1.GetPodList)  // 获取Pod列表
	}
}
