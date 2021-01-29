package models

import "time"

type SysMenu struct {
	ID         int       `gorm:"primary_key" json:"id"` //ID
	Pid        int       `json:"pid"`                   //上级菜单ID
	SubCount   int       `json:"sub_count"`             //子菜单数目
	Type       int       `json:"type"`                  //菜单类型
	Title      string    `json:"title"`                 //菜单标题
	Name       string    `json:"name"`                  //组件名称
	Component  string    `json:"component"`             //组件
	MenuSort   int       `json:"menu_sort"`             //排序
	Icon       string    `json:"icon"`                  //图标
	Path       string    `json:"path"`                  //链接地址
	IFrame     []byte    `json:"i_frame"`               //是否外链
	Cache      []byte    `json:"cache"`                 //缓存
	Hidden     []byte    `json:"hidden"`                //隐藏
	Permission string    `json:"permission"`            //权限
	CreateBy   int       `json:"create_by"`             //创建者
	UpdateBy   int       `json:"update_by"`             //更新者
	CreateTime time.Time `json:"create_time"`           //创建日期
	UpdateTime time.Time `json:"update_time"`           //更新时间
	IsDeleted  []byte    `json:"is_deleted"`
}
