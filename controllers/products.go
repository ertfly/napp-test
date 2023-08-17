package controllers

import (
	"encoding/json"
	"log"
	"napptest/business"
	"napptest/helpers"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func ProductsIndex(w http.ResponseWriter, r *http.Request) {
	rows := []interface{}{}

	for _, row := range business.ProductsAll() {
		lastStock := *row.GetLastStock()
		stock := map[string]interface{}{
			"stock_total":     *lastStock.StockTotal,
			"stock_cut":       *lastStock.StockCut,
			"stock_available": *lastStock.StockAvailable,
			"last_update":     *lastStock.CreatedAt,
		}
		newRow := map[string]interface{}{
			"id":         *row.Id,
			"sku":        *row.Sku,
			"name":       *row.Name,
			"price_unit": *row.PriceUnit,
			"stock":      stock,
		}

		rows = append(rows, newRow)
	}

	res := map[string]interface{}{
		"rows": rows,
	}

	helpers.ResponseOk(w, res)
}

func ProductsView(w http.ResponseWriter, r *http.Request) {
	record, err := business.ProductsById(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalln(err.Error())
		helpers.ResponseError(w, 2, "internal server error")
		return
	}

	if record.Id == nil {
		helpers.ResponseError(w, 1, "product not found")
		return
	}

	lastStock := *record.GetLastStock()
	stock := map[string]interface{}{
		"stock_total":     *lastStock.StockTotal,
		"stock_cut":       *lastStock.StockCut,
		"stock_available": *lastStock.StockAvailable,
		"last_update":     *lastStock.CreatedAt,
	}
	row := map[string]interface{}{
		"id":         *record.Id,
		"sku":        *record.Sku,
		"name":       *record.Name,
		"price_unit": *record.PriceUnit,
		"stock":      stock,
	}

	helpers.ResponseOk(w, row)
}

func ProductsPost(w http.ResponseWriter, r *http.Request) {
	var row ProductRequest
	json.NewDecoder(r.Body).Decode(&row)

	if !ProductsValidate(w, row) {
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

	createdAt = time.Now().UTC().Format("2006-01-02 15:04:05")
	stockAvailable := (*row.StockTotal - *row.StockCut)
	stock := business.Stock{
		ProductId:      record.Id,
		StockTotal:     row.StockTotal,
		StockCut:       row.StockCut,
		StockAvailable: &stockAvailable,
		CreatedAt:      &createdAt,
	}

	stock.Insert(helpers.DatabaseInstance())

	record.LastStockId = stock.Id
	record.Update(helpers.DatabaseInstance())

	res := map[string]interface{}{
		"id": *record.Id,
	}
	helpers.ResponseOk(w, res)
}

func ProductsPut(w http.ResponseWriter, r *http.Request) {
	var row ProductRequest
	json.NewDecoder(r.Body).Decode(&row)

	record, err := business.ProductsById(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalln(err.Error())
		helpers.ResponseError(w, 2, "internal server error")
		return
	}

	if record.Id == nil {
		helpers.ResponseError(w, 1, "product not found")
		return
	}

	row.Id = record.Id
	if !ProductsValidate(w, row) {
		return
	}

	updatedAt := time.Now().UTC().Format("2006-01-02 15:04:05")

	record.Sku = row.Sku
	record.Name = row.Name
	record.PriceUnit = row.PriceUnit
	record.UpdatedAt = &updatedAt

	createdAt := time.Now().UTC().Format("2006-01-02 15:04:05")
	stockAvailable := (*row.StockTotal - *row.StockCut)
	stock := business.Stock{
		ProductId:      record.Id,
		StockTotal:     row.StockTotal,
		StockCut:       row.StockCut,
		StockAvailable: &stockAvailable,
		CreatedAt:      &createdAt,
	}

	stock.Insert(helpers.DatabaseInstance())

	record.LastStockId = stock.Id
	record.Update(helpers.DatabaseInstance())

	res := map[string]interface{}{
		"id": *record.Id,
	}
	helpers.ResponseOk(w, res)
}

func ProductsDelete(w http.ResponseWriter, r *http.Request) {
	record, err := business.ProductsById(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalln(err.Error())
		helpers.ResponseError(w, 2, "internal server error")
		return
	}

	if record.Id == nil || *record.Trash {
		helpers.ResponseError(w, 1, "product not found")
		return
	}

	record.Delete(helpers.DatabaseInstance())

	res := map[string]interface{}{
		"id": *record.Id,
	}
	helpers.ResponseOk(w, res)
}

func ProductsValidate(w http.ResponseWriter, row ProductRequest) bool {
	if row.Sku == nil || strings.Trim(*row.Sku, " ") == "" {
		helpers.ResponseError(w, 1, "sku is required")
		return false
	}

	check, err := business.ProductsBySku(*row.Sku)
	if err != nil {
		log.Fatalln(err.Error())
		helpers.ResponseError(w, 2, "internal server error")
		return false
	}

	if check.Id != nil && row.Id == nil {
		helpers.ResponseError(w, 1, "sku already exists")
		return false
	}
	if check.Id != nil && row.Id != nil {
		if strconv.FormatInt(*check.Id, 10) != strconv.FormatInt(*row.Id, 10) {
			helpers.ResponseError(w, 1, "sku already exists")
			return false
		}
	}

	if row.Name == nil || strings.Trim(*row.Name, " ") == "" {
		helpers.ResponseError(w, 1, "name is required")
		return false
	}

	if row.StockTotal == nil {
		helpers.ResponseError(w, 1, "stock_total is required")
		return false
	}

	if row.StockCut == nil {
		helpers.ResponseError(w, 1, "stock_cut is required")
		return false
	}

	return true
}

type ProductRequest struct {
	Id         *int64
	Sku        *string  `json:"sku"`
	Name       *string  `json:"name"`
	PriceUnit  *float64 `json:"price_unit"`
	StockTotal *float64 `json:"stock_total"`
	StockCut   *float64 `json:"stock_cut"`
}
