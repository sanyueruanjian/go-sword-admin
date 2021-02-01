package service

import (
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/pkg/jwt"
	"project/utils"
)

type User struct {
}

// Login 返回json web token
func (u *User) Login(p *dto.UserLoginDto) (token string, err error) {
	user := new(models.SysUser)
	user.Username = p.Username

	user.Password = p.Password
	if err = user.Login(); err != nil {
		return "", err
	}
	return jwt.GenToken(user.ID, user.Username)
}

func (u *User) InsertUser(p *dto.InsertUserDto, userID int) (err error) {
	//初始化 user数据
	user := &models.SysUser{
		DeptId:   p.DeptId,
		Email:    p.Email,
		NickName: p.NickName,
		Phone:    utils.Int64ToString(p.Phone),
		Username: p.UserName,
		Enabled:  utils.StrBoolIntoByte(p.Enabled),
		Gender:   utils.StrBoolIntoByte(p.Gender),
		CreateBy: userID,
		UpdateBy: userID,
		IsAdmin:  []byte{0},
	}
	jobs := p.Jobs
	roles := p.Roles

	if err := user.InsertUser(jobs, roles); err != nil {
		return err
	}
	return nil
}

func (u *User) SelectUserInfoList(p *dto.SelectUserInfoArrayDto) (data []*bo.UserInfoListBo, err error) {
	//var users []*models.SysUser
	//user := new(models.SysUser)
	//users, err = user.SelectUserInfoList(p)

	//if err != nil {
	//	return nil, err
	//}
	return data, nil
}
