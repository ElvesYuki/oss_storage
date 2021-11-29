package dbgennerator

type TableCol struct {
	TableSchema            string `db:"table_schema"`
	TableName              string `db:"table_name"`
	ColumnName             string `db:"column_name"`
	OrdinalPosition        int    `db:"ordinal_position"`
	DataType               string `db:"data_type"`
	CharacterMaximumLength int    `db:"character_maximum_length"`
	ColumnComment          string `db:"column_comment"`
}

func ListTableColByTsANDTn(tableSchema string, tableName string) []*TableCol {
	// oss_storage
	sqlStr := `select * from information_schema.COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ?`

}
