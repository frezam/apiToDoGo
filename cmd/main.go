package main

import (
	"github.com/GbSouza15/apiToDoGo/internal/app/routers"
	"github.com/GbSouza15/apiToDoGo/internal/config/database"
	"github.com/GbSouza15/apiToDoGo/internal/config/database/schema"
)

func main() {
	db := database.InitDb()

	schema.CreateSchemaAndTable(db)

	database.CreateTables(db)

	routers.RoutesApi(db)

	database.CloseDb(db)
}
