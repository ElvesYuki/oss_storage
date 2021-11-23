package httperror

// 业务码
var (
	BIZ_DEFAULT_ERROR = bizcode("default_err", "默认错误", "")
	BIZ_System_ERROR  = bizcode("system_err", "系统错误", "")
	BIZ_Param_ERROR   = bizcode("param_err", "参数错误", "")
	BIZ_ARG_ERROR     = bizcode("arg_err", "参数错误", "")

	BIZ_SQL_NOT_EXIST_ERROR = bizcode("sql_not_exist_err", "系统错误", "")
	BIZ_SQL_INSERT_ERROR    = bizcode("sql_insert_err", "系统错误", "")
	BIZ_SQL_UPDATE_ERROR    = bizcode("sql_update_err", "系统错误", "")
	BIZ_SQL_DELETE_ERROR    = bizcode("sql_delete_err", "系统错误", "")
)

type BizCode struct {
	BizCode string
	Msg     string
	desc    string
}

func bizcode(bizCode string, msg string, desc string) *BizCode {
	return &BizCode{
		BizCode: bizCode,
		Msg:     msg,
		desc:    desc,
	}
}
