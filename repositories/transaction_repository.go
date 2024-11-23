package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepo interface {
	GetTransactions(userID string) (*[]models.Transaction, error)
	CreateTransaction(tx *gorm.DB, transaction *models.Transaction) (*models.Transaction, error)
}

type transactionRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewTransactionRepo(db *gorm.DB, globalRepo GlobalRepo) TransactionRepo {
	return &transactionRepo{db, globalRepo}
}

func (u *transactionRepo) GetTransactions(userID string) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	err := u.db.
		Order("created_date DESC").
		Preload(clause.Associations).
		Where("user_id = ?", userID).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return &transactions, err
}

func (u *transactionRepo) CreateTransaction(tx *gorm.DB, transaction *models.Transaction) (*models.Transaction, error) {
	db := tx
	if db == nil {
		db = u.db
	}

	err := db.Create(&transaction).Error
	return transaction, err
}
