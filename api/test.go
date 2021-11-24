package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"net/http"
	"oss_storage/common/httpresult"
	"oss_storage/pkg/oss"
	"oss_storage/service"
	"strconv"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test",
	})
}

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
