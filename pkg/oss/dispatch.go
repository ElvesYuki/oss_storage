package oss

import (
	"go.uber.org/zap"
	"oss_storage/dao"
	"oss_storage/entity/dto"
	"strconv"
)

var UO IUploadOperator

var UF IUploadFactory

var UH IUploadHandler

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

	objectTypeMap = make(map[string]*ObjectTypeItem)

	objectTypeMap[objectTypeEnum.OBJECT_TYPE_DEFAULT.ObjectType] = objectTypeEnum.OBJECT_TYPE_DEFAULT
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_IMAGE.ObjectType] = objectTypeEnum.OBJECT_TYPE_IMAGE
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_VIDEO.ObjectType] = objectTypeEnum.OBJECT_TYPE_VIDEO
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_AUDIO.ObjectType] = objectTypeEnum.OBJECT_TYPE_AUDIO
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_JSON.ObjectType] = objectTypeEnum.OBJECT_TYPE_JSON
	objectTypeMap[objectTypeEnum.OBJECT_TYPE_HTML.ObjectType] = objectTypeEnum.OBJECT_TYPE_HTML

	zap.L().Info("objectTypeMap的大小为" + strconv.Itoa(len(objectTypeMap)))

	// 初始化 UploadOperator
	UO = new(UploadOperator)

	// 初始化 UploadFactory
	UF = new(UploadFactory)

	// 初始化 UploadHandler
	UH = new(UploadHandler)
}
