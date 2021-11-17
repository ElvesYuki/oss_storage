package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"oss_storage/setting"
)

var rdb *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping(rdb.Context()).Result()
	return nil
}

func Close() {
	_ = rdb.Close()
}
