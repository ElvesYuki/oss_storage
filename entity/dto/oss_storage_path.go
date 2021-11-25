package dto

import (
	"encoding/json"
	"go.uber.org/zap"
	"oss_storage/entity/model"
)

type OssStoragePathDTO struct {
	OssStoragePathId int64             `json:"ossStoragePathId"` // id
	PathCode         string            `json:"pathCode"`         // 存储的路径 码，枚举类中 全局唯一
	ObjectType       string            `json:"objectType"`       // 上传的文件类型， 默认为default， 自行判断， 可手动填入， 如json、html
	BucketName       string            `json:"bucketName"`       // bucket name
	ObjectPath       string            `json:"objectPath"`       // 存储的对象路径，不包含文件名
	MaxSize          int64             `json:"maxSize"`          // 上传的最大大小 字节数 -1L代表不限制
	ObjectSuffix     map[string]string `json:"objectSuffix"`     // 允许的文件名后缀, 空代表允许所有,数组转成的字符串
	Enable           int               `json:"enable"`           // 是否启用
	SortNum          int64             `json:"sortNum"`          // 排序字段
	Status           int               `json:"status"`           // 状态字段
	Description      string            `json:"description"`      // 描述
}

func ToOssStoragePathDTO(src *model.OssStoragePath) *OssStoragePathDTO {

	var objectSuffixMap = make(map[string]string)

	// 将后缀 数组字符串 反序列化成 数组
	if src.ObjectSuffix.Valid {
		var objectSuffixArray []string
		err := json.Unmarshal([]byte(src.ObjectSuffix.String), &objectSuffixArray)
		if err != nil {
			zap.L().Error("Json序列化失败", zap.Error(err))
			return nil
		}
		for _, suffix := range objectSuffixArray {
			objectSuffixMap[suffix] = suffix
		}
	}

	return &OssStoragePathDTO{
		OssStoragePathId: src.Id.Int64,
		PathCode:         src.PathCode.String,
		ObjectType:       src.ObjectType.String,
		BucketName:       src.BucketName.String,
		ObjectPath:       src.ObjectPath.String,
		MaxSize:          src.MaxSize.Int64,
		ObjectSuffix:     objectSuffixMap,
		Enable:           int(src.Enable.Int16),
		SortNum:          src.SortNum.Int64,
		Status:           int(src.Status.Int16),
		Description:      src.Description.String,
	}

}
