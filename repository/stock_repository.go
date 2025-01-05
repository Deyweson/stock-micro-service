package repository

import (
	"database/sql"
)

type StockRepository struct {
	connection *sql.DB
}

func NewStockRepository(conn *sql.DB) StockRepository {
	return StockRepository{
		connection: conn,
	}
}
