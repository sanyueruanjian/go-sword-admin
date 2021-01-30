package service

import (
	"project/app/admin/models/dto"
)

type Menu struct {
}

func (m Menu) InsetMenu(p *dto.InsertMenuDto) error {
	//menu := &models.SysMenu{
	//	Cache:      utils.BoolIntoByte(p.Cache),
	//	Hidden:     utils.BoolIntoByte(p.Hidden),
	//	IFrame:     utils.BoolIntoByte(p.Iframe),
	//	MenuSort:   utils.BoolIntoInt(p.MenuSort),
	//	Icon:       p.Icon,
	//	Pid:        p.Pid,
	//	ID:         p.ID,
	//	Type:       p.Type,
	//	Component:  p.Component,
	//	Name:       p.Name,
	//	Path:       p.Path,
	//	Permission: p.Permission,
	//	Title:      p.Title,
	//}
	return nil
}
