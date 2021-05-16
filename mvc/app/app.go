package app

import (
	"net/http"

	"github.com/channrith/golang-microservice/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
