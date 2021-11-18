package service

import (
	"oss_storage/pkg/id_generator"
)

func Test() string {
	return "test"
}

func ListIdGenerate() (data []*id_generator.SysIdCount, err error) {

	data, err = id_generator.ListSysIdCount()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetIdGenerateById(id int64) (data *id_generator.SysIdCount, err error) {

	data, err = id_generator.GetSysIdCountById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetId(module string) (id int64) {
	return id_generator.GetIdByModule(module)
}
