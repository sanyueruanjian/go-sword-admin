package router

import (
	"github.com/gin-gonic/gin"
	"project/app/admin/apis"
	"strconv"
)

func init() {
	// 无需认证接口
	//routerNoCheckRole = append(routerNoCheckRole, roleAuthRouter)
	// 认证
	routerCheckRole = append(routerCheckRole, roleAuthRouter)
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
	}
}
