package tools

import "time"

type TS struct{}

func (t *TS) Now() int64 {
	return time.Now().Unix()
}
