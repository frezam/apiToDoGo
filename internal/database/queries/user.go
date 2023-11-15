package queries

import (
	"database/sql"
	"errors"
	"github.com/GbSouza15/apiToDoGo/internal/app/models"
)

var (
	ErrCreateUser      = errors.New("erro ao registrar usuário")
	ErrUserNoRegistry  = errors.New("nenhum registro desse usuário")
	ErrUserServerError = errors.New("erro no servidor")
)

func CreateUser(db *sql.DB, id string, name string, email string, password []byte) error {
	stmt, errPrepare := db.Prepare(`INSERT INTO tdlist.users (id, name, email, password) VALUES (?, ?, ?, ?)`)
	if errPrepare != nil {
		return ErrCreateUser
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(id, name, email, password)
	if errExec != nil {
		return ErrCreateUser
	}

	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	stmt, errPrepare := db.Prepare(`SELECT * FROM tdlist.users WHERE email = ?`)
	if errPrepare != nil {
		return nil, errPrepare
	}
	defer stmt.Close()

	var user *models.User
	errScan := stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if errors.Is(errScan, sql.ErrNoRows) {
		return nil, ErrUserNoRegistry
	} else if errScan != nil {
		return nil, ErrUserServerError
	}

	return user, nil
}
