package queries

import (
	"database/sql"
	"fmt"
	"github.com/GbSouza15/apiToDoGo/internal/migrations"
)

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(migrations.TablesScript)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela User: %s", err.Error())
	}

	return nil
}
