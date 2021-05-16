package main

import (
	"net/http"

	"github.com/channrith/golang-microservice/mvc/app"
)

func main() {
	app.StartApp()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
