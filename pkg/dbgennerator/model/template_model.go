package model

import "oss_storage/pkg/dbgennerator"

type templateModel struct {
	PackageName string           `json:"packageName"`
	ModelName   string           `json:"modelName"`
	TableCol    []*tableColModel `json:"TableColModel"`
}

type tableColModel struct {
	OrdinalPosition int16                 `json:"ordinalPosition"` // 属性位置序号
	FiledName       string                `json:"filedName"`       // 属性名
	FiledType       *dbgennerator.SqlType `json:"filedType"`       // 类型名
	FiledTag        string                `json:"filedTag"`        // 属性标签
	ColumnComment   string                `json:"columnComment"`   // 属性备注
}
