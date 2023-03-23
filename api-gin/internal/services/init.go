package services

import (
	"admin-api/internal/services/conf"
	"admin-api/internal/services/db"
	"os"
)

var Conf = conf.Conf

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	conf.InitConf("./configs", env)
	db.InitGorm()
}
