package router

import (
	"github.com/gin-gonic/gin"
	"project/app/admin/apis"
	"project/utils/app"
	"strconv"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, roleAuthRouter)
	// TODO 认证
	//routerCheckRole = append(routerCheckRole, roleAuthRouter)
}

func roleAuthRouter(v1 *gin.RouterGroup) {
	v1.GET("roles", apis.SelectRolesHandler)
	v1.POST("roles", apis.InsertRolesHandler)
	v1.PUT("roles", apis.UpdateRolesHandler)
	v1.DELETE("roles", apis.DeleteRolesHandler)
	r := v1.Group("/roles")
	{
		r.PUT("menu", apis.MenuRolesHandler)
		r.GET(":id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err == nil {
				c.String(int(app.CodeSuccess), apis.SelectRoleHandler(id))
			}
			if c.Param("id") == "all" {
				c.String(int(app.CodeSuccess), apis.SelectRolesAllHandler())
			}
			if c.Param("id") == "download" {
				c.String(int(app.CodeSuccess), apis.DownRolesHandler())
			}
			if c.Param("id") == "level" {
				c.String(int(app.CodeSuccess), apis.LevelRolesHandler())
			}
		})
	}
}
