package models

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	TransferID uuid.UUID `gorm:"type:uuid;primaryKey" json:"transfer_id"`
	FromUserID uuid.UUID `gorm:"type:uuid;not null" json:"from_user_id"`
	ToUserID   uuid.UUID `gorm:"type:uuid;not null" json:"to_user_id"`
	Amount     float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	Remarks    string    `gorm:"type:text" json:"remarks"`
	Status     string    `gorm:"type:enum('PENDING', 'SUCCESS', 'FAILED');not null" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
