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
