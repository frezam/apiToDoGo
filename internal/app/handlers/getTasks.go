package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/gorilla/mux"
)

func (h handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]

	res, err := h.DB.Query("SELECT * FROM tdlist.tasks WHERE user_id = $1", userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar tarefas"))
		return
	}

	tasks := []models.Task{}

	for res.Next() {
		var task models.Task

		if err := res.Scan(&task.ID, &task.Title, &task.Description, &task.UserID); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error in articles"))
			fmt.Println("Error scanning articles: ", err.Error())
			return
		}

		tasks = append(tasks, task)
	}

	responseJSON, err := json.Marshal(tasks)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao converter para JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
