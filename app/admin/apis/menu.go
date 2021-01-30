package apis

import (
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/common/api"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// InsertMenuHandler 新增菜单
// @Summary 新增菜单
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：系统授权接口 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertMenu
// @Router /api/menus [post]
func InsertMenuHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.InsertMenuDto)
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("InsertMenuHandler failed", zap.String("username", user.UserName), zap.Error(err))
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
	//业务逻辑处理
	m := new(service.Menu)
	if err := m.InsetMenu(p); err != nil {
		zap.L().Error("insert menu failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, nil)
}
