package api

import (
	"encoding/json"
	"errors"
	"math"
	"project/app/admin/models"
	"project/app/admin/service"
	"project/common/cache"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIdAndName = "user"
	CtxUserIDKey     = "user_id"
	CtxUserInfoKey   = "info"
	CtxUserOnline    = "user_online"
)

type UserMessage struct {
	UserId   int
	Username string
}

type UserInfo struct {
	Jobs           *[]models.SysJob
	Roles          *[]models.SysRole
	MenuPermission *[]string
	Dept           *models.SysDept
	DataScopes     *[]int
}

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUserId 获取当前登录的用户ID
func GetCurrentUserId(c *gin.Context) (userId int, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetUserMessage 获取当前登录的用户ID和用户名
func GetUserMessage(c *gin.Context) (*UserMessage, error) {
	res, ok := c.Get(CtxUserIdAndName)
	if !ok {
		err := ErrorUserNotLogin
		return nil, err
	}
	userMessage := res.(*UserMessage)
	return userMessage, nil
}

// 获取用户完整信息
func GetUserData(c *gin.Context) (user *UserInfo, err error) {

	userId, err := GetCurrentUserId(c)
	if err != nil {
		return
	}

	keys := new([]string)
	*keys = append(*keys, cache.KeyUserJob, cache.KeyUserRole, cache.KeyUserMenu, cache.KeyUserDept, cache.KeyUserDataScope)
	cacheMap := cache.GetUserCache(keys, userId)

	cacheJob, jobErr := cacheMap[cache.KeyUserJob].Result()
	cacheRole, rolesErr := cacheMap[cache.KeyUserRole].Result()
	cacheMenu, menuErr := cacheMap[cache.KeyUserMenu].Result()
	cacheDept, deptErr := cacheMap[cache.KeyUserDept].Result()
	cacheDataScopes, dataScopesErr := cacheMap[cache.KeyUserDataScope].Result()
	jobs := new([]models.SysJob)
	if err = service.GetUserJobData(cacheJob, jobErr, jobs, userId); err != nil {
		return nil, err
	}

	roles := new([]models.SysRole)
	if err = service.GetUserRoleData(cacheRole, rolesErr, roles, userId); err != nil {
		return nil, err
	}

	menuPermission := new([]string)
	if err = service.GetUserMenuData(cacheMenu, menuErr, userId, menuPermission, roles); err != nil {
		return nil, err
	}

	dept := new(models.SysDept)
	if err = service.GetUserDeptData(cacheDept, deptErr, dept, userId); err != nil {
		return nil, err
	}

	dataScopes := new([]int)
	if err = service.GetUserDataScopes(cacheDataScopes, dataScopesErr, dataScopes, userId, dept.ID, roles); err != nil {
		return nil, err
	}

	user = new(UserInfo)
	user.Jobs = jobs
	user.Roles = roles
	user.MenuPermission = menuPermission
	user.Dept = dept
	user.DataScopes = dataScopes
	return
}

// GetUserOnline 获取用户线上数据
func GetUserOnline(c *gin.Context) (userOnline *models.OnlineUser, err error) {
	res, ok := c.Get(CtxUserOnline)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userOnline = new(models.OnlineUser)
	err = json.Unmarshal([]byte(res.(string)), userOnline)
	return
}
func CheckDataScope(dataScope []int, deptID int) bool {
	if len(dataScope) == 0 {
		return true
	}
	for _, v := range dataScope {
		if v == deptID {
			return true
		}
	}
	return false
}

func CheckLevel(operatorRoles []models.SysRole, id int) bool {
	//查找操作者最高等级
	operatorMaxLevel := math.MaxInt64
	for _, role := range operatorRoles {
		if role.Level < operatorMaxLevel {
			operatorMaxLevel = role.Level
		}
	}
	//根据id查找用户角色最高等级
	byOperateRoles, err := models.SelectUserRole(id)
	if err != nil {
		return false
	}
	byOperatorMaxLevel := math.MaxInt64
	for _, role := range byOperateRoles {
		if role.Level < byOperatorMaxLevel {
			byOperatorMaxLevel = role.Level
		}
	}
	return operatorMaxLevel < byOperatorMaxLevel
}
