package dto

import "database/sql"

func SetNullInt16(v int16) sql.NullInt16 {
	if v == 0 {
		return sql.NullInt16{Int16: 0, Valid: false}
	}
	return sql.NullInt16{Int16: v, Valid: true}
}

func SetNullInt64(v int64) sql.NullInt64 {
	if v == 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: v, Valid: true}
}

func SetNullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: v, Valid: true}
}
