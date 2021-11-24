package oss

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"oss_storage/pkg/oss"
	"oss_storage/setting"
)

func Init(cfg *setting.OssConfig) (err error) {

	option := &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	}

	client, err := minio.New(cfg.Endpoint, option)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return nil
	}

	minioClient := &oss.Client{
		MinioClient: client,
	}

	oss.OC = &oss.MinioClient{MinioClient: minioClient}

	return nil
}
