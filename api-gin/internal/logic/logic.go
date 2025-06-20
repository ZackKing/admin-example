package logic

import (
	"admin-api/internal/constant"
	"fmt"
)

type LogicErr struct {
	Code int
	Msg  string
	Data any
}

type SearchConds struct {
	Limit  int
	Offset int
}

func (e *LogicErr) Error() string {
	return fmt.Sprintf("LogicErr{Code: %d, Msg: %s, Data: %v}", e.Code, e.Msg, e.Data)
}

func NewErr(code int, msg string, data any, a ...any) *LogicErr {
	if msg == "" {
		msg = constant.GetCodeMsg(code, a...)
	}
	return &LogicErr{code, msg, data}
}

func NewCodeErr(code int, a ...any) *LogicErr {
	msg := constant.GetCodeMsg(code, a...)
	return &LogicErr{code, msg, nil}
}

func NewMsgErr(code int, msg string) *LogicErr {
	return &LogicErr{code, msg, nil}
}
