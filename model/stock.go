package model

import "time"

type Stock struct {
	Id            int       `json:"id" db:"id"`
	ShopId        string    `json:"shop_id" db:"shop_id"`
	ProdId        string    `json:"prod_id" db:"prod_id"`
	OperationType string    `json:"operation_type" db:"operation_type"`
	Amount        float64   `json:"amount" db:"amount"`
	Description   string    `json:"description" db:"description"`
	CreatedAt     time.Time `json:"created_at" db:"created_at" `
}
