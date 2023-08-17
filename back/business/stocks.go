package business

import (
	"database/sql"
	"log"
	"napptest/helpers"

	_ "github.com/go-sql-driver/mysql"
)

func StockNew() Stock {
	row := Stock{}
	return row
}

func StockById(id string) (Stock, error) {
	db := helpers.DatabaseInstance()
	row := Stock{}
	res, err := db.Query("SELECT * FROM stock WHERE `id` = ? limit 1", id)
	if err != nil {
		return row, err
	}
	defer db.Close()

	for res.Next() {
		res.Scan(&row.Id, &row.ProductId, &row.StockTotal, &row.StockCut, &row.StockAvailable, &row.CreatedAt)
		return row, nil
	}

	return row, nil
}

type Stock struct {
	Id             *int64
	ProductId      *int64
	StockTotal     *float64
	StockCut       *float64
	StockAvailable *float64
	CreatedAt      *string
}

func (row *Stock) Insert(db *sql.DB) {
	res, err := db.Exec("INSERT INTO stock (`product_id`, `stock_total`, `stock_cut`, `stock_available`, `created_at` ) values (?, ?, ?, ?, ?)", row.ProductId, row.StockTotal, row.StockCut, row.StockAvailable, row.CreatedAt)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	lastId, _ := res.LastInsertId()
	row.Id = &lastId

	defer db.Close()
}

func (row Stock) Update(db *sql.DB) {
	_, err := db.Exec("UPDATE stock SET product_id=?, stock_total=?, stock_cut=?, stock_available=?, created_at=? WHERE `id` = ?", row.ProductId, row.StockTotal, row.StockCut, row.StockAvailable, row.CreatedAt, row.Id)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	defer db.Close()
}

func (row Stock) Delete(db *sql.DB) {
	_, err := db.Exec("DELETE FROM stock WHERE `id` = ?", true, row.Id)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	defer db.Close()
}

func (row Stock) HasNew() bool {
	return row.Id == nil
}

func StockAll(filter StockFilter) []Stock {
	db := helpers.DatabaseInstance()
	var bind []any
	where := " where 0 = 0 "
	if filter.ProductId != nil {
		where += " and product_id = ? "
		bind = append(bind, filter.ProductId)
	}
	res, err := db.Query("select * from stock "+where+" order by id desc", bind...)
	if err != nil {
		log.Fatalln(err.Error())
	}
	rows := []Stock{}
	for res.Next() {
		row := Stock{}
		res.Scan(&row.Id, &row.ProductId, &row.StockTotal, &row.StockCut, &row.StockAvailable, &row.CreatedAt)
		rows = append(rows, row)
	}

	return rows
}

type StockFilter struct {
	ProductId *int64
}
