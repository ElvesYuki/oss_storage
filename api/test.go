package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"net/http"
	"oss_storage/common/httpresult"
	"oss_storage/pkg/oss"
	"oss_storage/pkg/sensitiveword"
	"oss_storage/service"
	"strconv"
)

func ListIdGenerateHandler(c *gin.Context) {

	data, err := service.ListIdGenerate()
	if err != nil {

	}

	c.JSON(http.StatusOK, data)

}

func GetIdGenerateByIdHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	data, err := service.GetIdGenerateById(id)
	if err != nil {

	}

	c.JSON(http.StatusOK, data)

}

func GetIdHandler(c *gin.Context) {

	module := c.Query("module")
	num := c.Query("num")

	fmt.Print("第=", num, "=次请求====")

	id, err := service.GetId(module)

	if err != nil {
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	httpresult.OK.NewBuilder().Data(id).Build(c)

	return
}

func TestMinio(c *gin.Context) {

	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)

	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	uploadInfo, err := oss.OC.GetClient().MinioClient.PutObject(context.Background(),
		"oss-storage",
		file.Filename, src, -1,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		})
	if err != nil {
		return
	}
	fmt.Println(uploadInfo)
}

func TestMinioHandler(c *gin.Context) {

	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)

	object, err := oss.UploadObjectUtil("testoss", file)
	if err != nil {
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}
	fmt.Println(object)
	httpresult.OK.NewBuilder().Data(object).Build(c)

}

func TestSensitiveFilter(c *gin.Context) {
	text := c.PostForm("text")
	fmt.Println(text)
	textReturn := sensitiveword.SensitiveFilter(text)
	httpresult.OK.NewBuilder().Data(textReturn).Build(c)
}
