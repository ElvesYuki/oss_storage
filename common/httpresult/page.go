package httpresult

import (
	"github.com/gin-gonic/gin"
	"math"
	"oss_storage/common/httperror"
	"time"
)

type ResponsePage struct {
	Ts        int64       `json:"ts"`        // 响应时间戳
	Success   bool        `json:"success"`   // 响应是否成功
	Code      int         `json:"code"`      // 响应码
	Msg       string      `json:"msg"`       // 响应信息
	BizCode   string      `json:"bizCode"`   // 响应业务码
	Data      interface{} `json:"data"`      // 响应数据
	PageIndex int         `json:"pageIndex"` // 当前页
	PageSize  int         `json:"pageSize"`  // 页大小
	PageTotal int         `json:"pageTotal"` // 总页数
	Total     int64       `json:"total"`     // 总数量
}

func (resp *ResponsePage) Build(c *gin.Context) {
	resp.Ts = time.Now().UnixMilli()
	c.JSON(resp.Code, resp)
}

// ConvToPage 转换成分页返回
func (res *Response) ConvToPage(pageIndex int, pageSize int) *ResponsePage {
	return &ResponsePage{
		Success:   res.Success,
		Code:      res.Code,
		Msg:       res.Msg,
		BizCode:   res.BizCode,
		Data:      res.Data,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
}

// WithMsg 添加信息
func (resp *ResponsePage) WithMsg(msg string) *ResponsePage {
	resp.Msg = msg
	return resp
}

// WithData 添加数据
func (resp *ResponsePage) WithData(data interface{}) *ResponsePage {
	resp.Data = data
	return resp
}

// WithTotal 添加分页总数
func (resp *ResponsePage) WithTotal(total int64) *ResponsePage {
	pageTotal := int(math.Ceil(float64(total / int64(resp.PageSize))))
	resp.PageTotal = pageTotal
	resp.Total = total
	return resp
}

// WithError 添加错误
func (resp *ResponsePage) WithError(err error) *ResponsePage {
	xe := httperror.Create(err)
	resp.BizCode = xe.Biz.BizCode
	resp.Msg = xe.Biz.Msg
	return resp
}
