package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
)

type PaymentRepo interface {
	Payment(tx *gorm.DB, payment *models.Payment) (*models.Payment, error)
}

type paymentRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewPaymentRepo(db *gorm.DB, globalRepo GlobalRepo) PaymentRepo {
	return &paymentRepo{db, globalRepo}
}

func (u *paymentRepo) Payment(tx *gorm.DB, payment *models.Payment) (*models.Payment, error) {
	db := tx
	if db == nil {
		db = u.db
	}

	err := db.Create(&payment).Error
	return payment, err
}
