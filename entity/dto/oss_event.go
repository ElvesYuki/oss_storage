package dto

import (
	"oss_storage/entity/model"
)

type OssEventDTO struct {
	OssEventId  int64  `json:"ossEventId"`
	OssUrl      string `json:"ossUrl"`
	BucketName  string `json:"bucketName"`
	ObjectName  string `json:"objectName"`
	ContentType string `json:"contentType"`
	Size        int64  `json:"size"`
	Md5         string `json:"md5"`
	VersionId   string `json:"versionId"`
	GmtCreate   int64  `json:"gmtCreate"`
}

func ToOssEventDTO(src *model.OssEvent) *OssEventDTO {
	return &OssEventDTO{
		OssEventId:  src.Id.Int64,
		OssUrl:      src.OssUrl.String,
		BucketName:  src.BucketName.String,
		ObjectName:  src.ObjectName.String,
		ContentType: src.ContentType.String,
		Size:        src.Size.Int64,
		Md5:         src.Md5.String,
		VersionId:   src.VersionId.String,
		GmtCreate:   src.GmtCreate.Int64,
	}
}

func ToOssEvent(src *OssEventDTO) *model.OssEvent {
	return &model.OssEvent{
		Id:          SetNullInt64(src.OssEventId),
		OssUrl:      SetNullString(src.OssUrl),
		BucketName:  SetNullString(src.BucketName),
		ObjectName:  SetNullString(src.ObjectName),
		ContentType: SetNullString(src.ContentType),
		Size:        SetNullInt64(src.Size),
		Md5:         SetNullString(src.Md5),
		VersionId:   SetNullString(src.VersionId),
		GmtCreate:   SetNullInt64(src.GmtCreate),
	}
}
