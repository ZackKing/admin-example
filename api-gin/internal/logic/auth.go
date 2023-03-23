package logic

import (
	"admin-api/internal/models"
	"admin-api/tools"
	"fmt"
)

type AuthLogic struct{}

var Auth = &AuthLogic{}

func (l *AuthLogic) Login(name string, pwd string) (token string, le *LogicErr) {
	user := models.UserMdl.GetValidUserByName(name)
	if user.Id == 0 {
		return "", NewErr(10003, "", nil)
	}

	if !l.checkPwd(pwd, user.Salt, user.Password) {
		return "", NewErr(10003, "", nil)
	}

	token, err := Jwt.GenJwtToken(user.Id)
	if err != nil {
		return "", NewErr(-1, fmt.Sprintf("Jwt gen error: %s", err.Error()), nil)
	}
	return token, nil
}

func (a *AuthLogic) checkPwd(pwd string, salt string, dbPwd string) bool {
	spwd := tools.Crypto.Md5(salt + tools.Crypto.Md5(pwd))
	return spwd == dbPwd
}
