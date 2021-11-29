package oss

// UploadObjectUtil 上传文件工具类
func UploadObjectUtil(code string, object interface{}) (interface{}, error) {

	oss, err := uploadObjectHandler(code, object)
	if err != nil {
		return nil, err
	}
	return oss, nil
}

// CoverObjectUtil 覆盖上传文件工具类
func CoverObjectUtil(code string, url string, object interface{}) (interface{}, error) {

	oss, err := coverObjectHandler(code, url, object)
	if err != nil {
		return nil, err
	}
	return oss, nil
}
