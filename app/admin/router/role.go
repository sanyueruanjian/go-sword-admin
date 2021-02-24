package router

import (
	"project/app/admin/apis"
	"project/app/admin/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	// 无需认证接口
	//routerNoCheckRole = append(routerNoCheckRole, roleAuthRouter)
	// 认证
	routerCheckRole = append(routerCheckRole, roleAuthRouter)
}

func roleAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/roles")
	{
		r.PUT("menu", apis.MenuRolesHandler)
		r.GET(":id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err == nil {
				apis.SelectRoleHandler(c, id)
			}
			if c.Param("id") == "all" {
				apis.SelectRolesAllHandler(c)
			}
			if c.Param("id") == "download" {
				apis.DownRolesHandler(c)
			}
			if c.Param("id") == "level" {
				apis.LevelRolesHandler(c)
			}
		})
		r.Use(middleware.AuthCheckRole())
		r.GET("", apis.SelectRolesHandler)
		r.POST("", apis.InsertRolesHandler)
		r.PUT("", apis.UpdateRolesHandler)
		r.DELETE("", apis.DeleteRolesHandler)
	}

}
