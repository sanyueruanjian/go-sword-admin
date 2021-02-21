package apis

import (
	"errors"
	"fmt"
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/common/api"
	"project/common/global"
	"project/pkg/tools"
	"project/utils"
	"project/utils/app"
	"project/utils/config"

	"github.com/mojocn/base64Captcha"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

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
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 校验验证码
	//if !store.Verify(p.UuId, p.Code, true) {
	//	app.ResponseError(c, app.CodeLoginFailCode)
	//	return
	//}

	// 2.业务逻辑处理
	value, err := utils.RsaPriDecode(p.Password)
	if err != nil {
		zap.L().Error("ras decode fail", zap.Error(err))
		app.ResponseError(c, app.CodeLoginFailResCode)
		return
	}
	p.Password = value
	u := new(service.User)
	data, err := u.Login(c, p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get login user info message failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, models.ErrorInvalidPassword) || errors.Is(err, models.ErrorUserNotExist) {
			app.ResponseError(c, app.CodeLoginFailResCode)
			return
		} else if errors.Is(err, models.ErrorUserIsNotEnabled) {
			app.ResponseError(c, app.CodeUserIsNotEnabled)
		}
		app.ResponseError(c, app.CodeSeverError)
		return
	}

	// 3.返回响应
	app.ResponseSuccess(c, data)
}

// LogoutHandler 用户注销接口
// @Summary 用户注销接口
// @Description Author：JiaKunLi 2021/01/26 获得身份令牌
// @Tags 系统：系统授权接口 Authorization Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UserLoginDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/auth/logout [delete]
func LogoutHandler(c *gin.Context) {
	userOnline, err := api.GetUserOnline(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("获取线上用户数据失败", zap.Error(err))
		app.ResponseError(c, app.CodeBadRequest)
		return
	}

	global.Rdb.Del(fmt.Sprintf("%s%s%s", config.JwtConfig.RedisHeader, "-", userOnline.Token))
	app.ResponseSuccess(c, nil)
}

// InsertUserHandler 新增用户
// @Summary 新增用户
// @Description Author：Cgl 2021/02/01 获得身份令牌
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertUser
// @Router /api/users [post]
func InsertUserHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.InsertUserDto)
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("InsertUserHandler failed", zap.String("username", user.Username), zap.Error(err))
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
		if errors.Is(err, models.ErrorUserIsExist) {
			zap.L().Error("insert menu failed", zap.Error(err))
			app.ResponseErrorWithMsg(c, app.CodeInsertOperationFail, "用户已存在")
			return
		}
		zap.L().Error("insert menu failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, nil)
}

// SelectUserInfoListHandler 查询用户详细
// @Summary 查询用户详细
// @Description Author：Cgl 2021/02/01 获得身份令牌
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.SelectUserInfoArrayDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSelectUserInfoList
// @Router /api/users [get]
func SelectUserInfoListHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.SelectUserInfoArrayDto)
	//获取上下文中信息
	userInfo, err := api.GetUserData(c)
	userMessage := new(api.UserMessage)
	userMessage, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SelectUserInfoList failed", zap.String("username", userMessage.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	//menu业务逻辑处理
	m := new(service.User)
	var data *bo.UserInfoListBo
	data, err = m.SelectUserInfoList(p, &models.ModelUserMessage{
		UserId:         userMessage.UserId,
		Username:       userMessage.Username,
		DataScopes:     userInfo.DataScopes,
		Dept:           userInfo.Dept,
		Jobs:           userInfo.Jobs,
		Roles:          userInfo.Roles,
		MenuPermission: userInfo.MenuPermission,
	})
	if err != nil {
		zap.L().Error("select user failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, data)
}

// DeleteUserHandler 删除用户
// @Summary 删除用户
// @Description Author：Cgl 2021/02/02 获得身份令牌
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.DeleteUserDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseDeleteUser
// @Router /api/users [delete]
func DeleteUserHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	var ids []int
	if err := c.ShouldBind(&ids); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("DeleteMenuHandler failed", zap.String("username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	menu := new(service.User)
	if err := menu.DeleteUser(ids); err != nil {
		zap.L().Error("DeleteUser failed", zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// UpdateUserHandler 更新用户
// @Summary 更新用户
// @Description Author：Cgl 2021/02/02 获得身份令牌
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UpdateUserDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseUpdateUserCenter
// @Router /api/users/center [put]
func UpdateUserHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	p := new(dto.UpdateUserDto)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("UpdateUserHandler failed", zap.String("username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
	}
	//处理逻辑
	u := new(service.User)
	if err := u.UpdateUser(p, user.UserId); err != nil {
		zap.L().Error("UpdateUser failed", zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// UpdateUserCenterHandler 更新用户 个人中心
// @Summary 更新用户 个人中心
// @Description Author：Cgl 2021/02/02 获得身份令牌
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UpdateUserCenterDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseUpdateUser
// @Router /api/users [put]
func UpdateUserCenterHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	p := new(dto.UpdateUserCenterDto)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("UpdateUserCenterHandler failed", zap.String("username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
	}
	//处理逻辑
	u := new(service.User)
	if err := u.UpdateUserCenter(p, user.UserId); err != nil {
		zap.L().Error("UpdateUser failed", zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// SelectUserInfoHandler 查询用户详细
// @Summary 查询用户详细
// @Description Author：Cgl 2021/02/01 获得身份令牌
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSelectUserInfoList
// @Router /api/auth/info [get]
func SelectUserInfoHandler(c *gin.Context) {
	//获取上下文中信息
	userMessage, err := api.GetUserMessage(c)
	var userInfo *api.UserInfo
	userInfo, err = api.GetUserData(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//处理逻辑
	u := new(service.User)
	var data *bo.UserCenterInfoBo
	data, err = u.SelectUserInfo(&models.ModelUserMessage{
		Username:       userMessage.Username,
		UserId:         userMessage.UserId,
		MenuPermission: userInfo.MenuPermission,
	})
	if err != nil {
		zap.L().Error("SelectUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	app.ResponseSuccess(c, data)
}

// UpdatePassWordHandler 更新用户
// @Summary 更新用户
// @Description Author：Cgl 2021/02/02 获得身份令牌
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UpdateUserPassDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseUpdateUserCenter
// @Router /api/users/updatePass [post]
func UpdatePassWordHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	p := new(dto.UpdateUserPassDto)
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("UpdatePassWordHandler failed", zap.String("username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
	}
	p.NewPass, err = utils.RsaPubEncode(p.NewPass)
	//私钥解密
	valueNew, err := utils.RsaPriDecode(p.NewPass)
	if err != nil {
		zap.L().Error("ras decode fail", zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	p.NewPass = valueNew

	valueOld, err := utils.RsaPriDecode(p.OldPass)
	if err != nil {
		zap.L().Error("ras decode fail", zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	p.OldPass = valueOld
	if p.NewPass == "" {
		app.ResponseError(c, app.CodeParamIsBlank)
		return
	}
	//处理逻辑
	u := new(service.User)
	if err := u.UpdatePassWord(p, user.UserId); err != nil {
		zap.L().Error("UpdateUser failed", zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// UpdateAvatarHandler 更换头像（图片）
// @Summary 更换头像（图片）
// @Description Author：Cgl 2021/02/02
// @Tags 系统：用户管理 User Controller
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "file"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseFile
// @Router /api/users/updateAvatar [post]
func UpdateAvatarHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//上传图片
	files, err := c.FormFile("avatar")
	if err != nil {
		zap.L().Error("FormFile failed", zap.String("username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeImageIsNotNull)
		return
	}
	if utils.GetFileType(tools.GetExt(files.Filename)[1:]) != "image" {
		app.ResponseError(c, app.CodeFileImageFail)
		return
	}
	// 上传文件至指定目录
	guid := uuid.New().String()
	fileName := guid + tools.GetExt(files.Filename)
	singleFile := "static/uploadfile/" + fileName
	err = c.SaveUploadedFile(files, singleFile)
	if err != nil {
		app.ResponseError(c, app.CodeFileUploadFail)
		return
	}
	u := new(service.User)
	if err := u.UpdateAvatar(fileName, user.UserId); err != nil {
		return
	}
	app.ResponseSuccess(c, fileName)
}

// UserDownloadHandler 导出用户数据
// @Summary 导出用户数据
// @Description Author：JiaKunLi 2021/02/1
// @Tags 系统：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.DownloadUserInfoDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /api/user/download [get]
func UserDownloadHandler(c *gin.Context) {
	p := new(dto.DownloadUserInfoDto)
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("user bind params failed", zap.String("username", user.Username), zap.Error(err))
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
	s := new(service.User)
	res, err := s.UserDownload(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get user list failed", zap.String("username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	utils.ResponseXls(c, res, "用户数据")
}
