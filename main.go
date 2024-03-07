package main

import (
	"net/http"

	"github.com/MauricioUhlig/mvc_go/routers"
)

func main() {
	routers.Router()
	http.ListenAndServe(":8080", nil)
}
