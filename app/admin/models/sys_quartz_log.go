package models

import "time"

type SysQuartzLog struct {
	ID              int       `gorm:"primary_key" json:"id"` //ID
	BeanName        string    `json:"bean_name"`             //bean对象名称
	CronExpression  string    `json:"cron_expression"`       //cron表达式
	ExceptionDetail string    `json:"exception_detail"`      //异常详情
	IsSuccess       []byte    `json:"is_success"`            //状态（是否成功）1成功，0失败(默认)
	JobName         string    `json:"job_name"`              //任务名称
	MethodName      string    `json:"method_name"`           //执行方法
	Params          string    `json:"params"`                //方法参数
	Time            int       `json:"time"`                  //执行时间(ms)
	CreateTime      time.Time `json:"create_time"`           //创建时间
	UpdateTime      time.Time `json:"update_time"`           //更新时间
	CreateBy        int       `json:"create_by"`             //创建者
	UpdateBy        int       `json:"update_by"`             //更新者
	IsDeleted       []byte    `json:"is_deleted"`            //逻辑删除：0启用（默认）、1删除
}
