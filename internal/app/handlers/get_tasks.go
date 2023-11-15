package handlers

import (
	"encoding/json"
	"github.com/GbSouza15/apiToDoGo/internal/database/queries"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) GetTasksForUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	tasks, errGetTask := queries.GetTask(h.DB, userId)
	if errGetTask != nil {
		SendResponse(500, []byte(errGetTask.Error()), w)
		return
	}

	responseJSON, err := json.Marshal(tasks)
	if err != nil {
		SendResponse(500, []byte("Erro ao converter para JSON"), w)
		return
	}

	SendResponse(200, responseJSON, w)
}
