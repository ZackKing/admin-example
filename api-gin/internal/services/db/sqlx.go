package db

import (
	"fmt"
	"log"

	"admin-api/internal/services/conf"

	"github.com/jmoiron/sqlx"
)

func NewDBX(conf *conf.DbConf) *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Db,
		conf.Charset,
	))
	if err != nil {
		log.Fatal("sqlx mysql db open error: ", conf.Host, err)
		return nil
	}
	return db
}
