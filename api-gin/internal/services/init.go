package services

import (
	"admin-api/internal/services/cache"
	"admin-api/internal/services/conf"
	"admin-api/internal/services/db"
	"admin-api/internal/services/log"
	"admin-api/internal/services/uuid"
	"os"
)

var Conf = conf.Conf

func Init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	conf.InitConf("configs", env)
	log.InitLogger(conf.App.Log)
	uuid.Init()

	db.InitDb(conf.Db)
	cache.InitCache(conf.Redis)

}
