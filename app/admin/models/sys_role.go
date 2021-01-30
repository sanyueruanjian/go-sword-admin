package models

import (
	"fmt"
	"project/utils"
	"time"

	orm "project/common/global"
)

type SysRole struct {
	ID           int       `gorm:"primary_key" json:"id"` //ID
	Name         string    `json:"name"`                  //角色名称
	Level        int       `json:"level"`                 //角色级别（越小越大）
	Description  string    `json:"description"`           //描述
	DataScope    string    `json:"data_scope"`            //数据权限
	IsProtection []byte    `json:"is_protection"`         //是否受保护（内置角色，1为内置角色，默认值为0）
	CreateBy     int       `json:"create_by"`             //创建者id
	UpdateBy     int       `json:"update_by"`             //更新者id
	CreateTime   time.Time `json:"create_time"`           //创建日期
	UpdateTime   time.Time `json:"update_time"`           //更新时间
	IsDeleted    []byte    `json:"is_deleted"`            //软删除（默认值为0，1为删除）
}

func (e SysRole) SelectRoles() (err error) {
	return
}

func (e SysRole) InsertRole() (err error) {
	e.IsProtection = append(e.IsProtection, 1)
	// TODO 获取当前用户id
	e.CreateBy = 1
	e.UpdateBy = 1
	e.CreateTime = utils.GetCurrentTime()
	e.UpdateTime = utils.GetCurrentTime()
	e.IsDeleted = append(e.IsDeleted, 0)
	err = orm.Eloquent.Create(&e).Error
	return
}

func (e SysRole) UpdateRole() (err error) {
	e.CreateBy = 1
	e.UpdateBy = 1
	e.CreateTime = utils.GetCurrentTime()
	e.UpdateTime = utils.GetCurrentTime()
	err = orm.Eloquent.Model(&e).Updates(e).Error
	return
}

// 删除角色
func (e SysRole) DeleteRole(p []int) (err error) {
	err = orm.Eloquent.Table("sys_role").Where("id = ?", p).Update("is_deleted", 1).Error
	return
}

func (e SysRole) UpdateRoleMenu(id int, p []int) (err error) {
	// TODO
	fmt.Println(id)
	fmt.Println(p)
	return
}

func (e SysRole) SelectRoleOne() (role SysRole, err error) {
	// TODO
	//SelectRoleBo
	err = orm.Eloquent.First(&role, e.ID).Error
	return
}
