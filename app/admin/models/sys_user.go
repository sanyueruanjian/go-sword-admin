package models

import (
	"errors"
	orm "project/common/global"
	"project/utils"
	"time"

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
	Username string `json:"username"`
}

type PassWord struct {
	// 密码
	Password string `json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserId struct {
	ID int `gorm:"primary_key"  json:"id"` // ID
}

type SysUserB struct {
	DeptId       int       `json:"dept_id"`        //部门id
	NickName     string    `json:"nick_name"`      //
	Gender       uint8     `json:"gender"`         //性别（0为男默认，1为女）
	Phone        string    `json:"phone"`          //
	Email        string    `json:"email"`          //
	AvatarPath   string    `json:"avatar_path"`    //头像路径
	IsAdmin      uint8     `json:"is_admin"`       //是否为admin账号
	Enabled      uint8     `json:"enabled"`        //状态：1启用（默认）、0禁用
	CreateBy     string    `json:"create_by"`      //
	UpdateBy     string    `json:"update_by"`      //
	PwdResetTime time.Time `json:"pwd_reset_time"` //修改密码的时间
	CreateTime   time.Time `json:"create_time"`    //创建日期
	UpdateTime   time.Time `json:"update_time"`    //更新时间
	IsDeleted    uint8     `json:"is_deleted"`     //软删除（默认值为0，1为删除）
	CreatedAt    time.Time `json:"created_at"`     //
	Avatar       string    `json:"avatar"`         //
	Sex          string    `json:"sex"`            //
	Status       string    `json:"status"`         //
	DeletedAt    time.Time `json:"deleted_at"`     //
	Remark       string    `json:"remark"`         //
	UpdatedAt    time.Time `json:"updated_at"`     //
	Salt         string    `json:"salt"`           //
	PostId       int       `json:"post_id"`        //
	RoleId       int       `json:"role_id"`        //
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
