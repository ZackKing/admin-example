package cache

import (
	"admin-api/internal/services/conf"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var once sync.Once
var Rdb = make(map[string]*redis.Client)

func InitCache(conf map[string]*conf.RedisConf) {
	for k, v := range conf {
		Rdb[k] = newRdb(v)
		if ok := Rdb[k]; ok == nil {
			panic(fmt.Sprintf("Cache [%s] init error", k))
		}
		log.Default().Printf("Cache [%s] initialized", k)
	}
}
