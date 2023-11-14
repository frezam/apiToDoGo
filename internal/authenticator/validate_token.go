package authenticator

import (
	"github.com/GbSouza15/apiToDoGo/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func ValidatorToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.Secret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
