package config

import "github.com/spf13/viper"

type Mysql struct {
	User      string
	Password  string
	Host      string
	DbName    string
	Port      int
	DbMaxOpen int
	DbMaxIdle int
}

// InitMySql 初始化mysql配置
func InitMySql(cfg *viper.Viper) *Mysql {

	db := &Mysql{
		User:      cfg.GetString("user"),
		Password:  cfg.GetString("password"),
		Host:      cfg.GetString("host"),
		DbName:    cfg.GetString("dbname"),
		Port:      cfg.GetInt("port"),
		DbMaxOpen: cfg.GetInt("maxopen"),
		DbMaxIdle: cfg.GetInt("maxidle"),
	}
	return db
}

var MysqlConfig = new(Mysql)
