package queries

import (
	"database/sql"
	"errors"
	"github.com/GbSouza15/apiToDoGo/internal/app/models"
)

var (
	ErrCreateTasks  = errors.New("erro ao criar tarefa")
	ErrGetTasks     = errors.New("erro ao buscar tarefas")
	ErrArticleError = errors.New("error in articles")
)

func CreateTask(db *sql.DB, id string, title string, description string, userId string) error {
	stmt, errPrepare := db.Prepare(`INSERT INTO tdlist.tasks (id, title, description, user_id) VALUES (?, ?, ?, ?)`)
	if errPrepare != nil {
		return ErrCreateTasks
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(id, title, description, userId)
	if errExec != nil {
		return ErrCreateTasks
	}

	return nil
}

func GetTask(db *sql.DB, id string) ([]*models.Task, error) {
	stmt, errPrepare := db.Prepare(`SELECT title, description FROM tdlist.tasks WHERE user_id = ?`)
	if errPrepare != nil {
		return nil, errPrepare
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query(id)
	if errQuery != nil {
		return nil, ErrGetTasks
	}
	defer rows.Close()

	var tasks []*models.Task

	for rows.Next() {
		var task models.Task
		if errScan := rows.Scan(&task.ID, &task.Title, &task.Description, &task.UserID); errScan != nil {
			return nil, ErrArticleError
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}
