package routes

import (
	"github.com/gorilla/mux"
	"restProject/controllers"
)

func InitializeRoutes(router *mux.Router) { // notice how we are passing a pointer so we don't have to return an reassign

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//Either we create a subrouter for books
	bookSubrouter := subrouter.PathPrefix("/books").Subrouter()
	bookSubrouter.HandleFunc("", controllers.GetBooks).Methods("GET")
	bookSubrouter.HandleFunc("/{id}", controllers.GetBook).Methods("GET")
	bookSubrouter.HandleFunc("", controllers.PostBook).Methods("POST")

	//Or we manually type in author everytime instead of creating an authors subRouter
	subrouter.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
	subrouter.HandleFunc("/authors/{id}", controllers.GetAuthor).Methods("GET")
	subrouter.HandleFunc("/authors", controllers.PostAuthor).Methods("POST")

}
