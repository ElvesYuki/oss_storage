package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oss_storage/common/httpresult"
	"oss_storage/pkg/oss"
)

// OssSingleUploadHandler 单文件上传
func OssSingleUploadHandler(c *gin.Context) {

	file, _ := c.FormFile("file")
	code := c.PostForm("code")

	object, err := oss.UploadObjectHandler(code, file)
	if err != nil {
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	fmt.Println(object)
	httpresult.OK.NewBuilder().Data(object).Build(c)
}

// OssMultipleUploadHandler 多文件上传
func OssMultipleUploadHandler(c *gin.Context) {

	form, _ := c.MultipartForm()
	files, _ := form.File["files"]
	code := c.PostForm("code")

	objects := make([]interface{}, len(files))

	for i, file := range files {
		object, err := oss.UploadObjectHandler(code, file)
		if err != nil {
			httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
			return
		}
		objects[i] = object
	}

	httpresult.OK.NewBuilder().Data(objects).Build(c)
	return
}
