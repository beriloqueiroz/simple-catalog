package controllers

import (
	"html/template"
	"net/http"

	"github.com/beriloqueiroz/simple-go-catalog/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListAll()
	temp.ExecuteTemplate(w, "index", products)
}
