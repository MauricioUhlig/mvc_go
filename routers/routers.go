package routers

import (
	"net/http"

	"github.com/MauricioUhlig/mvc_go/controllers"
)

func Router() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}
