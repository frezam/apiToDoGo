package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/gorilla/mux"
)

func (h handler) GetTasksForUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	res, err := h.DB.Query("SELECT title, description FROM tdlist.tasks WHERE user_id = $1", userId)
	if err != nil {
		SendResponse(500, []byte("Erro ao buscar tarefas"), w)
		return
	}

	tasks := []models.Task{}

	for res.Next() {
		var task models.Task
		if err := res.Scan(&task.ID, &task.Title, &task.Description, &task.UserID); err != nil {
			SendResponse(500, []byte("Error in articles"), w)
			fmt.Println("Error scanning articles: ", err.Error())
			return
		}
		tasks = append(tasks, task)
	}

	responseJSON, err := json.Marshal(tasks)
	if err != nil {
		SendResponse(500, []byte("Erro ao converter para JSON"), w)
		return
	}

	SendResponse(200, responseJSON, w)
}
