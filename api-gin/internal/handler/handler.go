package handler

import (
	"admin-api/internal/logic"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Base struct{}

func (h *Base) okRes(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &Res{0, "ok", data})
}

func (h *Base) errRes(c *gin.Context, code int, msg string, data any, a ...any) {
	if cm, ok := CodeMsg[code]; ok && msg == "" {
		msg = fmt.Sprintf(cm, a...)
	}
	c.JSON(http.StatusOK, &Res{code, msg, data})
}

func (h *Base) parseLogicErr(c *gin.Context, le *logic.LogicErr) {
	msg := ""
	if le.Msg != "" {
		msg = le.Msg
	}
	h.errRes(c, le.Code, msg, le.Data)
}

func (h *Base) validData(c *gin.Context, data any) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		if ve, ok := err.(validator.ValidationErrors); !ok {
			h.errRes(c, 1000, "", nil)
		} else {
			msg := ""
			for _, v := range ve {
				msg += v.Error() + "; "
			}
			h.errRes(c, 1000, msg, nil, ve)
		}
		return false
	}
	return true
}

func (h *Base) validQuery(c *gin.Context, query any) bool {
	if err := c.ShouldBindQuery(&query); err != nil {
		if ve, ok := err.(validator.ValidationErrors); !ok {
			h.errRes(c, 1000, "", nil)
		} else {
			msg := ""
			for _, v := range ve {
				msg += v.Error() + "; "
			}
			h.errRes(c, 1000, msg, nil, ve)
		}
		return false
	}
	return true
}

func (h *Base) validUri(c *gin.Context, param any) bool {
	if err := c.ShouldBindUri(&param); err != nil {
		if ve, ok := err.(validator.ValidationErrors); !ok {
			h.errRes(c, 1000, "", nil)
		} else {
			msg := ""
			for _, v := range ve {
				msg += v.Error() + "; "
			}
			h.errRes(c, 1000, msg, nil, ve)
		}
		return false
	}
	return true
}
