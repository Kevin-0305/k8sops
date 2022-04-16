package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSharedStorageRouter(Router *gin.RouterGroup) {
	SharedStorageRouter := Router.Group("ShrSto").Use(middleware.OperationRecord())
	{
		SharedStorageRouter.POST("createSharedStorage", v1.CreateSharedStorage)   // 新建SharedStorage
		SharedStorageRouter.DELETE("deleteSharedStorage", v1.DeleteSharedStorage) // 删除SharedStorage
		SharedStorageRouter.DELETE("deleteSharedStorageByIds", v1.DeleteSharedStorageByIds) // 批量删除SharedStorage
		SharedStorageRouter.PUT("updateSharedStorage", v1.UpdateSharedStorage)    // 更新SharedStorage
		SharedStorageRouter.GET("findSharedStorage", v1.FindSharedStorage)        // 根据ID获取SharedStorage
		SharedStorageRouter.GET("getSharedStorageList", v1.GetSharedStorageList)  // 获取SharedStorage列表
	}
}
