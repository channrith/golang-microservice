package router

import (
	"github.com/channrith/golang-microservice/mvc/controller"
	"github.com/gorilla/mux"
)

func StartApp() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", controller.GetUserById).Methods("GET")
	r.HandleFunc("/add", controller.AddUser).Methods("POST")

	return r
}
