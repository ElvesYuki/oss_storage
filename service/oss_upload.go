package service

import (
	"mime/multipart"
	"oss_storage/pkg/oss"
)

// OssSingleUploadService Oss单个文件上传接口
// @Author luohuan
// @Description Oss文本上传接口
// @Param text formData string true "文本内容"
// @Param code formData string true "上传编码"
// @Return
func OssSingleUploadService(code string, file *multipart.FileHeader) (interface{}, error) {
	// 文件上传
	object, err := oss.UploadObjectUtil(code, file)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// OssMultipleUploadService Oss多文件上传接口
// @Author luohuan
// @Description Oss文本上传接口
// @Param text formData string true "文本内容"
// @Param code formData string true "上传编码"
// @Return
func OssMultipleUploadService(code string, files []*multipart.FileHeader) (interface{}, error) {

	objects := make([]interface{}, len(files))

	for i, file := range files {
		// 文件上传
		object, err := oss.UploadObjectUtil(code, file)
		if err != nil {
			return nil, err
		}
		objects[i] = object
	}
	return objects, nil
}

// OssTextUploadService Oss文本上传接口
// @Author luohuan
// @Description Oss文本上传接口
// @Param text formData string true "文本内容"
// @Param code formData string true "上传编码"
// @Return
func OssTextUploadService(code string, text string) (interface{}, error) {

	// 文本上传
	object, err := oss.UploadObjectUtil(code, text)
	if err != nil {
		return nil, err
	}
	return object, nil
}
