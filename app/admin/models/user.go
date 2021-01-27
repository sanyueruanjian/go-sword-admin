package models

import (
	"errors"
	orm "project/common/global"
	"project/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// User
type User struct {
	// key
	IdentityKey string
	// 用户名
	UserName  string
	FirstName string
	LastName  string
	// 角色
	Role string
}

type UserName struct {
	Username string `gorm:"size:64" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"size:128" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName  string `gorm:"size:128" json:"nickName"` // 昵称
	Phone     string `gorm:"size:11" json:"phone"`     // 手机号
	RoleId    int    `gorm:"" json:"roleId"`           // 角色编码
	Salt      string `gorm:"size:255" json:"salt"`     //盐
	Avatar    string `gorm:"size:255" json:"avatar"`   //头像
	Sex       string `gorm:"size:255" json:"sex"`      //性别
	Email     string `gorm:"size:128" json:"email"`    //邮箱
	DeptId    int    `gorm:"" json:"deptId"`           //部门编码
	PostId    int    `gorm:"" json:"postId"`           //职位编码
	CreateBy  string `gorm:"size:128" json:"createBy"` //
	UpdateBy  string `gorm:"size:128" json:"updateBy"` //
	Remark    string `gorm:"size:255" json:"remark"`   //备注
	Status    string `gorm:"size:4;" json:"status"`
	BaseModel

	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`
}

type SysUser struct {
	SysUserId
	LoginM
	SysUserB
}

func (SysUser) TableName() string {
	return "sys_user"
}

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorServerBusy      = errors.New("服务器繁忙")
)

// Login 查询用户是否存在，并验证密码
func (e SysUser) Login() (err error) {
	oPassword := e.Password
	err = orm.Eloquent.Table(e.TableName()).Where("username = ?", e.Username).First(&e).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户不存在", zap.Error(err))
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return ErrorServerBusy
	}
	if e.Password != utils.EncodeMD5(oPassword) {
		zap.L().Error("user account or password is error")
		return ErrorInvalidPassword
	}
	return
}

