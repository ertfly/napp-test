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
	router.HandleFunc("/products", controllers.ProductsIndex).Methods("GET")
	router.HandleFunc("/products", controllers.ProductsPost).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.ProductsPut).Methods("PUT")
	log.Fatal(http.ListenAndServe("localhost:"+apiPort, router))
}
