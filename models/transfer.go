package models

import (
	"time"
)

type Transfer struct {
	TransferID  string    `gorm:"type:uuid;primaryKey" json:"transfer_id"`
	FromUserID  string    `gorm:"type:uuid;not null" json:"from_user_id"`
	ToUserID    string    `gorm:"type:uuid;not null" json:"to_user_id"`
	Amount      float64   `gorm:"type:decimal(18,2);not null" json:"amount"`
	Remarks     string    `gorm:"type:text" json:"remarks"`
	Status      string    `gorm:"type:transfer_status;not null" json:"status"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}
