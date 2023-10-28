package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
)

func (h handler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserLogin
	var user models.User

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

	err = h.DB.QueryRow("SELECT * FROM tdlist.users WHERE email = $1", userLogin.Email).Scan(&user.ID, &user.Email, &user.Name, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Nenhum registro desse usuário"))
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Erro no servidor"))
			return
		}
	}

	responseJSON, err := json.Marshal(user)

	if err != nil {
		w.Write([]byte("Erro ao codificar o JSON"))
		return
	}

	fmt.Println(user)
	w.Write(responseJSON)
}
