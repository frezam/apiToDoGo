package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
)

func (h handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	fmt.Println(string(body))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		fmt.Println(err.Error())
		return
	}

	var newUser models.User

	if err := json.Unmarshal(body, &newUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao descodificar json"))
		fmt.Println(err.Error())
		return
	}

	_, err = h.DB.Exec("INSERT INTO tdlist.users (name, email, password) VALUES ($1, $2, $3)", newUser.Name, newUser.Email, newUser.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao registrar usuário."))
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário registrado com sucesso."))
}
