package apis

import (
	"github.com/gin-gonic/gin"

	"project/pkg/captcha"
	"project/utils"
	"project/utils/app"
)

// Captcha 获取图片验证码
// @Summary 获取图片验证码
// @Description Author：JiaKunLi 2021/01/26
// @Tags 系统：系统授权接口 Authorization Controller
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseCode
// @Router /api/auth/code [get]
func Captcha(c *gin.Context) {
	id, b64s, err := captcha.DriverMathFunc()
	utils.HasError(err, "验证码获取失败", 500)
	app.ResponseSuccess(c, gin.H{
		"code":    app.CodeSuccess,
		"message": app.CodeSuccess.Msg(),
		"img":     b64s,
		"uuid":    id,
	})
}
