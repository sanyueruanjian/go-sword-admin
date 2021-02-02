package main

import (
	"flag"
	"fmt"
	mycasbin "project/pkg/casbin"
	"project/utils"

	"project/common/database/mysql"
	"project/common/database/redis"
	"project/common/logger"
	"project/common/run"
	_ "project/docs"
	"project/utils/config"

	"go.uber.org/zap"
)

// @title go-sword项目接口文档
// @version 0.1.0
// @description 基于gin的后台通用框架

// @contact.name marchsoft@golang
// @contact.url http://marchsoft.cn/

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host 127.0.0.1:8000
// @BasePath
func main() {
	var configName string
	flag.StringVar(&configName, "o", "settings.dev.yml", "环境配置")
	flag.Parse()
	config.Setup("settings/" + configName)

	// 2. 初始化日志
	if err := logger.Init(config.LoggerConfig, config.ApplicationConfig.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug(utils.Green("logger init success..."))

	// 3. 初始化MySQL连接
	if err := mysql.Init(config.MysqlConfig); err != nil {
		zap.L().Error(fmt.Sprintf("init mysql failed, err:%v\n", err))
		return
	}
	defer mysql.Close()
	zap.L().Debug(utils.Green("mysql init success..."))

	// 4. 初始化Redis连接
	if err := redis.Init(config.RedisConfig); err != nil {
		zap.L().Error(fmt.Sprintf("init redis failed, err:%v\n", err))
		return
	}
	defer redis.Close()
	zap.L().Debug(utils.Green("redis init success..."))

	//初始化casbin
	if err := mycasbin.Setup(); err != nil {
		zap.L().Error("casbin failed set up", zap.Error(err))
	}
	// 5. 注册路由
	run.Run()

}
