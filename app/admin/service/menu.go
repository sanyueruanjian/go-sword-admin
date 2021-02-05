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

func (m *Menu) DeleteMenu(ids *[]int) (err error) {
	menu := new(models.SysMenu)
	return menu.DeleteMenu(ids)
}

func (m *Menu) UpdateMenu(p *dto.UpdateMenuDto, userId int) (err error) {
	menu := new(models.SysMenu)
	return menu.UpdateMenu(p, userId)
}

//TODO
func (m *Menu) SelectForeNeedMenu() (data []*bo.SelectForeNeedMenuBo, err error) {
	//var Menus []*models.SysMenu
	//menu := new(models.SysMenu)
	//Menus, err = menu.SelectForeNeedMenu()
	//for _, v := range Menus {
	//	tmp := &bo.SelectForeNeedMenuBo{}
	//	data = append(data, tmp)
	//}
	//
	//if err != nil {
	//	return nil, err
	//}
	return data, nil
}

func (m *Menu) SuperiorMenu(p *dto.DataMenuDto) (data []*bo.SelectSuperMenuBo, err error) {
	//var Menus []*models.SysMenu
	var children []*bo.SelectSuperMenuBo
	menu := new(models.SysMenu)
	Menus, child, superId, err := menu.SuperiorMenu(p)
	for _, v := range Menus {
		HasChildren := false
		if v.SubCount == 0 {
			HasChildren = false
		} else {
			HasChildren = true
		}
		tmp := &bo.SelectSuperMenuBo{
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
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Label:      v.Title,
			//TODO
			Children:    nil,
			Icon:        v.Icon,
			Component:   v.Component,
			Name:        v.Name,
			Path:        v.Path,
			Permission:  v.Permission,
			Title:       v.Title,
			HasChildren: HasChildren,
		}
		if v.ID == superId && child != nil {
			for _, j := range child {
				tmp := &bo.SelectSuperMenuBo{
					CreateBy:   j.CreateBy,
					UpdatedBy:  j.UpdateBy,
					SubCount:   j.SubCount,
					MenuSort:   j.MenuSort,
					ID:         j.ID,
					Pid:        j.Pid,
					Type:       j.Type,
					Cache:      utils.ByteIntoBool(j.Cache),
					Hidden:     utils.ByteIntoBool(j.Hidden),
					Leaf:       false,
					Iframe:     utils.ByteIntoBool(j.IFrame),
					CreateTime: j.CreateTime,
					UpdateTime: j.UpdateTime,
					Label:      j.Title,
					//TODO
					Children:    nil,
					Icon:        j.Icon,
					Component:   j.Component,
					Name:        j.Name,
					Path:        j.Path,
					Permission:  j.Permission,
					Title:       j.Title,
					HasChildren: HasChildren,
				}
				children = append(children, tmp)
			}
			tmp.Children = children
		}
		data = append(data, tmp)
	}
	return data, nil
}

func (m *Menu) ChildMenu(p int) (data []int, err error) {
	menu := new(models.SysMenu)
	data, err = menu.ChildMenu(p)
	return data, err
}
