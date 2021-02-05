package models

import (
	"encoding/json"
	"errors"
	"project/app/admin/models/bo"
	"project/common/global"
	"project/utils"
	"strconv"

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

type SysUser struct {
	*BaseModel
	Username     string `json:"username"`
	Password     string `json:"password"`
	DeptId       int    `json:"dept_id"`        //部门id
	PostId       int    `json:"post_id"`        //
	RoleId       int    `json:"role_id"`        //
	NickName     string `json:"nick_name"`      //
	Phone        string `json:"phone"`          //
	Email        string `json:"email"`          //
	AvatarPath   string `json:"avatar_path"`    //头像路径
	Avatar       string `json:"avatar"`         //
	Sex          string `json:"sex"`            //
	Status       string `json:"status"`         //
	Remark       string `json:"remark"`         //
	Salt         string `json:"salt"`           //
	Gender       []byte `json:"gender"`         //性别（0为男默认，1为女）
	IsAdmin      []byte `json:"is_admin"`       //是否为admin账号
	Enabled      []byte `json:"enabled"`        //状态：1启用（默认）、0禁用
	PwdResetTime int64  `json:"pwd_reset_time"` //修改密码的时间
	CreateBy     int    `json:"create_by"`      //
	UpdateBy     int    `json:"update_by"`      //
}

//redis 缓存model
type RedisUserInfo struct {
	UserId   int      `json:"user_id"`
	UserName string   `json:"user_name"`
	Role     []string `json:"role"`
	DeptId   int      `json:"dept_id"` //部门id
}

type RecordUser struct {
	Id           int    `json:"id"`
	DeptId       int    `json:"deptId"`
	PwdResetTime int    `json:"pwdResetTime"`
	CreateBy     int    `json:"createBy"`
	CreateTime   int    `json:"createTime"`
	UpdatedBy    int    `json:"updatedBy"`
	UpdateTime   int    `json:"updateTime"`
	AvatarName   string `json:"avatarName"`
	AvatarPath   string `json:"avatarPath"`
	Email        string `json:"email"`
	NickName     string `json:"nickName"`
	Phone        string `json:"phone"`
	Username     string `json:"username"`
	Enabled      bool   `json:"enabled"`
	Gender       bool   `json:"gender"`
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
func (u *SysUser) Login() (err error) {
	oPassword := u.Password
	err = global.Eloquent.Table(u.TableName()).Where("username = ?", u.Username).First(&u).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户不存在", zap.Error(err))
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return ErrorServerBusy
	}
	if u.Password != utils.EncodeMD5(oPassword) {
		zap.L().Error("user account or password is error")
		return ErrorInvalidPassword
	}

	//连表查询角色
	var roles []*bo.Role
	roles, err = u.SelectUserRole()
	//构造角色名字集合
	roleNames := make([]string, 0)
	for _, v := range roles {
		roleNames = append(roleNames, v.Name)
	}
	//初始化缓存模型
	var userInfo []byte
	userInfo, err = json.Marshal(RedisUserInfo{
		UserId:   u.ID,
		UserName: u.Username,
		DeptId:   u.DeptId,
		Role:     roleNames,
	})
	if err != nil {
		zap.L().Error("RedisUserInfo Marshal failed", zap.Error(err))
	}
	//添加缓存
	if err := global.Rdb.Set(strconv.Itoa(u.ID), userInfo, 0).Err(); err != nil {
		zap.L().Error("用户缓存错误", zap.Error(err))
		return err
	}
	return
}

func (u *SysUser) InsertUser(jobs []int, roles []int) (err error) {
	//创建事务
	tx := global.Eloquent.Begin()
	//用户表 增添
	err = tx.Table("sys_user").Create(u).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//维护 user role 关系表
	for _, role := range roles {
		roleUser := &SysUsersRoles{
			UserId: u.ID,
			RoleId: role,
		}
		err = tx.Table("sys_users_roles").Create(roleUser).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//维护 user job 关系表
	for _, job := range jobs {
		roleUser := &SysUsersJobs{
			UserId: u.ID,
			JobId:  job,
		}
		err = tx.Table("sys_users_jobs").Create(roleUser).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//提交事务
	return tx.Commit().Error
}

//TODO
//func (u *SysUser) SelectUserInfoList(p *dto.SelectUserInfoArrayDto) (data []*bo.UserInfoListBo, err error) {
//	//排序条件
//	var orderJson []bo.Order
//	orderJson, err = utils.OrderJson(p.Orders)
//	orderRule := utils.GetOrderRule(orderJson)
//	//查询用户基本信息
//	var users []*SysUser
//	err = global.Eloquent.Table("sys_user").Limit(p.Size).Offset(p.Current - 1*p.Size).Order(orderRule).Find(&users).Error
//	if err != nil {
//		return nil, err
//	}
//	for _, user := range users {
//		//查询 角色 部门 岗位
//		var roles []*bo.Role
//		var jobs []*bo.Job
//		var depts []*bo.DeptCommon
//		roles, err = user.SelectUserRole()
//		jobs, err = user.SelectUserJob()
//		depts, err = user.SelectUserDept()
//		userInfo := &bo.UserInfoListBo{
//			Records:
//		}
//	}
//}

func (u *SysUser) SelectUserRole() (role []*bo.Role, err error) {
	//连表查询角色
	err = global.Eloquent.Table("sys_role").Joins("left join sys_users_roles on "+
		"sys_users_roles.role_id = sys_role.id").Joins("left join sys_user on sys_user.id = sys_users_roles.user_id").Where("sys_user.id=?", u.ID).Find(&role).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户无角色", zap.Error(err))
		return nil, ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return nil, ErrorServerBusy
	}
	return
}

func (u *SysUser) SelectUserJob() (jobs []*bo.Job, err error) {
	//连表查询岗位
	err = global.Eloquent.Table("sys_job").Joins("left join sys_users_jobs on "+
		"sys_users_jobs.job_id = sys_job.id").Joins("left join sys_user on sys_user.id = sys_users_jobs.user_id").Where("sys_user.id=?", u.ID).Find(&jobs).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户无岗位", zap.Error(err))
		return nil, ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return nil, ErrorServerBusy
	}
	return
}

func (u *SysUser) SelectUserDept() (jobs []*bo.DeptCommon, err error) {
	//连表查询岗位
	err = global.Eloquent.Table("sys_dept").Joins("left join sys_users_depts on "+
		"sys_users_depts.dept_id = sys_dept.id").Joins("left join sys_user on sys_user.id = sys_users_depts.user_id").Where("sys_user.id=?", u.ID).Find(&jobs).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户无部门", zap.Error(err))
		return nil, ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return nil, ErrorServerBusy
	}
	return
}
