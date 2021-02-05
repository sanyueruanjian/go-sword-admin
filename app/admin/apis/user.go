package apis

import (
	"errors"
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/common/api"
	"project/utils"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// LoginHandler 登录授权接口
// @Summary 登录授权接口
// @Description Author：JiaKunLi 2021/01/26 获得身份令牌
// @Tags 系统：系统授权接口 Authorization Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/auth/login [post]
func LoginHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.UserLoginDto)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		//errs, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2.业务逻辑处理
	//TODO 方便postman测试 (模拟前端数据)
	p.Password, _ = utils.RsaPubEncode(p.Password)
	value, err := utils.RsaPriDecode(p.Password)
	if err != nil {
		zap.L().Error("ras decode fail", zap.Error(err))
		app.ResponseError(c, app.CodeLoginFailResCode)
		return
	}
	p.Password = value
	u := new(service.User)
	token, err := u.Login(p)
	if err != nil {
		c.Error(err)
		if errors.Is(err, models.ErrorInvalidPassword) {
			app.ResponseError(c, app.CodeSeverError)
			return
		}
		app.ResponseError(c, app.CodeLoginFailResCode)
		return
	}

	// 3.返回响应
	app.ResponseSuccess(c, token)
}

// InsertUserHandler 新增用户
// @Summary 新增用户
// @Description Author：Cgl 2021/02/01 获得身份令牌
// @Tags 系统：系统授权接口 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertUser
// @Router /api/menus [post]
func InsertUserHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.InsertUserDto)
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("InsertUserHandler failed", zap.String("username", user.UserName), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	//业务逻辑处理
	u := new(service.User)
	if err := u.InsertUser(p, user.UserId); err != nil {
		zap.L().Error("insert menu failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, nil)
}

// SelectUserInfoList 查询用户详细
// @Summary 查询用户详细
// @Description Author：Cgl 2021/02/01 获得身份令牌
// @Tags 系统：系统授权接口 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSelectUserInfoList
// @Router /api/menus [get]
func SelectUserInfoListHandle(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.SelectUserInfoArrayDto)
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SelectUserInfoList failed", zap.String("username", user.UserName), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	//业务逻辑处理
	m := new(service.User)
	var data []*bo.UserInfoListBo
	data, err = m.SelectUserInfoList(p)
	if err != nil {
		zap.L().Error("select menu failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, data)
}
