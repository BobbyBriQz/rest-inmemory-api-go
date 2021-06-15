package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"restProject/routes"
)

func main() {
	serverName, serverPort := loadEnv()
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	fmt.Printf("Starting %v Server on port %s", serverName, serverPort)
	http.Handle("/", router)
	log.Fatalln(http.ListenAndServe(serverPort, nil))
}

func loadEnv() (string, string) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	//Viper returns interface{} so type assertion must be done
	serverName, _ := viper.Get("SERVER_NAME").(string)
	serverPort, _ := viper.Get("SERVER_PORT").(string)

	return serverName, serverPort
}
