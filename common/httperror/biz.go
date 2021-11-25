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

	BIZ_OSS_PATH_CODE_NOT_EXIST_ERROR = bizcode("oss_path_code_not_exist_err", "存储编码不存在", "")
	BIZ_OSS_OBJECT_UNKNOWN_TYPE_ERROR = bizcode("oss_path_unknown_type_exist_err", "存储编码不存在", "")
	BIZ_OSS_UNKNOWN_TYPE_ERROR        = bizcode("oss_unknown_type_exist_err", "存储类型不存在", "")

	BIZ_ILLEGAL_FILE_TYPE_ERROR = bizcode("oss_illegal_file_type_err", "文件类型不合法", "")
	BIZ_ILLEGAL_FILE_SIZE_ERROR = bizcode("oss_illegal_file_size_err", "文件大小不合法", "")

	BIZ_OSS_OBJECT_NOT_EXIST_ERROR = bizcode("oss_object_not_exist_err", "请求文件资源不存在", "")
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
