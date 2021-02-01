package models

// BaseModel orm公有字段
type BaseModel struct {
	ID         int    `gorm:"primary_key" json:"id"`                   //ID
	IsDeleted  []byte `gorm:"default:[]byte{0}" json:"is_deleted"`     //默认为零
	CreateTime int64  `gorm:"autoCreateTime:milli" json:"create_time"` //创建日期 默认当前时间戳 毫秒
	UpdateTime int64  `gorm:"autoUpdateTime:milli" json:"update_time"` //更新时间 默认当前时间戳 毫秒
}
