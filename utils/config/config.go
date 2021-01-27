package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Mysql配置项
var cfgMysql *viper.Viper

// 应用配置项
var cfgApplication *viper.Viper

// Token配置项
var cfgJwt *viper.Viper

// Log配置项
var cfgLogger *viper.Viper

// Redis配置项
var cfgRedis *viper.Viper

//载入配置文件
func Setup(path string) {
	// 加载配置文件
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgLogger = viper.Sub("settings.logger")
	if cfgLogger == nil {
		panic("No found settings.logger in the configuration")
	}
	LoggerConfig = InitLogger(cfgLogger)

	cfgMysql = viper.Sub("settings.mysql")
	if cfgMysql == nil {
		panic("No found settings.mysql in the configuration")
	}
	MysqlConfig = InitMySql(cfgMysql)

	cfgJwt = viper.Sub("settings.jwt")
	if cfgJwt == nil {
		panic("No found settings.jwt in the configuration")
	}
	JwtConfig = InitJwt(cfgJwt)

	cfgRedis = viper.Sub("settings.redis")
	if cfgRedis == nil {
		panic("No found settings.redis in the configuration")
	}
	RedisConfig = InitRedis(cfgRedis)
}