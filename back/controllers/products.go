package controllers

import (
	"encoding/json"
	"log"
	"napptest/business"
	"napptest/helpers"
	"net/http"
	"strings"
	"time"
)

func ProductsIndex(w http.ResponseWriter, r *http.Request) {
	// res := map[string]interface{}{
	// 	"teste": "teste da silva",
	// }

	helpers.ResponseError(w, 1, "Eric deu merda")
}

func ProductsPost(w http.ResponseWriter, r *http.Request) {
	var row ProductRequest
	json.NewDecoder(r.Body).Decode(&row)

	if row.Sku == nil || strings.Trim(*row.Sku, " ") == "" {
		helpers.ResponseError(w, 1, "sku is required")
		return
	}

	if row.Name == nil || strings.Trim(*row.Name, " ") == "" {
		helpers.ResponseError(w, 1, "name is required")
		return
	}

	if row.StockTotal == nil {
		helpers.ResponseError(w, 1, "stock_total is required")
		return
	}

	if row.StockCut == nil {
		helpers.ResponseError(w, 1, "stock_cut is required")
		return
	}

	check, err := business.ProductsBySku(*row.Sku)
	if err != nil {
		log.Fatalln(err.Error())
		helpers.ResponseError(w, 2, "internal server error")
		return
	}

	if check.Id != nil {
		helpers.ResponseError(w, 1, "sku already exists")
		return
	}

	createdAt := time.Now().UTC().Format("2006-01-02 15:04:05")

	record := business.Products{
		Sku:       row.Sku,
		Name:      row.Name,
		PriceUnit: row.PriceUnit,
		CreatedAt: &createdAt,
	}

	record.Insert(helpers.DatabaseInstance())

	res := map[string]interface{}{
		"id": *record.Id,
	}
	helpers.ResponseOk(w, res)
}

type ProductRequest struct {
	Sku        *string  `json:"sku"`
	Name       *string  `json:"name"`
	PriceUnit  *float64 `json:"price_unit"`
	StockTotal *float64 `json:"stock_total"`
	StockCut   *float64 `json:"stock_cut"`
}
