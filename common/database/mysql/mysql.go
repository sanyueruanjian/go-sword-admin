package mysql

import (
	"database/sql"
	"fmt"

	"project/common/global"
	"project/utils"
	"project/utils/config"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Init 配置mysql gorm
func Init(cfg *config.Mysql) (err error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=2000ms",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	zap.L().Info(utils.Green(source))
	db, err := sql.Open("mysql", source)
	if err != nil {
		zap.L().Fatal(utils.Red("mysql connect error :"), zap.Error(err))
		return
	}

	db.SetMaxOpenConns(cfg.DbMaxOpen)
	db.SetMaxIdleConns(cfg.DbMaxIdle)

	global.Eloquent, err = open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Fatal(utils.Red("mysql connect error :"), zap.Error(err))
		return
	} else {
		zap.L().Info(utils.Green("mysql connect success !"))
	}

	if global.Eloquent.Error != nil {
		zap.L().Fatal(utils.Red(" database error :"), zap.Error(global.Eloquent.Error))
		return
	}

	err = migrateModel()
	if err != nil {
		zap.L().Fatal(utils.Red(" migrateModel error :"), zap.Error(err))
	}
	return
}

// Open 打开数据库连接
func open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// Close 关闭连接
func Close() {
	db, err := global.Eloquent.DB()
	if err != nil {
		zap.L().Error("db close err", zap.Error(err))
	}
	_ = db.Close()
}
