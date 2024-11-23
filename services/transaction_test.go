package services

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_TransactionService_GetTransactions(t *testing.T) {
	type testCase struct {
		name              string
		wantError         bool
		expectedResponse  *params.Response
		userID            string
		onGetTransactions func(mock *mocks.MockTransactionRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		userID:    "123e4567-e89b-12d3-a456-426614174000",
		onGetTransactions: func(mock *mocks.MockTransactionRepo) {
			mock.EXPECT().GetTransactions("123e4567-e89b-12d3-a456-426614174000").Return(&[]models.Transaction{
				{
					UserID:        "123e4567-e89b-12d3-a456-426614174000",
					Type:          models.TRANSACTION_TYPE_DEBIT,
					ReferenceType: models.TRANSACTION_REFERENCE_TYPE_TOPUP,
					ReferenceID:   "topup-123",
					Amount:        1000,
					Remarks:       "Top up test",
					BalanceBefore: 5000,
					BalanceAfter:  6000,
					CreatedDate:   time.Now(),
					Status:        "SUCCESS",
				},
			}, nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusOK,
			Payload: params.ResponseWithData{
				Status: "SUCCESS",
				Result: []map[string]interface{}{
					{
						"top_up_id":        "topup-123",
						"status":           "SUCCESS",
						"user_id":          "123e4567-e89b-12d3-a456-426614174000",
						"transaction_type": "DEBIT",
						"amount":           1000.0,
						"remarks":          "Top up test",
						"balance_before":   5000.0,
						"balance_after":    6000.0,
						"created_date":     helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, time.Now()),
					},
				},
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "error getting transactions",
		wantError: true,
		userID:    "123e4567-e89b-12d3-a456-426614174000",
		onGetTransactions: func(mock *mocks.MockTransactionRepo) {
			mock.EXPECT().GetTransactions("123e4567-e89b-12d3-a456-426614174000").Return(nil, errors.New("failed to retrieve transactions"))
		},
		expectedResponse: &params.Response{
			Status: http.StatusBadRequest,
			Payload: params.ResponseErrorMessage{
				Message: "failed to retrieve transactions",
			},
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			transactionRepo := mocks.NewMockTransactionRepo(mockCtrl)

			if tc.onGetTransactions != nil {
				tc.onGetTransactions(transactionRepo)
			}

			service := &TransactionService{
				transactionRepo: transactionRepo,
			}

			resp := service.GetTransactions(tc.userID)

			if tc.wantError {
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResponse.Status, resp.Status)
				assert.Equal(t, tc.expectedResponse.Payload, resp.Payload)
			} else {
				assert.Equal(t, http.StatusOK, resp.Status)
				assert.Equal(t, tc.expectedResponse.Payload, resp.Payload)
			}
		})
	}
}
