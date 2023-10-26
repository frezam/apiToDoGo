package database

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDb() *sql.DB {

	var err error

	godotenv.Load()

	var (
		user    = os.Getenv("USER")
		pass    = os.Getenv("PASS")
		dbname  = os.Getenv("DB")
		host    = os.Getenv("HOST")
		port    = os.Getenv("PORT")
		sslmode = os.Getenv("SSLMODE")
	)

	connectString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, dbname, pass, host, port, sslmode)

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		fmt.Println("Error opening database connection", err.Error())
		panic(err)
	}

	fmt.Println("Database connected!")

	return db
}

func CloseDb(db *sql.DB) {
	defer db.Close()
}

func CreateTables(db *sql.DB) {
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
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			user_id UUID REFERENCES tdlist.users(id)
		);
	`

	_, err := db.Exec(createUserTable)

	if err != nil {
		fmt.Println("Error creating user table", err.Error())
		panic(err)
	}

	_, err = db.Exec(createTaskTable)

	if err != nil {
		fmt.Println("Error creating task table", err.Error())
		panic(err)
	}
}
