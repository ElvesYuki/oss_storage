package oss

import (
	"fmt"
	"mime/multipart"
	"oss_storage/common"
	"oss_storage/common/httperror"
	"oss_storage/dao"
	"oss_storage/entity/dto"
	"oss_storage/entity/model"
)

func PackMultipartOssEventChan(uploadObject *uploadObject, object *multipart.FileHeader, logChan chan *dto.OssEventDTO) {

	srcReaderLog, err := object.Open()
	if err != nil {

	}
	defer srcReaderLog.Close()
	md5, err := common.Md5Util(srcReaderLog)
	if err != nil {

	}
	fmt.Println("Md5===>" + md5)

	ossEventDTO := &dto.OssEventDTO{
		Md5:         md5,
		ContentType: uploadObject.contentType,
		Size:        uploadObject.size,
	}

	logChan <- ossEventDTO
}

func PackStringOssEventChan(uploadObject *uploadObject, object string, logChan chan *dto.OssEventDTO) {

	md5, err := common.Md5StringUtil(object)
	if err != nil {

	}
	fmt.Println("Md5===>" + md5)

	ossEventDTO := &dto.OssEventDTO{
		Md5:         md5,
		ContentType: uploadObject.contentType,
		Size:        uploadObject.size,
	}

	logChan <- ossEventDTO
}

func LogRecordOssEvent(object interface{}, ossEventDTO *dto.OssEventDTO) (err error) {
	switch v := object.(type) {
	case *BaseObject:
		ossEventDTO.OssUrl = v.Url
		ossEventDTO.BucketName = v.Bucket
		ossEventDTO.ObjectName = v.Object
	case *HtmlObject:
		ossEventDTO.OssUrl = v.Url
		ossEventDTO.BucketName = v.Bucket
		ossEventDTO.ObjectName = v.Object
	case JsonObject:
		ossEventDTO.OssUrl = v.Url
		ossEventDTO.BucketName = v.Bucket
		ossEventDTO.ObjectName = v.Object
	default:
		return new(httperror.XmoError).WithBiz(httperror.BIZ_ARG_ERROR)
	}

	ossEvent := dto.ToOssEvent(ossEventDTO)

	err = AddOssEvent(ossEvent)
	if err != nil {
		return err
	}
	return nil
}

func AddOssEvent(event *model.OssEvent) error {
	err := dao.InsertOssEvent(event)
	if err != nil {
		return err
	}
	return nil
}
