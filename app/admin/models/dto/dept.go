package dto

// SelectDeptListDto 多条件查询部门
type SelectDeptDto struct {
	Pid       int    `form:"pid"`       // 上级部门（顶级部门为0，默认为0）
	Current   int    `form:"current"`   // 当前页
	Size      int    `form:"size"`      // 每页数据
	StartTime int64  `form:"startTime"` // 创建时间
	EndTime   int64  `form:"endTime"`   // 过期时间
	Enabled   bool   `form:"enabled"`   // 状态：1启用（默认）、0禁用
	Orders    string `form:"orders"`    // 排序规则
	Name      string `form:"name"`      // 模糊
	Sort      string `form:"sort"`      // id排序
}

// InsertDeptDto 新增部门
type InsertDeptDto struct {
	Pid      *int   `json:"pid" binding:"required"`      // 上级部门id
	SubCount *int   `json:"subCount" binding:"required"` // 子部门个数
	DeptSort int    `json:"deptSort" binding:"required"` // 部门排序
	IsTop    string `json:"isTop"`                       // 是否为顶级部门
	Enabled  string `json:"enabled" binding:"required"`  // 等级
	Name     string `json:"name" binding:"required"`     // 部门名称
}

// UpdateDeptDto 更新部门
type UpdateDeptDto struct {
	Pid         *int   `json:"pid" binding:"required"`      // 上级部门id
	SubCount    *int   `json:"subCount" binding:"required"` // 子部门个数
	DeptSort    *int   `json:"deptSort" binding:"required"` // 部门排序
	ID          int    `json:"id" binding:"required"`       // 部门id
	CreateBy    int    `json:"createBy"`
	UpdatedBy   int    `json:"updateBy"`
	CreateTime  int64  `json:"creatTime"`
	UpdateTime  int64  `json:"updateTime"`
	HasChildren bool   `json:"hasChildren"`                //是否有子节点
	Leaf        bool   `json:"leaf"`                       // 是否为子节点
	Enabled     string `json:"enabled" binding:"required"` // 等级
	IsTop       string `json:"isTop"`                      // 是否为根节点
	Name        string `json:"name" binding:"required"`    // 部门名称
	Label       string `json:"label"`                      // 原名称
}
