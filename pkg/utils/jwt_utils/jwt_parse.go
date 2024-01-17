package utils

// import (
// 	"os"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func JWTParser(tokenValue string, claims Claims) (jwt.Token, error){
// 	jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	})
// }
