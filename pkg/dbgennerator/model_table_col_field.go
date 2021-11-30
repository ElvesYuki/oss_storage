package dbgennerator

type TableFile struct {
	PackageName               string                    `json:"packageName"`               // 包名字
	TypeName                  string                    `json:"typeName"`                  // 结构体类型名字
	TableColArr               []*TableCol               `json:"tableColArr"`               // 列属性原数组 按属性未知排序
	TableColFieldArr          []*TableColField          `json:"tableColFieldArr"`          // 列属性数组 按属性未知排序
	TableColFieldMap          map[string]*TableColField `json:"tableColFieldMap"`          // 列属性Map [属性名]属性结构体
	TableColFieldWithColonArr []*TableColField          `json:"tableColFieldWithColonArr"` // 列属性数组 属性名称带冒号 按属性未知排序
	TableColFieldNotAlignArr  []*TableColField          `json:"TableColFieldNotAlignArr"`  // 列属性数组  不填充空格 按属性未知排序
}

type TableColField struct {
	OrdinalPosition  int16    `json:"ordinalPosition"`  // 属性位置序号
	FiledName        string   `json:"filedName"`        // 属性名
	FiledType        *SqlType `json:"filedType"`        // 类型名
	FiledTypeFinally string   `json:"filedTypeFinally"` // 类型名，最终
	FiledTag         string   `json:"filedTag"`         // 属性标签
	ColumnComment    string   `json:"columnComment"`    // 属性备注
}

type TemplateModel struct {
	TableFileDB  *TableFile
	TableFileDTO *TableFile
}

type SqlType struct {
	GoType             string
	GoTypeUcFirst      string
	SqlType            string
	SetSqlTypeNullFunc string
}
