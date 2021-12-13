package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"oss_storage/common/httperror"
	"oss_storage/common/httpresult"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 2
	CommonUserIdKey     = "userId"
)

var MySecret = []byte("elves")

type MyClaims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			errReturn := new(httperror.XmoError).WithBiz(httperror.BIZ_DEFAULT_ERROR)
			httpresult.ErrReturn.NewBuilder().Error(errReturn).Build(c)
			return
		}

		mc, err := ParseToken(authHeader)
		if err != nil {
			httpresult.ErrReturn.NewBuilder().Error(err).Build(c)
			return
		}
		c.Set(CommonUserIdKey, mc.UserId)
		c.Next()
	}
}

// GenToken 生成JWT
func GenToken(userId int64) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userId, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(TokenExpireDuration)),
			Issuer:    "elvesyuki",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, err
	}
	return nil, errors.New("invalid token")
}
