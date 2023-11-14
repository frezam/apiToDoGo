package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/GbSouza15/apiToDoGo/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func (h handler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		SendResponse(404, []byte("Erro ao ler o corpo da requisição."), w)
		return
	}

	var newTask models.TaskCreate
	var taskId = uuid.NewString()

	if err := json.Unmarshal(body, &newTask); err != nil {
		SendResponse(500, []byte("Erro ao decodificação do JSON"), w)
		return
	}

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			SendResponse(401, []byte("Não autorizado"), w)
			return
		}
		SendResponse(400, []byte("Erro no servidor"), w)
		return
	}

	tknStr := c.Value
	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Env.Secret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			SendResponse(401, []byte("Não autorizado"), w)
			return
		}
		SendResponse(400, []byte("Erro no servidor"), w)
		return
	}
	if !tkn.Valid {
		SendResponse(401, []byte("Não autorizado"), w)
		return
	}

	_, err = h.DB.Exec("INSERT INTO tdlist.tasks (id, title, description, user_id) VALUES ($1, $2, $3, $4)", taskId, newTask.Title, newTask.Description, claims.UserId)
	if err != nil {
		fmt.Println(err)
		fmt.Println(newTask)
		SendResponse(500, []byte("Erro ao criar tarefa."), w)
		return
	}

	SendResponse(201, []byte("Tarefa criada com sucesso."), w)
}
