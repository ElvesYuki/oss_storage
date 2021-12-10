package oss

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type iUploadOperator interface {
	putObject(object *uploadObject)
}

type uploadOperator struct{}

func (uo *uploadOperator) putObject(object *uploadObject) {
	uploadInfo, err := OC.GetClient().MinioClient.PutObject(context.Background(),
		object.bucketName, object.objectName,
		object.ossReader, object.size,
		minio.PutObjectOptions{
			ContentType: object.contentType,
		})
	if err != nil {
		zap.L().Error("对象上传出错", zap.Error(err))
	}
	fmt.Println(uploadInfo)
}
