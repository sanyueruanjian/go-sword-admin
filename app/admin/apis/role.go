package apis

import (
	"github.com/gin-gonic/gin"
)

// SelectRolesHandler 多条件查询角色
// @Summary 多条件查询角色
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles [get]
func SelectRolesHandler(c *gin.Context) {
	// 1.获取参数 校验参数
}

// SelectRolesHandler 新增角色
// @Summary 新增角色
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles [post]
func InsertRolesHandler(c *gin.Context) {

}

// SelectRolesHandler 修改角色
// @Summary 修改角色
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles [put]
func UpdateRolesHandler(c *gin.Context) {

}

// SelectRolesHandler 删除角色
// @Summary 删除角色
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles [delete]
func DeleteRolesHandler(c *gin.Context) {

}

// SelectRolesHandler 修改角色菜单
// @Summary 修改角色菜单
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles/menu [put]
func MenuRolesHandler(c *gin.Context) {

}

// SelectRolesHandler 获取单个角色
// @Summary 获取单个角色
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles/{id} [put]
func SelectRoleHandler(id int) string {
	return "ok"
}

// SelectRolesHandler 返回全部角色
// @Summary 返回全部角色
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles/all [get]
func SelectRolesAllHandler() string {
	return "ok"
}

// SelectRolesHandler 导出角色数据
// @Summary 导出角色数据
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles/download [get]
func DownRolesHandler() string {
	return "ok"
}

// SelectRolesHandler 获取当前登录用户级别
// @Summary 获取当前登录用户级别
// @Description Author：Ymq 2021/01/29 获得身份令牌
// @Tags 系统：系统授权接口 Role Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/roles/level [get]
func LevelRolesHandler() string {
	return "ok"
}
