package constant

import (
	"fmt"
	"strings"
)

var CodeMsg = map[int]string{
	-1:   "unkown",
	0:    "ok",
	1:    "Error",
	2:    "Validator Error",
	3:    "Locked",
	1000: "Parameter error %s",
	2000: "DB error !",

	// auth
	10001: "JWT Token Not Found!",
	10002: "Invalid JWT Token",
	10003: "Account not found / Password error!",
	10004: "Old password error! / password change error!",
	10005: "Account is disabled ! please contact administrator !",
	10006: "Add account error! ",
	10007: "Account user exist!",
	10008: "Account password wrong! will ban with %d retry !",
	10009: "Account password retry too much ! set disabled !",
}

func GetCodeMsg(code int, a ...any) string {
	msg, ok := CodeMsg[code]
	if !ok {
		return "unkown"
	}

	if len(a) > 0 && strings.Contains(msg, "%") {
		return fmt.Sprintf(msg, a...)
	}

	return msg
}
