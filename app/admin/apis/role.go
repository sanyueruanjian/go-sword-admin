package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"

	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/utils"
	"project/utils/app"
)

var r = new(service.Role)

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
	var role dto.SelectRoleArrayDto
	if err := c.ShouldBindQuery(&role); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	orderJsonData, err := utils.OrderJson(role.Orders)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2.参数正确执行响应
	roleData, err := r.SelectRoles(role, orderJsonData)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, roleData)
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
	// 1.获取参数 校验参数
	var insertrole dto.InsertRoleDto
	if err := c.ShouldBind(&insertrole); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	// 2.参数正确执行响应
	err := r.InsertRole(insertrole)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
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
	// 1.获取参数 校验参数
	var updaterole dto.UpdateRoleDto
	if err := c.ShouldBind(&updaterole); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	// 2.参数正确执行响应
	err := r.UpdateRole(updaterole)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
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
	// 1.获取参数 校验参数
	ids := []byte(c.PostForm("ids"))
	idsData := []int{}
	err := json.Unmarshal(ids, &idsData)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2.参数正确执行响应
	err = r.DeleteRole(idsData)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
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
	// 1.获取参数 校验参数
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	menus := []byte(c.PostForm("menus"))
	menusData := []int{}
	err = json.Unmarshal(menus, &menusData)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2.参数正确执行响应
	err = r.UpdateRoleMenu(id, menusData)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
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
func SelectRoleHandler(c *gin.Context, id int) {
	role, err := r.SelectRoleOne(id)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, role)
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
func SelectRolesAllHandler(c *gin.Context) {
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
func DownRolesHandler(c *gin.Context) {
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
func LevelRolesHandler(c *gin.Context) {
}
