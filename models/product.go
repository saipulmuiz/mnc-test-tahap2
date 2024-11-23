package models

import "time"

type Product struct {
	ProductID    int       `gorm:"not null;uniqueIndex;primaryKey;" json:"product_id"`
	ProductName  string    `json:"product_name"`
	Description  string    `json:"description"`
	ProductPrice float64   `json:"product_price"`
	Stock        int       `json:"stock"`
	MinStock     int       `json:"min_stock"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
