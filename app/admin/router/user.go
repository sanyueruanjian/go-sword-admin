package router

import (
	"project/utils/app"

	"project/app/admin/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, userRouter)
	routerCheckRole = append(routerCheckRole, userAuthRouter)
}

// 无需认证的路由代码
func userRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/auth")
	{
		r.POST("login", apis.LoginHandler)
	}
}

// 需认证的路由代码
func userAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/users")
	{
		r.GET("ping", func(c *gin.Context) {
			c.String(int(app.CodeSuccess), "ok")
		})
		r.POST("/", apis.InsertUserHandler)
		r.GET("/", apis.SelectUserInfoListHandle)
	}
}
