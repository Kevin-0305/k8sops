package k8s

import (
	v1 "gin-vue-admin/api/v1/k8s"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitPersistentVolumeRouter(Router *gin.RouterGroup) {
	PersistentVolumeRouter := Router.Group("pv").Use(middleware.OperationRecord())
	{
		PersistentVolumeRouter.POST("createPersistentVolume", v1.CreatePersistentVolume)   // 新建PersistentVolume
		PersistentVolumeRouter.DELETE("deletePersistentVolume", v1.DeletePersistentVolume) // 删除PersistentVolume
		PersistentVolumeRouter.PUT("updatePersistentVolume", v1.UpdatePersistentVolume)    // 更新PersistentVolume
		PersistentVolumeRouter.GET("findPersistentVolume", v1.FindPersistentVolume)        // 根据ID获取PersistentVolume
		PersistentVolumeRouter.GET("getPersistentVolumeList", v1.GetPersistentVolumeList)  // 获取PersistentVolume列表
	}
}
