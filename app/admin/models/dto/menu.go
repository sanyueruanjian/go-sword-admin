package dto

//SelectMenuDto 查询菜单
type SelectMenuDto struct {
	Current int `form:"current" binding:"required"`
	Size    int `form:"size" binding:"required"`
	Pid     int `form:"pid" binding:"required"` //父id
	//TODO
	Orders string `form:"orders" binding:"required"`
}

//InsertMenuDto 新增菜单
type InsertMenuDto struct {
	Cache      bool   `json:"cache" binding:"required"`
	Hidden     bool   `json:"hidden" binding:"required"`
	Iframe     bool   `json:"iframe" binding:"required"`
	MenuSort   bool   `json:"menu_sort" binding:"required"`
	ID         int    `json:"id" binding:"required"`
	Icon       int    `json:"icon" binding:"required"`
	Pid        int    `json:"pid" binding:"required"`
	Type       int    `json:"type" binding:"required"`
	Component  string `json:"component" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Path       string `json:"path" binding:"required"`
	Permission string `json:"permission" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Roles      []int  `json:"roles" binding:"required"`
}

//UpdateMenuDto 修改菜单
type UpdateMenuDto struct {
	Cache       bool   `json:"cache" binding:"required"`
	Hidden      bool   `json:"hidden" binding:"required"`
	Iframe      bool   `json:"iframe" binding:"required"`
	MenuSort    bool   `json:"menu_sort" binding:"required"`
	HasChildren bool   `json:"hasChildren" binding:"required"`
	Leaf        bool   `json:"leaf" binding:"required"`
	UpdatedBy   int    `json:"updatedBy" binding:"required"`
	SubCount    int    `json:"sub_count" binding:"required"`
	ID          int    `json:"id" binding:"required"`
	Icon        int    `json:"icon" binding:"required"`
	Pid         int    `json:"pid" binding:"required"`
	Type        int    `json:"type" binding:"required"`
	CreateBy    int    `json:"creatBy" binding:"required"`
	UpdateTime  int    `json:"updateTime" binding:"required"`
	Label       string `json:"label" binding:"required"`
	CreateTime  string `json:"creatTime" binding:"required"`
	Children    string `json:"children" binding:"required"`
	Component   string `json:"component" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Permission  string `json:"permission" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Roles       []int  `json:"roles" binding:"required"`
}

//删除菜单
type DeleteMenuDto struct {
	//	[]int解析即可
}

//SelectForeNeedMenuDto 查找前端所需菜单
type SelectForeNeedMenuDto struct {
	//	无参数
}

// SelectChildID 返回所有子节点ID
type SelectChildIdDto struct {
	ID int `form:"id"`
}

//DownloadMenuDto 导出菜单数据
type DownloadMenuDto struct {
	Current int `form:"current" binding:"required"`
	Size    int `form:"size" binding:"required"`
	Pid     int `form:"pid" binding:"required"` //父id
	//TODO
	Orders string `form:"orders" binding:"required"`
}

//SelectAllMenuDto 返回全部菜单
type SelectAllMenuDto struct {
	Pid int `form:"pid"`
}

//SelectCurrentLevelAndUpLevelInfo
type SelectCurrentLevelAndUpLevelInfo struct {
	//int[] 接受即可 id列表
}
