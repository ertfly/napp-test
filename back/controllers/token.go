package controllers

import (
	"encoding/json"
	"net/http"
)

func TokenPost(w http.ResponseWriter, r *http.Request) {
	var res ProductsIndexResponse
	res.Msg = "Deu certo"
	json.NewEncoder(w).Encode(res)
}

type ProductsPostResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
