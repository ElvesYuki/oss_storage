package route

import (
	"github.com/gin-gonic/gin"
	"oss_storage/api"
	"oss_storage/setting/logger"
)

func SetUp() *gin.Engine {

	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	v1 := r.Group("/v1")
	{
		v1.GET("/test", api.Test)
	}

	return r
}
