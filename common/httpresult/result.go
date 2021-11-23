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

// Build 请求返回
func (res *Response) Build(c *gin.Context) {
	res.Ts = time.Now().UnixMilli()
	c.JSON(res.Code, res)
}

// WithMsg 添加信息
func (res *Response) WithMsg(msg string) *Response {
	res.Msg = msg
	return res
}

// WithData 添加数据
func (res *Response) WithData(data interface{}) *Response {
	res.Data = data
	return res
}

// WithError 添加错误
func (res *Response) WithError(err error) *Response {
	xe := httperror.Create(err)
	res.BizCode = xe.Biz.BizCode
	res.Msg = xe.Biz.Msg
	return res
}
