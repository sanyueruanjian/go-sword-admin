package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

// orm对外全局变量
var Eloquent *gorm.DB

// redis对外全局变量
var Rdb *redis.Client

//Casbin对外全局边变量
var CasbinEnforcer *casbin.SyncedEnforcer
