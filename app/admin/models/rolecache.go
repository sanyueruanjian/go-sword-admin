package models

import (
	"encoding/json"
	"go.uber.org/zap"
	orm "project/common/global"
)

// 查询所有角色缓存
func RoleAllCache(sysRoleAll []SysRole) (err error) {
	for _, values := range sysRoleAll {
		roleByte, err := json.Marshal(values)
		roleString := string(roleByte)
		if err != nil {
			return
		}
		errRedis := orm.Rdb.Set("role::id:"+string(values.ID), roleString, 0).Err()
		if errRedis != nil {
			zap.L().Error("redis error: ", zap.Error(errRedis))
		}
	}
	return
}
