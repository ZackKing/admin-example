package middleware

import (
	"admin-api/internal/services/log"
	"admin-api/pkg/lib/def"
	"time"

	"github.com/gin-gonic/gin"
)

var Logger gin.HandlerFunc = gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

	logData := def.KV{
		"ip":        param.ClientIP,
		"time":      param.TimeStamp.Format(time.RFC3339),
		"method":    param.Method,
		"path":      param.Path,
		"proto":     param.Request.Proto,
		"status":    param.StatusCode,
		"latency":   param.Latency,
		"ua":        param.Request.UserAgent(),
		"error_msg": param.ErrorMessage,
	}

	log.Debug("gin_logger", logData)
	return ""
})
