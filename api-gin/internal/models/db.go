package models

import (
	"api-gin/internal/services/conf"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var conns map[string]*sqlx.DB
var once sync.Once

func GetConnect(k string) (conn *sqlx.DB) {
	once.Do(func() {
		conns = make(map[string]*sqlx.DB)
		initDb("default")
	})
	conn, ok := conns[k]
	if !ok {
		return nil
	}
	return conn
}

func initDb(k string) (err error) {
	f := conf.Db[k]
	conns[k], err = sqlx.Connect(
		f.Type,
		fmt.Sprintf(
			// "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
			"%s:%s@tcp(%s:%d)/%s?charset=%s",
			f.User,
			f.Password,
			f.Host,
			f.Port,
			f.Db,
			f.Charset,
		),
	)
	if err != nil {
		log.Fatal("default db connect init error: ", err)
	}
	conns[k].SetMaxOpenConns(f.MaxConns)
	conns[k].SetMaxIdleConns(f.MaxIdleConns)
	return
}
