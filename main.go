package main

import (
	"net/http"

	"github.com/beriloqueiroz/simple-go-catalog/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
