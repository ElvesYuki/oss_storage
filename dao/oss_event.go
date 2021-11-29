package dao

import (
	"database/sql"
	"go.uber.org/zap"
	"oss_storage/entity/dto"
	"oss_storage/entity/model"
	"oss_storage/pkg/idgenerator"
	"oss_storage/setting/mysql"
	"time"
)

func ListOssEvent() (data *model.OssEvent, err error) {

	sqlStr := `select id, oss_url, bucket_name, object_name, content_type, size, md5, version_id, gmt_create from oss_event`

	err = mysql.DB.Select(data, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no OssStoragePath")
		err = nil
	}
	return
}

func InsertOssEvent(row *model.OssEvent) (err error) {

	id, err := idgenerator.GetIdByModule(idgenerator.MODULE_OSS_EVENT)
	if err != nil {
		return err
	}
	row.Id = dto.SetNullInt64(id)

	row.GmtCreate = dto.SetNullInt64(time.Now().UnixMilli())

	sqlStr := `insert into oss_event 
    (id, oss_url, bucket_name, object_name, content_type, size, md5, version_id, gmt_create) 
    values 
    (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = mysql.DB.Exec(sqlStr, row.Id, row.OssUrl, row.BucketName, row.ObjectName, row.ContentType, row.Size, row.Md5, row.VersionId, row.GmtCreate)
	if err != nil {
		zap.L().Error("OssEvent新增错误", zap.Error(err))
		return err
	}
	return nil

}
