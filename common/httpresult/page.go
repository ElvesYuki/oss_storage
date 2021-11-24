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

type RespBuilder struct {
	resp *ResponsePage
}

// NewBuilder 建造者
func (resp *ResponsePage) NewBuilder() *RespBuilder {
	b := new(RespBuilder)
	b.resp = resp
	return b
}

// Msg 添加信息
func (b *RespBuilder) Msg(msg string) *RespBuilder {
	b.resp.Msg = msg
	return b
}

// Data 添加数据
func (b *RespBuilder) Data(data interface{}) *RespBuilder {
	b.resp.Data = data
	return b
}

// Total 添加分页总数
func (b *RespBuilder) Total(total int64) *RespBuilder {
	pageTotal := int(math.Ceil(float64(total / int64(b.resp.PageSize))))
	b.resp.PageTotal = pageTotal
	b.resp.Total = total
	return b
}

// Error 添加错误信息
func (b *RespBuilder) Error(err error) *RespBuilder {
	xe := httperror.Create(err)
	b.resp.BizCode = xe.Biz.BizCode
	b.resp.Msg = xe.Biz.Msg
	return b
}

// Build 构造返回
func (b *RespBuilder) Build(c *gin.Context) {
	b.resp.Ts = time.Now().UnixMilli()
	respReturn := b.resp
	c.JSON(respReturn.Code, respReturn)
}
