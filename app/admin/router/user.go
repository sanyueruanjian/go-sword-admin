package router

import (
	"project/app/admin/middleware"
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
		r.PUT("/center", apis.UpdateUserCenterHandler)
		r.POST("/updatePass", apis.UpdatePassWordHandler)
		r.POST("/updateAvatar", apis.UpdateAvatarHandler)
		r.GET("/download", apis.UserDownloadHandler)
		//casbin校验的接口
		r.Use(middleware.AuthCheckRole())
		r.GET("/", apis.SelectUserInfoListHandler)
		r.DELETE("/", apis.DeleteUserHandler)
		r.POST("/", apis.InsertUserHandler)
		r.PUT("/", apis.UpdateUserHandler)
	}
	r1 := v1.Group("/auth")
	{
		r1.DELETE("logout", apis.LogoutHandler)
		r1.GET("/info", apis.SelectUserInfoHandler)
	}
}
