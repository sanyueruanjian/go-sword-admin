package bo

type GetJob struct {
	Current int         `json:"current"`
	Size    int         `json:"size"`
	Pages   int         `json:"pages"`
	Total   int         `json:"total"`
	Orders  []Order       `json:"orders"`
	Records []*GetJobList `json:"records"`
}

type GetJobList struct {
	Id         int    `json:"id"`         //Id
	JobSort    int    `json:"jobSort"`    //排序
	CreateBy   int    `json:"createBy"`   //创建人
	UpdateBy   int    `json:"updateBy"`   //更新人
	CreateTime int64  `json:"createTime"` //创建时间
	UpdateTime int64  `json:"updateTime"` //更新时间
	Enabled    bool  `json:"enabled"`    //状态：1启用（默认）、0禁用
	Name       string `json:"name"`       //岗位名称
}

type JobListDownload struct {
	Name       string `json:"name"`     //岗位名称
	Enabled    string `json:"enabled"`  //状态：1启用（默认）、0禁用
	CreateTime string `json:"job_sort"` //排序
}
