package models

import "time"

type SysDept struct {
	ID         int       `gorm:"primary_key" json:"id"` //ID
	Pid        int       `json:"pid"`                   //上级部门（顶级部门为0，默认为0）
	SubCount   int       `json:"sub_count"`             //子部门数目
	Name       string    `json:"name"`                  //名称
	DeptSort   int       `json:"dept_sort"`             //排序
	Enabled    []byte    `json:"enabled"`               //状态：1启用（默认）、0禁用
	CreateBy   int       `json:"create_by"`             //创建者
	UpdateBy   int       `json:"update_by"`             //更新者
	CreateTime time.Time `json:"create_time"`           //创建日期
	UpdateTime time.Time `json:"update_time"`           //更新时间
	IsDeleted  []byte    `json:"is_deleted"`            //状态：1启用（默认）、0禁用
}
