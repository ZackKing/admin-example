package middleware

import (
	v1 "admin-api/internal/api/v1"
	"admin-api/internal/services/log"
	"admin-api/pkg/lib/def"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context) {
	defer func() {
		rec := recover()

		if rec == nil {
			return
		}

		stack := string(debug.Stack())

		log.Error("Panic recovered", nil, def.KV{
			"recover": rec,
			"stack":   stack,
		})

		c.JSON(http.StatusInternalServerError, v1.Res{
			Code:     -1,
			Msg:      "unknow",
			Data:     nil,
			Ts:       time.Now().Unix(),
			AccessID: c.GetString("sys_access_id"),
		})

		c.Abort()
	}()

	c.Next()
}
