package routes

import (
	"net/http"

	"github.com/beriloqueiroz/simple-go-catalog/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
}
