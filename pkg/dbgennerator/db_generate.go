package dbgennerator

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

// FieldTypeMap 类型Map[类型]*SqlType 一一对应
var FieldTypeMap map[string]*SqlType

// FieldSqlTypeMap sql类型map[sql类型]*SqlType 一一对应
var FieldSqlTypeMap map[string]*SqlType

// TableFileMap 文件名map[表名]文件结构体
var TableFileMap map[string]*TableFile

// TableFileDTOMap 文件名DTOmap[表名]文件结构体
var TableFileDTOMap map[string]*TableFile

func Init() {
	int16Type := &SqlType{
		GoType:             "int16",
		GoTypeUcFirst:      "Int16",
		SqlType:            "sql.NullInt16",
		SetSqlTypeNullFunc: "SetNullInt16",
	}

	int64Type := &SqlType{
		GoType:             "int64",
		GoTypeUcFirst:      "Int64",
		SqlType:            "sql.NullInt64",
		SetSqlTypeNullFunc: "SetNullInt64",
	}

	stringType := &SqlType{
		GoType:             "string",
		GoTypeUcFirst:      "String",
		SqlType:            "sql.NullString",
		SetSqlTypeNullFunc: "SetNullString",
	}

	FieldTypeMap = make(map[string]*SqlType)
	FieldTypeMap["int"] = int16Type
	FieldTypeMap["bigint"] = int64Type
	FieldTypeMap["varchar"] = stringType

	FieldSqlTypeMap = make(map[string]*SqlType)
	FieldSqlTypeMap["sql.NullInt16"] = int16Type
	FieldSqlTypeMap["sql.NullInt64"] = int64Type
	FieldSqlTypeMap["sql.NullString"] = stringType

	TableFileMap = make(map[string]*TableFile)

	TableFileDTOMap = make(map[string]*TableFile)
}

// GenerateTableMap 获取 数据库结构体模板对象 并放入map
func GenerateTableMap(tableSchema string, tableName string, packageName string) (tableFile *TableFile, err error) {

	tableFile, hasTableName := TableFileMap[tableName]
	if hasTableName {
		return tableFile, nil
	}

	// 查询对应表需要的所有列属性
	tableColArr, err := ListTableColByTsANDTn(tableSchema, tableName)
	if err != nil {
		fmt.Println("数据查询错误")
		return nil, err
	}

	// 构造 数据库结构体模板对象 初步
	tableFile = &TableFile{
		PackageName:      packageName,
		TypeName:         UcFirst(Case2Camel(tableName)),
		TableColFieldMap: make(map[string]*TableColField),
	}

	// 将原始数据属性放入
	tableFile.TableColArr = tableColArr

	// 循环原始数据属性
	for _, col := range tableColArr {
		// 构造 单个数据属性 结构体对象
		tableColField := &TableColField{
			OrdinalPosition:  col.OrdinalPosition.Int16,
			FiledName:        UcFirst(Case2Camel(col.ColumnName.String)),
			FiledType:        FieldTypeMap[col.DataType.String],
			FiledTypeFinally: FieldTypeMap[col.DataType.String].SqlType,
			FiledTag:         "`db:\"" + col.ColumnName.String + "\"`",
		}
		// 检查common属性是否为空
		if col.ColumnComment.Valid {
			tableColField.ColumnComment = "// " + col.ColumnComment.String
		} else {
			tableColField.ColumnComment = "// "
		}
		// 放进数组
		tableFile.TableColFieldArr = append(tableFile.TableColFieldArr, tableColField)
		// 放进Map
		tableFile.TableColFieldMap[tableColField.FiledName] = tableColField
	}

	// 循环原始数据属性 构造带: 的结构体
	for _, col := range tableColArr {
		tableColField := &TableColField{
			OrdinalPosition:  col.OrdinalPosition.Int16,
			FiledName:        UcFirst(Case2Camel(col.ColumnName.String)),
			FiledType:        FieldTypeMap[col.DataType.String],
			FiledTypeFinally: FieldTypeMap[col.DataType.String].SqlType,
			FiledTag:         "`db:\"" + col.ColumnName.String + "\"`",
		}
		if col.ColumnComment.Valid {
			tableColField.ColumnComment = "// " + col.ColumnComment.String
		}
		tableFile.TableColFieldWithColonArr = append(tableFile.TableColFieldWithColonArr, tableColField)
	}

	// 循环原始数据属性 构造不对齐 的结构体
	for _, col := range tableColArr {
		tableColField := &TableColField{
			OrdinalPosition:  col.OrdinalPosition.Int16,
			FiledName:        UcFirst(Case2Camel(col.ColumnName.String)),
			FiledType:        FieldTypeMap[col.DataType.String],
			FiledTypeFinally: FieldTypeMap[col.DataType.String].SqlType,
			FiledTag:         "`db:\"" + col.ColumnName.String + "\"`",
		}
		if col.ColumnComment.Valid {
			tableColField.ColumnComment = "// " + col.ColumnComment.String
		}
		tableFile.TableColFieldNotAlignArr = append(tableFile.TableColFieldNotAlignArr, tableColField)
	}

	// 对齐
	tableFile.TableColFieldArr = AlignTableColField(tableFile.TableColFieldArr)
	// 对齐
	tableFile.TableColFieldWithColonArr = AlignTableColFieldWithColon(tableFile.TableColFieldWithColonArr)

	//for _, colField := range tableFile.TableColFieldArr {
	//	fmt.Println(colField)
	//}
	TableFileMap[tableName] = tableFile
	return tableFile, nil
}

// GenerateTableDTOMap 获取 DTO结构体模板对象 并放入map
func GenerateTableDTOMap(tableSchema string, tableName string, packageName string) (tableFileDTO *TableFile, err error) {

	tableFileDTO, hasTableName := TableFileDTOMap[tableName]
	if hasTableName {
		return tableFileDTO, nil
	}

	tableFile, err := GenerateTableMap(tableSchema, tableName, "model")
	if err != nil {
		return nil, err
	}

	// 深拷贝
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(tableFile); err != nil {
		return nil, err
	}
	tableFileDTO = &TableFile{}
	if err = gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(tableFileDTO); err != nil {
		return nil, err
	}

	// 新的结构体对象
	// 替换
	tableFileDTO.PackageName = packageName
	tableFileDTO.TypeName = tableFile.TypeName + "DTO"

	tableFileDTO.TableColFieldMap = make(map[string]*TableColField)

	// 循环属性数组，替换
	for i, colFieldDTO := range tableFileDTO.TableColFieldArr {
		// 剔除从一 中添加的空格
		colFieldDTO.FiledName = strings.Replace(colFieldDTO.FiledName, " ", "", -1)
		colFieldDTO.FiledTypeFinally = strings.Replace(colFieldDTO.FiledTypeFinally, " ", "", -1)
		colFieldDTO.FiledTag = strings.Replace(colFieldDTO.FiledTag, " ", "", -1)
		// 拼接 结构体名称+Id
		if i == 0 {
			colFieldDTO.FiledName = tableFileDTO.TypeName + colFieldDTO.FiledName
			colFieldDTO.FiledName = strings.Replace(colFieldDTO.FiledName, "DTO", "", -1)
		}
		// 更新类型 sql类型转go类型
		colFieldDTO.FiledType = FieldSqlTypeMap[colFieldDTO.FiledType.SqlType]
		colFieldDTO.FiledTypeFinally = FieldSqlTypeMap[colFieldDTO.FiledType.SqlType].GoType
		colFieldDTO.FiledTag = "`json:\"" + LcFirst(colFieldDTO.FiledName) + "\"`"
		tableFileDTO.TableColFieldMap[colFieldDTO.FiledName] = colFieldDTO
	}

	// 循环属性数组 带：，替换
	for i, colFieldDTO := range tableFileDTO.TableColFieldWithColonArr {
		colFieldDTO.FiledName = strings.Replace(colFieldDTO.FiledName, " ", "", -1)
		colFieldDTO.FiledTypeFinally = strings.Replace(colFieldDTO.FiledTypeFinally, " ", "", -1)
		colFieldDTO.FiledTag = strings.Replace(colFieldDTO.FiledTag, " ", "", -1)
		if i == 0 {
			colFieldDTO.FiledName = tableFileDTO.TypeName + colFieldDTO.FiledName
			colFieldDTO.FiledName = strings.Replace(colFieldDTO.FiledName, "DTO", "", -1)
		}
		colFieldDTO.FiledType = FieldSqlTypeMap[colFieldDTO.FiledType.SqlType]
		colFieldDTO.FiledTypeFinally = FieldSqlTypeMap[colFieldDTO.FiledType.SqlType].GoType
		colFieldDTO.FiledTag = "`json:\"" + LcFirst(colFieldDTO.FiledName) + "\"`"
	}

	// 循环属性数组 带：，替换 不对其
	for i, colFieldDTO := range tableFileDTO.TableColFieldNotAlignArr {
		colFieldDTO.FiledName = strings.Replace(colFieldDTO.FiledName, " ", "", -1)
		colFieldDTO.FiledTypeFinally = strings.Replace(colFieldDTO.FiledTypeFinally, " ", "", -1)
		colFieldDTO.FiledTag = strings.Replace(colFieldDTO.FiledTag, " ", "", -1)
		if i == 0 {
			colFieldDTO.FiledName = tableFileDTO.TypeName + colFieldDTO.FiledName
			colFieldDTO.FiledName = strings.Replace(colFieldDTO.FiledName, "DTO", "", -1)
		}
		colFieldDTO.FiledType = FieldSqlTypeMap[colFieldDTO.FiledType.SqlType]
		colFieldDTO.FiledTypeFinally = FieldSqlTypeMap[colFieldDTO.FiledType.SqlType].GoType
		colFieldDTO.FiledTag = "`json:\"" + LcFirst(colFieldDTO.FiledName) + "\"`"
	}

	tableFileDTO.TableColFieldArr = AlignTableColField(tableFileDTO.TableColFieldArr)
	tableFileDTO.TableColFieldWithColonArr = AlignTableColFieldWithColon(tableFileDTO.TableColFieldWithColonArr)

	for _, colFieldDTO := range tableFileDTO.TableColFieldArr {
		fmt.Println(colFieldDTO)
	}
	TableFileDTOMap[tableName] = tableFileDTO
	return tableFileDTO, nil
}

func AlignTableColField(colArr []*TableColField) []*TableColField {
	// 对齐
	filedNameMaxLength := 1
	filedTypeMaxLength := 1
	filedTagMaxLength := 1

	for _, colField := range colArr {

		colField.FiledName = strings.Replace(colField.FiledName, " ", "", -1)
		if temp := len(colField.FiledName); temp > filedNameMaxLength {
			filedNameMaxLength = temp
		}
		colField.FiledTypeFinally = strings.Replace(colField.FiledTypeFinally, " ", "", -1)
		if temp := len(colField.FiledTypeFinally); temp > filedTypeMaxLength {
			filedTypeMaxLength = temp
		}
		colField.FiledTag = strings.Replace(colField.FiledTag, " ", "", -1)
		if temp := len(colField.FiledTag); temp > filedTagMaxLength {
			filedTagMaxLength = temp
		}
	}

	// 对齐
	for _, colField := range colArr {
		tempFiledNameByte := []byte(colField.FiledName)
		filedNameBytes := make([]byte, filedNameMaxLength, filedNameMaxLength)
		for i := 0; i < filedNameMaxLength; i++ {
			if i < len(tempFiledNameByte) {
				filedNameBytes[i] = tempFiledNameByte[i]
			} else {
				filedNameBytes[i] = ' '
			}
		}
		colField.FiledName = string(filedNameBytes)

		tempFiledTypeByte := []byte(colField.FiledTypeFinally)
		filedTypeBytes := make([]byte, filedTypeMaxLength, filedTypeMaxLength)
		for i := 0; i < filedTypeMaxLength; i++ {
			if i < len(tempFiledTypeByte) {
				filedTypeBytes[i] = tempFiledTypeByte[i]
			} else {
				filedTypeBytes[i] = ' '
			}
		}
		colField.FiledTypeFinally = string(filedTypeBytes)

		tempFiledTagByte := []byte(colField.FiledTag)
		filedTagBytes := make([]byte, filedTagMaxLength, filedTagMaxLength)
		for i := 0; i < filedTagMaxLength; i++ {
			if i < len(tempFiledTagByte) {
				filedTagBytes[i] = tempFiledTagByte[i]
			} else {
				filedTagBytes[i] = ' '
			}
		}
		colField.FiledTag = string(filedTagBytes)
	}
	return colArr
}

func AlignTableColFieldWithColon(colArr []*TableColField) []*TableColField {
	// 对齐
	filedNameMaxLength := 1
	filedTypeMaxLength := 1
	filedTagMaxLength := 1

	for _, colField := range colArr {

		colField.FiledName = strings.Replace(colField.FiledName, ":", "", -1)
		colField.FiledName = strings.Replace(colField.FiledName, " ", "", -1)
		colField.FiledName = colField.FiledName + ":" + " "
		if temp := len(colField.FiledName); temp > filedNameMaxLength {
			filedNameMaxLength = temp
		}
		colField.FiledTypeFinally = strings.Replace(colField.FiledTypeFinally, " ", "", -1)
		if temp := len(colField.FiledTypeFinally); temp > filedTypeMaxLength {
			filedTypeMaxLength = temp
		}
		colField.FiledTag = strings.Replace(colField.FiledTag, " ", "", -1)
		if temp := len(colField.FiledTag); temp > filedTagMaxLength {
			filedTagMaxLength = temp
		}
	}

	// 对齐
	for _, colField := range colArr {
		tempFiledNameByte := []byte(colField.FiledName)
		filedNameBytes := make([]byte, filedNameMaxLength, filedNameMaxLength)
		for i := 0; i < filedNameMaxLength; i++ {
			if i < len(tempFiledNameByte) {
				filedNameBytes[i] = tempFiledNameByte[i]
			} else {
				filedNameBytes[i] = ' '
			}
		}
		colField.FiledName = string(filedNameBytes)

		tempFiledTypeByte := []byte(colField.FiledTypeFinally)
		filedTypeBytes := make([]byte, filedTypeMaxLength, filedTypeMaxLength)
		for i := 0; i < filedTypeMaxLength; i++ {
			if i < len(tempFiledTypeByte) {
				filedTypeBytes[i] = tempFiledTypeByte[i]
			} else {
				filedTypeBytes[i] = ' '
			}
		}
		colField.FiledTypeFinally = string(filedTypeBytes)

		tempFiledTagByte := []byte(colField.FiledTag)
		filedTagBytes := make([]byte, filedTagMaxLength, filedTagMaxLength)
		for i := 0; i < filedTagMaxLength; i++ {
			if i < len(tempFiledTagByte) {
				filedTagBytes[i] = tempFiledTagByte[i]
			} else {
				filedTagBytes[i] = ' '
			}
		}
		colField.FiledTag = string(filedTagBytes)
	}
	return colArr
}
