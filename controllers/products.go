package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/beriloqueiroz/simple-go-catalog/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListAll()
	temp.ExecuteTemplate(w, "index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro ao converter preço!", err)
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro ao converter quantidade!", err)
		}

		product := models.Product{
			Name:        name,
			Description: description,
			Price:       priceFloat,
			Quantity:    quantityInt,
		}

		models.Insert(product)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	models.Delete(id)

	http.Redirect(w, r, "/", 301)

}
