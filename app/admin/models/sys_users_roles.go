package models

type SysUsersRoles struct {
	ID     int `gorm:"primary_key" json:"id"` //id
	UserId int `json:"user_id"`               //用户ID
	RoleId int `json:"role_id"`               //角色ID
}

//表名
func (SysUsersRoles) TableName() string {
	return "sys_users_roles"
}
