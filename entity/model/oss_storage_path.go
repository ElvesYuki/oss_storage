package model

import "database/sql"

type OssStoragePath struct {
	Id           sql.NullInt64  `db:"id"`            // id
	PathCode     sql.NullString `db:"path_code"`     // 存储的路径 码，枚举类中 全局唯一
	ObjectType   sql.NullString `db:"object_type"`   // 上传的文件类型， 默认为default， 自行判断， 可手动填入， 如json、html
	BucketName   sql.NullString `db:"bucket_name"`   // bucket name
	ObjectPath   sql.NullString `db:"object_path"`   // 存储的对象路径，不包含文件名
	MaxSize      sql.NullInt64  `db:"max_size"`      // 上传的最大大小 字节数 -1L代表不限制
	ObjectSuffix sql.NullString `db:"object_suffix"` // 允许的文件名后缀, 空代表允许所有,数组转成的字符串
	Enable       sql.NullInt16  `db:"enable"`        // 是否启用
	SortNum      sql.NullInt64  `db:"sort_num"`      // 排序字段
	Status       sql.NullInt16  `db:"status"`        // 状态字段
	Description  sql.NullString `db:"description"`   // 描述
}
