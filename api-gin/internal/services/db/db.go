package db

import (
	"admin-api/internal/services/conf"
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var GdbMap = make(map[string]*gorm.DB)

func InitDb(conf map[string]*conf.DbConf) {
	once.Do(func() {
		for k, v := range conf {
			GdbMap[k] = NewGorm(v)
			if ok := GdbMap[k]; ok == nil {
				panic(fmt.Sprintf("DB [%s] init error", k))
			}
			log.Default().Printf("DB [%s] initialized", k)
		}
	})
}

func GetGDB(k string) (db *gorm.DB) {
	db, ok := GdbMap[k]
	if !ok {
		return nil
	}
	return db
}
