package bo

type SelectMenuBo struct {
	CreateBy    int         `json:"createBy"`
	UpdatedBy   int         `json:"updatedBy"`
	SubCount    int         `json:"subCount"`
	MenuSort    int         `json:"menuSort"`
	ID          int         `json:"id"`
	Pid         int         `json:"pid"`
	Type        int         `json:"type"`
	Cache       bool        `json:"cache"`
	Hidden      bool        `json:"hidden"`
	Leaf        bool        `json:"leaf"`
	Iframe      bool        `json:"iframe"`
	HasChildren bool        `json:"hasChildren"`
	CreateTime  string      `json:"createTime"`
	UpdateTime  string      `json:"updateTime"`
	Label       string      `json:"label"`
	Icon        string      `json:"icon"`
	Component   string      `json:"component"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Permission  string      `json:"permission"`
	Title       string      `json:"title"`
	Children    []*Children `json:"children"`
}

type Meta struct {
	Icon    string `json:"icon"`
	NoCache bool   `json:"noCache"`
	Title   string `json:"title"`
}

type Children struct {
	Component string `json:"component"`
	Name      string `json:"name"`
	Path      string `json:"path"`
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

type ReturnToAllMenusBo struct {
	Cache       bool        `json:"cache"`
	Children    interface{} `json:"children"`
	Component   string      `json:"component"`
	CreateBy    int         `json:"createBy"`
	CreateTime  int64       `json:"createTime"`
	HasChildren bool        `json:"hasChildren"`
	Hidden      bool        `json:"hidden"`
	Icon        string      `json:"icon"`
	ID          int         `json:"id"`
	Iframe      bool        `json:"iframe"`
	Label       string      `json:"label"`
	Leaf        bool        `json:"leaf"`
	MenuSort    int         `json:"menuSort"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Permission  string      `json:"permission"`
	Pid         int         `json:"pid"`
	SubCount    int         `json:"subCount"`
	Title       string      `json:"title"`
	Type        int         `json:"type"`
	UpdatedBy   int         `json:"updatedBy"`
	UpdateTime  int64       `json:"updateTime"`
}

type DownloadMenuInfoBo struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	Permission string `json:"permission"`
	IFrame     string `json:"i_frame"`
	Hidden     string `json:"hidden"`
	Cache      string `json:"cache"`
	CreateTime string `json:"create_time"`
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
