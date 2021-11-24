package httpresult

import (
	"github.com/gin-gonic/gin"
	"oss_storage/common/httperror"
	"time"
)

type Response struct {
	Ts      int64       `json:"ts"`      // 响应时间戳
	Success bool        `json:"success"` // 响应是否成功
	Code    int         `json:"code"`    // 响应码
	Msg     string      `json:"msg"`     // 响应信息
	BizCode string      `json:"bizCode"` // 响应业务码
	Data    interface{} `json:"data"`    // 响应数据
}

func response(success bool, code int, bizCode string, msg string) *Response {
	return &Response{
		Success: success,
		Code:    code,
		Msg:     msg,
		BizCode: bizCode,
	}
}

type ResBuilder struct {
	res *Response
}

// NewBuilder 建造者
func (res *Response) NewBuilder() *ResBuilder {
	b := new(ResBuilder)
	b.res = res
	return b
}

// Msg 添加信息
func (b *ResBuilder) Msg(msg string) *ResBuilder {
	b.res.Msg = msg
	return b
}

// Data 添加数据
func (b *ResBuilder) Data(data interface{}) *ResBuilder {
	b.res.Data = data
	return b
}

// Error 添加错误信息
func (b *ResBuilder) Error(err error) *ResBuilder {
	xe := httperror.Create(err)
	b.res.BizCode = xe.Biz.BizCode
	b.res.Msg = xe.Biz.Msg
	return b
}

// Build 构造返回
func (b *ResBuilder) Build(c *gin.Context) {
	b.res.Ts = time.Now().UnixMilli()
	resReturn := b.res
	c.JSON(resReturn.Code, resReturn)
}
