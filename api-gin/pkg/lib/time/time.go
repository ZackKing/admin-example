package time

import "time"

type TS struct{}

var T = &TS{}

func (t *TS) Now() int64 {
	return time.Now().Unix()
}

func (t *TS) NowNano() int64 {
	return time.Now().UnixNano()
}
