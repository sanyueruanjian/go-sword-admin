package redis

import (
	"fmt"
	"project/common/global"
	"time"

	"project/utils/config"

	"github.com/go-redis/redis/v7"
)

// Init 初始化redis连接
func Init(cfg *config.Redis) (err error) {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
		IdleTimeout: time.Duration(cfg.IdleTimeOutSec),
	})

	_, err = global.Rdb.Ping().Result()
	return
}

func Close() {
	_ = global.Rdb.Close()
}
