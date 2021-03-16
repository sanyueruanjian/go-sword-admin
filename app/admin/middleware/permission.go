package middleware

import (
	"project/common/api"
	mycasbin "project/pkg/casbin"
	"project/utils"
	"project/utils/app"

	"github.com/gin-gonic/gin"
)

//权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := new(api.UserInfo)
		var role []string
		var err error
		userInfo, err := api.GetUserData(c)
		data = userInfo
		roles := data.Roles
		for _, v := range *roles {
			role = append(role, utils.IntToString(v.ID))
		}
		if err != nil {
			c.Abort()
			return
		}
		e, err := mycasbin.LoadPolicy()
		if err != nil {
			c.Abort()
			return
		}
		//检查权限
		//此处为多角色 要在做处理
		var res bool
		for _, roleID := range role {
			res, err = e.Enforce(roleID, c.Request.URL.Path, c.Request.Method)
			if err != nil {
				c.Abort()
				return
			}
			if res {
				break
			}
		}

		if res {
			c.Next()
		} else {
			app.ResponseError(c, app.CodeIdentityNotRow)
			c.Abort()
			return
		}
	}
}

func CheckLevel() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := new(api.UserInfo)
		var role []string
		var err error
		userInfo, err := api.GetUserData(c)
		data = userInfo
		roles := data.Roles
		for _, v := range *roles {
			role = append(role, utils.IntToString(v.ID))
		}
		if err != nil {
			c.Abort()
			return
		}
		e, err := mycasbin.LoadPolicy()
		if err != nil {
			c.Abort()
			return
		}
		//检查权限
		//此处为多角色 要在做处理
		var res bool
		for _, roleID := range role {
			res, err = e.Enforce(roleID, c.Request.URL.Path, c.Request.Method)
			if err != nil {
				c.Abort()
				return
			}
			if res {
				break
			}
		}

		if res {
			c.Next()
		} else {
			app.ResponseError(c, app.CodeIdentityNotRow)
			c.Abort()
			return
		}
	}
}
