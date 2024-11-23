package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID      string    `gorm:"type:uuid;primaryKey" json:"user_id"`
	FirstName   string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName    string    `gorm:"type:varchar(100);not null" json:"last_name"`
	PhoneNumber string    `gorm:"type:varchar(16);unique;not null" json:"phone_number"`
	Address     string    `gorm:"type:text" json:"address"`
	PIN         string    `gorm:"type:varchar;not null" json:"-"`
	Balance     float64   `gorm:"type:decimal(18,2);default:0" json:"balance"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.PIN), 8)
	u.PIN = string(hash)
	return
}
