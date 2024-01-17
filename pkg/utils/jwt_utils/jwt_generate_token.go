package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWTGenerateToken(claims *Claims, expirationTime time.Time) (JWTOutput, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return JWTOutput{}, err
	}

	return JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}, err
}
