package tools

//import (
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/google/uuid"
//	"mime/multipart"
//	"project/app/admin/dao/db"
//	"project/pkg/config"
//	"project/utils"
//	"strconv"
//	"strings"
//)
//
//func UploadPicture(form *multipart.Form, c *gin.Context) (imageUrl []string, idStringList string, err error) {
//	files := form.File["file"]
//	if len(files) > 0 {
//		// 文件解析存储
//		for _, f := range files {
//			// 文件信息
//			var tooFile ToolFile
//
//			tooFile.RealName = f.Filename
//
//			fileSuffix := GetExt(f.Filename)[1:]
//			tooFile.FileSuffix = fileSuffix
//
//			fh, _ := f.Open()
//			fileByte, _ := GetSize(fh)
//			tooFile.FileSize = FormatFileSize(int64(fileByte))
//
//			guid := uuid.New().String()
//			tooFile.FileUuid = guid
//
//			fileType := GetFileType(fileSuffix)
//			tooFile.FileType = fileType
//			if fileType != "image" {
//				err = errors.New("type error")
//				return
//			}
//
//			fileName := guid + "." + fileSuffix
//			tooFile.FileName = fileName
//
//			dir := fmt.Sprintf("%s/%s", config.ApplicationConfig.StaticPath, fileType)
//			path := fmt.Sprintf("%s/%s", dir, fileName)
//			tooFile.FilePath = path
//
//			err = IsNotExistMkDir(dir)
//			if err != nil {
//				return
//			}
//
//			// 上传文件到指定的目录
//			err = c.SaveUploadedFile(f, path)
//			if err != nil {
//				return
//			}
//			imageUrl = append(imageUrl, "/static/image/"+fileName)
//		}
//		createTime := utils.NowTime()
//		var (
//			idList []string
//			result sql.Result
//			id     int64
//		)
//		sqlStr := "insert into picture (picture, createTime) values (?, ?)"
//		for _, i := range imageUrl {
//			result, err = db.DB.Exec(sqlStr, i, createTime)
//			if err != nil {
//				return
//			}
//			id, err = result.LastInsertId()
//			if err != nil {
//				return
//			}
//			idList = append(idList, strconv.Itoa(int(id)))
//		}
//		idStringList = strings.Join(idList, ",")
//	}
//	return
//}
