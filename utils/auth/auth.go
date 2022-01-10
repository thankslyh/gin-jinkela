package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	Email string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenToken(email, password string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer: "thankslyh@gmail.com",
		},
	})
	return token.SignedString([]byte("saksnoashcoao"))
}

func VerifyToken(token string) {

}