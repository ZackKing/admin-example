package handler

import (
	"admin-api/internal/logic"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Base
}

type LoginReq struct {
	Account  string `json:"account" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRes struct {
	Token string `json:"token"`
}

func (h *Auth) Login(c *gin.Context) {
	data := &LoginReq{}
	if ok := h.validData(c, &data); !ok {
		return
	}
	token, err := logic.Auth.Login(data.Account, data.Password)
	if err != nil {
		h.parseLogicErr(c, err)
		return
	}
	h.okRes(c, &LoginRes{token})
}

func (h *Auth) RenewToken(c *gin.Context) {
	uid := c.GetInt("uid")
	token, ok := logic.Jwt.GenJwtToken(uid)
	if ok != nil {
		h.errRes(c, -1, "", nil)
	}
	h.okRes(c, &LoginRes{token})
}
