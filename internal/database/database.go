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

	db, errOpen := sql.Open("postgres", connectString)
	if errOpen != nil {
		return nil, errOpen
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}

	fmt.Println("Database connected!")
	return db, nil
}
