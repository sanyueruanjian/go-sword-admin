package dto

//SelectRoleDto 多条件查询角色
type SelectRoleArrayDto struct {
	Current int `json:"current" binding:"required"` //当前页
	Size    int `json:"size" binding:"required"`    //每页数据
	//TODO
	//Orders    `json:"orders" binding:"required"` //排序规则
	Blurry    string `json:"blurry"`    //模糊
	EndTime   string `json:"endTime"`   //过期时间
	StartTime string `json:"startTime"` //创建时间
}

//InsertRoleDto 新增角色
type InsertRoleDto struct {
	ID          int    `json:"id"`                           //id
	Level       int    `json:"level" binding:"required"`     //等级
	Name        string `json:"name" binding:"required"`      //姓名
	DataScope   string `json:"dataScope" binding:"required"` //数据权限
	Description string `json:"description" default:""`       //描述
	Depts       []int  `json:"depts"`                        //部门
}

//UpdateRoleDto 修改角色
type UpdateRoleDto struct {
	ID          int    `json:"id"`                            //id
	Level       int    `json:"level" binding:"required"`      //等级
	CreateBy    int    `json:"createBy"`                      //创建者 有值
	UpdatedBy   int    `json:"updatedBy"`                     //更新者 一般为空
	Name        string `json:"name" binding:"required"`       //姓名
	DataScope   string `json:"dataScope" binding:"required"`  //数据权限
	Description string `json:"description" default:""`        //描述
	UpdateTime  string `json:"updateTime" binding:"required"` //更新时间
	CreateTime  string `json:"createTime" binding:"required"` //创建时间
	Protection  bool   `json:"protection"`                    //
	Depts       []int  `json:"depts" binding:"required"`      //部门 (一般用为空)
	Menus       []int  `json:"menus"`                         //菜单列表
}

//DeleteRoleDto 删除角色
type DeleteRoleDto struct {
	//	直接用[]int 解析即可
}

//SelectRoleDto 查询单个角色
type SelectRoleDto struct {
	ID int `form:"id"` //role id
}

//SelectAllRoleDto 查询所有角色
type SelectAllRoleDto struct {
	//	无参数
}

//DownloadRoleDto 导出角色数据
type DownloadRoleDto struct {
	Current int `form:"current"`
	Size    int `form:"size"`
	//TODO
	//Orders `form:"orders"`
}

//SelectCurrentUserLevelDto 获取当前登录用户级别
type SelectCurrentUserLevelDto struct {
	//token
}

//UpdateRoleMenu 修改角色菜单
type UpdateRoleMenu struct {
	ID    int   `json:"id"`    //id
	Menus []int `json:"menus"` //菜单
}
