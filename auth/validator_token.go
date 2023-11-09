package auth

import (
	"net/http"
	"os"

	"github.com/GbSouza15/apiToDoGo/internal/app/handlers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func ValidatorToken(tokenString string, w http.ResponseWriter) (*jwt.Token, error) {
	godotenv.Load()
	secret := os.Getenv("SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		handlers.SendResponse(401, []byte("Token inválido"), w)
		return nil, err
	}

	return token, nil
}
