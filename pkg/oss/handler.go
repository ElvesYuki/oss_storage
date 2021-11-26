package oss

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"oss_storage/common/httperror"
	"oss_storage/entity/dto"
	"strings"
)

var ossStoragePathDTOMap map[string]*dto.OssStoragePathDTO

// UploadObject 封装的上传对象
type uploadObject struct {
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

// uploadObjectHandler 上传处理器
func uploadObjectHandler(code string, object interface{}) (objectReturn interface{}, err error) {

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

func uploadMultipartFile(path *dto.OssStoragePathDTO, oType *objectTypeItem, object *multipart.FileHeader) (objectReturn interface{}, err error) {

	// 封装上传对象
	srcReader, err := object.Open()
	defer srcReader.Close()
	if err != nil {
		return nil, err
	}

	uploadObject := &uploadObject{
		cover:       false,
		ossType:     oType.ObjectType,
		ossContent:  object,
		ossReader:   srcReader,
		ossName:     object.Filename,
		format:      GetFileSuffix(object.Filename),
		contentType: object.Header.Get("Content-Type"),
		size:        object.Size,
	}

	// 调度器 分配 对象工厂
	uploadFactory, err := uploadObjectDispatcher(oType)
	if err != nil {
		return nil, err
	}

	// 工厂上传
	objectReturn, err = uploadFactory.getOssObject(path, oType, uploadObject)
	if err != nil {
		return nil, err
	}

	fmt.Println(uploadObject)
	fmt.Println(objectReturn)

	return objectReturn, err
}

func uploadString(path *dto.OssStoragePathDTO, oType *objectTypeItem, object string) (objectReturn interface{}, err error) {

	ossName, hasName := defaultOssNameMap[oType.ObjectType]
	if !hasName {
		return nil, new(httperror.XmoError).WithBiz(httperror.BIZ_ARG_ERROR)
	}

	uploadObject := &uploadObject{
		cover:       false,
		ossType:     oType.ObjectType,
		ossContent:  object,
		ossReader:   strings.NewReader(object),
		ossName:     ossName,
		format:      GetFileSuffix(ossName),
		contentType: oType.ContentType,
		size:        int64(len(object)),
	}

	// 调度器 分配 对象工厂
	uploadFactory, err := uploadObjectDispatcher(oType)
	if err != nil {
		return nil, err
	}

	// 工厂上传
	objectReturn, err = uploadFactory.getOssObject(path, oType, uploadObject)
	if err != nil {
		return nil, err
	}

	fmt.Println(uploadObject)
	fmt.Println(objectReturn)

	return objectReturn, err
}
