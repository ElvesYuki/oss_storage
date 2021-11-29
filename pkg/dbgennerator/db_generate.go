package dbgennerator

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
)

func GenerateDbModel(tableName string) {

	file, err := os.OpenFile("./oss_event_demo.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	tableColArray, err := ListTableColByTsANDTn("oss_storage", tableName)
	if err != nil {
		fmt.Println("数据查询错误")
	}

	tmpl, err := template.ParseFiles("./template_model.tmpl")
	if err != nil {
		fmt.Println(err)
		fmt.Println("模板初始化错误")
	}

	root := &templateModel{
		PackageName: "model",
		ModelName:   "OssEvent",
	}

	// 表列
	var tableColModelArray []*tableColModel

	for _, col := range tableColArray {

		tableColModelStruct := &tableColModel{
			OrdinalPosition: col.OrdinalPosition.Int16,
			FiledName:       UcFirst(Case2Camel(col.ColumnName.String)),
			FiledType:       filedTypeMap[col.DataType.String],
			FiledTag:        "`db:\"" + col.ColumnName.String + "\"`",
		}

		if col.ColumnComment.Valid {
			tableColModelStruct.ColumnComment = "// " + col.ColumnComment.String
		}

		tableColModelArray = append(tableColModelArray, tableColModelStruct)
	}

	// 对齐
	filedNameMaxLength := 1
	filedTypeMaxLength := 1
	filedTagMaxLength := 1

	for _, colModel := range tableColModelArray {
		if temp := len(colModel.FiledName); temp > filedNameMaxLength {
			filedNameMaxLength = temp
		}
		if temp := len(colModel.FiledType); temp > filedTypeMaxLength {
			filedTypeMaxLength = temp
		}
		if temp := len(colModel.FiledTag); temp > filedTagMaxLength {
			filedTagMaxLength = temp
		}
	}

	// 对齐
	for _, colModel := range tableColModelArray {
		tempFiledNameByte := []byte(colModel.FiledName)
		filedNameBytes := make([]byte, filedNameMaxLength, filedNameMaxLength)
		for i := 0; i < filedNameMaxLength; i++ {
			if i < len(tempFiledNameByte) {
				filedNameBytes[i] = tempFiledNameByte[i]
			} else {
				filedNameBytes[i] = ' '
			}
		}
		colModel.FiledName = string(filedNameBytes)

		tempFiledTypeByte := []byte(colModel.FiledType)
		filedTypeBytes := make([]byte, filedTypeMaxLength, filedTypeMaxLength)
		for i := 0; i < filedTypeMaxLength; i++ {
			if i < len(tempFiledTypeByte) {
				filedTypeBytes[i] = tempFiledTypeByte[i]
			} else {
				filedTypeBytes[i] = ' '
			}
		}
		colModel.FiledType = string(filedTypeBytes)

		tempFiledTagByte := []byte(colModel.FiledTag)
		filedTagBytes := make([]byte, filedTagMaxLength, filedTagMaxLength)
		for i := 0; i < filedTagMaxLength; i++ {
			if i < len(tempFiledTagByte) {
				filedTagBytes[i] = tempFiledTagByte[i]
			} else {
				filedTagBytes[i] = ' '
			}
		}
		colModel.FiledTag = string(filedTagBytes)
	}
	root.TableCol = tableColModelArray

	write := bufio.NewWriter(file)

	tmpl.Execute(write, root)
	write.Flush()

}
