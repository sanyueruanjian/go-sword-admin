package models

import (
	"encoding/json"
	"errors"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
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

type GenderEnabled struct {
	Gender  []byte `json:"gender"`  //性别（0为男默认，1为女）
	Enabled []byte `json:"enabled"` //状态：1启用（默认）、0禁用
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

func (SysUser) TableName() string {
	return "sys_user"
}

var (
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
	roles, err = SelectUserRole(u.ID)
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

func (u *SysUser) SelectUserInfoList(p *dto.SelectUserInfoArrayDto) (data *bo.UserInfoListBo, err error) {
	//排序条件
	var orderJson []bo.Order
	orderJson, err = utils.OrderJson(p.Orders)
	orderRule := utils.GetOrderRule(orderJson)
	//查询用户基本信息
	var userHalfs []*bo.RecordUserHalf
	var users []*bo.RecordUser
	err = global.Eloquent.Table("sys_user").Where("is_deleted=?", []byte{0}).Limit(p.Size).Offset(p.Current - 1*p.Size).Order(orderRule).Find(&userHalfs).Error
	if err != nil {
		return nil, err
	}
	for _, userHalf := range userHalfs {
		//查询 角色 部门 岗位
		var roles []*bo.Role
		var jobs []*bo.Job
		dept := new(bo.DeptCommon)
		user := new(bo.RecordUser)
		user.RecordUserHalf = new(bo.RecordUserHalf)
		user.RoleDeptJobBool = new(bo.RoleDeptJobBool)
		genderEnabled := new(GenderEnabled)
		//查询角色
		roles, err = SelectUserRole(userHalf.Id)
		if err != nil {
			zap.L().Debug("查询角色", zap.Error(err))
			return nil, err
		}
		//查询岗位
		jobs, err = SelectUserJob(userHalf.Id)
		if err != nil {
			zap.L().Debug("查询岗位", zap.Error(err))
			return nil, err
		}
		//查询部门
		err = global.Eloquent.Table("sys_dept").Joins("left join sys_user "+
			"on sys_user.dept_id = sys_dept.id").Where("sys_user.id=?", userHalf.Id).Scan(dept).Error
		if err != nil {
			zap.L().Debug("查询部门", zap.Error(err))
			return nil, err
		}
		//查询性别
		err = global.Eloquent.Table("sys_user").Select("gender", "enabled").Where("id=?", userHalf.Id).First(genderEnabled).Error
		if err != nil {
			zap.L().Debug("查询性别", zap.Error(err))
			return nil, err
		}
		user.Role = roles
		user.Jobs = jobs
		user.Dept = dept
		user.Id = userHalf.Id
		user.Phone = userHalf.Phone
		user.DeptId = userHalf.DeptId
		user.PwdResetTime = userHalf.PwdResetTime
		user.CreateBy = userHalf.CreateBy
		user.CreateTime = userHalf.CreateTime
		user.UpdateBy = userHalf.UpdateBy
		user.UpdateTime = userHalf.UpdateTime
		user.AvatarName = userHalf.AvatarName
		user.AvatarPath = userHalf.AvatarPath
		user.Email = userHalf.Email
		user.NickName = userHalf.NickName
		user.Phone = userHalf.Phone
		user.Username = userHalf.Username
		user.Enabled = utils.ByteIntoBool(genderEnabled.Enabled)
		user.Gender = utils.ByteIntoBool(genderEnabled.Gender)
		users = append(users, user)
	}
	data = &bo.UserInfoListBo{Records: users}
	return data, nil

}

func SelectUserRole(userId int) (role []*bo.Role, err error) {
	//连表查询角色
	err = global.Eloquent.Table("sys_role").Joins("left join sys_users_roles on "+
		"sys_users_roles.role_id = sys_role.id").Joins("left join sys_user on sys_user.id = sys_users_roles.user_id").Where("sys_user.id=?", userId).Find(&role).Error
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

func SelectUserJob(jobId int) (jobs []*bo.Job, err error) {
	//连表查询岗位
	err = global.Eloquent.Table("sys_job").Joins("left join sys_users_jobs on "+
		"sys_users_jobs.job_id = sys_job.id").Joins("left join sys_user on sys_user.id = sys_users_jobs.user_id").Where("sys_user.id=?", jobId).Find(&jobs).Error
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

func (u *SysUser) DeleteUser(ids *[]int) (err error) {
	return global.Eloquent.Table("sys_user").Where("id IN (?)", *ids).Updates(map[string]interface{}{"is_deleted": []byte{1}}).Error
}

func (u *SysUser) UpdateUser(p *dto.UpdateUserDto, optionId int) (err error) {
	//开始事务
	tx := global.Eloquent.Begin()
	//校验用户是否存在
	test := new(SysUser)
	err = tx.Table("sys_user").Where("id=? AND is_delete=?", p.ID, []byte{0}).First(test).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//更新 用户表
	err = tx.Table("sys_user").Where("id=?", p.ID).Updates(map[string]interface{}{
		"dept_id":     p.DeptId,
		"email":       p.Email,
		"nick_name":   p.NickName,
		"phone":       p.Phone,
		"username":    p.UserName,
		"avatar_path": p.AvatarPath,
		"enabled":     utils.StrBoolIntoByte(p.Enabled),
		"gender":      utils.StrBoolIntoByte(p.Gender),
		"update_by":   optionId,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//更新 角色用户 关系表
	//1删除原有关系
	err = tx.Table("sys_users_roles").Unscoped().Where("user_id=?", p.ID).Delete(&SysUsersRoles{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("sys_users_jobs").Unscoped().Where("user_id=?", p.ID).Delete(&SysUsersJobs{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//2增添现有关系
	//2.1 角色关系
	for _, role := range p.Roles {
		err = tx.Table("sys_users_roles").Create(&SysUsersRoles{
			RoleId: role,
			UserId: p.ID,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//2.1 岗位关系
	for _, job := range p.Jobs {
		err = tx.Table("sys_users_jobs").Create(&SysUsersJobs{
			JobId:  job,
			UserId: p.ID,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//提交事务
	return tx.Commit().Error
}

func (u *SysUser) UpdateUserCenter(p *dto.UpdateUserCenterDto, optionId int) (err error) {
	return global.Eloquent.Table("sys_user").Where("id=?", p.Id).Updates(map[string]interface{}{
		"gender":    utils.BoolIntoByte(p.Gender),
		"phone":     p.Phone,
		"nick_name": p.NickName,
		"update_by": optionId,
	}).Error

}

func (u *SysUser) SelectUserInfo(p *RedisUserInfo) (data *bo.UserCenterInfoBo, err error) {
	//查询用户基本信息
	var userHalf bo.RecordUserHalf
	err = global.Eloquent.Table("sys_user").Where("is_deleted=? AND id=?", []byte{0}, p.UserId).First(&userHalf).Error
	if err != nil {
		zap.L().Debug("查询基本信息", zap.Error(err))
		return nil, err
	}
	//查询 角色 部门 岗位
	var role []*bo.Role
	var job []*bo.Job
	dept := new(bo.DeptCommon)
	user := new(bo.RecordUser)
	user.RecordUserHalf = new(bo.RecordUserHalf)
	user.RoleDeptJobBool = new(bo.RoleDeptJobBool)
	genderEnabled := new(GenderEnabled)
	//查询角色
	role, err = SelectUserRole(p.UserId)
	if err != nil {
		zap.L().Debug("查询角色", zap.Error(err))
		return nil, err
	}
	//查询岗位
	job, err = SelectUserJob(p.UserId)
	if err != nil {
		zap.L().Debug("查询岗位", zap.Error(err))
		return nil, err
	}
	//查询部门
	err = global.Eloquent.Table("sys_dept").Joins("left join sys_user "+
		"on sys_user.dept_id = sys_dept.id").Where("sys_user.id=?", userHalf.Id).Scan(dept).Error
	if err != nil {
		zap.L().Debug("查询部门", zap.Error(err))
		return nil, err
	}
	//查询性别
	err = global.Eloquent.Table("sys_user").Select("gender", "enabled").Where("id=?", userHalf.Id).First(genderEnabled).Error
	if err != nil {
		zap.L().Debug("查询性别", zap.Error(err))
		return nil, err
	}
	//初始化bo
	user.Role = role
	user.Jobs = job
	user.Dept = dept
	user.Id = userHalf.Id
	user.Phone = userHalf.Phone
	user.DeptId = userHalf.DeptId
	user.PwdResetTime = userHalf.PwdResetTime
	user.CreateBy = userHalf.CreateBy
	user.CreateTime = userHalf.CreateTime
	user.UpdateBy = userHalf.UpdateBy
	user.UpdateTime = userHalf.UpdateTime
	user.AvatarName = userHalf.AvatarName
	user.AvatarPath = userHalf.AvatarPath
	user.Email = userHalf.Email
	user.NickName = userHalf.NickName
	user.Phone = userHalf.Phone
	user.Username = userHalf.Username
	user.Enabled = utils.ByteIntoBool(genderEnabled.Enabled)
	user.Gender = utils.ByteIntoBool(genderEnabled.Gender)
	dataScopes := make([]string, 0)
	roleNames := make([]string, 0)
	for _, v := range role {
		roleNames = append(roleNames, v.Name)
		dataScopes = append(dataScopes, v.DataScope)
	}
	data = &bo.UserCenterInfoBo{
		DataScopes: dataScopes,
		User:       user,
		Roles:      roleNames,
	}
	return data, nil
}

func (u *SysUser) UpdatePassWord(p *dto.UpdateUserPassDto, optionId int) (err error) {
	//md5加密
	pwd := utils.EncodeMD5(p.NewPass)
	return global.Eloquent.Table("sys_user").Where("id=?", optionId).Updates(map[string]interface{}{
		"password":       pwd,
		"update_by":      optionId,
		"pwd_reset_time": utils.GetCurrentTimeUnix(),
	}).Error
}

func (u *SysUser) UpdateAvatar(path string, userId int) (err error) {
	return global.Eloquent.Table("sys_user").Where("id=?", userId).Updates(map[string]interface{}{
		"avatar_path": path,
	}).Error
}
