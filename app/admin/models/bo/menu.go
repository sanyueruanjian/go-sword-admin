package bo

type SelectMenuBo struct {
	CreateBy    int    `json:"createBy"`
	UpdatedBy   int    `json:"updatedBy"`
	SubCount    int    `json:"subCount"`
	MenuSort    int    `json:"menuSort"`
	ID          int    `json:"id"`
	Pid         int    `json:"pid"`
	Type        int    `json:"type"`
	Cache       bool   `json:"cache"`
	Hidden      bool   `json:"hidden"`
	Leaf        bool   `json:"leaf"`
	Iframe      bool   `json:"iframe"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	Label       string `json:"label"`
	Children    string `json:"children"`
	Icon        string `json:"icon"`
	Component   string `json:"component"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Permission  string `json:"permission"`
	Title       string `json:"title"`
	HasChildren string `json:"hasChildren"`
}

type meta struct {
	Icon    string `json:"icon"`
	NoCache bool   `json:"noCache"`
	Title   string `json:"title"`
}

type children struct {
	Component string  `json:"component"`
	Hidden    string  `json:"hidden"`
	Meta      []*meta `json:"meta"`
}

type SelectForeNeedMenuBo struct {
	AlwaysShow bool        `json:"alwaysShow"`
	Hidden     bool        `json:"hidden"`
	Component  string      `json:"component"`
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Redirect   string      `json:"redirect"`
	Meta       *meta       `json:"meta"`
	Children   []*children `json:"children"`
}

type SelectSuperMenuBo struct {
	CreateBy    int                  `json:"createBy"`
	UpdatedBy   int                  `json:"updatedBy"`
	SubCount    int                  `json:"subCount"`
	MenuSort    int                  `json:"menuSort"`
	ID          int                  `json:"id"`
	Pid         int                  `json:"pid"`
	Type        int                  `json:"type"`
	Cache       bool                 `json:"cache"`
	Hidden      bool                 `json:"hidden"`
	Leaf        bool                 `json:"leaf"`
	Iframe      bool                 `json:"iframe"`
	CreateTime  int64                `json:"createTime"`
	UpdateTime  int64                `json:"updateTime"`
	Label       string               `json:"label"`
	Children    []*SelectSuperMenuBo `json:"children"`
	Icon        string               `json:"icon"`
	Component   string               `json:"component"`
	Name        string               `json:"name"`
	Path        string               `json:"path"`
	Permission  string               `json:"permission"`
	Title       string               `json:"title"`
	HasChildren bool                 `json:"hasChildren"`
}
