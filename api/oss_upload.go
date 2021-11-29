package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"oss_storage/common/httperror"
	"oss_storage/common/httpresult"
	"oss_storage/service"
)

// OssSingleUploadHandler Oss单个文件上传接口
// @Summary Oss单个文件上传接口
// @Description Oss单个文件上传接口
// @Tags Oss上传相关接口
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "对象文件"
// @Param code formData string true "上传编码"
// @Success 200 {object} oss.BaseObject
// @Router /v1/oss/single/upload [post]
func OssSingleUploadHandler(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		err := new(httperror.XmoError).WithBiz(httperror.BIZ_ARG_ERROR)
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	code := c.PostForm("code")

	// 文件上传
	object, err := service.OssSingleUploadService(code, file)
	if err != nil {
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	httpresult.OK.NewBuilder().Data(object).Build(c)
	return
}

// OssMultipleUploadHandler Oss多文件上传接口
// @Summary Oss多文件上传接口
// @Description Oss多文件上传接口
// @Tags Oss上传相关接口
// @Accept multipart/form-data
// @Produce application/json
// @Param files formData file true "对象文件数组"
// @Param code formData string true "上传编码"
// @Success 200 {object} httpresult.Response
// @Router /v1/oss/multiple/upload [post]
func OssMultipleUploadHandler(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		zap.L().Error("上传文件异常", zap.Error(err))
		err = new(httperror.XmoError).WithBiz(httperror.BIZ_ARG_ERROR)
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	files, hasFile := form.File["files"]
	if !hasFile {
		err := new(httperror.XmoError).WithBiz(httperror.BIZ_ARG_ERROR)
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	code := c.PostForm("code")

	// 文件上传
	objects, err := service.OssMultipleUploadService(code, files)
	if err != nil {
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	httpresult.OK.NewBuilder().Data(objects).Build(c)
	return
}

// OssTextUploadHandler Oss文本上传接口
// @Summary Oss文本上传接口
// @Description Oss文本上传接口
// @Tags Oss上传相关接口
// @Accept multipart/form-data
// @Produce application/json
// @Param text formData string true "文本内容"
// @Param code formData string true "上传编码"
// @Success 200 {object} oss.BaseObject
// @Router /v1/oss/text/upload [post]
func OssTextUploadHandler(c *gin.Context) {

	text := c.PostForm("text")
	code := c.PostForm("code")

	// 文本上传
	object, err := service.OssTextUploadService(code, text)
	if err != nil {
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	httpresult.OK.NewBuilder().Data(object).Build(c)
}
