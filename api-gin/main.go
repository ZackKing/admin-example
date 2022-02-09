package main

import (
	"api-gin/internal/routers"
	"api-gin/internal/services/conf"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	conf.InitConf("./configs", env)
	// logger.InitLogger("./logs")
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routers.Init(r)

	r.Run(fmt.Sprintf("%s:%d", conf.Common.Host, conf.Common.Port))
}
