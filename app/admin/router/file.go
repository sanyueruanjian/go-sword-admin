package router

import (
	"github.com/gin-gonic/gin"
	"project/app/admin/apis"
)

func init()  {
	routerNoCheckRole = append(routerNoCheckRole, fileRouter)
}

// 无需认证的路由代码
func fileRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/file")
	{
		r.POST("uploadImage", apis.UploadImage)
		r.POST("uploadFile", apis.UploadFile)
	}
}
