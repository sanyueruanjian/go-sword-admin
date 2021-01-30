package models

import (
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
	Img     string      `json:"data"`    // base64验证码
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
