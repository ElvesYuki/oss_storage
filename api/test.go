package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oss_storage/common/httpresult"
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
		httpresult.ErrReturn.WithError(err).Build(c)
		return
	}
	httpresult.OK.WithData(id).Build(c)
	return

	//httpresult.OK.ConvToPage(1,5).
	//	WithBiz(httpresult.Err_Param_Big_Code).
	//	WithTotal(14).
	//	WithData(id).
	//	Build(c)

}
