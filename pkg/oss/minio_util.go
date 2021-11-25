package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
)

//var name = OC.GetClient().MinioClient

// CheckObjectExist  检查minio对象是否存在
func CheckObjectExist(bucketName string, objectName string) bool {

	_, err := OC.GetClient().MinioClient.StatObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return false
	}
	return true

}
