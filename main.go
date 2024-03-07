package main

import (
	"net/http"

	"github.com/MauricioUhlig/mvc/routers"
)

func main() {
	routers.Router()
	http.ListenAndServe(":8080", nil)
}
