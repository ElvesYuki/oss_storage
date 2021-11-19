package httpresult

import (
	"github.com/gin-gonic/gin"
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

// WithBiz 添加错误码
func (res *Response) WithBiz(biz BizCode) *Response {
	res.BizCode = biz.BizCode
	res.Msg = biz.Msg
	return res
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
