package routers

import (
	"api-gin/internal/handler"
	"api-gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.Static("/public", "./public")

	r.POST("/login", handler.Login)

	userApi := r.Group("/user")
	userApi.Use(middleware.Auth)
	userApi.GET("/", handler.Self)

}
