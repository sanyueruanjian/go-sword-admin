package middleware

//import (
//	"github.com/casbin/casbin"
//	"github.com/gin-gonic/gin"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func CasBinMiddleWare(e *casbin.Enforcer) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		//获取请求的URI
//		obj := c.Request.URL.RequestURI()
//		//获取请求方法
//		act := c.Request.Method
//		//获取用户的角色
//		sub := "admin"
//
//		//判断策略中是否存在
//		if e.Enforce(sub, obj, act) {
//			c.Next()
//		} else {
//			c.Abort()
//		}
//	}
//}

//"github.com/casbin/gorm-adapter"
//a := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb", true)
//e := casbin.NewEnforcer("rabc.conf", a)
////从DB加载策略
//e.LoadPolicy()

//保存策略
//e.SavePolicy()

// e.AddPolicy(...)
// e.RemovePolicy(...)