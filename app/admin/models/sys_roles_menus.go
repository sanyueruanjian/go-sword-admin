package models

type SysRolesMenus struct {
	ID     int `gorm:"primary_key" json:"id"` //id
	MenuId int `json:"menu_id"`               //菜单ID
	RoleId int `json:"role_id"`               //角色ID
}
