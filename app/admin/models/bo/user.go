package bo

type RecordUserHalf struct {
	Id           int    `json:"id"`
	DeptId       int    `json:"deptId"`
	PwdResetTime int    `json:"pwdResetTime"`
	CreateBy     int    `json:"createBy"`
	CreateTime   int    `json:"createTime"`
	UpdateBy     int    `json:"updatedBy"`
	UpdateTime   int    `json:"updateTime"`
	AvatarName   string `json:"avatarName"`
	AvatarPath   string `json:"avatarPath"`
	Email        string `json:"email"`
	NickName     string `json:"nickName"`
	Phone        string `json:"phone"`
	Username     string `json:"username"`
}

type RoleDeptJobBool struct {
	Enabled bool        `json:"enabled"`
	Gender  bool        `json:"gender"`
	Jobs    []*Job      `json:"jobs"`
	Role    []*Role     `json:"role"`
	Dept    *DeptCommon `json:"dept"`
}

type RecordUser struct {
	*RecordUserHalf
	*RoleDeptJobBool
}

type authority struct {
	Authority string `json:"authority"`
}

//返回用户详细列表
type UserInfoListBo struct {
	*paging               //分页器相关
	Records []*RecordUser `json:"records"` //记录查询
}

//更新头像
type UpdateAvatarBo struct {
	Avatar string `json:"avatar" example:"xxx.png"`
}

//个人中心详细信息
type UserCenterInfoBo struct {
	DataScopes []string    `json:"dataScopes"`
	Roles      []string    `json:"roles"`
	User       *RecordUser `json:"user"`
}
