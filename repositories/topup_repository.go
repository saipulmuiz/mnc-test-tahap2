package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
)

type TopupRepo interface {
	Topup(tx *gorm.DB, topup *models.Topup) (*models.Topup, error)
}

type topupRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewTopupRepo(db *gorm.DB, globalRepo GlobalRepo) TopupRepo {
	return &topupRepo{db, globalRepo}
}

func (u *topupRepo) Topup(tx *gorm.DB, topup *models.Topup) (*models.Topup, error) {
	db := tx
	if db == nil {
		db = u.db
	}

	err := db.Create(&topup).Error
	return topup, err
}
