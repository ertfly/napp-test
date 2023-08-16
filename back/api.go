package main

import (
	"log"
	"napptest/controllers"
	"napptest/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	Migration()

	apiPort := helpers.GetEnv("API_PORT")

	router := mux.NewRouter()
	router.HandleFunc("/token", controllers.TokenPost).Methods("POST")
	router.HandleFunc("/products", controllers.ProductsIndex).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:"+apiPort, router))
}
