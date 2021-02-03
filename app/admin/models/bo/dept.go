package bo

type RecordDept struct {
	ID          int    `json:"id"`
	DeptSort    int    `json:"deptSort"`
	Pid         int    `json:"pid"`
	CreateBy    int    `json:"createBy"`
	UpdateBy    int    `json:"updateBy"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
	HasChildren bool   `json:"hasChildren"`
	Enabled     bool   `json:"enabled"`
	Leaf        bool   `json:"leaf"`
	Label       string `json:"label"`
	Name        string `json:"name"`
}

type SelectDeptListBo struct {
	Records []*RecordDept `json:"records"`
	Paging  paging
}

type DownloadDeptList struct {
	Name       string `json:"name"`    //部门名称
	Enabled    string `json:"enabled"` //状态：1启用（默认）、0禁用
	CreateTime string `json:"更新时间"`    //排序
}
