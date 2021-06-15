package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"restProject/controllers"
	"restProject/services"
)

func InitializeRoutes(router *mux.Router) { // notice how we are passing a pointer so we don't have to return an reassign

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	var mbs = services.NewMyBookService() // Create service here to inject into controller
	bc := controllers.NewBookController(mbs)
	//Either we create a subrouter for books
	bookSubrouter := subrouter.PathPrefix("/books").Subrouter()
	bookSubrouter.HandleFunc("", bc.GetBooks).Methods(http.MethodGet)
	bookSubrouter.HandleFunc("/{id}", bc.GetBook).Methods(http.MethodGet)
	bookSubrouter.HandleFunc("", bc.PostBook).Methods(http.MethodPost)

	var mas = services.NewMyAuthorService() // Create service here to inject into controller
	ac := controllers.NewAuthorController(mas)
	//Or we manually type in 'author' path everytime instead of creating an authors subRouter
	subrouter.HandleFunc("/authors", ac.GetAuthors).Methods("GET")
	subrouter.HandleFunc("/authors/{id}", ac.GetAuthor).Methods("GET")
	subrouter.HandleFunc("/authors", ac.PostAuthor).Methods("POST")

}
