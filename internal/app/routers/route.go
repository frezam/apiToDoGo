package routers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/handlers"
	"github.com/gorilla/mux"
)

func RoutesApi(db *sql.DB) error {

	r := mux.NewRouter()
	h := handlers.New(db)

	r.HandleFunc("/{userId}/tasks", h.GetTasksForUserHandler).Methods(http.MethodGet)

	r.HandleFunc("/register", h.RegisterUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/login", h.LoginUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/tasks", h.CreateTasks).Methods(http.MethodPost)

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
