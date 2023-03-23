package db

import (
	"admin-api/internal/services/conf"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbs = make(map[string]*gorm.DB)
var once sync.Once

func GetGormDb(k string) (db *gorm.DB) {
	db, ok := dbs[k]
	if !ok {
		return nil
	}
	return db
}

func InitGorm() {
	once.Do(func() {
		for k, v := range conf.Db {
			sqlDB, err := sql.Open(
				v.Type,
				fmt.Sprintf(
					"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
					v.User,
					v.Password,
					v.Host,
					v.Port,
					v.Db,
					v.Charset,
				),
			)
			if err != nil {
				log.Fatal("sql open error: ", k, err)
			}
			sqlDB.SetMaxOpenConns(v.MaxConns)
			sqlDB.SetMaxIdleConns(v.MaxIdleConns)

			gc, err := gorm.Open(mysql.New(mysql.Config{
				Conn: sqlDB,
			}), &gorm.Config{})
			if err != nil {
				log.Fatal("gorm open error: ", k, err)
			}

			dbs[k] = gc
		}
	})
}
