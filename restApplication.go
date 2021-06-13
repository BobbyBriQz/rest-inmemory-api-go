package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restProject/routes"
)

const port = ":8080"

func main() {
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	fmt.Println("Starting Server on port", port)
	http.Handle("/", router)
	log.Fatalln(http.ListenAndServe(port, nil))
}
