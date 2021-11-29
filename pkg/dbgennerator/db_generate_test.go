package dbgennerator

import (
	"fmt"
	"github.com/spf13/viper"
	"oss_storage/setting/mysql"
	"oss_storage/setting/redis"
	"oss_storage/test"
	"testing"
)

func Init() {
	viper.AddConfigPath("../../")
	test.InitTest()
}
func Close() {
	defer mysql.Close()
	defer redis.Close()
}

func TestListTableColByTsANDTn(t *testing.T) {
	Init()
	tableColArray, err := ListTableColByTsANDTn("oss_storage", "oss_event")
	if err != nil {
		fmt.Println(err)
	}

	for _, tableCol := range tableColArray {
		fmt.Println(*tableCol)
		s := Case2Camel(tableCol.ColumnName.String)
		fmt.Println(UcFirst(s))
	}
	Close()
}

func TestGenerateDbModel(t *testing.T) {
	Init()
	filedTypeMapInit()
	GenerateDbModel("oss_event")
	Close()
}
