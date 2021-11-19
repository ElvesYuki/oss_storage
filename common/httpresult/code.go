package httpresult

// 响应
var (
	// OK
	OK = response(true, 200, "", "请求成功") // 通用成功
	//Err = response(false, 500, "", "请求失败")   // 通用错误
	//
	//// 服务级错误码
	//Err_Param     = response(false, 500, Err_Param_Big_Code.BizCode, "")
	//ErrSignParam = response(false,500, "签名参数有误", "")
)

// 业务码
var (
	Err_Param_Big_Code      = bizcode("err_param", "参数错误", "")
	Err_Sign_Param_Big_Code = bizcode("err_sign_param", "签名参数错误", "")
)

type BizCode struct {
	BizCode string
	Msg     string
	desc    string
}

func bizcode(bizCode string, msg string, desc string) BizCode {
	return BizCode{
		BizCode: bizCode,
		Msg:     msg,
		desc:    desc,
	}
}
