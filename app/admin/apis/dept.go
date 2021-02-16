package apis

import (
	"net/http"
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

// SelectDept 查询部门
// @Summary 查询部门
// @Description Author：Lbl 2021/02/2 获得身份令牌
// @Tags 系统：部门管理 Dept Controller
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
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("SelectDeptHandler GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(dept); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SelectDeptHandler params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	if dept.Size == 0 {
		dept.Size = 10
	}
	if dept.Current == 0 {
		dept.Current = 1
	}

	// 排序规则
	var orders string
	if dept.Orders != "" {
		orders = dept.Orders
	} else if dept.Sort != "" {
		orders = `[{"column": "id", "asc": "false"}]`
	}
	orderJson, err := utils.OrderJson(orders)
	if err != nil {
		zap.L().Error("SelectDeptHandler orderString to orderJson failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	// 参数正确执行相应业务
	data, err := d.SelectDeptList(dept, orderJson)
	if err != nil {

		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, data)
}

// InsertDept 添加部门
// @Summary 添加部门
// @Description Author：Lbl 2021/02/2 获得身份令牌
// @Tags 系统：部门管理 Dept Controller
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
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("InsertDeptHandler GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindJSON(dept); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("InsertDeptHandler params failed", zap.String("Username", user.Username), zap.Error(err))
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
		zap.L().Error("InsertDeptDao Insert failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
}

// UpdateDept 修改部门
// @Summary 修改部门
// @Description Author：Lbl 2021/02/2 获得身份令牌
// @Tags 系统：部门管理 Dept Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UpdateDeptDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseDept
// @Router /api/dept [post]
func UpdateDeptHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("UpdateDeptHandler GetUserMsg failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	dept := new(dto.UpdateDeptDto)
	if err := c.ShouldBindJSON(dept); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("UpdateDeptHandler ShouldBindJson failed", zap.String("Username", user.Username), zap.Error(err))
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
	if err := d.UpdateDept(dept); err != nil {
		zap.L().Error("UpdateDeptHandler UpdateSQL failed", zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// DeleteDept 删除部门
// @Summary 删除部门
// @Description Author：Lbl 2021/02/2 获得身份令牌
// @Tags 系统：部门管理 Dept Controller
// @Accept application/json
// @Produce application/json
// @Param object body []int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseDept
// @Router /api/menus [delete]
func DeleteDeptHandle(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("DeleteDeptHandle failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	var ids []int
	if err := c.ShouldBind(&ids); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("DeleteDeptHandle params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	count, err := d.DeleteDept(&ids)
	if err != nil {
		zap.L().Error("DeleteDeptDao failed", zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}
	if count > 0 {
		app.ResponseErrorWithMsg(c, http.StatusBadRequest, "所选部门存在用户关联，请解除后再试！")
	} else {
		app.ResponseSuccess(c, nil)
	}
}

// SuperiorDept 查询部门:根据id
// @Summary 查询部门
// @Description Author：Lbl 2021/02/3 获得身份令牌
// @Tags 系统：部门管理 Dept Controller
// @Accept application/json
// @Produce application/json
// @Param object body []int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSelectDeptList
// @Router /api/dept/superior [post]
func SuperiorDeptHandler(c *gin.Context) {
	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("SelectDeptHandler GetUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	//	绑定校验参数
	var ids []int
	if err := c.ShouldBind(&ids); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SuperiorDeptHandler params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	// 参数正确执行相应业务
	data, err := d.SuperiorDept(&ids)
	if err != nil {
		zap.L().Error("SuperiorDeptHandler Select failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, data)
}

// JobDownload 导出部门数据
// @Summary 导出部门数据
// @Description Author：Lbl 2021/02/3
// @Tags 系统：部门管理 Dept Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.SelectDeptDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /api/dept/download [get]
func DownloadDeptHandler(c *gin.Context) {
	// 声明dto
	dept := new(dto.SelectDeptDto)
	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("DownloadDeptHandler GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(dept); err != nil {
		//请求参数有误， 直接返回响应
		zap.L().Error("DownloadDeptHandler params failed", zap.String("Username", user.Username), zap.Error(err))
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
		zap.L().Error("DownloadDeptHandler orderString to orderJson failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	// 参数正确执行相应业务
	ioRead, err := d.DownloadDeptList(dept, orderJson)
	if err != nil {
		zap.L().Error("DownloadDeptHandler Select failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 3.返回数据
	utils.ResponseXls(c, ioRead, `部门数据`)
}
