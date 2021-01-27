package router

import (
	"project/app/admin/apis"

	"github.com/gin-gonic/gin"
)

func init()  {
	routerNoCheckRole = append(routerNoCheckRole, getCaptchaRouter)
}

// 无需认证的路由代码
func getCaptchaRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/auth")
	{
		r.GET("code", apis.Captcha)
	}
}
