package middleware

import (
	"net/http"
	"project/common/api"
	mycasbin "project/pkg/casbin"
	"project/utils"

	"github.com/gin-gonic/gin"
)

//权限检查中间件
// TODO 前端需要修改
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
		res, err = e.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.Abort()
			return
		}

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
