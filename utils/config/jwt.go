package config

import "github.com/spf13/viper"

type Jwt struct {
	Header      string
	RedisHeader string
	Secret      string
	Timeout     int64
}

// InitJwt 初始化jwt配置
func InitJwt(cfg *viper.Viper) *Jwt {
	return &Jwt{
		Header:      cfg.GetString("header"),
		RedisHeader: cfg.GetString("redis_header"),
		Secret:      cfg.GetString("secret"),
		Timeout:     cfg.GetInt64("timeout"),
	}
}

var JwtConfig = new(Jwt)
