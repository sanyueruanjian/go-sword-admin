package router

import (
	"project/app/admin/apis/public"

	"github.com/gin-gonic/gin"
)

func init()  {
	routerNoCheckRole = append(routerNoCheckRole, fileRouter)
}

// 无需认证的路由代码
func fileRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/file")
	{
		r.POST("uploadImage", public.UploadImage)
		r.POST("uploadFile", public.UploadFile)
	}
}
