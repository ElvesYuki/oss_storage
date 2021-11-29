package dbgennerator

type templateModel struct {
	PackageName string           `json:"packageName"`
	ModelName   string           `json:"modelName"`
	TableCol    []*tableColModel `json:"TableColModel"`
}

type tableColModel struct {
	OrdinalPosition int16  `json:"ordinalPosition"` // 属性位置序号
	FiledName       string `json:"filedName"`       // 属性名
	FiledType       string `json:"filedType"`       // 类型名
	FiledTag        string `json:"filedTag"`        // 属性标签
}

var filedTypeMap map[string]string

func filedTypeMapInit() {
	filedTypeMap = make(map[string]string)
	filedTypeMap["bigint"] = "sql.NullInt64"
	filedTypeMap["varchar"] = "sql.NullString"
}
