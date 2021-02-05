package dto

//SelectRoleDto 多条件查询角色
type SelectRoleArrayDto struct {
	Current   int    `form:"current" binding:"required"` //当前页
	Size      int    `form:"size" binding:"required"`    //每页数据
	Orders    string `form:"orders" binding:"required"`  //排序规则
	Blurry    string `form:"blurry"`                     //模糊
	EndTime   string `form:"endTime"`                    //过期时间
	StartTime string `form:"startTime"`                  //创建时间
}

//InsertRoleDto 新增角色
type InsertRoleDto struct {
	ID          int    `form:"id"`                           //id
	Level       int    `form:"level" binding:"required"`     //等级
	Name        string `form:"name" binding:"required"`      //姓名
	DataScope   string `form:"dataScope" binding:"required"` //数据权限
	Description string `form:"description" default:""`       //描述
	Depts       []int  `form:"depts"`                        //部门
}

//UpdateRoleDto 修改角色
type UpdateRoleDto struct {
	ID          int    `form:"id"`                            //id
	Level       int    `form:"level" binding:"required"`      //等级
	CreateBy    int    `form:"createBy"`                      //创建者 有值
	UpdatedBy   int    `form:"updatedBy"`                     //更新者 一般为空
	Name        string `form:"name" binding:"required"`       //姓名
	DataScope   string `form:"dataScope" binding:"required"`  //数据权限
	Description string `form:"description" default:""`        //描述
	UpdateTime  int64  `form:"updateTime" binding:"required"` //更新时间
	CreateTime  int64  `form:"createTime" binding:"required"` //创建时间
	Protection  string `form:"protection"`                    //
	Depts       []int  `form:"depts" binding:"required"`      //部门 (一般用为空)
	Menus       []int  `form:"menus"`                         //菜单列表
}

// 角色菜单分配
type RoleMenus struct {
	ID    int   `form:"id"`
	Menus []int `form:"menus"`
}
