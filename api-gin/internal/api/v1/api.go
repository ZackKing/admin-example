package v1

import (
	"admin-api/internal/constant"
	"admin-api/internal/logic"
	"admin-api/internal/services/conf"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Res struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Data     any    `json:"data,omitempty"`
	Ts       int64  `json:"ts"`
	AccessID string `json:"access_id"`
}

type Api struct{}

func (a *Api) GetUid(c *gin.Context) int32 {
	v, _ := c.Get("uid")
	return v.(int32)
}

func (a *Api) OkRes(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &Res{0, "ok", data, time.Now().Unix(), c.GetString("sys_access_id")})
}

func (a *Api) ErrRes(c *gin.Context, code int, msg string, data any, av ...any) {
	if msg == "" {
		msg = constant.GetCodeMsg(code, av...)
	}
	c.JSON(http.StatusOK, &Res{code, msg, data, time.Now().Unix(), c.GetString("sys_access_id")})
}

func (a *Api) ErrCodeRes(c *gin.Context, code int, av ...any) {
	c.JSON(http.StatusOK, &Res{code, constant.GetCodeMsg(code, av...), nil, time.Now().Unix(), c.GetString("sys_access_id")})
}

func (a *Api) errValidRes(c *gin.Context, err error) {
	if ve, ok := err.(validator.ValidationErrors); !ok {
		a.ErrRes(c, 1000, "", nil, err.Error())
	} else {
		msg := ""
		for _, v := range ve {
			msg += v.Error() + "; "
		}
		a.ErrRes(c, 1000, msg, nil, ve)
	}
}

func (a *Api) LogicErrRes(c *gin.Context, e error) {

	var (
		code     = 1
		msg      = ""
		data any = nil
	)

	if le, ok := e.(*logic.LogicErr); ok {
		code = le.Code
		msg = le.Msg
		data = le.Data
	} else if conf.App.Debug {
		msg = e.Error()
	}

	a.ErrRes(c, code, msg, data)
}

func (a *Api) ValidJson(c *gin.Context, data any) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		a.errValidRes(c, err)
		return false
	}
	return true
}

func (a *Api) ValidQuery(c *gin.Context, query any) bool {
	if err := c.ShouldBindQuery(&query); err != nil {
		a.errValidRes(c, err)
		return false
	}
	return true
}

func (a *Api) ValidUri(c *gin.Context, param any) bool {
	fmt.Println(param)
	fmt.Println(&param)
	if err := c.ShouldBindUri(param); err != nil {
		a.errValidRes(c, err)
		return false
	}
	return true
}
