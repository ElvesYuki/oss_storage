package sensitiveword

import (
	"go.uber.org/zap"
	"math"
	"oss_storage/setting/mysql"
)

type sensitiveWord struct {
	Id      int64  `db:"id"`
	Content string `db:"content"`
}

// listSensitiveWord 初始化SensitiveFilter 专用
func listSensitiveWord(wordChan chan *sensitiveWord, partSize int64) error {
	var count int64
	sqlStr := `select count(*) from sensitive_word`

	err := mysql.DB.Get(&count, sqlStr)
	if err != nil {
		zap.L().Error("敏感词查询数目错误", zap.Error(err))
	}

	// 分批次查询
	// 每次查询大小
	queryCount := int64(math.Ceil(float64(count/partSize)) + 1)
	start := int64(1)
	for i := int64(0); i < queryCount; i++ {
		wordPartArray, err := listSensitiveWordPart(start, partSize)
		if err != nil {
			return err
		}
		for i, word := range wordPartArray {
			wordChan <- word
			if int64(i+1) == partSize {
				start = word.Id
			}
		}
	}
	wordChan <- &sensitiveWord{
		Id: int64(-1),
	}
	return nil
}

// listSensitiveWordPart 初始化SensitiveFilter 专用
func listSensitiveWordPart(start int64, size int64) ([]*sensitiveWord, error) {

	var result []*sensitiveWord

	sqlStr := `select id, content from sensitive_word where id > ? order by id asc limit ?`

	err := mysql.DB.Select(&result, sqlStr, start, size)
	if err != nil {
		zap.L().Error("敏感词查询内容错误", zap.Error(err))
		return nil, err
	}
	return result, nil
}

// insertSensitiveWord 初始化SensitiveFilter 专用
func insertSensitiveWord(row *sensitiveWord) error {

	sqlStr := `insert into sensitive_word (id, content) values (?, ?)`

	_, err := mysql.DB.Exec(sqlStr, row.Id, row.Content)
	if err != nil {
		return err
	}
	return nil
}
