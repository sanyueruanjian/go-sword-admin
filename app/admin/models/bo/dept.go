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
