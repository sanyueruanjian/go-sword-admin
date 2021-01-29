package models

type SysRolesDepts struct {
	ID     int `gorm:"primary_key" json:"id"` //id
	RoleId int `json:"role_id"`               //角色id
	DeptId int `json:"dept_id"`               //部门id
}
