package main

import (
	"admin-api/internal/routers"
	"admin-api/internal/services"
	"admin-api/internal/services/conf"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	start()
}

func start() {
	services.Init()

	r := gin.New()
	// r.Use(gin.Recovery())
	if !conf.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	routers.Init(r)

	addr := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	endless.ListenAndServe(addr, r)
}
