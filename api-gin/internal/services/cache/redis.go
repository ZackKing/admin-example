package cache

import (
	"admin-api/internal/services/conf"

	"github.com/redis/go-redis/v9"
)

func newRdb(conf *conf.RedisConf) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	return rdb
}
