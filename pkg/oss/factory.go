package oss

import (
	"oss_storage/common/httperror"
	"oss_storage/entity/dto"
)

type iUploadFactory interface {
	getFacType() string
	getOssObject(path *dto.OssStoragePathDTO, oType *objectTypeItem, object *uploadObject) (interface{}, error)
}

type uploadFactory struct{}

func (uf *uploadFactory) getFacType() string {
	return objectTypeEnum.OBJECT_TYPE_DEFAULT.ObjectType
}

func (uf *uploadFactory) getOssObject(path *dto.OssStoragePathDTO, oType *objectTypeItem, object *uploadObject) (interface{}, error) {

	// 检验格式是否正确
	_, hasFormat := path.ObjectSuffix[object.format]
	if !hasFormat {
		err := new(httperror.XmoError).WithBiz(httperror.BIZ_ILLEGAL_FILE_TYPE_ERROR)
		return nil, err
	}

	// 校验文件大小
	if path.MaxSize > 0 && object.size > path.MaxSize {
		err := new(httperror.XmoError).WithBiz(httperror.BIZ_ILLEGAL_FILE_SIZE_ERROR)
		return nil, err
	}

	// 调用上传
	// 如果cover为true bucketName 和ObjectName 不为空 ，这是覆盖
	if object.cover && len(object.bucketName) != 0 && len(object.objectName) != 0 {
		// 覆盖
		if !CheckObjectExist(object.bucketName, object.objectName) {
			err := new(httperror.XmoError).WithBiz(httperror.BIZ_OSS_OBJECT_NOT_EXIST_ERROR)
			return nil, err
		}
	} else {
		// 写入上传路径
		object.bucketName = path.BucketName
		object.objectName = GenerateFileUrl(path.ObjectPath, object.ossName)
	}

	// 上传文件
	UO.putObject(object)

	// 生成对象
	objectReturn := &BaseObject{
		UploadStatus: 1,
		Url:          "/" + object.bucketName + "/" + object.objectName,
		FileName:     object.ossName,
		ContentType:  object.contentType,
		Bucket:       object.bucketName,
		Object:       object.objectName,
		Size:         object.size,
		Format:       object.format,
	}

	return objectReturn, nil
}

type jsonFactory struct{}

func (jf *jsonFactory) getFacType() string {
	return objectTypeEnum.OBJECT_TYPE_JSON.ObjectType
}

func (jf *jsonFactory) getOssObject(path *dto.OssStoragePathDTO, oType *objectTypeItem, object *uploadObject) (interface{}, error) {

	ossReturn, err := defaultUploadFactory.getOssObject(path, oType, object)
	if err != nil {
		return nil, err
	}

	jsonObject := &JsonObject{}

	switch v := ossReturn.(type) {
	case *BaseObject:
		jsonObject.BaseObject = *v
	default:
		return nil, new(httperror.XmoError).WithBiz(httperror.BIZ_OSS_UNKNOWN_TYPE_ERROR)
	}
	jsonObject.ContentExcerpt = "文本节选"

	return jsonObject, nil
}

type htmlFactory struct{}

func (jf *htmlFactory) getFacType() string {
	return objectTypeEnum.OBJECT_TYPE_HTML.ObjectType
}

func (jf *htmlFactory) getOssObject(path *dto.OssStoragePathDTO, oType *objectTypeItem, object *uploadObject) (interface{}, error) {

	ossReturn, err := defaultUploadFactory.getOssObject(path, oType, object)
	if err != nil {
		return nil, err
	}

	htmlObject := &HtmlObject{}

	switch v := ossReturn.(type) {
	case *BaseObject:
		htmlObject.BaseObject = *v
	default:
		return nil, new(httperror.XmoError).WithBiz(httperror.BIZ_OSS_UNKNOWN_TYPE_ERROR)
	}

	htmlObject.ContentExcerpt = "文本节选"

	return htmlObject, nil
}
