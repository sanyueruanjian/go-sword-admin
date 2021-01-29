package models

import "time"

type SysLog struct {
	ID              int       `gorm:"primary_key" json:"id"` //日志id
	UserId          int       `json:"user_id"`               //操作用户id
	Description     string    `json:"description"`           //描述
	LogType         int       `json:"log_type"`              //日志类型
	Method          string    `json:"method"`                //方法名
	Params          string    `json:"params"`                //参数
	RequestIp       string    `json:"request_ip"`            //请求ip
	RequestTime     int       `json:"request_time"`          //请求耗时（毫秒值）
	Address         string    `json:"address"`               //地址
	Browser         string    `json:"browser"`               //浏览器
	ExceptionDetail string    `json:"exception_detail"`      //详细异常
	CreateBy        int       `json:"create_by"`             //创建人id
	UpdateBy        int       `json:"update_by"`             //更新人id
	CreateTime      time.Time `json:"create_time"`           //创建时间
	UpdateTime      time.Time `json:"update_time"`           //更新时间
	IsDeleted       []byte    `json:"is_deleted"`            //软删除（默认值为0，1为删除）
}
