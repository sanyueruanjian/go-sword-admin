package apis

import (
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"project/common/api"
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
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	err = r.InsertRole(insertrole, user.UserId)
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
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	err = r.UpdateRole(updaterole, user.UserId)
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
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	err = r.DeleteRole(idsData, user.UserId)
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
	// TODO 加缓存
	role, err := r.SelectRoleAll()
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
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
	xlsx := excelize.NewFile()
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+utils.GetCurrentTimeStr()+"角色数据.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	//回写到web 流媒体 形成下载
	xlsx.SetCellValue("Sheet1", "A1", "角色名称")
	xlsx.SetCellValue("Sheet1", "B1", "角色级别")
	xlsx.SetCellValue("Sheet1", "C1", "描述")
	xlsx.SetCellValue("Sheet1", "D1", "创建日期")
	j := 0
	for i := 2; i < len(roleData); i++ {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(i), roleData[j].Name)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(i), roleData[j].Level)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(i), roleData[j].Description)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(i), roleData[j].CreateTime)
		j++
	}
	_ = xlsx.Write(c.Writer)
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
	//user, err := api.GetCurrentUserInfo(c)
	//if err != nil {
	//	app.ResponseError(c, app.CodeParamNotComplete)
	//	return
	//}
	//level, err := r.SelectRoleLevel(user.Role)
	// TODO
	level, err := r.SelectRoleLevel([]string{"超级管理员", "普通用户"})
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	// 3.返回数据
	app.ResponseSuccess(c, level)
}
