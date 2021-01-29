package models

type FileResponse struct {
	Size     int64  `json:"size"`      //文件大小
	Path     string `json:"path"`      // 文件相对地址
	FullPath string `json:"full_path"` // 文件完整地址
	Name     string `json:"name"`      // 文件名
	Type     string `json:"type"`      // 文件类型
}
