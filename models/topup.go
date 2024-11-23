package models

import (
	"time"

	"github.com/google/uuid"
)

type TopUp struct {
	TopUpID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"top_up_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Amount    float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
