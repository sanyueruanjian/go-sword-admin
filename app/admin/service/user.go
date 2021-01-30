package service

import (
	"project/app/admin/models"
	"project/app/admin/models/dto"
	"project/pkg/jwt"
)

type User struct {
}

// Login 返回json web token
func (e User) Login(p *dto.UserLoginDto) (token string, err error) {
	user := new(models.SysUser)
	user.Username = p.Username
	user.Password = p.Password
	if err = user.Login(); err != nil {
		return "", err
	}
	return jwt.GenToken(user.ID, user.Username)
}
