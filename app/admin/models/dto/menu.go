package dto

//SelectMenuDto 查询菜单
type SelectMenuDto struct {
	Current   int    `form:"current"`
	Size      int    `form:"size"`
	Pid       int    `form:"pid"` //父id
	StartTime int    `form:"startTime"`
	EndTime   int    `form:"endTime"`
	Blurry    string `form:"blurry"`
	Orders    string `form:"orders"`
}

//InsertMenuDto 新增菜单
type InsertMenuDto struct {
	Cache      bool        `json:"cache"`
	MenuSort   int         `json:"menuSort"`
	ID         int         `json:"id"`
	Pid        int         `json:"pid"`
	Icon       string      `json:"icon"`
	Component  string      `json:"component"`
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Permission string      `json:"permission"`
	Title      string      `json:"title"`
	Address    string      `json:"address"`
	Action     string      `json:"action"`
	Roles      []int       `json:"roles"`
	Hidden     interface{} `json:"hidden"`
	Iframe     interface{} `json:"iframe"`
	Type       interface{} `json:"type"`
}

type InsertFlexMenuDto struct {
	Hidden []byte `json:"hidden"`
	Iframe []byte `json:"iframe"`
	Type   int    `json:"type"`
}

//UpdateMenuDto 修改菜单
type UpdateMenuDto struct {
	Cache       bool   `json:"cache"`
	Hidden      bool   `json:"hidden"`
	Iframe      bool   `json:"iframe"`
	HasChildren bool   `json:"hasChildren"`
	Leaf        bool   `json:"leaf"`
	MenuSort    int    `json:"menuSort" binding:"required"`
	UpdatedBy   int    `json:"updatedBy"`
	SubCount    int    `json:"subCount"`
	ID          int    `json:"id" binding:"required"`
	Pid         int    `json:"pid"`
	Type        int    `json:"type"`
	CreateBy    int    `json:"creatBy"`
	UpdateTime  string `json:"updateTime" binding:"required"`
	CreateTime  string `json:"creatTime"`
	Icon        string `json:"icon" binding:"required"`
	Label       string `json:"label"`
	Children    string `json:"children"`
	Component   string `json:"component"`
	Name        string `json:"name"`
	Path        string `json:"path" binding:"required"`
	Permission  string `json:"permission"`
	Title       string `json:"title" binding:"required"`
	Roles       []int  `json:"roles"`
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
	Current int    `form:"current" binding:"required"`
	Size    int    `form:"size" binding:"required"`
	Pid     int    `form:"pid" binding:"required"` //父id
	Orders  string `form:"orders" binding:"required"`
}

//SelectAllMenuDto 返回全部菜单
type SelectAllMenuDto struct {
	Pid int `form:"pid"`
}

//SelectCurrentLevelAndUpLevelInfo
type SelectCurrentLevelAndUpLevelInfo struct {
	//int[] 接受即可 id列表
}

//DataMenuDto 同级上级菜单数据获取
type DataMenuDto struct {
	//int[] 接受即可 id列表`
}
