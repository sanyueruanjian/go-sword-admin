package bo

type recordRole struct {
	CreateBy    int     `json:"createBy"`
	ID          int     `json:"id"`
	Level       int     `json:"level"`
	UpdateBy    int     `json:"updateBy"`
	CreateTime  string  `json:"createTime"`
	DataScope   string  `json:"dataScope"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	UpdateTime  string  `json:"updateTime"`
	Protection  bool    `json:"protection"`
	Depts       []int   `json:"depts"`
	Menus       []*menu `json:"menus"`
}

type menu struct {
	CreateBy    int    `json:"createBy"`
	Icon        int    `json:"icon"`
	ID          int    `json:"id"`
	MenuSort    int    `json:"menuSort"`
	Pid         int    `json:"pid"`
	SubCount    int    `json:"subCount"`
	Type        int    `json:"type"`
	UpdateBy    int    `json:"updateBy"`
	Children    string `json:"children"`
	Component   string `json:"component"`
	CreateTime  string `json:"createTime"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Permission  string `json:"permission"`
	Title       string `json:"title"`
	UpdateTime  string `json:"updateTime"`
	Cache       bool   `json:"cache"`
	HasChildren bool   `json:"hasChildren"`
	Hidden      bool   `json:"hidden"`
	Iframe      bool   `json:"iframe"`
	Leaf        bool   `json:"leaf"`
}

//SelectRoleArrayBo 多条件查询 角色列表
type SelectRoleArrayBo struct {
	Records []*recordRole `json:"records"`
	Orders  []*order      `json:"orders"`
	*paging
}

//SelectRoleBo 查询单个角色
type SelectRoleBo struct {
	Menus []*menu `json:"menus"`
	*recordRole
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
