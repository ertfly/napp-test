package controllers

import (
	"napptest/business"
	"napptest/helpers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StockIndex(w http.ResponseWriter, r *http.Request) {
	rows := []interface{}{}

	productId, _ := strconv.ParseInt(mux.Vars(r)["productId"], 10, 0)
	filter := business.StockFilter{}
	filter.ProductId = &productId
	for _, row := range business.StockAll(filter) {

		newRow := map[string]interface{}{
			"id":              *row.Id,
			"product_id":      *row.ProductId,
			"stock_total":     *row.StockTotal,
			"stock_cut":       *row.StockCut,
			"stock_available": *row.StockAvailable,
			"created_at":      *row.CreatedAt,
		}

		rows = append(rows, newRow)
	}

	res := map[string]interface{}{
		"rows": rows,
	}

	helpers.ResponseOk(w, res)
}
