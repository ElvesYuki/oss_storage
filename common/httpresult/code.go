package httpresult

// 响应
var (
	// OK 正常返回
	OK = response(true, 200, "", "请求成功") // 通用成功
	// ErrReturn 异常返回
	ErrReturn = response(false, 500, "", "") // 异常返回
)
