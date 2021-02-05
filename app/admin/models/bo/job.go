package bo

type GetJobList struct {
	Name    string `json:"name"`     //岗位名称
	Enabled uint8  `json:"enabled"`  //状态：1启用（默认）、0禁用
	JobSort int    `json:"job_sort"` //排序
}

type JobListDownload struct {
	Name       string `json:"name"`     //岗位名称
	Enabled    string `json:"enabled"`  //状态：1启用（默认）、0禁用
	CreateTime string `json:"job_sort"` //排序
}
