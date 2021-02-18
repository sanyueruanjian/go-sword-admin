package dto

//GetJobList 岗位列表数据入参
type GetJobList struct {
	EndTime   int    `form:"endTime"`   //结束时间
	StartTime int    `form:"startTime"` //创建时间
	Size      int    `form:"size"`      //每页数据
	Current   int    `form:"current"`   //当前页
	Name      string `form:"name"`      //模糊查询
	Orders    string `form:"orders"`    //排序规则
	Enabled   bool   `form:"enabled"`   //状态：1启用（默认）、0禁用
	Page      int    `form:"page"`      //页数
}

type AddJob struct {
	ID      int    `json:"id"`                         //ID
	JobSort int    `json:"jobSort" binding:"required"` //排序
	Name    string `json:"name" binding:"required"`    //岗位名称
	Enabled bool   `json:"enabled" binding:"required"` //岗位状态
}

type UpdateJob struct {
	ID      int    `json:"id"`                         //ID
	JobSort int    `json:"jobSort" binding:"required"` //排序
	Enabled bool  `json:"enabled" binding:"required"` //岗位状态
	Name    string `json:"name" binding:"required"`    //岗位名称
}
