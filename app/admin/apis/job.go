package apis

import (
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/common/api"
	"project/utils"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)


// InsertMenuHandler 查询岗位
// @Summary 查询岗位
// @Description Author：JiaKunLi 2021/02/1
// @Tags 系统：岗位管理 job Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetJobList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseGetJobList
// @Router /api/job [get]
func GetJobList(c *gin.Context) {
	p := new(dto.GetJobList)

	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}

	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("job bind params failed", zap.String("Username", user.Username), zap.Error(err))
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
	s := new(service.Job)
	res, err := s.GetJobList(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get job list failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 返回响应
	app.ResponseSuccess(c, res)
}

// DelJobById 删除岗位
// @Summary 删除岗位
// @Description Author：JiaKunLi 2021/02/1
// @Tags 系统：岗位管理 job Controller
// @Accept application/json
// @Produce application/json
// @Param object body []int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertMenu
// @Router /api/job [delete]
func DelJobById(c *gin.Context) {
	var ids []int

	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(&ids); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("del job bind ids failed", zap.String("Username", user.Username), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	if len(ids) == 0 {
		app.ResponseError(c, app.CodeParamIsBlank)
		return
	}
	job := new(service.Job)
	count, err := job.DelJobById(user.UserId, &ids)
	if err != nil {
		c.Error(err)
		zap.L().Error("delete job service failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}
	if *count > 0 {
		app.ResponseErrorWithMsg(c, app.CodeOperationFail, "请解除该岗位用户关联后再试！")
		return
	}

	app.ResponseSuccess(c, nil)

}

// AddJob 新增岗位
// @Summary 新增岗位
// @Description Author：JiaKunLi 2021/02/1
// @Tags 系统：岗位管理 job Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.AddJob false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertMenu
// @Router /api/job [post]
func AddJob(c *gin.Context) {
	p := new(dto.AddJob)

	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}

	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("del job bind ids failed", zap.String("Username", user.Username), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	job := new(service.Job)
	if err := job.AddJob(user.UserId, p); err != nil {
		c.Error(err)
		zap.L().Error("add job service failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// UpdateJob 修改岗位
// @Summary 修改岗位
// @Description Author：JiaKunLi 2021/02/1
// @Tags 系统：岗位管理 job Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UpdateJob false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertMenu
// @Router /api/job [put]
func UpdateJob(c *gin.Context) {
	p := new(dto.UpdateJob)

	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}

	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("update job bind body failed", zap.String("Username", user.Username), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	job := new(service.Job)
	if err = job.Update(user.UserId, p); err != nil {
		c.Error(err)
		zap.L().Error("update job failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)

}

// JobDownload 导出岗位数据
// @Summary 导出岗位数据
// @Description Author：JiaKunLi 2021/02/1
// @Tags 系统：岗位管理 job Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetJobList false "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /api/job/download [get]
func JobDownload(c *gin.Context) {
	p := new(dto.GetJobList)

	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}

	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("JobDownload bind params failed", zap.String("Username", user.Username), zap.Error(err))
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
	s := new(service.Job)
	res, err := s.JobListDownload(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("JobDownload service failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	utils.ResponseXls(c, res, `岗位数据`)
}
