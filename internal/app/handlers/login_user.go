package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserLogin
	var user models.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		SendResponse(500, []byte("Erro ao fazer o login"), w)
	}

	if err := json.Unmarshal(body, &userLogin); err != nil {
		SendResponse(500, []byte("Erro na decodificação do json"), w)
		return
	}

	err = h.DB.QueryRow("SELECT * FROM tdlist.users WHERE email = $1", userLogin.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			SendResponse(404, []byte("Nenhum registro desse usuário"), w)
			return
		}
		SendResponse(401, []byte("Erro no servidor"), w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		SendResponse(401, []byte("Senha incorreta"), w)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	secret := os.Getenv("SECRET")
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(secret))
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
