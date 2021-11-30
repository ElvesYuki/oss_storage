package model

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"oss_storage/pkg/dbgennerator"
	"oss_storage/setting/mysql"
	"oss_storage/setting/redis"
	"oss_storage/test"
	"testing"
	"text/template"
)

func InitTest() {
	viper.AddConfigPath("../../../")
	test.InitTest()
}
func CloseTest() {
	defer mysql.Close()
	defer redis.Close()
}

func TestListTableColByTsANDTn(t *testing.T) {
	InitTest()
	tableColArray, err := dbgennerator.ListTableColByTsANDTn("oss_storage", "oss_event")
	if err != nil {
		fmt.Println(err)
	}

	for _, tableCol := range tableColArray {
		fmt.Println(*tableCol)
		s := dbgennerator.Case2Camel(tableCol.ColumnName.String)
		fmt.Println(dbgennerator.UcFirst(s))
	}
	CloseTest()
}

func TestGenerateDbModel(t *testing.T) {
	InitTest()
	dbgennerator.Init()
	GenerateDbModel("oss_event")
	CloseTest()
}

func TestGenerateTableModelFile(t *testing.T) {
	InitTest()
	dbgennerator.Init()
	tableFile, err := dbgennerator.GenerateTableMap("oss_storage", "oss_event", "model")
	if err != nil {
		fmt.Println(err)
	}

	fileName := "./" + "oss_event" + ".go"

	if _, err := os.Stat(fileName); err != nil {
		fmt.Println(err)
	}
	if err := os.Remove(fileName); err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	tmpl, err := template.ParseFiles("./template_model.tmpl")
	if err != nil {
		fmt.Println(err)
		fmt.Println("模板初始化错误")
	}

	write := bufio.NewWriter(file)

	tmpl.Execute(write, tableFile)
	write.Flush()

	CloseTest()
}
