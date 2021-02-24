package router

import (
	"project/app/admin/apis"
	"project/app/admin/middleware"
	"project/utils/app"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, jobRouter)
	routerCheckRole = append(routerCheckRole, jobAuthRouter)
}

// 无需认证的路由代码
func jobRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/job")
	{
		r.GET("ping", func(c *gin.Context) {
			c.String(int(app.CodeSuccess), "ok")
		})
	}
}

// 需认证的路由代码
func jobAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/job")
	{
		//权限认证的接口
		r.Use(middleware.AuthCheckRole())
		r.GET("download", apis.JobDownload)
		r.GET("", apis.GetJobList)
		r.DELETE("", apis.DelJobById)
		r.POST("", apis.AddJob)
		r.PUT("", apis.UpdateJob)

	}
}
