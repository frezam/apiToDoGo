package handlers

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}

func New(db *sql.DB) Handler {
	return Handler{db}
}

func SendResponse(code int, data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
