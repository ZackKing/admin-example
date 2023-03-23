package handler

import "fmt"

type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

var CodeMsg = map[int]string{
	-1:   "unkown",
	0:    "ok",
	1000: "Parameter error: %s",
	2000: "Mysql error !",

	// auth
	10001: "JWT Token Not Found!",
	10002: "Invalid JWT Token",
	10003: "Account not found / Password error!",
	10004: "Old password error! / password change error!",
	10005: "Account is disabled ! please contact administrator !",
	10006: "Add account error! ",
	10007: "Account user exist!",
	10008: "Account password wrong! will ban with %s chance left !",
	10009: "Account password retry too much ! set disabled !",
}

func GetCodeMsg(code int, a ...any) string {
	msg, ok := CodeMsg[code]
	if !ok {
		return "unkown"
	} else {
		return fmt.Sprintf(msg, a...)
	}
}
