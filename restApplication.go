package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restProject/routes"
)

func main() {
	router := mux.NewRouter()
	router = routes.InitializeRoutes(router)

	fmt.Println("Starting Server")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
