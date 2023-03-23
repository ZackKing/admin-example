package handler

import "github.com/gin-gonic/gin"

type User struct {
	Base
	Path string
}

func (h *User) Self(c *gin.Context) {
	v, ok := c.Get("uid")
	if !ok {
		v = nil
	}
	h.okRes(c, v)
}

func (h *User) Info(c *gin.Context) {
	v, ok := c.Get("uid")
	if !ok {
		v = nil
	}
	h.okRes(c, v)
}
