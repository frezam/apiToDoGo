package authenticator

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ValidatorToken(tokenString string) (bool, error) {
	secret := os.Getenv("SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
