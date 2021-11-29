package dbgennerator

import (
	"database/sql"
	"oss_storage/setting/mysql"
)

type TableCol struct {
	TableSchema            sql.NullString `db:"table_schema"`
	TableName              sql.NullString `db:"table_name"`
	ColumnName             sql.NullString `db:"column_name"`
	OrdinalPosition        sql.NullInt16  `db:"ordinal_position"`
	DataType               sql.NullString `db:"data_type"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	ColumnComment          sql.NullString `db:"column_comment"`
}

func ListTableColByTsANDTn(tableSchema string, tableName string) (data []*TableCol, err error) {
	// oss_storage
	sqlStr := `select table_schema, table_name, column_name, ordinal_position, data_type, character_maximum_length, column_comment
       from information_schema.COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ? order by ordinal_position asc`

	err = mysql.DB.Select(&data, sqlStr, tableSchema, tableName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
