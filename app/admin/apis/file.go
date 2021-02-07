package apis

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"project/app/admin/models"
	"project/pkg/tools"
	"project/utils"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadFile 文件上传（任意类型文件）
// @Summary 文件上传（任意类型文件）
// @Description Author：JiaKunLi 2021/01/27
// @Tags 文件：文件管理 File Controller
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "file"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseFile
// @Router /api/file/uploadFile [post]
func UploadFile(c *gin.Context) {
	urlPerfix := fmt.Sprintf("http://%s/", c.Request.Host)
	var fileResponse models.FileResponse
	fileResponse, done := singleFile(c, fileResponse, urlPerfix, false)
	if done {
		return
	}
	app.ResponseSuccess(c, fileResponse)
}

// UploadFileImage 文件上传（图片）
// @Summary 文件上传（图片）
// @Description Author：JiaKunLi 2021/01/27
// @Tags 文件：文件管理 File Controller
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "file"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseFile
// @Router /api/file/uploadImage [post]
func UploadImage(c *gin.Context) {
	urlPerfix := fmt.Sprintf("http://%s/", c.Request.Host)
	var fileResponse models.FileResponse
	fileResponse, done := singleFile(c, fileResponse, urlPerfix, true)
	if done {
		return
	}
	app.ResponseSuccess(c, fileResponse)
}

//func UploadFile(c *gin.Context) {
//	tag, _ := c.GetPostForm("type")
//	urlPerfix := fmt.Sprintf("http://%s/", c.Request.Host)
//	var fileResponse FileResponse
//	if tag == "" {
//		app.ResponseErrorWithMsg(c, 200, "缺少标识")
//		return
//	} else {
//		switch tag {
//		case "1": // 单图
//			fileResponse, done := singleFile(c, fileResponse, urlPerfix)
//			if done {
//				return
//			}
//			app.ResponseSuccess(c, fileResponse)
//			return
//		case "2": // 多图
//			multipartFile := multipleFile(c, urlPerfix)
//			app.ResponseSuccess(c, multipartFile)
//			return
//		case "3": // base64
//			fileResponse = baseImg(c, fileResponse, urlPerfix)
//			app.ResponseSuccess(c, fileResponse)
//		}
//	}
//}

func baseImg(c *gin.Context, fileResponse models.FileResponse, urlPerfix string) models.FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	guid := uuid.New().String()
	fileName := guid + ".jpg"
	base64File := "static/uploadfile/" + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse = models.FileResponse{
		Size:     utils.GetFileSize(base64File),
		Path:     base64File,
		FullPath: urlPerfix + base64File,
		Name:     "",
		Type:     typeStr,
	}
	return fileResponse
}

func multipleFile(c *gin.Context, urlPerfix string) []models.FileResponse {
	files := c.Request.MultipartForm.File["file"]
	var multipartFile []models.FileResponse
	for _, f := range files {
		guid := uuid.New().String()
		fileName := guid + utils.GetExt(f.Filename)
		multipartFileName := "static/uploadfile/" + fileName
		e := c.SaveUploadedFile(f, multipartFileName)
		fileType, _ := tools.GetType(multipartFileName)
		if e == nil {
			fileResponse := models.FileResponse{
				Size:     utils.GetFileSize(multipartFileName),
				Path:     multipartFileName,
				FullPath: urlPerfix + multipartFileName,
				Name:     f.Filename,
				Type:     fileType,
			}
			multipartFile = append(multipartFile, fileResponse)
		}
	}
	return multipartFile
}

func singleFile(c *gin.Context, fileResponse models.FileResponse, urlPerfix string, image bool) (models.FileResponse, bool) {
	files, err := c.FormFile("file")

	if err != nil {
		app.ResponseError(c, app.CodeImageIsNotNull)
		return models.FileResponse{}, true
	}

	if image && utils.GetFileType(tools.GetExt(files.Filename)[1:]) != "image" {
		app.ResponseError(c, app.CodeFileImageFail)
		return models.FileResponse{}, true
	}

	// 上传文件至指定目录
	guid := uuid.New().String()
	fileName := guid + tools.GetExt(files.Filename)
	singleFile := "static/uploadfile/" + fileName
	err = c.SaveUploadedFile(files, singleFile)
	if err != nil {
		app.ResponseError(c, app.CodeFileUploadFail)
		return models.FileResponse{}, true
	}
	fileType, _ := tools.GetType(singleFile)
	fileResponse = models.FileResponse{
		Size:     utils.GetFileSize(singleFile),
		Path:     fileName,
		FullPath: urlPerfix + "api/file/download/" + fileName,
		Name:     files.Filename,
		Type:     fileType,
	}
	return fileResponse, false
}
