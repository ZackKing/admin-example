package db

import (
	"admin-api/internal/services/conf"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGorm(conf *conf.DbConf) *gorm.DB {
	sqlDB, err := sql.Open(
		conf.Type,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Db,
			conf.Charset,
		),
	)
	if err != nil {
		log.Fatal("sql open error: ", conf.Host, err)
	}

	sqlDB.SetMaxOpenConns(conf.MaxConns)
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)

	gconf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	}

	gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gconf)
	if err != nil {
		log.Fatal("gorm open error: ", conf.Host, err)
	}

	return gdb
}
