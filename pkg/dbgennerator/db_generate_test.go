package dbgennerator

import (
	"fmt"
	"github.com/spf13/viper"
	"oss_storage/setting/mysql"
	"oss_storage/setting/redis"
	"oss_storage/test"
	"testing"
)

func InitTest() {
	viper.AddConfigPath("../../")
	test.InitTest()
}
func CloseTest() {
	defer mysql.Close()
	defer redis.Close()
}

func TestListTableColByTsANDTn(t *testing.T) {
	InitTest()
	tableColArray, err := ListTableColByTsANDTn("oss_storage", "oss_event")
	if err != nil {
		fmt.Println(err)
	}

	for _, tableCol := range tableColArray {
		fmt.Println(*tableCol)
		s := Case2Camel(tableCol.ColumnName.String)
		fmt.Println(UcFirst(s))
	}
	CloseTest()
}

func TestGenerateTableMap(t *testing.T) {
	InitTest()
	Init()
	GenerateTableMap("oss_storage", "oss_event", "model")
	CloseTest()
}
