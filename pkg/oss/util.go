package oss

import (
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"path"
)

// GenerateUUID 生成UUID
func GenerateUUID() string {
	u4 := uuid.NewV4()
	return hex.EncodeToString(u4[:])
}

// GetFileSuffix 获取对象后缀名，对象格式,不带 点 .
func GetFileSuffix(fileName string) string {
	fileSuffix := path.Ext(fileName)
	return fileSuffix[1:]
}

// GetFileSuffixWithDot 获取对象后缀名，对象格式,带 点 .
func GetFileSuffixWithDot(fileName string) string {
	fileSuffix := path.Ext(fileName)
	return fileSuffix
}

// GenerateFileUrl 生成对象存储路径
func GenerateFileUrl(filePath string, fileName string) string {
	return filePath + GenerateUUID() + GetFileSuffixWithDot(fileName)
}
