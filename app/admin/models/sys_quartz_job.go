package models

import "time"

type SysQuartzJob struct {
	ID                int       `gorm:"primary_key" json:"id"` //ID
	BeanName          string    `json:"bean_name"`             //Spring Bean名称
	CronExpression    string    `json:"cron_expression"`       //cron 表达式
	IsPause           []byte    `json:"is_pause"`              //状态：0暂停、1启用
	JobName           string    `json:"job_name"`              //任务名称
	MethodName        string    `json:"method_name"`           //方法名称
	Params            string    `json:"params"`                //参数
	Description       string    `json:"description"`           //备注
	PersonInCharge    string    `json:"person_in_charge"`      //负责人
	Email             string    `json:"email"`                 //报警邮箱
	SubTask           string    `json:"sub_task"`              //子任务ID
	PauseAfterFailure []byte    `json:"pause_after_failure"`   //任务失败后是否暂停,0是暂停，1是不暂停
	CreateBy          int       `json:"create_by"`             //创建者
	UpdateBy          int       `json:"update_by"`             //更新者
	CreateTime        time.Time `json:"create_time"`           //创建日期
	UpdateTime        time.Time `json:"update_time"`           //更新时间
	IsDeleted         []byte    `json:"is_deleted"`            //逻辑删除：0启用（默认）、1删除
}
