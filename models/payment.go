package models

import (
	"time"
)

type Payment struct {
	PaymentID   string    `gorm:"type:uuid;primaryKey" json:"payment_id"`
	UserID      string    `gorm:"type:uuid;not null" json:"user_id"`
	Amount      float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	Remarks     string    `gorm:"type:text" json:"remarks"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}
