package handlers

import (
	"database/sql"
	"net/http"
)

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}

func SendResponse(code int, data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
