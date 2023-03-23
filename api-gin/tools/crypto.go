package tools

import (
	"crypto/md5"
	"encoding/hex"
)

type crypto struct{}

var Crypto = &crypto{}

func (c crypto) Md5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

func (c crypto) Md516(str string) string {
	s := c.Md5(str)
	return s[8:24]
}
