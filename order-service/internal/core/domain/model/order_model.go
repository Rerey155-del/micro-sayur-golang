package model

import "time"

type Order struct {
	ID           uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderCode    string      `gorm:"type:varchar(50);uniqueIndex;not null" json:"order_code"`
	UserID       int64       `gorm:"not null;index" json:"user_id"`
	CustomerName string      `gorm:"type:varchar(120);not null" json:"customer_name"`
	Status       string      `gorm:"type:varchar(30);not null;default:'Pending'" json:"status"`
	TotalAmount  float64     `gorm:"type:decimal(12,2);not null" json:"total_amount"`
	Address      string      `gorm:"type:text" json:"address"`
	Notes        string      `gorm:"type:text" json:"notes"`
	Items        []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"items"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID           uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID      uint    `gorm:"index;not null" json:"order_id"`
	ProductID    int64   `gorm:"not null" json:"product_id"`
	ProductName  string  `gorm:"type:varchar(120);not null" json:"product_name"`
	ProductPrice float64 `gorm:"type:decimal(12,2);not null" json:"product_price"`
	Quantity     int     `gorm:"not null" json:"quantity"`
	Subtotal     float64 `gorm:"type:decimal(12,2);not null" json:"subtotal"`
}
