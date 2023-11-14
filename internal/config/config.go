package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Environment struct {
	Secret  string
	Schema  string
	User    string
	Pass    string
	DB      string
	Host    string
	Port    string
	SSLMode string
}

var Env Environment

func init() {
	// Overload to prevent conflicts with the system environment
	err := godotenv.Overload(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Env = Environment{
		Secret:  os.Getenv("SECRET"),
		Schema:  os.Getenv("SCHEMA"),
		User:    os.Getenv("USER"),
		Pass:    os.Getenv("PASS"),
		DB:      os.Getenv("DB"),
		Host:    os.Getenv("HOST"),
		Port:    os.Getenv("PORT"),
		SSLMode: os.Getenv("SSLMODE"),
	}
}
