package main

import (
	"fmt"
	"github.com/GbSouza15/apiToDoGo/internal/database/queries"
	"net/http"
	"os"

	"github.com/GbSouza15/apiToDoGo/internal/app/routers"
	"github.com/GbSouza15/apiToDoGo/internal/database"
)

func main() {
	db, err := database.InitDb()
	checkError(err)
	defer db.Close()

	err = queries.CreateSchemaAndTable(db)
	checkError(err)

	err = queries.CreateTables(db)
	checkError(err)

	r := routers.RoutesApi(db)

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	errListen := http.ListenAndServe(":8080", nil)
	checkError(errListen)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
