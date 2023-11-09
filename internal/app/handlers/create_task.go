package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/google/uuid"
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

	_, err = h.DB.Exec("INSERT INTO tdlist.tasks (id, title, description, user_id) VALUES ($1, $2, $3, $4)", taskId, newTask.Title, newTask.Description)
	if err != nil {
		SendResponse(500, []byte("Erro ao criar tarefa."), w)
		return
	}
}
