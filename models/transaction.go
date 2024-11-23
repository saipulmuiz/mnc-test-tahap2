package models

import (
	"time"
)

type Transaction struct {
	TransactionID string    `gorm:"type:uuid;primaryKey" json:"transaction_id"`
	UserID        string    `gorm:"type:uuid;not null" json:"user_id"`
	Type          string    `gorm:"type:transaction_type;not null" json:"type"`
	ReferenceType string    `gorm:"type:uuid;not null" json:"reference_type"`
	ReferenceID   string    `gorm:"type:uuid;not null" json:"reference_id"`
	Amount        float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	Remarks       string    `gorm:"type:text" json:"remarks"`
	BalanceBefore float64   `gorm:"type:decimal(18,2);not null" json:"balance_before"`
	BalanceAfter  float64   `gorm:"type:decimal(18,2);not null" json:"balance_after"`
	Status        string    `gorm:"type:transaction_status;not null" json:"status"`
	CreatedDate   time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}
