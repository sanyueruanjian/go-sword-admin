package service

import (
	"bytes"
	"encoding/json"
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/cache"
	"project/app/admin/models/dto"
	"project/utils"
	"strconv"
)

type Menu struct {
}

func (m *Menu) InsertMenu(p *dto.InsertMenuDto, userID int) error {
	typeInt, err := utils.StringToInt(p.Type)
	if err != nil {
		return err
	}
	menu := &models.SysMenu{
		Cache:      utils.BoolIntoByte(p.Cache),
		Hidden:     utils.BoolIntoByte(p.Hidden),
		IFrame:     utils.BoolIntoByte(p.Iframe),
		MenuSort:   p.MenuSort,
		Icon:       p.Icon,
		Pid:        p.Pid,
		Type:       typeInt,
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
	////查询缓存
	var redisData []byte
	redisData, err = cache.GetMenuListCache("menu::pid:" + strconv.Itoa(p.Pid))
	if redisData != nil && len(redisData) != 0 {
		err = json.Unmarshal(redisData, &data)
		return data, nil
	}
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
			Children:   nil,
			Icon:       v.Icon,
			Component:  v.Component,
			Name:       v.Name,
			Path:       v.Path,
			Permission: v.Permission,
			Title:      v.Title,
		}
		tmp.HasChildren = v.SubCount != 0
		data = append(data, tmp)
	}
	//添加缓存
	var dataByte []byte
	dataByte, err = json.Marshal(data)
	err = cache.SetMenuListCache(dataByte, p.Pid)
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

func (m *Menu) SelectForeNeedMenu(user *models.ModelUserMessage) (data []*bo.SelectForeNeedMenuBo, err error) {
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
			ID:          v.ID,
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
