package oss

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"oss_storage/setting"
)

var Client *minio.Client

func Init(cfg *setting.OssConfig) (err error) {

	option := &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	}

	Client, err = minio.New(cfg.Endpoint, option)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return nil
	}
	return nil
}
