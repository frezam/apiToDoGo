package schema

import (
	"database/sql"
	"fmt"
	"github.com/GbSouza15/apiToDoGo/internal/config"
)

func CreateSchemaAndTable(db *sql.DB) error {
	var createSchemaStr = fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", config.Env.Schema)

	_, err := db.Exec(createSchemaStr)
	if err != nil {
		return err
	}
	fmt.Println("Schema created successfully")
	return nil
}
