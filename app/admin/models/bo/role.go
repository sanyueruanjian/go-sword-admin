package bo

type RecordRole struct {
	CreateBy    int    `json:"createBy"`
	ID          int    `json:"id"`
	Level       int    `json:"level"`
	UpdateBy    int    `json:"updateBy"`
	CreateTime  int64  `json:"createTime"`
	DataScope   string `json:"dataScope"`
	Description string `json:"description"`
	Name        string `json:"name"`
	UpdateTime  int64  `json:"updateTime"`
	Protection  bool   `json:"protection"`
	Depts       []Dept `json:"depts"`
	Menus       []Menu `json:"menus"`
}

type Dept struct {
	CreateBy    int    `json:"createBy"`
	CreateTime  int64  `json:"createTime"`
	DeptSort    int    `json:"deptSort"`
	Enabled     bool   `json:"enabled"`
	HasChildren bool   `json:"hasChildren"`
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Leaf        bool   `json:"leaf"`
	Name        string `json:"name"`
	Pid         int    `json:"pid"`
	SubCount    int    `json:"subCount"`
	UpdateTime  int64  `json:"updateTime"`
	UpdateBy    int    `json:"updateBy"`
}

type Menu struct {
	CreateBy    int    `json:"createBy"`
	Icon        string `json:"icon"`
	ID          int    `json:"id"`
	MenuSort    int    `json:"menuSort"`
	Pid         int    `json:"pid"`
	SubCount    int    `json:"subCount"`
	Type        int    `json:"type"`
	UpdateBy    int    `json:"updateBy"`
	Children    string `json:"children"`
	Component   string `json:"component"`
	CreateTime  int64  `json:"createTime"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Permission  string `json:"permission"`
	Title       string `json:"title"`
	UpdateTime  int64  `json:"updateTime"`
	Cache       bool   `json:"cache"`
	HasChildren bool   `json:"hasChildren"`
	Hidden      bool   `json:"hidden"`
	Iframe      bool   `json:"iframe"`
	Leaf        bool   `json:"leaf"`
}

//SelectRoleArrayBo 多条件查询 角色列表
type SelectRoleArrayBo struct {
	Records []RecordRole `json:"records"`
	Paging  paging
}

//SelectRoleBo 查询单个角色
type SelectRoleBo struct {
	RecordRole
}

//SelectAllRoleBo 查询所有角色
type SelectAllRoleBo struct {
}

//导出角色数据
type DownloadRoleInfoBo struct {
	//	输出文件
}

type SelectCurrentUserLevel struct {
	Level int `json:"level"`
}
