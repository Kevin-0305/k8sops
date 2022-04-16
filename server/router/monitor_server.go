package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitMonitorServerRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("monitorServer")
	//ServerRouter := Router.Group("monitorServer").Use(middleware.OperationRecord())
	{
		ServerRouter.POST("createServer", v1.CreateServer)
		ServerRouter.GET("getServerList/:page/:pageSize/", v1.GetServerList)
		ServerRouter.PUT("putServerInfo/:id/", v1.PutServerInfo)
		ServerRouter.GET("getServerInfo/:id/", v1.GetServerInfo)
		ServerRouter.DELETE("deleteServer/:id/", v1.DeleteServer)
		ServerRouter.GET("getServerStatList", v1.GetServerStatList)
		ServerRouter.GET("getZabbixStat/:id/", v1.GetZabbixStat)
	}
}
