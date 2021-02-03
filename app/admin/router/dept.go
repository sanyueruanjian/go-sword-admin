package router

import (
	"project/app/admin/apis"
	"project/utils/app"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, deptRouter)
	routerCheckRole = append(routerCheckRole, deptAuthRouter)
}

// 无需认证的路由代码
func deptRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/dept")
	{
		r.GET("ping", func(c *gin.Context) {
			c.String(int(app.CodeSuccess), "ok")
		})
	}
}

// 需认证的路由代码
func deptAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/dept")
	{
		r.GET("/download", apis.DownloadDeptHandler)
		r.GET("/", apis.SelectDeptHandler)
		r.POST("/", apis.InsertDeptHandler)
		r.DELETE("/", apis.DeleteDeptHandle)
		r.PUT("/", apis.UpdateDeptHandler)
		r.POST("/superior", apis.SuperiorDeptHandler)
	}
}
