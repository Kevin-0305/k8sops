package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitBuildTaskRouter(Router *gin.RouterGroup) {
	BuildTaskRouter := Router.Group("buildTask").Use(middleware.OperationRecord())
	{
		//BuildTaskRouter.POST("createBuildTask", v1.CreateBuildTask)   // 新建BuildTask
		BuildTaskRouter.DELETE("deleteBuildTask", v1.DeleteBuildTask)           // 删除BuildTask
		BuildTaskRouter.DELETE("deleteBuildTaskByIds", v1.DeleteBuildTaskByIds) // 批量删除BuildTask
		BuildTaskRouter.PUT("updateBuildTask", v1.UpdateBuildTask)              // 更新BuildTask
		BuildTaskRouter.GET("findBuildTask", v1.FindBuildTask)                  // 根据ID获取BuildTask
		BuildTaskRouter.GET("getBuildTaskList", v1.GetBuildTaskList)            // 获取BuildTask列表
	}
}
