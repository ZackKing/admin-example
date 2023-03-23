package logic

type LogicErr struct {
	Code int
	Msg  string
	Data any
}

func NewErr(Code int, Msg string, Data any) *LogicErr {
	return &LogicErr{Code, Msg, Data}
}
