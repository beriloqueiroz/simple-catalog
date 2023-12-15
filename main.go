package main

import (
	"net/http"

	"github.com/beriloqueiroz/simple-go-catalog/db"
	"github.com/beriloqueiroz/simple-go-catalog/routes"
)

func main() {
	db.CreateTables()
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
