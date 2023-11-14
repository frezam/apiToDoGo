package database

import (
	"database/sql"
	"fmt"
	"github.com/GbSouza15/apiToDoGo/internal/config"

	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {

	connectString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		config.Env.User,
		config.Env.DB,
		config.Env.Pass,
		config.Env.Host,
		config.Env.Port,
		config.Env.SSLMode)

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}

	fmt.Println("Database connected!")
	return db, nil
}

func CreateTables(db *sql.DB) error {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS tdlist.users (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL 
		);
	`

	createTaskTable := `
		CREATE TABLE IF NOT EXISTS tdlist.tasks (
			id UUID PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			user_id UUID REFERENCES tdlist.users(id)
		);
	`

	_, err := db.Exec(createUserTable)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela User: %s", err.Error())
	}

	_, err = db.Exec(createTaskTable)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela Task: %s", err.Error())
	}

	return nil
}
