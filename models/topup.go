package models

import (
	"time"
)

type Topup struct {
	TopUpID     string    `gorm:"type:uuid;primaryKey" json:"top_up_id"`
	UserID      string    `gorm:"type:uuid;not null" json:"user_id"`
	Amount      float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
}
