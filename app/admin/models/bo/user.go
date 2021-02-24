package bo

//返回用户详细列表
type RecordUserHalf struct {
	Id           int    `json:"id"`
	DeptId       int    `json:"deptId"`
	CreateBy     int    `json:"createBy"`
	UpdateBy     int    `json:"updatedBy"`
	PwdResetTime int64  `json:"pwdResetTime"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	AvatarName   string `json:"avatarName"`
	AvatarPath   string `json:"avatarPath"`
	Email        string `json:"email"`
	NickName     string `json:"nickName"`
	Phone        string `json:"phone"`
	Username     string `json:"username"`
}

//返回用户详细列表
type RoleDeptJobBool struct {
	Enabled bool        `json:"enabled"`
	Gender  bool        `json:"gender"`
	Jobs    []*Job      `json:"jobs"`
	Role    []*Role     `json:"role"`
	Dept    *DeptCommon `json:"dept"`
}

//返回用户登录详细列表
type LoginData struct {
	Token string     `json:"token"`
	User  *LoginUser `json:"user"`
}

type LoginUser struct {
	DataScopes []int       `json:"dataScopes"`
	Roles      []string    `json:"roles"`
	User       *RecordUser `json:"user"`
}

type RecordUser struct {
	*RecordUserHalf
	*RoleDeptJobBool
}

//返回用户详细列表
type UserInfoListBo struct {
	Current          int           `json:"current"`
	CountID          int           `json:"count_id"`
	MaxLimit         int           `json:"maxLimit"`
	Pages            int           `json:"pages"`
	SearchCount      bool          `json:"searchCount"`
	Size             int           `json:"size"`
	Total            int           `json:"total"`
	HitCount         bool          `json:"hitCount"`
	OptimizeCountSql bool          `json:"optimizeCountSql"`
	Orders           []Order       `json:"orders"`
	Records          []*RecordUser `json:"records"` //记录查询
}

type authority struct {
	Authority string `json:"authority"`
}

//更新头像
type UpdateAvatarBo struct {
	Avatar string `json:"avatar" example:"xxx.png"`
}

//个人中心详细信息
type UserCenterInfoBo struct {
	DataScopes []int       `json:"dataScopes"`
	Roles      []string    `json:"roles"`
	User       *RecordUser `json:"user"`
}

type DownloadUserBo struct {
	Username     string `json:"username"`
	Role         string `json:"role"`
	Dept         string `json:"dept"`
	Jobs         string `json:"jobs"`
	Email        string `json:"email"`
	Enabled      string `json:"enabled"`
	Phone        string `json:"phone"`
	PwdResetTime string `json:"pwdResetTime"`
	CreateTime   string `json:"createTime"`
}
