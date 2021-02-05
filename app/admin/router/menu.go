package router

import (
	"project/app/admin/apis"
	"project/utils/app"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, menuRouter)
	routerCheckRole = append(routerCheckRole, menuAuthRouter)
}

// 无需认证的路由代码
func menuRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/menus")
	{
		r.GET("ping", func(c *gin.Context) {
			c.String(int(app.CodeSuccess), "ok")
		})
	}
}

// 需认证的路由代码
func menuAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/menus")
	{
		r.GET("/build", apis.SelectForeNeedMenuHandler)
		r.GET("/", apis.SelectMenuHandler)
		r.POST("/", apis.InsertMenuHandler)
		r.DELETE("/", apis.DeleteMenuHandle)
		r.PUT("/", apis.UpdateMenuHandler)

	}
}
