package sstr

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

var DefStrSet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var LowStrSet string = "abcdefghijklmnopqrstuvwxyz0123456789"

func LenRune(s string) int {
	if s == "" {
		return 0
	}

	return utf8.RuneCountInString(s)
}

func Random(length int, base string) string {
	if length < 1 {
		return ""
	}

	if base == "" {
		base = DefStrSet
	}

	var str = make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range str {
		str[i] = base[r.Intn(len(base))]
	}

	return string(str)
}
