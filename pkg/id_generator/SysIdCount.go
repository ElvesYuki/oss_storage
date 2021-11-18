package id_generator

import (
	"database/sql"
	"go.uber.org/zap"
	"oss_storage/setting/mysql"
)

type SysIdCount struct {
	Id      int64  `json:"id"`
	Module  string `json:"module"`
	Step    int64  `json:"step"`
	Counter int64  `json:"counter"`
}

func ListSysIdCount() (data []*SysIdCount, err error) {

	sqlStr := `select id, module, step, counter from sys_id_count`

	err = mysql.DB.Select(&data, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no SysIdCount")
		err = nil
	}
	return
}

func GetSysIdCountById(id int64) (data *SysIdCount, err error) {

	data = new(SysIdCount)

	sqlStr := `select id, module, step, counter from sys_id_count where id = ? limit 1`

	err = mysql.DB.Get(data, sqlStr, id)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no SysIdCount")
		err = nil
	}
	return
}

func UpdateCounterSysIdCountById(id int64, counter int64) (err error) {

	sqlStr := `update sys_id_count set counter = ? where id = ? limit 1`

	_, err = mysql.DB.Exec(sqlStr, counter, id)
	if err != nil {
		return err
	}
	return
}
