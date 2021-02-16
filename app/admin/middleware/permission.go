package middleware

import (
	"fmt"
	"net/http"
	"project/app/admin/models"
	"project/common/api"
	mycasbin "project/pkg/casbin"

	"github.com/gin-gonic/gin"
)

//权限检查中间件
// TODO 需要修改
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get(api.CtxUserInfoKey)
		v := data.(models.RedisUserInfo)
		e, err := mycasbin.Casbin()
		if err != nil {
			c.Abort()
			return
		}
		//检查权限
		//此处为多角色 要在做处理
		var res bool
		res, err = e.Enforce(v.Role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.Abort()
			return
		}

		fmt.Printf("%s [INFO] %s %s \r\n",
			//tools.GetCurrentTimeStr(),
			c.Request.Method,
			c.Request.URL.Path,
			v.Role,
		)

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
