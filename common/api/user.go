package api

import (
	"encoding/json"
	"errors"
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
	Jobs           *[]models.SysJob  //用户岗位
	Roles          *[]models.SysRole //用户角色
	MenuPermission *[]string         //菜单权限
	Dept           *models.SysDept   //部门
	DataScopes     *[]int            //数据权限
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

//TODO GetCurrentUserInfo 获取当前登录的用户信息  需要删除
func GetCurrentUserInfo(c *gin.Context) (*models.RedisUserInfo, error) {
	res, ok := c.Get(CtxUserInfoKey)
	if !ok {
		err := ErrorUserNotLogin
		return nil, err
	}
	userInfo := res.(*models.RedisUserInfo)
	return userInfo, nil
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
