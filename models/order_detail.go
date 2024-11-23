package models

import "time"

type OrderDetail struct {
	ID          int       `gorm:"not null;uniqueIndex;primaryKey;" json:"id"`
	OrderID     int       `json:"order_id"`
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Qty         int       `json:"qty"`
	Total       float64   `json:"total"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
