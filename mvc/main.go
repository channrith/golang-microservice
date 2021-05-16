package main

import (
	"net/http"

	app "github.com/channrith/golang-microservice/mvc/router"
)

func main() {
	r := app.StartApp()

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
