package schema

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateSchemaAndTable(db *sql.DB) {
	var schemaName = os.Getenv("SCHEMA")
	var createSchemaStr = fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName)

	_, err := db.Exec(createSchemaStr)
	if err != nil {
		fmt.Println("Error creating schema", err.Error())
		return
	}
	fmt.Println("Schema created successfully")
}
