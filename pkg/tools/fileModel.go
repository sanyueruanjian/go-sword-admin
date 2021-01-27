package tools

import "time"

type ToolFile struct {
	// Id
	Id int64 `json:"id" db:"id"`
	// 原文件名称
	RealName string `json:"real_name" db:"real_name"`
	//文件别名
	FileName string `json:"file_name" db:"name"`
	//文件名的uuid
	FileUuid string `json:"file_uuid" db:"uuid"`
	//后缀
	FileSuffix string `json:"file_suffix" db:"suffix"`
	//路径
	FilePath string `json:"file_path" db:"path"`
	//类型
	FileType string `json:"file_type" db:"type"`
	//大小
	FileSize string `json:"file_size" db:"size"`
	//创建者
	CreateBy int64 `json:"create_by" db:"create_by"`
	//更新者
	UpDataBy int64 `json:"updata_by" db:"update_by"`
	//创建日期
	CreatTime time.Time `json:"creat_time" db:"create_time"`
	//更新时间
	UpDataTime time.Time `json:"updata_time" db:"update_time"`
}

type DownFile struct {
	// 文件名
	RealName string `json:"real_name" db:"real_name"`
	// 文件路径
	FilePath string `json:"file_path" db:"path"`
}