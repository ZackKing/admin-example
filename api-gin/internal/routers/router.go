package routers

import (
	"admin-api/internal/handler"
	"admin-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	authHandler := &handler.Auth{}
	r.POST("/api/login", authHandler.Login)

	api := r.Group("/api")
	api.Use(middleware.Auth)

	userApi := api.Group("/user")
	userHandler := &handler.User{}
	userApi.GET("/", userHandler.Self)
	userApi.GET("/:id", userHandler.Info)
}
