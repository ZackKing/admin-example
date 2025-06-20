package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
)

var Cors = cors.New(cors.Config{
	AllowAllOrigins:  true,
	AllowWildcard:    true,
	AllowMethods:     []string{"*"},
	AllowHeaders:     []string{"*"},
	AllowCredentials: true,
	MaxAge:           24 * time.Hour,
})
