package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID uuid.UUID `gorm:"type:uuid;primaryKey" json:"transaction_id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Type          string    `gorm:"type:enum('credit', 'debit');not null" json:"type"`
	Amount        float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	Remarks       string    `gorm:"type:text" json:"remarks"`
	BalanceBefore float64   `gorm:"type:decimal(18,2);not null" json:"balance_before"`
	BalanceAfter  float64   `gorm:"type:decimal(18,2);not null" json:"balance_after"`
	Status        string    `gorm:"type:enum('PENDING', 'SUCCESS', 'FAILED');not null" json:"status"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}
