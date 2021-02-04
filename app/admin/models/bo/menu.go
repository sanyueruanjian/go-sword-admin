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

type Meta struct {
	Icon    string `json:"icon"`
	NoCache bool   `json:"noCache"`
	Title   string `json:"title"`
}

type Children struct {
	Component string `json:"component"`
	Hidden    bool   `json:"hidden"`
	Meta      *Meta  `json:"meta"`
}

type SelectForeNeedMenuBo struct {
	AlwaysShow bool        `json:"alwaysShow"`
	Hidden     bool        `json:"hidden"`
	Component  string      `json:"component"`
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Redirect   string      `json:"redirect"`
	Meta       *Meta       `json:"meta"`
	Children   []*Children `json:"children"`
}
