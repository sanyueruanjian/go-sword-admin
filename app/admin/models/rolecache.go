package models

import (
	"encoding/json"
	"go.uber.org/zap"

	"project/app/admin/models/bo"
	orm "project/common/global"
	"strconv"
)

// -----------------------------/api/roles/all-----------------------------
// 查找所有角色缓存
func SelectRoleAllCache() (roleAll []bo.RecordRole, err error) {
	// 1. 先找redis all
	val, err := orm.Rdb.Get("rolesAll").Result()
	if val != "" && err == nil {
		err = json.Unmarshal([]byte(val), &roleAll)
		if err == nil {
			return
		}
	}
	// 2. 再找redis单个组织数据
	values, _ := orm.Rdb.Keys("role::id:*").Result()
	value, err := orm.Rdb.MGet(values...).Result()
	if err == nil {
		for _, val := range value {
			var role bo.RecordRole
			err = json.Unmarshal([]byte(val.(string)), &role)
			if err != nil {
				return
			}
			roleAll = append(roleAll, role)
		}
	}
	err = InsertRoleAll(roleAll)
	return
}

// 添加RoleAll缓存
func InsertRoleAll(sysRoleAll []bo.RecordRole) (err error) {
	roleByte, err := json.Marshal(sysRoleAll)
	roleString := string(roleByte)
	if err != nil {
		return
	}
	errRedis := orm.Rdb.Set("rolesAll", roleString, 0).Err()
	if errRedis != nil {
		err = errRedis
		zap.L().Error("redis error: ", zap.Error(errRedis))
	}
	return
}

// 添加所有角色缓存
func InsertRoleAllCache(sysRoleAll []bo.RecordRole) (err error) {
	for _, values := range sysRoleAll {
		roleByte, errValue := json.Marshal(values)
		roleString := string(roleByte)
		if errValue != nil {
			err = errValue
			return
		}
		errRedis := orm.Rdb.Set("role::id:"+strconv.Itoa(values.ID), roleString, 0).Err()
		if errRedis != nil {
			err = errRedis
			zap.L().Error("redis error: ", zap.Error(errRedis))
		}
	}
	return
}

// -----------------------------/api/roles/{Post Delete Put}-----------------------------
func DeleteRoleCache(roleId int) (err error) {
	_, err = orm.Rdb.Do("DEL", "role::id:"+strconv.Itoa(roleId)).Result()
	return
}

// 但角色存入缓存
func InsertRoleId(roleId int) (err error) {

	var roleData bo.RecordRole
	role := new(SysRole)
	role.ID = roleId
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

	// 缓存
	roleIdData, err := role.SelectRoleOne()
	roleByte, errValue := json.Marshal(roleIdData)
	roleString := string(roleByte)
	if errValue != nil {
		err = errValue
		return
	}
	errRedis := orm.Rdb.Set("role::id:"+strconv.Itoa(roleIdData.ID), roleString, 0).Err()
	if errRedis != nil {
		err = errRedis
		zap.L().Error("redis error: ", zap.Error(errRedis))
	}
	return
}

// 删除RoleAll缓存
func DeleteRoleAll() (err error) {
	_, err = orm.Rdb.Do("DEL", "rolesAll").Result()
	return
}
