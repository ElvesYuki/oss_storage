package oss

import (
	"go.uber.org/zap"
	"oss_storage/dao"
	"oss_storage/entity/dto"
	"strconv"
)

var UO iUploadOperator

var defaultUploadFactory iUploadFactory

// 对象工厂Map
var uploadFactoryMap map[string]iUploadFactory

// 对象默认名称Map
var defaultOssNameMap map[string]string

func Init() {
	ossStoragePath, err := dao.ListOssStoragePath()
	if err != nil {
		zap.L().Error("初始化ossStoragePathMap失败", zap.Error(err))
	}

	ossStoragePathDTOMap = make(map[string]*dto.OssStoragePathDTO)
	for _, path := range ossStoragePath {
		ossStoragePathDTOMap[path.PathCode.String] = dto.ToOssStoragePathDTO(path)
	}

	zap.L().Info("ossStoragePathDTOMap的大小为" + strconv.Itoa(len(ossStoragePathDTOMap)))

	// 初始化ObjectTypeMap

	objectTypeMap = make(map[string]*objectTypeItem)

	objectTypeMap[objectTypeEnum.OBJECT_TYPE_DEFAULT.ObjectType] = objectTypeEnum.OBJECT_TYPE_DEFAULT
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_IMAGE.ObjectType] = objectTypeEnum.OBJECT_TYPE_IMAGE
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_VIDEO.ObjectType] = objectTypeEnum.OBJECT_TYPE_VIDEO
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_AUDIO.ObjectType] = objectTypeEnum.OBJECT_TYPE_AUDIO
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_JSON.ObjectType] = objectTypeEnum.OBJECT_TYPE_JSON
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_HTML.ObjectType] = objectTypeEnum.OBJECT_TYPE_HTML

	zap.L().Info("objectTypeMap的大小为" + strconv.Itoa(len(objectTypeMap)))

	// 初始化 UploadOperator
	UO = new(uploadOperator)

	// 初始化 UploadFactory
	defaultUploadFactory = new(uploadFactory)

	// 初始化 UploadFactory
	uploadFactoryMap = make(map[string]iUploadFactory)
	uploadFactoryMap[objectTypeEnum.OBJECT_TYPE_DEFAULT.ObjectType] = new(uploadFactory)
	uploadFactoryMap[objectTypeEnum.OBJECT_TYPE_JSON.ObjectType] = new(jsonFactory)
	uploadFactoryMap[objectTypeEnum.OBJECT_TYPE_HTML.ObjectType] = new(htmlFactory)

	// 初始化 defaultOssNameMap
	defaultOssNameMap = make(map[string]string)
	defaultOssNameMap[objectTypeEnum.OBJECT_TYPE_JSON.ObjectType] = "default.json"
	defaultOssNameMap[objectTypeEnum.OBJECT_TYPE_HTML.ObjectType] = "default.html"

}
