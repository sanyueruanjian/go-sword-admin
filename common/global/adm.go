package global

import (
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

// orm对外全局变量
var Eloquent *gorm.DB
// redis对外全局变量
var Rdb *redis.Client
