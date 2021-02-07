package cache

import (
	"fmt"
	"project/app/admin/models"
	"project/utils"

	"project/common/global"
)

const (
	KeyUserJob = "job::user:"
	KeyUserRole = "role::user:"
	KeyUserMenu = "menu::user:"
	KeyUserDept = "dept::user:"
	KeyUserDataScope = "data::user:"
)

func GetUserCache(userId int, cacheKey string) (cache string, err error) {
	return global.Rdb.Get(fmt.Sprintf("%s%d", cacheKey, userId)).Result()
}

func SetUserJobCache(userId int, jobs *[]*models.SysJob, cacheKey string) {
	s, e := global.Rdb.Set(fmt.Sprintf("%s%d", cacheKey, userId), *jobs, 0).Result()
	fmt.Println(s, e)
}

func SetUserCache(userId int, data interface{}, cacheKey string) {
	res, err := utils.StructToJson(data)
	if err != nil {
		return
	}
	global.Rdb.Set(fmt.Sprintf("%s%d", cacheKey, userId), res, 0)
}
