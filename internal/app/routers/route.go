package routers

import (
	"database/sql"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/handlers"
	"github.com/GbSouza15/apiToDoGo/internal/app/middleware"
	"github.com/gorilla/mux"
)

func RoutesApi(db *sql.DB) *mux.Router {

	r := mux.NewRouter()
	h := handlers.New(db)

	r.HandleFunc("/{userId}/tasks", h.GetTasksForUserHandler).Methods(http.MethodGet)

	r.HandleFunc("/register", h.RegisterUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/login", h.LoginUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/tasks", middleware.CheckTokenIsValid(h.CreateTasks)).Methods(http.MethodPost)

	return r
}
