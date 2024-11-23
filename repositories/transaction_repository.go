package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepo interface {
	GetTransactions(page, size int) (*[]models.Transaction, int64, error)
	FindById(transactionId int) (*models.Transaction, error)
	CreateTransaction(tx *gorm.DB, transaction *models.Transaction) (*models.Transaction, error)
}

type transactionRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewTransactionRepo(db *gorm.DB, globalRepo GlobalRepo) TransactionRepo {
	return &transactionRepo{db, globalRepo}
}

func (u *transactionRepo) GetTransactions(page, size int) (*[]models.Transaction, int64, error) {
	var (
		transactions []models.Transaction
		count        int64
	)
	err := u.db.
		Order("created_at DESC").
		Preload(clause.Associations).
		Scopes(u.globalRepo.Paginate(page, size)).
		Find(&transactions).Error

	if err != nil {
		return nil, count, err
	}

	err = u.db.
		Model(&transactions).
		Count(&count).Error

	return &transactions, count, err
}

func (u *transactionRepo) FindById(transactionId int) (*models.Transaction, error) {
	var transaction models.Transaction
	err := u.db.Where("transaction_id = ?", transactionId).First(&transaction).Error
	return &transaction, err
}

func (u *transactionRepo) CreateTransaction(tx *gorm.DB, transaction *models.Transaction) (*models.Transaction, error) {
	db := tx
	if db == nil {
		db = u.db
	}

	err := db.Create(&transaction).Error
	return transaction, err
}
