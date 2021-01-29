package bo

type recordUser struct {
	Id           int    `json:"id"`
	DeptId       int    `json:"deptId"`
	PwdResetTime int    `json:"pwdResetTime"`
	CreateBy     int    `json:"createBy"`
	CreateTime   int    `json:"createTime"`
	UpdatedBy    int    `json:"updatedBy"`
	UpdateTime   int    `json:"updateTime"`
	AvatarName   string `json:"avatarName"`
	AvatarPath   string `json:"avatarPath"`
	Email        string `json:"email"`
	NickName     string `json:"nickName"`
	Phone        string `json:"phone"`
	Username     string `json:"username"`
	Enabled      bool   `json:"enabled"`
	Gender       bool   `json:"gender"`
	Jobs         *job   `json:"jobs"`
	Dept         *dept  `json:"dept"`
	Role         *role  `json:"role"`
}

type authority struct {
	Authority string `json:"authority"`
}

//返回用户详细列表
type UserInfoListBo struct {
	*paging           //分页器相关
	Records []*record `json:"records"` //记录查询
}

//更新头像
type UpdateAvatarBo struct {
	Avatar string `json:"avatar" example:"xxx.png"`
}

//个人中心详细信息
type UserCenterInfoBo struct {
	AccountNonExpired     bool         `json:"accountNonExpired"`
	AccountNonLocked      bool         `json:"accountNonLocked"`
	CredentialsNonExpired bool         `json:"credentialsNonExpired"`
	Enabled               bool         `json:"enabled"`
	Password              string       `json:"password"`
	Username              string       `json:"username"`
	Authorities           []*authority `json:"authorities"`
}
