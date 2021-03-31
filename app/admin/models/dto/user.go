package dto

// 定义请求的参数结构体

// UserLoginDto 登录请求参数
type UserLoginDto struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	Code     string `json:"code" binding:"required"`     // 验证码
	UuId     string `json:"uuid" binding:"required"`     // 验证码id
}

//SelectUserInfoArrayDto 查询用户详细列表
type SelectUserInfoArrayDto struct {
	Current   int    `form:"current"` //current 当前页
	Size      int    `form:"size"`    // 当前页最大数据量
	DepID     int    `form:"deptId"`  //本门id
	Orders    string `form:"orders"`  //排序规则与字段
	Blurry    string `form:"blurry"`  //模糊
	Enabled   bool   `form:"enabled"` //是否激活
	StartTime int64  `form:"startTime"`
	EndTime   int64  `form:"endTime"`
}

//SelectUserInfoDto 查询"单个"用户详细 个人中心信息使用
type SelectUserInfoDto struct {
	//token
}

//InsertUserDto 新增用户
type InsertUserDto struct {
	DeptId   int    `json:"deptid" binding:"required"`   //部门
	ID       int    `json:"id"`                          //id (目前不传)
	Phone    int64  `json:"phone" binding:"required"`    //手机号
	Email    string `json:"email"`                       //邮箱 (目前不传)
	NickName string `json:"nickName" binding:"required"` //昵称
	UserName string `json:"username" binding:"required"` //用户名
	Gender   string `json:"gender" binding:"required"`   //性别
	Jobs     []int  `json:"jobs" binding:"required"`     //职位
	Roles    []int  `json:"roles" binding:"required"`    //角色
}

//UpdateUserDto 修改用户
type UpdateUserDto struct {
	DeptId     int    `json:"deptid" binding:"required"`   //部门
	ID         int    `json:"id" binding:"required"`       //用户id
	Enabled    bool   `json:"enabled"`                     //是否激活
	Email      string `json:"email"`                       //邮箱 (目前不传)
	NickName   string `json:"nickName" binding:"required"` //昵称
	Phone      string `json:"phone" binding:"required"`    //手机号
	UserName   string `json:"username" binding:"required"` //用户名
	AvatarName string `json:"avatarName"`                  //头像名字
	AvatarPath string `json:"avatarPath"`                  //头像地址
	Gender     string `json:"gender"`                      //性别
	Jobs       []int  `json:"jobs"`                        //职位
	Roles      []int  `json:"roles"`                       //角色
}

//DeleteUserDto 删除用户
type DeleteUserDto struct {
	//直接用[]int 去接受用户id列表即可
}

//UpdateUserCenterDto 修改用户个人信息
type UpdateUserCenterDto struct {
	NickName string `json:"nickName" binding:"required"` //昵称
	Phone    string `json:"phone" binding:"required"`    //手机号
	Gender   bool   `json:"gender"`                      //性别
	Id       int    `json:"id" binding:"required"`       //id
}

//DownloadUserInfoDto 导出用户数据
type DownloadUserInfoDto struct {
	Current int    `form:"current"` //当前页
	Size    int    `form:"size"`    //展示数量
	Orders  string `form:"orders"`  //排序规则
	Enabled bool   `form:"enabled"` //是否激活
}

//UpdateUserAvatarDto 修改用户头像
type UpdateUserAvatarDto struct {
	//	from 文件
}

//UpdateUserPassDtp 修改用户密码
type UpdateUserPassDto struct {
	OldPass string `json:"oldPass" binding:"required"` //旧密码
	NewPass string `json:"newPass" binding:"required"` //新密码
}
