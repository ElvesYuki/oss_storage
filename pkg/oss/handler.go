package oss

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"oss_storage/common/httperror"
	"oss_storage/entity/dto"
)

var ossStoragePathDTOMap map[string]*dto.OssStoragePathDTO

// UploadObject 封装的上传对象
type UploadObject struct {
	cover       bool
	ossType     string
	ossContent  interface{}
	ossReader   io.Reader
	ossName     string
	format      string
	contentType string
	size        int64
	bucketName  string
	objectName  string
	versionId   string
}

type IUploadHandler interface {
	UploadObject(code string, object interface{}) (interface{}, error)
}

type UploadHandler struct{}

func (handler *UploadHandler) UploadObject(code string, object interface{}) (objectReturn interface{}, err error) {

	path, hasCode := ossStoragePathDTOMap[code]
	if !hasCode {
		zap.L().Error("没有对应的存储编码")
		err = new(httperror.XmoError).WithBiz(httperror.BIZ_OSS_PATH_CODE_NOT_EXIST_ERROR)
		return nil, err
	}

	oType, hasType := objectTypeMap[path.ObjectType]
	if !hasType {
		zap.L().Error("没有对应的操作类型")
		err = new(httperror.XmoError).WithBiz(httperror.BIZ_OSS_UNKNOWN_TYPE_ERROR)
		return nil, err
	}

	switch o := object.(type) {
	case *multipart.FileHeader:
		// 上传文件
		objectReturn, err = uploadMultipartFile(path, oType, o)
		if err != nil {
			return nil, err
		}
	case string:
		// 上传字符串，可能是json或者html
		objectReturn, err = uploadString(path, oType, o)
		if err != nil {
			return nil, err
		}
	default:
		// 未知的文件类型
		err = new(httperror.XmoError).WithBiz(httperror.BIZ_OSS_OBJECT_UNKNOWN_TYPE_ERROR)
		return nil, err
	}
	return objectReturn, nil

}

func uploadMultipartFile(path *dto.OssStoragePathDTO, oType *ObjectTypeItem, object *multipart.FileHeader) (objectReturn interface{}, err error) {

	// 封装上传对象

	srcReader, err := object.Open()
	defer srcReader.Close()
	if err != nil {
		return nil, err
	}

	uploadObject := &UploadObject{
		cover:       false,
		ossType:     oType.ObjectType,
		ossContent:  object,
		ossReader:   srcReader,
		ossName:     object.Filename,
		format:      GetFileSuffix(object.Filename),
		contentType: object.Header.Get("Content-Type"),
		size:        object.Size,
	}

	// 分配器分配对象工厂

	// 工厂上传
	objectReturn, err = UF.getOssObject(path, oType, uploadObject)
	if err != nil {
		return nil, err
	}

	fmt.Println(uploadObject)
	fmt.Println(objectReturn)

	return objectReturn, err
}

func uploadString(path *dto.OssStoragePathDTO, oType *ObjectTypeItem, object string) (objectReturn interface{}, err error) {

	return nil, err
}
