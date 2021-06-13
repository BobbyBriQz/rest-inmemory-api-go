package routes

import (
	"github.com/gorilla/mux"
	"restProject/controllers"
)

func InitializeRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.PostBook).Methods("POST")

	router.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
	router.HandleFunc("/authors/{id}", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/authors", controllers.PostAuthor).Methods("POST")

	return router
}
