package api

import (
	"encoding/json"
	"errors"
	"project/app/admin/models"
	"project/app/admin/service"

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
	jobs           *[]models.SysJob
	roles          *[]models.SysRole
	menuPermission *[]string
	dept           *models.SysDept
	dataScopes     *[]int
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

//TODO GetCurrentUserInfo 获取当前登录的用户信息
func GetCurrentUserInfo(c *gin.Context) (*models.RedisUserInfo, error) {
	res, ok := c.Get(CtxUserInfoKey)
	if !ok {
		err := ErrorUserNotLogin
		return nil, err
	}
	userInfo := res.(*models.RedisUserInfo)
	return userInfo, nil
}

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

	jobs := new([]models.SysJob)
	if err = service.GetUserJobData(jobs, userId); err != nil {
		return nil, err
	}

	roles := new([]models.SysRole)
	if err = service.GetUserRoleData(roles, userId); err != nil {
		return nil, err
	}

	menuPermission := new([]string)
	if err = service.GetUserMenuData(userId, menuPermission, roles); err != nil {
		return nil, err
	}

	dept := new(models.SysDept)
	if err = service.GetUserDeptData(dept, userId); err != nil {
		return nil, err
	}

	dataScopes := new([]int)
	if err = service.GetUserDataScopes(dataScopes, userId, dept.ID, roles); err != nil {
		return nil, err
	}

	user = new(UserInfo)
	user.jobs = jobs
	user.roles = roles
	user.menuPermission = menuPermission
	user.dept = dept
	user.dataScopes = dataScopes
	return
}

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
