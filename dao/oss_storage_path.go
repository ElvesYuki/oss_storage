package dao

import (
	"database/sql"
	"go.uber.org/zap"
	"oss_storage/entity/model"
	"oss_storage/setting/mysql"
)

func ListOssStoragePath() (data []*model.OssStoragePath, err error) {

	sqlStr := `select id, path_code, object_type, bucket_name, object_path, max_size, object_suffix, enable, sort_num, status, description from oss_storage_path`

	err = mysql.DB.Select(&data, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no OssStoragePath")
		err = nil
	}

	return
}

func GetOssStoragePathByCode(code string) (data *model.OssStoragePath, err error) {

	data = new(model.OssStoragePath)

	sqlStr := `select id, path_code, object_type, bucket_name, object_path, max_size, object_suffix, enable, sort_num, status, description from oss_storage_path where path_code = ? limit 1`

	err = mysql.DB.Get(data, sqlStr, code)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no OssStoragePath")
		err = nil
	}
	return
}
