package config

import "github.com/spf13/viper"

type Redis struct {
	PoolSize      int
	IdleTimeOutSec int
	DB             int
	Port           int
	Password       string
	Host           string
}

// InitRedis 初始化redis配置
func InitRedis(cfg *viper.Viper) *Redis {

	db := &Redis{
		PoolSize:        cfg.GetInt("poolsize"),
		IdleTimeOutSec: cfg.GetInt("idletimeoutsec"),
		DB:             cfg.GetInt("db"),
		Port:           cfg.GetInt("port"),
		Host:           cfg.GetString("host"),
		Password:       cfg.GetString("password"),
	}
	return db
}

var RedisConfig = new(Redis)
