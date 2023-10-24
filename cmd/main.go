package main

import (
	"fmt"

	"github.com/GbSouza15/apiToDoGo/internal/config/database"
	"github.com/GbSouza15/apiToDoGo/internal/config/database/schema"

	"net/http"
)

func main() {
	db := database.InitDb()

	schema.CreateSchemaAndTable(db)

	database.CloseDb(db)

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
