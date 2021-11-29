package dbgennerator

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateDbModel(tableName string) {

	tableColArray, err := ListTableColByTsANDTn("oss_storage", tableName)
	if err != nil {
		fmt.Println("数据查询错误")
	}

	tmpl, err := template.ParseFiles("./template_model.tmpl")
	if err != nil {
		fmt.Println(err)
		fmt.Println("模板初始化错误")
	}

	root := templateModel{
		PackageName: "model",
		ModelName:   "OssEvent",
	}

	var tableColModelArray []tableColModel

	for _, col := range tableColArray {
		tableColModelArray = append(tableColModelArray, tableColModel{
			FiledName: Ucfirst(Case2Camel(col.ColumnName.String)),
		})
	}
	root.TableCol = tableColModelArray

	tmpl.Execute(os.Stdout, root)

}
