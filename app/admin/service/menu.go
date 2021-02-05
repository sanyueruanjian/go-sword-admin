package service

import (
	"bytes"
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/utils"
)

type Menu struct {
}

func (m *Menu) InsertMenu(p *dto.InsertMenuDto, userID int) error {
	menu := &models.SysMenu{
		Cache:      utils.BoolIntoByte(p.Cache),
		Hidden:     utils.BoolIntoByte(p.Hidden),
		IFrame:     utils.BoolIntoByte(p.Iframe),
		MenuSort:   p.MenuSort,
		Icon:       p.Icon,
		Pid:        p.Pid,
		Type:       p.Type,
		Component:  p.Component,
		Name:       p.Name,
		Path:       p.Path,
		Permission: p.Permission,
		Title:      p.Title,
		CreateBy:   userID,
		UpdateBy:   userID,
	}
	if err := menu.InsertMenu(); err != nil {
		return err
	}
	return nil
}

func (m *Menu) SelectMenu(p *dto.SelectMenuDto) (data []*bo.SelectMenuBo, err error) {
	var Menus []*models.SysMenu
	menu := new(models.SysMenu)
	Menus, err = menu.SelectMenu(p)
	for _, v := range Menus {
		tmp := &bo.SelectMenuBo{
			CreateBy:   v.CreateBy,
			UpdatedBy:  v.UpdateBy,
			SubCount:   v.SubCount,
			MenuSort:   v.MenuSort,
			ID:         v.ID,
			Pid:        v.Pid,
			Type:       v.Type,
			Cache:      utils.ByteIntoBool(v.Cache),
			Hidden:     utils.ByteIntoBool(v.Hidden),
			Leaf:       false,
			Iframe:     utils.ByteIntoBool(v.IFrame),
			CreateTime: utils.UnixTimeToString(v.CreateTime),
			UpdateTime: utils.UnixTimeToString(v.UpdateTime),
			Label:      "",
			//TODO
			Children:    "",
			Icon:        v.Icon,
			Component:   v.Component,
			Name:        v.Name,
			Path:        v.Path,
			Permission:  v.Permission,
			Title:       v.Title,
			HasChildren: "",
		}
		data = append(data, tmp)
	}

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Menu) DeleteMenu(ids []int) (err error) {
	menu := new(models.SysMenu)
	return menu.DeleteMenu(ids)
}

func (m *Menu) UpdateMenu(p *dto.UpdateMenuDto, userId int) (err error) {
	menu := new(models.SysMenu)
	return menu.UpdateMenu(p, userId)
}

func (m *Menu) SelectForeNeedMenu(user *models.RedisUserInfo) (data []*bo.SelectForeNeedMenuBo, err error) {
	menu := new(models.SysMenu)
	return menu.SelectForeNeedMenu(user)
}

// TODO
func (m *Menu) ReturnToAllMenus(pid int) (data []*bo.ReturnToAllMenusBo, err error) {

	menu := new(models.SysMenu)
	menus, err := menu.ReturnToAllMenus(pid)
	if err != nil {
		return nil, err
	}
	for _, v := range menus {
		tmp := &bo.ReturnToAllMenusBo{
			Cache:       utils.ByteIntoBool(v.Cache),
			Children:    nil,
			Component:   v.Component,
			CreateBy:    v.CreateBy,
			CreateTime:  v.CreateTime,
			HasChildren: false,
			Hidden:      utils.ByteIntoBool(v.Hidden),
			Icon:        v.Icon,
			ID:          0,
			Iframe:      utils.ByteIntoBool(v.IFrame),
			Label:       v.Title,
			Leaf:        true,
			MenuSort:    v.MenuSort,
			Name:        v.Name,
			Path:        v.Path,
			Permission:  v.Permission,
			Pid:         v.Pid,
			SubCount:    v.SubCount,
			Title:       v.Title,
			Type:        v.Type,
			UpdatedBy:   v.UpdateBy,
			UpdateTime:  v.UpdateTime,
		}
		if tmp.SubCount != 0 {
			tmp.HasChildren = true
			tmp.Leaf = false
		}
		data = append(data, tmp)
	}
	return data, nil
}

func (m *Menu) DownloadMenuInfoBo(p *dto.DownloadMenuDto, orderData []bo.Order) (menuData []*bo.DownloadMenuInfoBo, err error) {
	menu := new(models.SysMenu)
	sysMenu, err := menu.DownloadMenu(*p, orderData)
	if err != nil {
		return
	}
	for _, v := range sysMenu {
		tmp := &bo.DownloadMenuInfoBo{
			Title:      v.Title,
			Type:       "按钮",
			Permission: v.Permission,
			IFrame:     "否",
			Hidden:     "否",
			Cache:      "否",
			CreateTime: utils.Int64ToString(v.CreateTime),
		}
		switch {
		case bytes.Equal(v.IFrame, []byte{1}):
			tmp.IFrame = "是"
		case bytes.Equal(v.Hidden, []byte{1}):
			tmp.Hidden = "是"
		case bytes.Equal(v.Cache, []byte{1}):
			tmp.Cache = "是"
		}
		menuData = append(menuData, tmp)
	}
	return
}
