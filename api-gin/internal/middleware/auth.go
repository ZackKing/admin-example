package middleware

import (
	v1 "admin-api/internal/api/v1"
	"admin-api/internal/constant"
	"admin-api/internal/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("ADMIN-TOKEN")
	if token == "" {
		c.AbortWithStatusJSON(
			200,
			&v1.Res{Code: 10001, Msg: constant.GetCodeMsg(10001), Data: nil},
		)
		return
	}
	claims, err := logic.Jwt.ParseJwtToken(token)
	if err != nil {
		fmt.Printf("jwt middleware parse error: %v", err.Error())
		c.AbortWithStatusJSON(200, &v1.Res{Code: 10002, Msg: constant.GetCodeMsg(10002), Data: nil})
		return
	}
	c.Set("uid", claims.Uid)

	c.Next()

	// TODO: after auth middleware

}
