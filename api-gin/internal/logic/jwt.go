package logic

import (
	"admin-api/internal/services/conf"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtLogic struct{}

var Jwt = &jwtLogic{}

type JwtClaims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

// Gen jwt token string
func (j *jwtLogic) GenJwtToken(uid int) (string, error) {
	now := time.Now().Unix()
	exp := now + int64(conf.Jwt.Exp)
	claims := &JwtClaims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    conf.Jwt.Iss,
		},
	}
	tokenClaims := jwt.NewWithClaims(getJwtAlgorithmsMethod(conf.Jwt.Algorithms), claims)
	token, err := tokenClaims.SignedString([]byte(conf.Jwt.SecretKey))
	return token, err
}

// Parse token string
func (j *jwtLogic) ParseJwtToken(token string) (*JwtClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(conf.Jwt.SecretKey), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// Get method of algorithms
func getJwtAlgorithmsMethod(algorithms string) *jwt.SigningMethodHMAC {
	var method *jwt.SigningMethodHMAC
	switch algorithms {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	case "HS512":
		method = jwt.SigningMethodHS512
	default:
		method = nil
	}
	return method
}
