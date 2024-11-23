package services

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories"
	"gorm.io/gorm"
)

type TransactionService struct {
	transactionRepo repositories.TransactionRepo
	topupRepo       repositories.TopupRepo
	paymentRepo     repositories.PaymentRepo
	transferRepo    repositories.TransferRepo
	userRepo        repositories.UserRepo
	db              *gorm.DB
}

func NewTransactionService(
	transactionRepo repositories.TransactionRepo,
	topupRepo repositories.TopupRepo,
	paymentRepo repositories.PaymentRepo,
	transferRepo repositories.TransferRepo,
	userRepo repositories.UserRepo,
	db *gorm.DB,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		topupRepo:       topupRepo,
		paymentRepo:     paymentRepo,
		transferRepo:    transferRepo,
		userRepo:        userRepo,
		db:              db,
	}
}

func (u *TransactionService) Topup(userID string, request params.TopupRequest) *params.Response {
	// Start transaction
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, _ := u.userRepo.FindById(userID)
	if user.UserID == "" {
		return helpers.HandleErrorService(http.StatusNotFound, "User not found")
	}

	topup := &models.Topup{
		TopUpID:     uuid.NewString(),
		UserID:      userID,
		Amount:      request.Amount,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	topup, err := u.topupRepo.Topup(tx, topup)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	balanceAfter := user.Balance + request.Amount
	transaction := models.Transaction{
		TransactionID: uuid.NewString(),
		UserID:        userID,
		Type:          models.TRANSACTION_TYPE_CREDIT,
		ReferenceType: models.TRANSACTION_REFERENCE_TYPE_TOPUP,
		ReferenceID:   topup.TopUpID,
		Amount:        request.Amount,
		BalanceBefore: user.Balance,
		BalanceAfter:  balanceAfter,
		Status:        models.TRANSACTION_STATUS_SUCCESS,
		CreatedDate:   time.Now(),
		UpdatedDate:   time.Now(),
	}

	_, err = u.transactionRepo.CreateTransaction(tx, &transaction)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	_, err = u.userRepo.UpdateBalance(tx, user.UserID, balanceAfter)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, "Failed to commit transaction")
	}

	result := params.ResponseSuccess{
		Status: "SUCCESS",
		Data: params.TopupResponse{
			TopupID:       topup.TopUpID,
			AmountTopup:   request.Amount,
			BalanceBefore: user.Balance,
			BalanceAfter:  balanceAfter,
			CreatedDate:   helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, topup.CreatedDate),
		},
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *TransactionService) Payment(userID string, request params.PaymentRequest) *params.Response {
	// Start transaction
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, _ := u.userRepo.FindById(userID)
	if user.UserID == "" {
		return helpers.HandleErrorService(http.StatusNotFound, "User not found")
	}

	if user.Balance < request.Amount {
		return helpers.HandleErrorService(http.StatusBadRequest, "Balance is not enough")
	}

	payment := &models.Payment{
		PaymentID:   uuid.NewString(),
		UserID:      userID,
		Amount:      request.Amount,
		Remarks:     request.Remarks,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	payment, err := u.paymentRepo.Payment(tx, payment)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	balanceAfter := user.Balance - request.Amount
	transaction := models.Transaction{
		TransactionID: uuid.NewString(),
		UserID:        userID,
		Type:          models.TRANSACTION_TYPE_DEBIT,
		ReferenceType: models.TRANSACTION_REFERENCE_TYPE_PAYMENT,
		ReferenceID:   payment.PaymentID,
		Amount:        request.Amount,
		Remarks:       request.Remarks,
		BalanceBefore: user.Balance,
		BalanceAfter:  balanceAfter,
		Status:        models.TRANSACTION_STATUS_SUCCESS,
		CreatedDate:   time.Now(),
		UpdatedDate:   time.Now(),
	}

	_, err = u.transactionRepo.CreateTransaction(tx, &transaction)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	_, err = u.userRepo.UpdateBalance(tx, user.UserID, balanceAfter)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, "Failed to commit transaction")
	}

	result := params.ResponseSuccess{
		Status: "SUCCESS",
		Data: params.PaymentResponse{
			PaymentID:     payment.PaymentID,
			Amount:        request.Amount,
			Remarks:       request.Remarks,
			BalanceBefore: user.Balance,
			BalanceAfter:  balanceAfter,
			CreatedDate:   helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, payment.CreatedDate),
		},
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}
