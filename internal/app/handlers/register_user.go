package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GbSouza15/apiToDoGo/internal/database/queries"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrFailedRequest    = errors.New("erro ao ler o corpo da requisição")
	ErrFailedDecodeJSON = errors.New("erro ao descodificar json")
	ErrHashError        = errors.New("Error in Hash.")
)

func (h Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		SendResponse(500, []byte(ErrFailedRequest.Error()), w)
		fmt.Println(err.Error())
		return
	}

	var newUser models.User
	userId := uuid.NewString()

	if err := json.Unmarshal(body, &newUser); err != nil {
		SendResponse(500, []byte(ErrFailedDecodeJSON.Error()), w)
		fmt.Println(err.Error())
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		fmt.Println(ErrHashError.Error())
	}

	err = queries.CreateUser(h.DB, userId, newUser.Name, newUser.Email, bytes)
	if err != nil {
		SendResponse(500, []byte(err.Error()), w)
		fmt.Println(err.Error())
		return
	}

	SendResponse(200, []byte("Usuário registrado com sucesso."), w)
}
