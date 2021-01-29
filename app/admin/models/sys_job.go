package models

import "time"

type SysJob struct {
	ID         int       `gorm:"primary_key" json:"int"` //ID
	Name       string    `json:"name"`                   //岗位名称
	Enabled    []byte    `json:"enabled"`                //状态：1启用（默认）、0禁用
	JobSort    int       `json:"job_sort"`               //排序
	CreateBy   int       `json:"create_by"`              //创建者id
	UpdateBy   int       `json:"update_by"`              //更新者id
	CreateTime time.Time `json:"create_time"`            //创建日期
	UpdateTime time.Time `json:"update_time"`            //更新时间
	IsDeleted  []byte    `json:"is_deleted"`             //软删除（默认值为0，1为删除）
}
