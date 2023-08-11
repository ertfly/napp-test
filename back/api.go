package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dotenv := goDotEnvVariable("STRONGEST_AVENGER")
	fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)

	router := mux.NewRouter()
	router.HandleFunc("/products", CheckFaceID).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
func CheckFaceID(w http.ResponseWriter, r *http.Request) {
	var res ResponseCheck
	res.Msg = "Deu certo"
	json.NewEncoder(w).Encode(res)
}

type ResponseCheck struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
