package routes

import (
	"github.com/baalamurgan/crud-golang/pkg/controllers"
	"github.com/gorilla/mux"
)

var PostRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{postID}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/posts/users/{userID}", controllers.GetPostByUserId).Methods("GET")

}
