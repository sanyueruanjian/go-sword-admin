package models

import (
	"time"
)

type SysRole struct {
	ID           int       `gorm:"primary_key" json:"id"` //ID
	Name         string    `json:"name"`                  //角色名称
	Level        int       `json:"level"`                 //角色级别（越小越大）
	Description  string    `json:"description"`           //描述
	DataScope    string    `json:"data_scope"`            //数据权限
	IsProtection []byte    `json:"is_protection"`         //是否受保护（内置角色，1为内置角色，默认值为0）
	CreateBy     int       `json:"create_by"`             //创建者id
	UpdateBy     int       `json:"update_by"`             //更新者id
	CreateTime   time.Time `json:"create_time"`           //创建日期
	UpdateTime   time.Time `json:"update_time"`           //更新时间
	IsDeleted    []byte    `json:"is_deleted"`            //软删除（默认值为0，1为删除）
}
