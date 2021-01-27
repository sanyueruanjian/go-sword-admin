package config

import (
	"github.com/spf13/viper"
)

type Logger struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Filename   string
	Level      string
	Stdout     bool
}

// InitLogger 初始化日志配置
func InitLogger(cfg *viper.Viper) *Logger {
	return &Logger{
		MaxSize:    cfg.GetInt("max_size"),
		MaxAge:     cfg.GetInt("max_age"),
		MaxBackups: cfg.GetInt("max_backups"),
		Filename:   cfg.GetString("filename"),
		Level:      cfg.GetString("level"),
		Stdout:     cfg.GetBool("stdout"),
	}
}

var LoggerConfig = new(Logger)
