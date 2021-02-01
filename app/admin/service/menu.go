package service

import (
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
			CreateBy:    v.CreateBy,
			UpdatedBy:   v.UpdateBy,
			SubCount:    v.SubCount,
			MenuSort:    v.MenuSort,
			ID:          v.ID,
			Pid:         v.Pid,
			Type:        v.Type,
			Cache:       utils.ByteIntoBool(v.Cache),
			Hidden:      utils.ByteIntoBool(v.Hidden),
			Leaf:        false,
			Iframe:      utils.ByteIntoBool(v.IFrame),
			CreateTime:  utils.UnixTimeToString(v.CreateTime),
			UpdateTime:  utils.UnixTimeToString(v.UpdateTime),
			Label:       "",
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

func (m *Menu) DeleteMenu(ids *[]int) (err error) {
	menu := new(models.SysMenu)
	return menu.DeleteMenu(ids)
}

func (m *Menu) UpdateMenu(p *dto.UpdateMenuDto, userId int) (err error) {
	menu := new(models.SysMenu)
	return menu.UpdateMenu(p, userId)
}
