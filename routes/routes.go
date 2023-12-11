package routes

import (
	"net/http"

	"github.com/beriloqueiroz/simple-go-catalog/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
