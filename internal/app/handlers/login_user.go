package handlers

import (
	"encoding/json"
	"errors"
	"github.com/GbSouza15/apiToDoGo/internal/config"
	"github.com/GbSouza15/apiToDoGo/internal/database/queries"
	"io"
	"net/http"
	"time"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h Handler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserLogin

	body, err := io.ReadAll(r.Body)
	if err != nil {
		SendResponse(500, []byte("Erro ao fazer o login"), w)
	}

	if err := json.Unmarshal(body, &userLogin); err != nil {
		SendResponse(500, []byte("Erro na decodificação do json"), w)
		return
	}

	user, errGetUserByEmail := queries.GetUserByEmail(h.DB, userLogin.Email)
	if errGetUserByEmail != nil {
		if errors.Is(errGetUserByEmail, queries.ErrUserNoRegistry) {
			SendResponse(404, []byte(queries.ErrUserNoRegistry.Error()), w)
			return
		}
		SendResponse(401, []byte(queries.ErrUserServerError.Error()), w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		SendResponse(401, []byte("Senha incorreta"), w)
		return
	}

	claims := &models.Claims{UserId: user.ID.String(), RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Env.Secret))
	if err != nil {
		SendResponse(500, []byte("Erro ao gerar o token"), w)
		return
	}

	response := map[string]string{"token": tokenString}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		SendResponse(500, []byte("Erro ao codificar o JSON"), w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
	})
	SendResponse(200, responseJSON, w)
}
