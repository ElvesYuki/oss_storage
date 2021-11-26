package oss

// UploadObjectUtil 上传文件工具类
func UploadObjectUtil(code string, object interface{}) (interface{}, error) {

	oss, err := uploadObjectHandler(code, object)
	if err != nil {
		return nil, err
	}
	return oss, nil
}
