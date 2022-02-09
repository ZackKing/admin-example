package tools

import "time"

type TS struct{}

func (t *TS) now() int64 {
	return time.Now().Unix()
}
