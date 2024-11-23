package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
)

type TransferRepo interface {
	Transfer(tx *gorm.DB, transfer *models.Transfer) (*models.Transfer, error)
}

type transferRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewTransferRepo(db *gorm.DB, globalRepo GlobalRepo) TransferRepo {
	return &transferRepo{db, globalRepo}
}

func (u *transferRepo) Transfer(tx *gorm.DB, transfer *models.Transfer) (*models.Transfer, error) {
	db := tx
	if db == nil {
		db = u.db
	}

	err := db.Create(&transfer).Error
	return transfer, err
}
