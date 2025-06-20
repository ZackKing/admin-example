package sstr

import (
	"strconv"
	"strings"
)

type IntType interface {
	~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
}

func Split(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	if sep == "" {
		return []string{s}
	}
	return strings.Split(s, sep)
}

func Split2IntT[T IntType](s string, sep string) []T {
	if s == "" {
		return []T{}
	}

	ss := strings.Split(s, sep)
	res := make([]T, len(ss))

	for i, v := range ss {
		switch any(res[0]).(type) {
		case int, int32, int64:
			n, _ := strconv.ParseInt(v, 10, 64)
			res[i] = T(n)
		case uint, uint32, uint64:
			n, _ := strconv.ParseUint(v, 10, 64)
			res[i] = T(n)
		}
	}
	return res
}

func Split2Int(s string, sep string) []int {
	return Split2IntT[int](s, sep)
}

func Split2Int32(s string, sep string) []int32 {
	return Split2IntT[int32](s, sep)
}
