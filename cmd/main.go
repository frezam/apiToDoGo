package main

import (
	"fmt"
	"os"

	"github.com/GbSouza15/apiToDoGo/internal/app/routers"
	"github.com/GbSouza15/apiToDoGo/internal/database"
	"github.com/GbSouza15/apiToDoGo/internal/database/schema"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		fmt.Printf("Erro ao iniciar o banco de dados, %v", err.Error())
		os.Exit(1)
	}

	defer db.Close()

	err = schema.CreateSchemaAndTable(db)
	if err != nil {
		fmt.Printf("Erro no schema: %v", err)
	}

	err = database.CreateTables(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = routers.RoutesApi(db)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v", err)
		os.Exit(1)
	}
}
