package entity

import "time"

type OrderEntity struct {
	ID           int64              `json:"id"`
	OrderCode    string             `json:"order_code"`
	UserID       int64              `json:"user_id"`
	CustomerName string             `json:"customer_name"`
	Status       string             `json:"status"`
	TotalAmount  float64            `json:"total_amount"`
	Address      string             `json:"address"`
	Notes        string             `json:"notes"`
	Items        []*OrderItemEntity `json:"items"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

type OrderItemEntity struct {
	ID           int64   `json:"id"`
	OrderID      int64   `json:"order_id"`
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
	Subtotal     float64 `json:"subtotal"`
}
