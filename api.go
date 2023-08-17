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
	router.HandleFunc("/products/{id}", controllers.ProductsView).Methods("GET")
	router.HandleFunc("/products", controllers.ProductsPost).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.ProductsPut).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.ProductsDelete).Methods("DELETE")
	router.HandleFunc("/stock/{productId}", controllers.StockIndex).Methods("GET")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+apiPort, router))
}
