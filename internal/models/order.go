package models

import (
	"time"
)

type Order struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	ProductID  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice string    `json:"total_price"`
	OrderDate  time.Time `json:"order_date"`
}
