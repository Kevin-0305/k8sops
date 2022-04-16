package k8s

import (
	v1 "gin-vue-admin/api/v1/k8s"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitPersistentVolumeClaimRouter(Router *gin.RouterGroup) {
	PersistentVolumeClaimRouter := Router.Group("pvc").Use(middleware.OperationRecord())
	{
		PersistentVolumeClaimRouter.POST("createPersistentVolumeClaim", v1.CreatePersistentVolumeClaim)   // 新建PersistentVolumeClaim
		PersistentVolumeClaimRouter.DELETE("deletePersistentVolumeClaim", v1.DeletePersistentVolumeClaim) // 删除PersistentVolumeClaim
		PersistentVolumeClaimRouter.PUT("updatePersistentVolumeClaim", v1.UpdatePersistentVolumeClaim)    // 更新PersistentVolumeClaim
		PersistentVolumeClaimRouter.GET("findPersistentVolumeClaim", v1.FindPersistentVolumeClaim)        // 根据ID获取PersistentVolumeClaim
		PersistentVolumeClaimRouter.GET("getPersistentVolumeClaimList", v1.GetPersistentVolumeClaimList)  // 获取PersistentVolumeClaim列表
	}
}
