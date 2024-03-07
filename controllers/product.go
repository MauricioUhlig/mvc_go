package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/MauricioUhlig/mvc_go/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	products := models.GetProducts()
	tmpl.ExecuteTemplate(w, "Index", products)
}
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceString := r.FormValue("price")
		quantityString := r.FormValue("quantity")

		price, err := strconv.ParseFloat(priceString, 32)
		if err != nil {
			panic(err.Error())
		}
		quantity, err := strconv.Atoi(quantityString)
		if err != nil {
			panic(err.Error())
		}
		models.InsertProduct(name, description, quantity, float32(price))
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err.Error())
	}
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err.Error())
	}
	product := models.GetProduct(id)
	tmpl.ExecuteTemplate(w, "Edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idString := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceString := r.FormValue("price")
		quantityString := r.FormValue("quantity")

		price, err := strconv.ParseFloat(priceString, 32)
		if err != nil {
			panic(err.Error())
		}
		quantity, err := strconv.Atoi(quantityString)
		if err != nil {
			panic(err.Error())
		}
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err.Error())
		}
		models.UpdateProduct(id, name, description, quantity, float32(price))
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
