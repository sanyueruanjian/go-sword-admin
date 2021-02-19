package service

import (
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/utils"
	"strconv"
	"sync"
)

type Role struct {
}

// 多条件查询角色
func (e Role) SelectRoles(p dto.SelectRoleArrayDto, orderData []bo.Order) (roleData bo.SelectRoleArrayBo, err error) {
	role := new(models.SysRole)
	sysRole, status, err := role.SelectRoles(p, orderData)
	if err != nil {
		return
	}
	// 1. 查询缓存All
	if status != 1 {
		roleAll, err := models.SelectRoleAllCaches()
		if err == nil && len(roleAll) > 0 {
			// 分页
			if p.Current*p.Size > len(roleAll) {
				roleData.Records = roleAll[(p.Current-1)*p.Size : len(roleAll)]
			} else {
				roleData.Records = roleAll[(p.Current-1)*p.Size : p.Current*p.Size]
			}
			sysRole = nil
		}
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

			// 查询缓存
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				recordRole.Depts, err = models.InsertDept(role, value.ID)
				if err != nil {
					return
				}
				wg.Done()
			}()
			go func() {
				recordRole.Menus, err = models.InsertMenu(role, value.ID)
				if err != nil {
					return
				}
				wg.Done()
			}()
			wg.Wait()
			if recordRole.Depts == nil {
				recordRole.Depts = make([]bo.Dept, 0)
			}
			if recordRole.Menus == nil {
				recordRole.Menus = make([]bo.Menu, 0)
			}
			roleData.Records = append(roleData.Records, recordRole)
		}

		// 存入RoleAll缓存
		if status != 1 {
			_ = models.InsertRoleAlls(roleData.Records)
		}
	}

	roleData.Current = p.Current
	if status == 1 {
		roleData.Page = utils.PagesCount(len(roleData.Records), p.Size)
		roleData.Total = len(roleData.Records)
	} else {
		roleData.Page = utils.PagesCount(int(role.RoleAllNum()), p.Size)
		roleData.Total = int(role.RoleAllNum())
	}
	roleData.SearchCount = true
	roleData.Size = p.Size
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

	// 删除缓存
	if err = models.DeleteRoleAll(); err != nil {
		return
	}
	err = models.DeleteRoleAlls()
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
	if err = models.DeleteDeptCache(role.ID); err != nil {
		return
	}
	if err = models.DeleteRoleAll(); err != nil {
		return
	}
	if err = models.DeleteRoleAlls(); err != nil {
		return
	}
	// 更新单个role缓存
	err = models.InsertRoleId(p.ID)
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

	// 删除缓存
	if err = models.DeleteRoleAll(); err != nil {
		return
	}
	if err = models.DeletMenuCache(id); err != nil {
		return
	}
	err = models.DeleteRoleAlls()
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
	deptList, menuList, err := role.GetDeptsMenus(sysDept, sysMenu)
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
	sysRole, _, err := role.SelectRoles(p, orderData)
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
