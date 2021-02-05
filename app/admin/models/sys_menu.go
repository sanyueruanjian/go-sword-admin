package models

import (
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/common/global"
	"project/utils"
)

type SysMenu struct {
	*BaseModel
	Pid        int    `json:"pid"`        //上级菜单ID
	SubCount   int    `json:"sub_count"`  //子菜单数目
	Type       int    `json:"type"`       //菜单类型
	Title      string `json:"title"`      //菜单标题
	Name       string `json:"name"`       //组件名称
	Component  string `json:"component"`  //组件
	MenuSort   int    `json:"menu_sort"`  //排序
	Icon       string `json:"icon"`       //图标
	Path       string `json:"path"`       //链接地址
	IFrame     []byte `json:"i_frame"`    //是否外链
	Cache      []byte `json:"cache"`      //缓存
	Hidden     []byte `json:"hidden"`     //隐藏
	Permission string `json:"permission"` //权限
	CreateBy   int    `json:"create_by"`  //
	UpdateBy   int    `json:"update_by"`  //
}

func (m *SysMenu) InsertMenu() error {
	err := global.Eloquent.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *SysMenu) SelectMenu(p *dto.SelectMenuDto) (data []*SysMenu, err error) {
	//排序条件
	var orderJson []bo.Order
	orderJson, err = utils.OrderJson(p.Orders)
	orderRule := utils.GetOrderRule(orderJson)
	//模糊条件
	blurry := "%" + p.Blurry + "%"
	//时间条件
	if p.EndTime != 0 && p.StatTime != 0 {
		if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND create_time > ? AND create_time < ? AND title like ?", []byte{0}, p.StatTime, p.EndTime, blurry).
			Limit(p.Size).Offset(p.Current - 1*p.Size).Order(orderRule).Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND title like ?", []byte{0}, blurry).
			Limit(p.Size).Offset(p.Current - 1*p.Size).Order(orderRule).Find(&data).Error; err != nil {
			return nil, err
		}
	}
	return data, err
}

func (m *SysMenu) DeleteMenu(ids *[]int) (err error) {
	return global.Eloquent.Table("sys_menu").Where("id IN (?)", *ids).Updates(map[string]interface{}{"is_deleted": []byte{1}}).Error
}

func (m *SysMenu) UpdateMenu(p *dto.UpdateMenuDto, userId int) (err error) {
	return global.Eloquent.Table("sys_menu").Where("id=?", p.ID).Updates(map[string]interface{}{
		"pid":        p.Pid,
		"sub_count":  p.SubCount,
		"type":       p.Type,
		"title":      p.Title,
		"name":       p.Name,
		"component":  p.Component,
		"menu_sort":  p.MenuSort,
		"icon":       p.Icon,
		"permission": p.Permission,
		"path":       p.Path,
		"update_by":  userId,
		"i_frame":    utils.BoolIntoByte(p.Iframe),
		"cache":      utils.BoolIntoByte(p.Cache),
		"hidden":     utils.BoolIntoByte(p.Iframe),
	}).Error
}

//TODO
func (m *SysMenu) SelectForeNeedMenu() (err error) {
	//global.Eloquent.Find
	//寻找子集
	return nil
}
