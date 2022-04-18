package token

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func GetToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("secretkey")), nil
	})

	return parsedToken, err
}

func CheckValid(token *jwt.Token) bool {
	if !token.Valid {
		return false
	}
	return true
}

func GetPayload(token *jwt.Token) jwt.MapClaims {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	}
	return nil
}
