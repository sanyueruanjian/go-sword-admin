package apis

import (
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/common/api"
	"project/utils"
	"project/utils/app"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var d = new(service.Dept)

// SelectMenuHandler 查询部门
// @Summary 查询部门
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：系统授权接口 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.SelectDeptDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSelectDeptList
// @Router /api/dept [get]
func SelectDeptHandler(c *gin.Context) {
	// 声明dto
	dept := new(dto.SelectDeptDto)

	// 获取缓存信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("SelectDeptHandler GetCurrentUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(dept); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SelectDeptHandler params failed", zap.String("username", user.UserName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	// 排序规则
	orderJson, err := utils.OrderJson(dept.Orders)
	if err != nil {
		zap.L().Error("SelectDeptHandler orderString to orderJson failed", zap.String("username", user.UserName), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	// 参数正确执行相应业务
	data, err := d.SelectDeptList(dept, orderJson)
	if err != nil {
		zap.L().Error("SelectDeptDao Select failed", zap.String("username", user.UserName), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, data)
}

// SelectMenuHandler 添加部门
// @Summary 添加部门
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：系统授权接口 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertDeptDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseDept
// @Router /api/dept [post]
func InsertDeptHandler(c *gin.Context) {
	// 声明dto
	dept := new(dto.InsertDeptDto)

	// 获取缓存信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("InsertDeptHandler GetCurrentUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindJSON(dept); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("InsertDeptHandler params failed", zap.String("username", user.UserName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	// 参数正确执行相应业务
	err = d.InsertDept(dept, user.UserId)
	if err != nil {
		zap.L().Error("InsertDeptDao Insert failed", zap.String("username", user.UserName), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
}

// SelectMenuHandler 修改部门
// @Summary 修改部门
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：系统授权接口 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UpdateDeptDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseDept
// @Router /api/dept [post]
func UpdateDeptHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("UpdateDeptHandler GetUserMsg failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	dept := new(dto.UpdateDeptDto)
	if err := c.ShouldBindJSON(dept); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("UpdateDeptHandler SouBindJson failed", zap.String("username", user.UserName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	// 替换更新者
	dept.UpdatedBy = user.UserId
	// 业务处理
	//TODO updateTime时间同步
	if err := d.UpdateDept(dept); err != nil {
		zap.L().Error("UpdateDeptHandler UpdateSQL failed", zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// DeleteMenuHandle 删除部门
// @Summary 删除部门
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：系统授权接口 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body []int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseDept
// @Router /api/menus [delete]
func DeleteDeptHandle(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("DeleteDeptHandle failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	var ids []int
	if err := c.ShouldBind(&ids); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("DeleteDeptHandle failed", zap.String("username", user.UserName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	if err := d.DeleteDept(&ids); err != nil {
		zap.L().Error("DeleteDeptDao failed", zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}
