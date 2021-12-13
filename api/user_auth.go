package api

import (
	"github.com/gin-gonic/gin"
	"oss_storage/common/httpresult"
	"oss_storage/entity/param"
	"oss_storage/service"
)

func UserSignupHandler(c *gin.Context) {
	var userSignup param.UserSignup
	if err := c.ShouldBindJSON(&userSignup); err != nil {
		httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
		return
	}

	service.UserSignup(&userSignup)

}

func UserLoginHandler(c *gin.Context) {

}

func UserLogoutHandler(c *gin.Context) {

}
