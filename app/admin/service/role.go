package service

import (
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/utils"
)

type Role struct {
}

// 多条件查询角色
func (e Role) SelectRoles(p dto.SelectRoleArrayDto, orderData []bo.Order) (roleData bo.SelectRoleArrayBo, err error) {
	role := new(models.SysRole)
	if err = role.SelectRoles(); err != nil {
		// TODO
		//return "", err
	}
	return
}

// 新增角色
func (e Role) InsertRole(p dto.InsertRoleDto) (err error) {
	role := new(models.SysRole)
	role.Level = p.Level
	role.Name = p.Name
	role.DataScope = p.DataScope
	role.Description = p.Description
	if err = role.InsertRole(); err != nil {
		return err
	}
	return
}

// 修改角色
func (e Role) UpdateRole(p dto.UpdateRoleDto) (err error) {
	role := new(models.SysRole)
	role.ID = p.ID
	role.Level = p.Level
	role.CreateBy = p.CreateBy
	role.UpdateBy = p.UpdatedBy
	role.Name = p.Name
	role.DataScope = p.DataScope
	role.Description = p.Description
	datatime, err := utils.UnixToTime(p.UpdateTime)
	role.UpdateTime = datatime
	if err != nil {
		return
	}
	datatime, err = utils.UnixToTime(p.CreateTime)
	role.UpdateTime = datatime
	if err != nil {
		return
	}
	if p.Protection == "true" {
		role.IsProtection = append(role.IsProtection, 1)
	} else {
		role.IsProtection = append(role.IsProtection, 0)
	}
	role.IsDeleted = append(role.IsDeleted, 0)
	if err = role.UpdateRole(); err != nil {
		return
	}
	return
}

// 删除角色
func (e Role) DeleteRole(p []int) (err error) {
	role := new(models.SysRole)
	if err = role.DeleteRole(p); err != nil {
		return
	}
	return
}

// 修改角色菜单
func (e Role) UpdateRoleMenu(id int, p []int) (err error) {
	role := new(models.SysRole)
	if err = role.UpdateRoleMenu(id, p); err != nil {
		return
	}
	return
}

// 获取单个角色
func (e Role) SelectRoleOne(id int) (roleone models.SysRole, err error) {
	role := new(models.SysRole)
	role.ID = id
	if roleone, err = role.SelectRoleOne(); err != nil {
		return
	}
	return
}
