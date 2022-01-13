package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

var tokenKey = []byte("thankslyh-jinkela")

func GenToken(userId int, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer: "thankslyh@gmail.com",
		},
	})
	return token.SignedString(tokenKey)
}

func VerifyToken(token string) (*jwt.Token, error) {
	var tokenClaims CustomClaims
	tokenC, err := jwt.ParseWithClaims(token, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		return nil, errors.New("token解析错误")
	}
	if tokenC.Valid {
		return tokenC, nil
	}
	return tokenC, errors.New("token解析错误")
}