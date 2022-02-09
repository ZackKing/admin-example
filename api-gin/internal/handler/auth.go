package handler

import (
	"api-gin/internal/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Account  string `json:"account" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRes struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	data := &LoginReq{}
	if ok := getData(c, &data); !ok {
		return
	}
	token, e := logic.GenJwtToken(1)
	if e != nil {
		fmt.Printf("error: %v", e.Error())
		errorRes(c, -1, "", nil)
		return
	}
	successRes(c, &LoginRes{token})
}

func Self(c *gin.Context) {
	v, ok := c.Get("uid")
	if !ok {
		v = nil
	}
	successRes(c, v)
}
