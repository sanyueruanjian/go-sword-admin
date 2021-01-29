package models

type SysUsersJobs struct {
	ID     int `gorm:"primary_key" json:"id"` //id
	UserId int `json:"user_id"`               //用户ID
	JobId  int `json:"job_id"`                //岗位ID
}
