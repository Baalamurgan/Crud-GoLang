package routes

import (
	"github.com/baalamurgan/crud-golang/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterWeatherRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetUserbyId).Methods("GET")
}
