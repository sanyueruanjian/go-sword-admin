package apis

import (
	"errors"
	"project/app/admin/models"
	"project/app/admin/models/dto"
	"project/app/admin/service"
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

