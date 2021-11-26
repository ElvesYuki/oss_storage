package oss

import "oss_storage/common/httperror"

// uploadObjectDispatcher 上传处理器
func uploadObjectDispatcher(oType *objectTypeItem) (iUploadFactory, error) {

	uF, hasFactory := uploadFactoryMap[oType.ObjectType]
	if !hasFactory {
		// 返回默认上传对象工厂
		if defaultUF, hasFactory := uploadFactoryMap[objectTypeEnum.OBJECT_TYPE_DEFAULT.ObjectType]; hasFactory {
			return defaultUF, nil
		} else {
			return nil, new(httperror.XmoError).WithBiz(httperror.BIZ_OSS_UNKNOWN_TYPE_ERROR)
		}
	}

	return uF, nil
}
