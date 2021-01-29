package router

import (
	"github.com/gin-gonic/gin"
	"project/app/admin/apis"
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
		r.GET(":id", apis.SelectRoleHandler)
		//r.GET("all", apis.SelectRolesAllHandler)
		//r.GET("download", apis.DownRolesHandler)
		//r.GET("level", apis.LevelRolesHandler)
	}

}
