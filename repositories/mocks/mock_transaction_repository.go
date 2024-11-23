// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/saipulmuiz/mnc-test-tahap2/repositories (interfaces: TransactionRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/saipulmuiz/mnc-test-tahap2/models"
	gorm "gorm.io/gorm"
)

// MockTransactionRepo is a mock of TransactionRepo interface.
type MockTransactionRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepoMockRecorder
}

// MockTransactionRepoMockRecorder is the mock recorder for MockTransactionRepo.
type MockTransactionRepoMockRecorder struct {
	mock *MockTransactionRepo
}

// NewMockTransactionRepo creates a new mock instance.
func NewMockTransactionRepo(ctrl *gomock.Controller) *MockTransactionRepo {
	mock := &MockTransactionRepo{ctrl: ctrl}
	mock.recorder = &MockTransactionRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepo) EXPECT() *MockTransactionRepoMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionRepo) CreateTransaction(arg0 *gorm.DB, arg1 *models.Transaction) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0, arg1)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionRepoMockRecorder) CreateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionRepo)(nil).CreateTransaction), arg0, arg1)
}

// GetTransactions mocks base method.
func (m *MockTransactionRepo) GetTransactions(arg0 string) (*[]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", arg0)
	ret0, _ := ret[0].(*[]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockTransactionRepoMockRecorder) GetTransactions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockTransactionRepo)(nil).GetTransactions), arg0)
}