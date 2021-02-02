package models

import (
	"project/app/admin/models/bo"
	"project/utils/app"
)

// _ResponseLogin swagger登录授权响应结构体
type _ResponseLogin struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
	Data    struct {
		Token string `json:"token"` // 授权令牌
	} `json:"data"` // 数据
}

// _ResponseCode 短信验证码响应结构体
type _ResponseCode struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
	Img     string      `json:"img"`     // base64验证码
	UuId    string      `json:"uuid"`    // 验证码id
}

// _ResponseFile 文件上传响应结构体
type _ResponseFile struct {
	Code    app.ResCode  `json:"code"`    // 业务响应状态码
	Message string       `json:"message"` // 提示信息
	Data    FileResponse `json:"data"`    // 数据
}

//_ResponseInsertMenu 新增菜单
type _ResponseInsertMenu struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
}

//_ResponseInsertMenu 查询菜单
type _ResponseSelectMenu struct {
	Code    app.ResCode      `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
	Data    *bo.SelectMenuBo `json:"data"`    // 数据
}

//_ResponseDeleteMenu 删除菜单
type _ResponseDeleteMenu struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
}

//_ResponseUpdateMenu 删除菜单
type _ResponseUpdateMenu struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
}

//_ResponseInsertUser 新增用户
type _ResponseInsertUser struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
}

//_ResponseSelectForeNeedMenu
type _ResponseSelectForeNeedMenu struct {
	Code    app.ResCode              `json:"code"`    // 业务响应状态码
	Message string                   `json:"message"` // 提示信息
	Data    *bo.SelectForeNeedMenuBo `json:"data"`    // 数据
}

//_ResponseSelectUserInfoList
type _ResponseSelectUserInfoList struct {
	Code    app.ResCode        `json:"code"`    // 业务响应状态码
	Message string             `json:"message"` // 提示信息
	Data    *bo.UserInfoListBo `json:"data"`    // 数据
}

//_ResponseSelectUserInfoList
type _ResponseGetJobList struct {
	Code    app.ResCode      `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
	Data    []*bo.GetJobList `json:"data"`    // 数据
}
