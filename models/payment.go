package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PaymentID uuid.UUID `gorm:"type:uuid;primaryKey" json:"payment_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Amount    float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	Remarks   string    `gorm:"type:text" json:"remarks"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
