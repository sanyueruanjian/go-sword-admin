package service

import (
	"encoding/json"
	"io"
	"project/app/admin/models/cache"
	"strconv"

	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/common/global"
	"project/pkg/jwt"
	"project/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type User struct {
}

// Login 返回json web token
func (u *User) Login(c *gin.Context, p *dto.UserLoginDto) (loginData *bo.LoginData, err error) {
	user := new(models.SysUser)
	user.Username = p.Username

	user.Password = p.Password
	r, err := user.Login()
	if err != nil {
		return nil, err
	}

	//var dept *bo.DeptCommon
	if r.Jobs, err = models.SelectUserJob(r.Id); err != nil {
		return nil, err
	}
	if r.Role, err = models.SelectUserRole(r.Id); err != nil {
		return nil, err
	}
	if r.Dept, err = models.SelectUserDept(r.Id); err != nil {
		return nil, err
	}
	loginUser := new(bo.LoginUser)
	loginUser.User = r
	// 获取菜单权限
	if utils.ByteIntoInt(user.IsAdmin) == 1 {
		loginUser.Roles = append(loginUser.Roles, `admin`)
	} else {
		if loginUser.Roles, err = models.SelectUserMenuPermission(r.Role); err != nil {
			return nil, err
		}
	}

	// 获取部门权限
	var dataScopesRoleIds []int
	var allScopes bool
	for _, role := range r.Role {
		switch role.DataScope {
		case `全部`:
			allScopes = true
			loginUser.DataScopes = []int{}
			break
		case `本级`:
			loginUser.DataScopes = append(loginUser.DataScopes, user.DeptId)
		default:
			dataScopesRoleIds = append(dataScopesRoleIds, role.ID)
		}
	}

	if !allScopes {
		deptIds, err := models.SelectUserDeptIdByRoleId(dataScopesRoleIds)
		if err != nil {
			return nil, err
		}
		loginUser.DataScopes = append(loginUser.DataScopes, deptIds...)
	}

	token, err := jwt.GenToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}
	loginData = new(bo.LoginData)
	loginData.Token = "Bearer " + token
	loginData.User = loginUser

	err = u.RedisUserMessage(c, r)
	return
}

func (u *User) RedisUserMessage(c *gin.Context, r *bo.RecordUser) (err error) {
	//构造角色名字集合
	roleNames := make([]string, 0)
	for _, v := range r.Role {
		roleNames = append(roleNames, v.Name)
	}
	//初始化缓存模型
	var userInfo []byte
	userInfo, err = json.Marshal(models.RedisUserInfo{
		UserId:   r.Id,
		UserName: r.Username,
		DeptId:   r.DeptId,
		Role:     roleNames,
	})
	//json.Marshal(models.OnlineUser{
	//	LoginTime: utils.NowUnix(),
	//	Browser:   c.,
	//	Dept:      "",
	//	Ip:        "",
	//	Nickname:  "",
	//	Username:  "",
	//})
	if err != nil {
		zap.L().Error("RedisUserInfo Marshal failed", zap.Error(err))
		return
	}
	//添加缓存
	if err := global.Rdb.Set(strconv.Itoa(r.Id), userInfo, 0).Err(); err != nil {
		zap.L().Error("用户缓存错误", zap.Error(err))
		return err
	}
	return
}

func (u *User) InsertUser(p *dto.InsertUserDto, userID int) (err error) {
	//设置默认密码123456
	defaultPass := "123456"
	pass := utils.EncodeMD5(defaultPass)
	//初始化 user数据
	user := &models.SysUser{
		DeptId:       p.DeptId,
		Email:        p.Email,
		NickName:     p.NickName,
		Phone:        utils.Int64ToString(p.Phone),
		Username:     p.UserName,
		Enabled:      utils.StrBoolIntoByte(p.Enabled),
		Gender:       utils.StrBoolIntoByte(p.Gender),
		CreateBy:     userID,
		UpdateBy:     userID,
		IsAdmin:      []byte{0},
		Password:     pass,
		PwdResetTime: utils.GetCurrentTimeUnix(),
	}
	jobs := p.Jobs
	roles := p.Roles

	if err := user.InsertUser(jobs, roles); err != nil {
		return err
	}
	return nil
}

func (u *User) SelectUserInfoList(p *dto.SelectUserInfoArrayDto) (data *bo.UserInfoListBo, err error) {
	user := new(models.SysUser)
	data, err = user.SelectUserInfoList(p)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *User) DeleteUser(ids []int) error {
	user := new(models.SysUser)
	return user.DeleteUser(ids)
}

func (u *User) UpdateUser(p *dto.UpdateUserDto, optionId int) error {
	user := new(models.SysUser)
	return user.UpdateUser(p, optionId)
}

func (u *User) UpdateUserCenter(p *dto.UpdateUserCenterDto, optionId int) (err error) {
	user := new(models.SysUser)
	return user.UpdateUserCenter(p, optionId)
}

func (u *User) SelectUserInfo(p *models.RedisUserInfo) (data *bo.UserCenterInfoBo, err error) {
	//读取缓存
	if data, err = cache.GetUserCenterCache(p.UserId); err != nil && data != nil {
		return data, nil
	}
	user := new(models.SysUser)
	data, err = user.SelectUserInfo(p)
	if err != nil {
		return nil, err
	}
	//	redis缓存
	err = cache.SetUserCenterListCache(data)
	if err != nil {
		zap.L().Error("SetUserCenterListCache failed", zap.Error(err))
	}
	return data, nil
}

func (u *User) UpdatePassWord(p *dto.UpdateUserPassDto, optionId int) (err error) {
	user := new(models.SysUser)
	return user.UpdatePassWord(p, optionId)
}

func (u *User) UpdateAvatar(path string, userId int) (err error) {
	user := new(models.SysUser)
	return user.UpdateAvatar(path, userId)
}

func (u *User) UserDownload(p *dto.DownloadUserInfoDto) (content io.ReadSeeker, err error) {
	user := new(models.SysUser)
	var downloadUsers []interface{}
	userList, err := user.UserDownload(p)
	if err != nil {
		return nil, err
	}
	//数据整合
	recordUsers := userList.Records
	for _, v := range recordUsers {
		jobs := ""
		roles := ""
		for k, job := range v.Jobs {
			if k != len(v.Jobs)-1 {
				jobs += job.Name + ","
			} else {
				jobs += job.Name
			}
		}
		for k, role := range v.Role {
			if k != len(v.Role)-1 {
				roles += role.Name + ","
			} else {
				roles += role.Name
			}
		}
		tmp := &bo.DownloadUserBo{
			Username:     v.Username,
			Dept:         v.Dept.Name,
			Jobs:         jobs,
			Role:         roles,
			Email:        v.Email,
			Phone:        v.Phone,
			PwdResetTime: utils.UnixTimeToString(v.PwdResetTime),
			CreateTime:   utils.UnixTimeToString(v.CreateTime),
		}
		if v.Enabled {
			tmp.Enabled = "启用"
		} else {
			tmp.Enabled = "未启用"
		}
		downloadUsers = append(downloadUsers, tmp)
	}
	content = utils.ToExcel([]string{"用户名", "角色", "部门", "岗位", "邮箱", "状态", "手机号码", "修改密码的时间", "创建日期"}, downloadUsers)
	return
}
