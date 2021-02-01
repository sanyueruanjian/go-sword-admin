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
	Depts       string `form:"depts"`                        //部门
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
	UpdateTime  string `form:"updateTime" binding:"required"` //更新时间
	CreateTime  string `form:"createTime" binding:"required"` //创建时间
	Protection  string `form:"protection"`                    //
	Depts       string `form:"depts" binding:"required"`      //部门 (一般用为空)
	Menus       string `form:"menus"`                         //菜单列表
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
