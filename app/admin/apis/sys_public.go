package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/golang-module/carbon/v2"
	"go-admin/app/admin/service"
	"os"
	"path"
	"strconv"
	"strings"
)

type SysPublic struct {
	api.Api
}

// UploadFile
// @Summary 上传文件
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/public/upload [post]
// @Security Bearer
func (e SysPublic) UploadFile(c *gin.Context) {
	s := service.SysPublic{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var filePath string
	//只有一个
	for _, file := range files {
		// 上传文件至指定目录
		//目录不存在创建目录
		uploadDir := "static/uploadfile/" + carbon.Now().Format("Ymd") + "/"
		_, dirErr := os.Stat(uploadDir)
		if os.IsNotExist(dirErr) {
			err := os.Mkdir(uploadDir, os.ModePerm)
			if err != nil {
				return
			}
		}
		filePath = uploadDir + file.Filename

		//文件存在重命名文件
		filePath = e.getRealFilePath(filePath, 0)

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			e.Logger.Errorf("save file error, %s", err.Error())
			e.Error(500, err, "")
			return
		}
	}
	e.OK(filePath, "上传成功")
}

func (e SysPublic) getRealFilePath(filePath string, i int) string {
	endFilePath := filePath
	//文件存在重命名文件
	openFile, openErr := os.Open(filePath)
	if openErr != nil {
		if os.IsNotExist(openErr) {
			endFilePath = filePath
		}
	} else {
		err := openFile.Close()
		if err != nil {
		}
		fileSuffix := path.Ext(filePath)
		var replace string
		if i == 0 {
			replace = strings.Replace(filePath, fileSuffix, "("+strconv.Itoa(i+1)+")"+fileSuffix, -1)
		} else {
			replace = strings.Replace(filePath, "("+strconv.Itoa(i)+")"+fileSuffix, "("+strconv.Itoa(i+1)+")"+fileSuffix, -1)
		}
		return e.getRealFilePath(replace, i+1)
	}
	return endFilePath
}
