package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
)

func (h handler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login user")

	var userLogin models.UserLogin
	//var user models.User

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao fazer login."))
	}

	if err := json.Unmarshal(body, &userLogin); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro na decodificação do json"))
		return
	}

	// Lógica para o login

	fmt.Println(userLogin)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Nenhum registro desse usuário"))
		return
	}
}
