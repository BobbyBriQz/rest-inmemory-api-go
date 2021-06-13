package routes

import (
	"github.com/gorilla/mux"
	"restProject/controllers"
)

func InitializeRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.PostBook).Methods("POST")

	return router
}
