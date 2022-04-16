package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitServerGroupRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("serverGroup")
	{
		ServerRouter.POST("createServerGroup", v1.CreateServerGroup)
		ServerRouter.PUT("putServerGroup/:id/", v1.PutServerGroup)
		ServerRouter.GET("getServerGroupList", v1.GetServerGroupList)
	}
}
