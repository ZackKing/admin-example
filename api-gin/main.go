package main

import (
	"admin-api/internal/routers"
	_ "admin-api/internal/services"
	"admin-api/internal/services/conf"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routers.Init(r)

	addr := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	endless.ListenAndServe(addr, r)
}
