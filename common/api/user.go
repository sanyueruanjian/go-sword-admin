package api

import (
	"errors"
	"project/app/admin/models"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "user_id"

const CtxUserInfoKey = "info"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
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
