package model

import "database/sql"

type OssEvent struct {
	Id          sql.NullInt64  `db:"id"`
	OssUrl      sql.NullString `db:"oss_url"`
	BucketName  sql.NullString `db:"bucket_name"`
	ObjectName  sql.NullString `db:"object_name"`
	ContentType sql.NullString `db:"content_type"`
	Size        sql.NullInt64  `db:"size"`
	Md5         sql.NullString `db:"md5"`
	VersionId   sql.NullString `db:"version_id"`
	GmtCreate   sql.NullInt64  `db:"gmt_create"`
}
