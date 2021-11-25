package oss

import (
	"oss_storage/common/httperror"
	"oss_storage/entity/dto"
)

type IUploadFactory interface {
	getFacType() string
	getOssObject(path *dto.OssStoragePathDTO, oType *ObjectTypeItem, object *UploadObject) (interface{}, error)
}

type UploadFactory struct{}

func (uf *UploadFactory) getFacType() string {
	return objectTypeEnum.OBJECT_TYPE_DEFAULT.ObjectType
}

func (uf *UploadFactory) getOssObject(path *dto.OssStoragePathDTO, oType *ObjectTypeItem, object *UploadObject) (interface{}, error) {

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
