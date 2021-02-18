package apis

import (
	"github.com/gin-gonic/gin"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/common/api"
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
	err := c.ShouldBind(&insertrole)

	// 2.参数正确执行响应
	user, err := api.GetUserMessage(c)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	err = r.InsertRole(insertrole, user.UserId)
	if err != nil {
		if err.Error()[0:10] == "Error 1062" {
			app.ResponseError(c, app.CodeUserNameExist)
			return
		}
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
	var updateRole dto.UpdateRoleDto
	err := c.ShouldBind(&updateRole)

	// 2.参数正确执行响应
	user, err := api.GetUserMessage(c)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	err = r.UpdateRole(updateRole, user.UserId)
	if err != nil {
		if err.Error()[0:10] == "Error 1062" {
			app.ResponseError(c, app.CodeUserNameExist)
			return
		}
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
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2.参数正确执行响应
	user, err := api.GetUserMessage(c)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	err = r.DeleteRole(ids, user.UserId)
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
	var roleMenus dto.RoleMenus
	err := c.ShouldBind(&roleMenus)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2.参数正确执行响应
	err = r.UpdateRoleMenu(roleMenus.ID, roleMenus.Menus)
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
	// 1.获取数据
	role, err := r.SelectRoleAll()
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2.返回数据
	app.ResponseSuccess(c, role)
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
	roleData, err := r.DownloadRoleInfoBo(role, orderJsonData)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回文件数据
	var res []interface{}
	for _, role := range roleData {
		res = append(res, &bo.DownloadRoleInfoBo{
			Name:        role.Name,
			Level:       role.Level,
			Description: role.Description,
			CreateTime:  role.CreateTime,
		})
	}
	content := utils.ToExcel([]string{`角色名称`, `角色级别`, `描述`, `创建日期`}, res)
	utils.ResponseXls(c, content, "角色数据")
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
	user, err := api.GetUserData(c)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	var level bo.SelectCurrentUserLevel
	for _, values := range *user.Roles {
		if level.Level < values.Level {
			level.Level = values.Level
		}
	}

	// 3.返回数据
	app.ResponseSuccess(c, level)
}
