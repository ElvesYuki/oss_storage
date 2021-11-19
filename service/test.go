package service

import (
	"oss_storage/pkg/idgenerator"
)

func Test() string {
	return "test"
}

func ListIdGenerate() (data []*idgenerator.SysIdCount, err error) {

	data, err = idgenerator.ListSysIdCount()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetIdGenerateById(id int64) (data *idgenerator.SysIdCount, err error) {

	data, err = idgenerator.GetSysIdCountById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetId(module string) (id int64) {
	return idgenerator.GetIdByModule(module)
}
