package mycasbin

import (
	"project/common/global"

	"go.uber.org/zap"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup() (err error) {
	//var err error
	var Apter *gormAdapter.Adapter
	var m model.Model
	var e *casbin.SyncedEnforcer

	// Initialize a gorm adapter with MySQL database. 使用MySQL数据库初始化gorm适配器。
	Apter, err = gormAdapter.NewAdapterByDB(global.Eloquent)
	if err != nil {
		zap.L().Error("NewAdapterByDB()", zap.Error(err))
		return err
	}

	// NewModelFromString从包含模型文本的字符串创建模型
	m, err = model.NewModelFromString(text)
	if err != nil {
		zap.L().Error("NewModelFromString()", zap.Error(err))
		return err
	}

	//NewSyncedEnforcer通过file或DB创建一个同步强制器
	e, err = casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		zap.L().Error("NewSyncedEnforcer()", zap.Error(err))
		return err
	}

	global.CasbinEnforcer = e
	return nil
}

func Casbin() (*casbin.SyncedEnforcer, error) {
	if err := global.CasbinEnforcer.LoadPolicy(); err == nil {
		return global.CasbinEnforcer, err
	} else {
		zap.L().Error("casbin rbac_model or policy init error, message: %v \r\n", zap.Error(err))
		return nil, err
	}
}
