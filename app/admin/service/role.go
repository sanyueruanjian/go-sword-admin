package service

import (
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/utils"
	"strconv"
)

type Role struct {
}

// 多条件查询角色
func (e Role) SelectRoles(p dto.SelectRoleArrayDto, orderData []bo.Order) (roleData bo.SelectRoleArrayBo, err error) {
	role := new(models.SysRole)
	sysRole, err := role.SelectRoles(p, orderData)
	if err != nil {
		return
	}
	if len(sysRole) > 0 {
		for _, value := range sysRole {
			var recordRole bo.RecordRole
			recordRole.CreateBy = value.CreateBy
			recordRole.ID = value.ID
			recordRole.Level = value.Level
			recordRole.UpdateBy = value.UpdateBy
			recordRole.CreateTime = value.CreateTime
			recordRole.DataScope = value.DataScope
			recordRole.Description = value.Description
			recordRole.Name = value.Name
			recordRole.UpdateTime = value.UpdateTime
			if value.IsProtection[0] == 1 {
				recordRole.Protection = true
			} else {
				recordRole.Protection = false
			}
			sysDept, sysMenu, errSysMenu := role.SysDeptAndMenu(value.ID)
			if errSysMenu != nil {
				err = errSysMenu
				return
			}
			deptList, menuList, errGetMenu := getDeptsMenus(sysDept, sysMenu)
			if errGetMenu != nil {
				err = errGetMenu
				return
			}
			recordRole.Depts = deptList
			recordRole.Menus = menuList
			if recordRole.Depts == nil {
				recordRole.Depts = make([]bo.Dept, 0)
			}
			if recordRole.Menus == nil {
				recordRole.Menus = make([]bo.Menu, 0)
			}
			roleData.Records = append(roleData.Records, recordRole)
		}
	}
	roleData.Current = p.Current
	roleData.Page = utils.PagesCount(int(role.RoleAllNum()), p.Size)
	roleData.SearchCount = true
	roleData.Size = p.Size
	roleData.Total = int(role.RoleAllNum())
	roleData.HitCount = false
	roleData.OptimizeCountSql = true
	for _, value := range orderData {
		var roleOrder bo.RoleOrder
		if value.Asc == "true" {
			roleOrder.Asc = true
		} else {
			roleOrder.Asc = false
		}
		roleOrder.Column = value.Column
		roleData.Orders = append(roleData.Orders, roleOrder)
	}
	return
}

func getDeptsMenus(sysDept []models.SysDept, sysMenu []models.SysMenu) (deptList []bo.Dept, menuList []bo.Menu, err error) {
	// Dept
	for _, value := range sysDept {
		var dept bo.Dept
		dept.CreateBy = value.CreateBy
		dept.CreateTime = value.CreateTime
		dept.DeptSort = value.DeptSort
		if value.Enabled[0] == 1 {
			dept.Enabled = true
		} else {
			dept.Enabled = false
		}
		if value.SubCount > 0 {
			dept.HasChildren = true
		} else {
			dept.HasChildren = false
		}
		dept.ID = value.ID
		dept.Name = value.Name
		dept.Pid = value.Pid
		dept.SubCount = value.SubCount
		dept.UpdateTime = value.UpdateTime
		dept.UpdateBy = value.UpdateBy
		deptList = append(deptList, dept)
	}
	// Menu
	for _, value := range sysMenu {
		var menu bo.Menu
		menu.CreateBy = value.CreateBy
		menu.Icon = value.Icon
		menu.ID = value.ID
		menu.MenuSort = value.MenuSort
		menu.Pid = value.Pid
		menu.SubCount = value.SubCount
		menu.Type = value.Type
		menu.UpdateBy = value.UpdateBy
		menu.Component = value.Component
		menu.CreateTime = value.CreateTime
		menu.Name = value.Name
		menu.Path = value.Path
		menu.Permission = value.Permission
		menu.Title = value.Title
		menu.UpdateTime = value.UpdateTime
		menu.Label = menu.Title
		if value.Cache[0] == 1 {
			menu.Cache = true
		} else {
			menu.Cache = false
		}
		if value.Hidden[0] == 1 {
			menu.Hidden = true
		} else {
			menu.Hidden = false
		}
		if value.IFrame[0] == 1 {
			menu.Iframe = true
		} else {
			menu.Iframe = false
		}
		menuList = append(menuList, menu)
	}
	return
}

// 新增角色
func (e Role) InsertRole(p dto.InsertRoleDto, userId int) (err error) {
	role := new(models.SysRole)
	role.Level = p.Level
	role.Name = p.Name
	role.DataScope = p.DataScope
	role.Description = p.Description
	role.CreateBy = userId
	role.UpdateBy = userId
	if err = role.InsertRole(p.Depts); err != nil {
		return
	}
	return
}

// 修改角色
func (e Role) UpdateRole(p dto.UpdateRoleDto, userId int) (err error) {
	role := new(models.SysRole)
	role.ID = p.ID
	role.Level = p.Level
	role.CreateBy = p.CreateBy
	role.UpdateBy = p.UpdatedBy
	role.Name = p.Name
	role.DataScope = p.DataScope
	role.Description = p.Description
	role.UpdateBy = userId
	role.UpdateTime = p.UpdateTime
	if p.Protection == "true" {
		role.IsProtection = append(role.IsProtection, 1)
	} else {
		role.IsProtection = append(role.IsProtection, 0)
	}
	role.IsDeleted = append(role.IsDeleted, 0)
	if err = role.UpdateRole(p.Depts, p.Menus); err != nil {
		return
	}

	// 删除缓存
	if err = models.DeleteRoleCache(role.ID); err != nil {
		return
	}
	if err = models.DeleteRoleAll(); err != nil {
		return
	}
	// 更新单个role缓存
	// TODO

	return
}

// 删除角色
func (e Role) DeleteRole(p []int, userId int) (err error) {
	role := new(models.SysRole)
	role.ID = userId
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
func (e Role) SelectRoleOne(id int) (roleData bo.RecordRole, err error) {
	role := new(models.SysRole)
	role.ID = id
	roleOne, err := role.SelectRoleOne()
	if err != nil {
		return
	}
	sysDept, sysMenu, err := role.SysDeptAndMenu(roleOne.ID)
	if err != nil {
		return
	}
	roleData.CreateBy = roleOne.CreateBy
	roleData.ID = roleOne.ID
	roleData.Level = roleOne.Level
	roleData.UpdateBy = roleOne.UpdateBy
	roleData.CreateTime = roleOne.CreateTime
	roleData.DataScope = roleOne.DataScope
	roleData.Description = roleOne.Description
	roleData.Name = roleOne.Name
	roleData.UpdateTime = roleOne.UpdateTime
	if roleOne.IsProtection[0] == 1 {
		roleData.Protection = true
	} else {
		roleData.Protection = false
	}
	deptList, menuList, err := getDeptsMenus(sysDept, sysMenu)
	if err != nil {
		return
	}
	// Depts
	roleData.Depts = deptList
	// Menu
	roleData.Menus = menuList
	return
}

// 获取所有角色
func (e Role) SelectRoleAll() (roleAll []bo.RecordRole, err error) {
	// 1.查找redis缓存
	roleAll, err = models.SelectRoleAllCache()
	if err == nil && len(roleAll) > 0 {
		return
	}

	// 2. 查找mysql数据
	role := new(models.SysRole)
	roleAll, err = role.SelectRoleAll()

	// 3.加入Redis缓存 all
	err = models.InsertRoleAll(roleAll)
	return
}

// 获取当前登录用户级别
func (e Role) SelectRoleLevel(roleName []string) (level bo.SelectCurrentUserLevel, err error) {
	role := new(models.SysRole)
	level, err = role.SelectRoleLevel(roleName)
	return
}

// 导出角色数据
func (e Role) DownloadRoleInfoBo(p dto.SelectRoleArrayDto, orderData []bo.Order) (roleData []bo.DownloadRoleInfoBo, err error) {
	role := new(models.SysRole)
	sysRole, err := role.SelectRoles(p, orderData)
	if err != nil {
		return
	}
	for _, values := range sysRole {
		var role bo.DownloadRoleInfoBo
		createTime, errTime := utils.UnixToTime(strconv.FormatInt(values.CreateTime, 10))
		if errTime != nil {
			err = errTime
			return
		}
		role.CreateTime = createTime
		role.Name = values.Name
		role.Level = values.Level
		role.Description = values.Description
		roleData = append(roleData, role)
	}
	return
}
