package controllers

import (
	"encoding/json"
	"net/http"
)

func ProductsIndex(w http.ResponseWriter, r *http.Request) {
	var res ProductsIndexResponse
	res.Msg = "Deu certo 2222"
	json.NewEncoder(w).Encode(res)
}

type ProductsIndexResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
