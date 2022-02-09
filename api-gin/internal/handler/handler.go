package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func successRes(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Res{0, "ok", data})
	return
}

func errorRes(c *gin.Context, code int, msg string, data interface{}, a ...interface{}) {
	if cm, ok := ErrorCodeMsg[code]; ok && msg == "" {
		msg = fmt.Sprintf(cm, a...)
	}
	c.JSON(http.StatusOK, &Res{code, msg, data})
	return
}

func getData(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		if ve, ok := err.(validator.ValidationErrors); !ok {
			errorRes(c, 1000, "", nil)
		} else {
			msg := ""
			for _, v := range ve {
				msg += v.Error() + "; "
			}
			errorRes(c, 1000, msg, nil, ve)
		}
		return false
	}
	return true
}
