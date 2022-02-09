package middleware

import (
	"api-gin/internal/handler"
	"api-gin/internal/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("ADMIN-TOKEN")
	if token == "" {
		c.AbortWithStatusJSON(200, &handler.Res{Code: 10001, Msg: handler.GetErrorCodeMsg(10001), Data: nil})
		return
	}
	claims, err := logic.ParseJwtToken(token)
	if err != nil {
		fmt.Printf("jwt middleware parse error: %v", err.Error())
		c.AbortWithStatusJSON(200, &handler.Res{Code: 10002, Msg: handler.GetErrorCodeMsg(10002), Data: nil})
		return
	}
	c.Set("uid", claims.Uid)
	c.Next()
	println("after auth middleware")
}
