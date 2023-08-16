package controllers

import (
	"encoding/json"
	"napptest/helpers"
	"net/http"
)

func ProductsIndex(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"teste": "teste da silva",
	}

	json.NewEncoder(w).Encode(helpers.ResponseOk(res))
}

func ProductsPost(w http.ResponseWriter, r *http.Request) {

	// var res ProductsIndexResponse
	// res.Msg = "Deu certo 2222"
	// json.NewEncoder(w).Encode(res)
}
