package main

import (
	"fmt"

	"github.com/GbSouza15/apiToDoGo/internal/app/routers"
	"github.com/GbSouza15/apiToDoGo/internal/database"
	"github.com/GbSouza15/apiToDoGo/internal/database/schema"
)

func main() {
	db := database.InitDb()

	schema.CreateSchemaAndTable(db)

	database.CreateTables(db)

	err := routers.RoutesApi(db)

	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v", err)
	}

	database.CloseDb(db)
}
