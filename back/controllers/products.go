package controllers

import (
	"napptest/helpers"
	"net/http"
)

func ProductsIndex(w http.ResponseWriter, r *http.Request) {
	// res := map[string]interface{}{
	// 	"teste": "teste da silva",
	// }

	helpers.ResponseError(w, 1, "Eric deu merda")
}

func ProductsPost(w http.ResponseWriter, r *http.Request) {

	helpers.ResponseError(w, 1, "Eric deu merda")
}
