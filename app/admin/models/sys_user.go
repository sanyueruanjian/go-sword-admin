package models

import (
	"encoding/json"
	"errors"
	"project/common/global"
	"project/utils"
	"strconv"
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
	PostId       int       `json:"post_id"`        //
	RoleId       int       `json:"role_id"`        //
	NickName     string    `json:"nick_name"`      //
	Phone        string    `json:"phone"`          //
	Email        string    `json:"email"`          //
	AvatarPath   string    `json:"avatar_path"`    //头像路径
	CreateBy     string    `json:"create_by"`      //
	UpdateBy     string    `json:"update_by"`      //
	Avatar       string    `json:"avatar"`         //
	Sex          string    `json:"sex"`            //
	Status       string    `json:"status"`         //
	Remark       string    `json:"remark"`         //
	Salt         string    `json:"salt"`           //
	Gender       []byte    `json:"gender"`         //性别（0为男默认，1为女）
	IsAdmin      []byte    `json:"is_admin"`       //是否为admin账号
	IsDeleted    []byte    `json:"is_deleted"`     //软删除（默认值为0，1为删除）
	Enabled      []byte    `json:"enabled"`        //状态：1启用（默认）、0禁用
	CreatedAt    time.Time `json:"created_at"`     //
	PwdResetTime time.Time `json:"pwd_reset_time"` //修改密码的时间
	CreateTime   time.Time `json:"create_time"`    //创建日期
	UpdateTime   time.Time `json:"update_time"`    //更新时间
	DeletedAt    time.Time `json:"deleted_at"`     //
	UpdatedAt    time.Time `json:"updated_at"`     //
}

type SysUser struct {
	SysUserId
	LoginM
	SysUserB
}

//redis 缓存model
type RedisUserInfo struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	DeptId   int    `json:"dept_id"` //部门id
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
	role := new(SysRole)
	err = global.Eloquent.Table(e.TableName()).Where("username = ?", e.Username).First(&e).Error
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

	//连表查询角色
	err = global.Eloquent.Table("sys_role").Select("name").Joins("left join sys_users_roles on "+
		"sys_users_roles.role_id = sys_role.id").Joins("left join sys_user on sys_user.id = sys_users_roles.user_id").Where("sys_user.id=?", e.ID).Scan(&role.Name).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户无角色", zap.Error(err))
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return ErrorServerBusy
	}

	var userInfo []byte
	userInfo, err = json.Marshal(RedisUserInfo{
		UserId:   e.ID,
		UserName: e.Username,
		DeptId:   e.DeptId,
		Role:     role.Name,
	})
	if err != nil {
		zap.L().Error("RedisUserInfo Marshal failed", zap.Error(err))
	}
	//添加缓存
	if err := global.Rdb.Set(strconv.Itoa(e.ID), userInfo, 0).Err(); err != nil {
		zap.L().Error("用户缓存错误", zap.Error(err))
		return err
	}
	return
}
