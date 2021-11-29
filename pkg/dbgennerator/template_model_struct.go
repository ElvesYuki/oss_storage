package dbgennerator

type templateModel struct {
	PackageName string          `json:"packageName"`
	ModelName   string          `json:"modelName"`
	TableCol    []tableColModel `json:"TableColModel"`
}

type tableColModel struct {
	FiledName string `json:"filedName"` // 属性名
	SqlType   string `json:"sqlType"`   // 类型名
	FiledTag  string `json:"filedTag"`  // 属性标签
}
