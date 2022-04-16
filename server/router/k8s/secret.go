package k8s

import (
	v1 "gin-vue-admin/api/v1/k8s"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSecretRouter(Router *gin.RouterGroup) {
	SecretRouter := Router.Group("sct").Use(middleware.OperationRecord())
	{
		SecretRouter.POST("createSecret", v1.CreateSecret)   // 新建Secret
		SecretRouter.DELETE("deleteSecret", v1.DeleteSecret) // 删除Secret
		SecretRouter.PUT("updateSecret", v1.UpdateSecret)    // 更新Secret
		SecretRouter.GET("findSecret", v1.FindSecret)        // 根据ID获取Secret
		SecretRouter.GET("getSecretList", v1.GetSecretList)  // 获取Secret列表
	}
}
