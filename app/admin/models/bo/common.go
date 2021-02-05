package bo

//order 排序规则
type Order struct {
	Column string `json:"column"`
	Asc    string `json:"asc"`
}

type MenuPermission struct {
	Permission string `json:"permission"`
}

type DeptCommon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Job struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID        int    `json:"id"`
	Level     int    `json:"level"`
	Name      string `json:"name"`
	DataScope string `json:"dataScope"`
}

//paging 分页器所含字段 (公共父类)
type paging struct {
	Current          int     `json:"current"`
	CountID          int     `json:"count_id"`
	MaxLimit         int     `json:"maxLimit"`
	Page             int     `json:"page"`
	SearchCount      bool    `json:"searchCount"`
	Size             int     `json:"size"`
	Total            int     `json:"total"`
	HitCount         bool    `json:"hitCount"`
	OptimizeCountSql bool    `json:"optimizeCountSql"`
	Orders           []Order `json:"orders"`
}
