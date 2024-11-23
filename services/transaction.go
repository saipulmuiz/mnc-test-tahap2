package services

import (
	"net/http"
	"strings"
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

func (u *TransactionService) GetTransactions(userID string) *params.Response {
	var (
		transactions *[]models.Transaction
		err          error
	)

	transactions, err = u.transactionRepo.GetTransactions(userID)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	var transactionData []map[string]interface{}
	for _, transaction := range *transactions {
		var dynamicKey string

		switch transaction.ReferenceType {
		case models.TRANSACTION_REFERENCE_TYPE_TOPUP:
			dynamicKey = "top_up_id"
		case models.TRANSACTION_REFERENCE_TYPE_PAYMENT:
			dynamicKey = "payment_id"
		case models.TRANSACTION_REFERENCE_TYPE_TRANSFER:
			dynamicKey = "transfer_id"
		default:
			dynamicKey = "unknown_id"
		}

		data := map[string]interface{}{
			dynamicKey:         transaction.ReferenceID,
			"status":           transaction.Status,
			"user_id":          transaction.UserID,
			"transaction_type": strings.ToUpper(transaction.Type),
			"amount":           transaction.Amount,
			"remarks":          transaction.Remarks,
			"balance_before":   transaction.BalanceBefore,
			"balance_after":    transaction.BalanceAfter,
			"created_date":     helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, transaction.CreatedDate),
		}

		transactionData = append(transactionData, data)
	}

	result := params.ResponseWithData{
		Status: "SUCCESS",
		Result: transactionData,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
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

func (u *TransactionService) Transfer(userID string, request params.TransferRequest) *params.Response {
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

	targetUser, _ := u.userRepo.FindById(request.TargetUser)
	if targetUser.UserID == "" {
		return helpers.HandleErrorService(http.StatusNotFound, "Target user not found")
	}

	transfer := &models.Transfer{
		TransferID:  uuid.NewString(),
		FromUserID:  userID,
		ToUserID:    request.TargetUser,
		Amount:      request.Amount,
		Remarks:     request.Remarks,
		Status:      models.TRANSFER_STATUS_SUCCESS,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	transfer, err := u.transferRepo.Transfer(tx, transfer)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	balanceAfter := user.Balance - request.Amount
	transaction := models.Transaction{
		TransactionID: uuid.NewString(),
		UserID:        userID,
		Type:          models.TRANSACTION_TYPE_DEBIT,
		ReferenceType: models.TRANSACTION_REFERENCE_TYPE_TRANSFER,
		ReferenceID:   transfer.TransferID,
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

	// update balance user sender
	_, err = u.userRepo.UpdateBalance(tx, user.UserID, balanceAfter)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	balanceAfterTarget := targetUser.Balance + request.Amount
	transactionTarget := models.Transaction{
		TransactionID: uuid.NewString(),
		UserID:        targetUser.UserID,
		Type:          models.TRANSACTION_TYPE_CREDIT,
		ReferenceType: models.TRANSACTION_REFERENCE_TYPE_TRANSFER,
		ReferenceID:   transfer.TransferID,
		Amount:        request.Amount,
		Remarks:       request.Remarks,
		BalanceBefore: targetUser.Balance,
		BalanceAfter:  balanceAfterTarget,
		Status:        models.TRANSACTION_STATUS_SUCCESS,
		CreatedDate:   time.Now(),
		UpdatedDate:   time.Now(),
	}

	_, err = u.transactionRepo.CreateTransaction(tx, &transactionTarget)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	// update balance user receiver
	_, err = u.userRepo.UpdateBalance(tx, targetUser.UserID, balanceAfterTarget)
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
		Data: params.TransferResponse{
			TransferID:    transfer.TransferID,
			Amount:        request.Amount,
			Remarks:       request.Remarks,
			BalanceBefore: user.Balance,
			BalanceAfter:  balanceAfter,
			CreatedDate:   helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, transfer.CreatedDate),
		},
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}
