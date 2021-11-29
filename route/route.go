package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"oss_storage/api"
	_ "oss_storage/docs"
	"oss_storage/route/middleware"
)

func SetUp() *gin.Engine {

	r := gin.New()

	r.Use(middleware.GinLogger())

	r.Use(middleware.GinRecovery(true))

	// 注册业务路由
	v1 := r.Group("/v1")
	{

		v1.POST("/oss/single/upload", api.OssSingleUploadHandler)
		v1.POST("/oss/multiple/upload", api.OssMultipleUploadHandler)
		v1.POST("/oss/text/upload", api.OssTextUploadHandler)
		v1.POST("/oss/text/cover", api.OssTextCoverHandler)

		v1.GET("/test/idGenerate", api.ListIdGenerateHandler)
		v1.GET("/test/idGenerate/:id", api.GetIdGenerateByIdHandler)
		v1.GET("/test/get/id", api.GetIdHandler)
		v1.POST("/test/minio", api.TestMinio)
		v1.POST("/test/minio/handler", api.TestMinioHandler)
		v1.POST("/test/sensitive/handler", api.TestSensitiveFilter)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
