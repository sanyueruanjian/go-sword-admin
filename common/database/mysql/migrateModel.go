package mysql

import (
	"project/app/admin/models"
	orm "project/common/global"
)

func migrateModel() error {
	err := orm.Eloquent.AutoMigrate(&models.SysUser{})
	return err
}