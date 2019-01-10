package utils

import (
	"github.com/dgrijalva/jwt-go"
	"Kiddy/setting"
	"time"
)

func SetToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(15)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(setting.JwtSecret))
	if err != nil {
		return ""
	}
	return tokenString
}
