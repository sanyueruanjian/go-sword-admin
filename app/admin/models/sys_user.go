package models

import (
	"errors"
	"fmt"
	"project/app/admin/models/bo"
	"project/app/admin/models/cache"
	"project/app/admin/models/dto"
	"project/common/global"
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
	IsAdmin []byte `json:"is_admin"`
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

// OnlineUser 用户线上数据
type OnlineUser struct {
	LoginTime     int64  `json:"loginTime"`     //登录时间
	LoginLocation string `json:"loginLocation"` // 归属地
	Browser       string `json:"browser"`       // 浏览器
	Dept          string `json:"dept"`          //部门
	Ip            string `json:"ip"`            //ip地址
	Nickname      string `json:"nickname"`      //昵称
	Username      string `json:"username"`      //用户名
	Token         string `json:"key"`           // token
}

type ModelUserMessage struct {
	UserId         int
	Username       string
	DataScopes     *[]int
	MenuPermission *[]string
	Roles          *[]SysRole
	Dept           *SysDept
	Jobs           *[]SysJob
}

type Admin struct {
	IsAdmin []byte `json:"is_admin"` //是否为admin账号
}

func (SysUser) TableName() string {
	return "sys_user"
}

var (
	ErrorUserNotExist     = errors.New("用户不存在")
	ErrorInvalidPassword  = errors.New("用户名或密码错误")
	ErrorServerBusy       = errors.New("服务器繁忙")
	ErrorUserIsNotEnabled = errors.New("用户未激活")
	ErrorUserIsExist      = errors.New("用户已存在")
)

func (a *Admin) GetIsAdmin(userId int) error {
	return global.Eloquent.Table("sys_user").Where("id = ?", userId).First(a).Error
}

func (u *SysUser) GetUser(userId int) error {
	return global.Eloquent.Table(u.TableName()).Where("id = ?", userId).First(u).Error
}

// Login 查询用户是否存在，并验证密码
func (u *SysUser) Login() error {
	oPassword := u.Password
	err := global.Eloquent.Table(u.TableName()).Where("username = ?", u.Username).First(u).Error
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
	if u.Enabled[0] == 0 {
		return ErrorUserIsNotEnabled
	}
	return nil
}

func (u *SysUser) InsertUser(jobs []int, roles []int) (err error) {
	var num int64
	global.Eloquent.Table("sys_user").Where("username=?", u.Username).Count(&num)
	if num != 0 {
		zap.L().Info(fmt.Sprintf("%d", num))
		return ErrorUserIsExist
	}
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
	//删除用户列表缓存
	if err := cache.DelAllUserRecordsCache(); err != nil {
		zap.L().Error("DelAllUserCenterCache", zap.Error(err))
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error
}

func (u *SysUser) SelectUserInfoList(p *dto.SelectUserInfoArrayDto, currentUser *ModelUserMessage) (data *bo.UserInfoListBo, err error) {
	//排序条件
	var orderJson []bo.Order
	orderJson, err = utils.OrderJson(p.Orders)
	orderRule := utils.GetOrderRule(orderJson)

	//查询缓存
	if p.StartTime == 0 && p.EndTime == 0 && p.Blurry == "" {
		var userRecords []*bo.RecordUser
		userRecords, err = cache.GetUserRecordsCache(currentUser.UserId)
		if err != nil {
			zap.L().Error("GetUserRecordsCache failed", zap.Error(err))
		}
		//筛选部门
		records := make([]*bo.RecordUser, 0, 0)
		if p.DepID == 0 {
			records = append(records, userRecords...)
		} else {
			for _, record := range userRecords {
				if record.DeptId == p.DepID {
					records = append(records, record)
				}
			}
		}
		//分页
		var total int
		total = len(records)
		start := p.Size * (p.Current - 1)
		end := start + p.Size - 1
		resRecords := make([]*bo.RecordUser, 0, 0)
		if end >= total-1 {
			resRecords = records[start:]
		} else if len(records) != 0 {
			resRecords = records[start:end]
		}
		//查询页数
		data = &bo.UserInfoListBo{Records: resRecords}
		data.Orders = orderJson
		data.Size = p.Size
		data.Current = p.Current
		data.Pages = (total + p.Size - 1) / p.Size
		data.Total = total
		data.SearchCount = true
		data.OptimizeCountSql = true
		if resRecords != nil && total > 0 {
			return data, nil
		}
	}
	//查询用户基本信息
	var usersHalf []*bo.RecordUserHalf

	//模糊查询
	blurry := "%" + p.Blurry + "%"
	table := global.Eloquent.Table("sys_user").Where("is_deleted=? AND enabled=? AND (username like ? or nick_name like ? or email like ?)", []byte{0}, 1, blurry, blurry, blurry)

	//部门筛选
	if len(*currentUser.DataScopes) != 0 {
		table = table.Where("dept_id in (?)", *currentUser.DataScopes)
	}

	zap.L().Info(fmt.Sprintf("%v", *currentUser.DataScopes))

	//日期筛选
	if p.EndTime != 0 && p.StartTime != 0 {
		table = table.Where("create_time > ? AND create_time < ?", p.StartTime, p.EndTime)
	}

	//排序
	var total int64
	err = table.Count(&total).Order(orderRule).Find(&usersHalf).Error
	pages := (int(total) + p.Size - 1) / p.Size
	if err != nil {
		return nil, err
	}

	var users []*bo.RecordUser
	for _, userHalf := range usersHalf {
		//查询角色
		var roles []*bo.Role
		roles, err = SelectUserRole(userHalf.Id)
		if err != nil {
			zap.L().Debug("查询角色", zap.Error(err))
			return nil, err
		}

		//查询岗位
		var jobs []*bo.Job
		jobs, err = SelectUserJob(userHalf.Id)
		if err != nil {
			zap.L().Debug("查询岗位", zap.Error(err))
			return nil, err
		}

		//查询部门
		dept := new(bo.DeptCommon)
		err = global.Eloquent.Table("sys_dept").Joins("left join sys_user "+
			"on sys_user.dept_id = sys_dept.id").Where("sys_user.id=? AND sys_dept.is_deleted=?", userHalf.Id, []byte{0}).Scan(dept).Error
		if err != nil {
			zap.L().Debug("查询部门", zap.Error(err))
			return nil, err
		}

		//查询性别
		genderEnabled := new(GenderEnabled)
		err = global.Eloquent.Table("sys_user").Select("gender", "enabled", "is_admin").Where("id=?", userHalf.Id).First(genderEnabled).Error
		if err != nil {
			zap.L().Debug("查询性别", zap.Error(err))
			return nil, err
		}
		user := new(bo.RecordUser)
		user.RecordUserHalf = new(bo.RecordUserHalf)
		user.RoleDeptJobBool = new(bo.RoleDeptJobBool)
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

	//设置records缓存
	if p.StartTime == 0 && p.EndTime == 0 && p.Blurry == "" {
		err = cache.SetUserRecordsCache(users, currentUser.UserId)
		if err != nil {
			zap.L().Error("SetUserRecordsCache failed", zap.Error(err))
		}
	}
	start := p.Size * (p.Current - 1)
	end := start + p.Size - 1
	if end > int(total) {
		end = int(total) - 1
	}
	records := make([]*bo.RecordUser, 0, 0)
	if len(users) != 0 {
		records = users[start:end]
	}

	data = &bo.UserInfoListBo{Records: records}
	data.Orders = orderJson
	data.Size = p.Size
	data.Current = p.Current
	data.Pages = pages
	data.Total = int(total)
	data.SearchCount = true
	data.OptimizeCountSql = true
	return data, nil
}

func SelectUserRole(userId int) (role []*bo.Role, err error) {
	//连表查询角色
	err = global.Eloquent.Table("sys_role").
		Joins("left join sys_users_roles on sys_users_roles.role_id = sys_role.id").
		Joins("left join sys_user on sys_user.id = sys_users_roles.user_id").
		Where("sys_role.is_deleted=? and sys_user.id=?", []byte{0}, userId).
		Find(&role).Error
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

func GetUserJob(jobs *[]SysJob, userId int) (err error) {
	//连表查询岗位
	err = global.Eloquent.Table("sys_job").
		Joins("left join sys_users_jobs on sys_users_jobs.job_id = sys_job.id").
		Joins("left join sys_user on sys_user.id = sys_users_jobs.user_id").
		Where("sys_job.is_deleted=? and sys_user.id=?", []byte{0}, userId).
		Find(jobs).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户无岗位", zap.Error(err))
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return ErrorServerBusy
	}
	return
}

func GetUserRole(role *[]SysRole, userId int) (err error) {
	//连表查询角色
	err = global.Eloquent.Table("sys_role").
		Select("sys_role.id, sys_role.level, sys_role.create_by, sys_role.update_by, sys_role.create_time, sys_role.update_time, sys_role.is_protection, sys_role.is_deleted, sys_role.name, sys_role.description, sys_role.data_scope").
		Joins("left join sys_users_roles on sys_users_roles.role_id = sys_role.id").
		Joins("left join sys_user on sys_user.id = sys_users_roles.user_id").
		Where("sys_role.is_deleted=? and sys_user.id=?", []byte{0}, userId).
		Find(role).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户无角色", zap.Error(err))
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return ErrorServerBusy
	}
	return
}

func SelectUserJob(userId int) (jobs []*bo.Job, err error) {
	//连表查询岗位
	err = global.Eloquent.Table("sys_job").
		Select("sys_job.id, sys_job.name").
		Joins("left join sys_users_jobs on sys_users_jobs.job_id = sys_job.id").
		Joins("left join sys_user on sys_user.id = sys_users_jobs.user_id").
		Where("sys_job.is_deleted=? and sys_user.id=?", []byte{0}, userId).
		Find(&jobs).Error
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

// SelectUserDept 查询部门
func SelectUserDept(dept *SysDept, userId int) (err error) {
	fmt.Println(userId)
	fmt.Printf("%#v\n", dept)
	err = global.Eloquent.Table("sys_dept").
		Select("sys_dept.name, sys_dept.pid, sys_dept.sub_count, sys_dept.dept_sort, sys_dept.create_by, sys_dept.update_by, sys_dept.enabled, sys_dept.id, sys_dept.is_deleted, sys_dept.create_time, sys_dept.update_time").
		Joins("left join sys_user on sys_user.dept_id = sys_dept.id").
		Where("sys_user.id=? AND sys_dept.is_deleted=?", userId, []byte{0}).
		Find(dept).Error
	fmt.Printf("%#v\n", dept)
	zap.L().Info(fmt.Sprintf("%s%d  %d", "dept:id", dept.ID, userId))
	return
}

// SelectUserDeptIdByRoleId 根据角色id查询部门
func SelectUserDeptIdByRoleId(roleId []int) (deptIds []int, err error) {
	err = global.Eloquent.Table("sys_roles_depts").Where("role_id in (?)", roleId).
		Select([]string{"dept_id"}).Scan(&deptIds).Error
	return
}

// SelectUserMenuPermission 查询菜单权限
func SelectUserMenuPermission(menus *[]SysMenu, roles *[]SysRole) (err error) {
	var rolesId []int
	for _, role := range *roles {
		rolesId = append(rolesId, role.ID)
	}
	err = global.Eloquent.Table("sys_menu").
		Joins("left join sys_roles_menus on sys_roles_menus.menu_id = sys_menu.id").
		Where("sys_roles_menus.role_id in (?)", rolesId).Find(menus).Error
	return
}

func (u *SysUser) DeleteUser(ids []int) (err error) {
	tx := global.Eloquent.Begin()
	err = tx.Table("sys_user").Where("id IN (?)", ids).Updates(map[string]interface{}{"is_deleted": []byte{1}}).Error
	if err != nil {
		return err
	}
	//	删除userCenter缓存
	if err := cache.DelManyUserCenterCache(ids); err != nil {
		tx.Rollback()
		return err
	}
	//删除用户相关缓存
	if err := cache.DelUsersAboutCache(ids); err != nil {
		tx.Rollback()
		return err
	}
	//删除用户列表缓存
	if err := cache.DelAllUserRecordsCache(); err != nil {
		zap.L().Error("DelAllUserCenterCache", zap.Error(err))
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *SysUser) UpdateUser(p *dto.UpdateUserDto, optionId int) (err error) {
	//开始事务
	tx := global.Eloquent.Begin()
	//校验用户是否存在
	test := new(SysUser)
	err = tx.Table("sys_user").Where("id=? AND is_deleted=?", p.ID, []byte{0}).First(test).Error
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
		"enabled":     p.Enabled,
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
	//	删除用户缓存
	if err := cache.DelUserCenterCache(p.ID); err != nil {
		tx.Rollback()
		return err
	}
	//删除用户列表缓存
	if err := cache.DelAllUserRecordsCache(); err != nil {
		zap.L().Error("DelAllUserCenterCache", zap.Error(err))
		tx.Rollback()
		return err
	}
	//删除用户相关缓存
	if err := cache.DelUserAboutCache(p.ID); err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error
}

func (u *SysUser) UpdateUserCenter(p *dto.UpdateUserCenterDto, optionId int) (err error) {
	//创建事务
	tx := global.Eloquent.Begin()
	err = tx.Table("sys_user").Where("id=?", p.Id).Updates(map[string]interface{}{
		"gender":    utils.BoolIntoByte(p.Gender),
		"phone":     p.Phone,
		"nick_name": p.NickName,
		"update_by": optionId,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//	删除个人中心缓存
	if err := cache.DelUserCenterCache(p.Id); err != nil {
		tx.Rollback()
		return err
	}
	//  删除用户列表缓存
	if err := cache.DelAllUserRecordsCache(); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *SysUser) SelectUserInfo(p *ModelUserMessage) (data *bo.UserCenterInfoBo, err error) {
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
	err = global.Eloquent.Table("sys_dept").Where("id=?", userHalf.DeptId).Find(dept).Error
	if err != nil {
		zap.L().Debug("查询部门", zap.Error(err))
		return nil, err
	}
	//查询性别
	err = global.Eloquent.Table("sys_user").Select("gender", "enabled", "is_admin").Where("id=?", userHalf.Id).First(genderEnabled).Error
	if err != nil {
		zap.L().Debug("查询性别", zap.Error(err))
		return nil, err
	}
	//查询操作权限
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
	data = &bo.UserCenterInfoBo{
		DataScopes: *p.DataScopes,
		User:       user,
		Roles:      *p.MenuPermission,
	}
	return data, nil
}

func (u *SysUser) UpdatePassWord(p *dto.UpdateUserPassDto) (err error) {
	//md5加密
	newPwd := utils.EncodeMD5(p.NewPass)
	oldPwd := utils.EncodeMD5(p.OldPass)
	//查询当前用户密码
	err = global.Eloquent.Table(u.TableName()).Where("username = ?", u.Username).First(u).Error
	if err != nil {
		return
	}
	if oldPwd != u.Password {
		zap.L().Error("user account or password is error")
		return ErrorInvalidPassword
	}
	//验证旧密码
	tx := global.Eloquent.Begin()
	err = tx.Table("sys_user").Where("id=?", u.ID).Updates(map[string]interface{}{
		"password":       newPwd,
		"update_by":      u.ID,
		"pwd_reset_time": utils.GetCurrentTimeUnix(),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//	删除缓存
	if err := cache.DelUserCenterCache(u.ID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *SysUser) UpdateAvatar(path string, userId int) (err error) {
	err = global.Eloquent.Table("sys_user").Where("id=?", userId).Updates(map[string]interface{}{
		"avatar_path": path,
	}).Error
	if err != nil {
		return err
	}
	if err := cache.DelUserCenterCache(userId); err != nil {
		return err
	}
	return nil
}

// UserDownload 导出用户数据
func (u *SysUser) UserDownload(p *dto.DownloadUserInfoDto) (data *bo.UserInfoListBo, err error) {
	//排序条件
	var orderJson []bo.Order
	orderJson, err = utils.OrderJson(p.Orders)
	orderRule := utils.GetOrderRule(orderJson)
	//查询用户基本信息
	var usersHalf []*bo.RecordUserHalf
	//分页
	var total int64
	err = global.Eloquent.Table("sys_user").Limit(p.Size).Offset((p.Current - 1) * p.Size).Count(&total).Order(orderRule).Find(&usersHalf).Error
	pages := (int(total) + p.Size - 1) / p.Size
	if err != nil {
		return nil, err
	}
	var users []*bo.RecordUser
	for _, userHalf := range usersHalf {
		//查询角色
		var roles []*bo.Role
		roles, err = SelectUserRole(userHalf.Id)
		if err != nil {
			zap.L().Debug("查询角色", zap.Error(err))
			return nil, err
		}
		//查询岗位
		var jobs []*bo.Job
		jobs, err = SelectUserJob(userHalf.Id)
		if err != nil {
			zap.L().Debug("查询岗位", zap.Error(err))
			return nil, err
		}
		//查询部门
		dept := new(bo.DeptCommon)
		err = global.Eloquent.Table("sys_dept").Joins("left join sys_user "+
			"on sys_user.dept_id = sys_dept.id").Where("sys_user.id=? AND sys_dept.is_deleted=?", userHalf.Id, []byte{0}).Scan(dept).Error
		if err != nil {
			zap.L().Debug("查询部门", zap.Error(err))
			return nil, err
		}
		//查询性别
		genderEnabled := new(GenderEnabled)
		err = global.Eloquent.Table("sys_user").Select("gender", "enabled", "is_admin").Where("id=?", userHalf.Id).First(genderEnabled).Error
		if err != nil {
			zap.L().Debug("查询性别", zap.Error(err))
			return nil, err
		}
		user := new(bo.RecordUser)
		user.RecordUserHalf = new(bo.RecordUserHalf)
		user.RoleDeptJobBool = new(bo.RoleDeptJobBool)
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
	data.Orders = orderJson
	data.Size = p.Size
	data.Current = p.Current
	data.Pages = pages
	data.Total = int(total)
	data.SearchCount = true
	data.OptimizeCountSql = true
	return data, nil
}
