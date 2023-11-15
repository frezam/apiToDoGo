package queries

import (
	"database/sql"
	"fmt"
	"github.com/GbSouza15/apiToDoGo/internal/config"
)

func CreateSchemaAndTable(db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", config.Env.Schema))
	if err != nil {
		return fmt.Errorf("erro ao criar tabela User: %s", err.Error())
	}

	fmt.Println("Schema created successfully")
	return nil
}
