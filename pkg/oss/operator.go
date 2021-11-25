package oss

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

type IUploadOperator interface {
	putObject(object *UploadObject)
}

type UploadOperator struct{}

func (uo *UploadOperator) putObject(object *UploadObject) {
	uploadInfo, err := OC.GetClient().MinioClient.PutObject(context.Background(),
		object.bucketName, object.objectName,
		object.ossReader, -1,
		minio.PutObjectOptions{
			ContentType: object.contentType,
		})
	if err != nil {
		return
	}
	fmt.Println(uploadInfo)
}
