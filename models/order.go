package models

import "time"

type Order struct {
	ID         int       `gorm:"not null;uniqueIndex;primaryKey;" json:"id"`
	UserID     int       `json:"user_id"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
