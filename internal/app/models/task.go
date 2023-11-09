package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

type TaskCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
