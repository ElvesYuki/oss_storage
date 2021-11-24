package oss

import "github.com/minio/minio-go/v7"

var OC IClient

type IClient interface {
	GetClient() *Client
}

type Client struct {
	MinioClient *minio.Client
}

type MinioClient struct {
	MinioClient *Client
}

func (oc MinioClient) GetClient() *Client {
	return oc.MinioClient
}
