// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/saipulmuiz/mnc-test-tahap2/repositories (interfaces: ProductRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/saipulmuiz/mnc-test-tahap2/models"
	gorm "gorm.io/gorm"
)

// MockProductRepo is a mock of ProductRepo interface.
type MockProductRepo struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepoMockRecorder
}

// MockProductRepoMockRecorder is the mock recorder for MockProductRepo.
type MockProductRepoMockRecorder struct {
	mock *MockProductRepo
}

// NewMockProductRepo creates a new mock instance.
func NewMockProductRepo(ctrl *gomock.Controller) *MockProductRepo {
	mock := &MockProductRepo{ctrl: ctrl}
	mock.recorder = &MockProductRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepo) EXPECT() *MockProductRepoMockRecorder {
	return m.recorder
}

// CheckProductByID mocks base method.
func (m *MockProductRepo) CheckProductByID(arg0 int, arg1 *models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckProductByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckProductByID indicates an expected call of CheckProductByID.
func (mr *MockProductRepoMockRecorder) CheckProductByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckProductByID", reflect.TypeOf((*MockProductRepo)(nil).CheckProductByID), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockProductRepo) CreateProduct(arg0 *models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductRepoMockRecorder) CreateProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductRepo)(nil).CreateProduct), arg0)
}

// DeleteProduct mocks base method.
func (m *MockProductRepo) DeleteProduct(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductRepoMockRecorder) DeleteProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductRepo)(nil).DeleteProduct), arg0)
}

// FindById mocks base method.
func (m *MockProductRepo) FindById(arg0 int) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockProductRepoMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockProductRepo)(nil).FindById), arg0)
}

// GetProducts mocks base method.
func (m *MockProductRepo) GetProducts(arg0, arg1 int) (*[]models.Product, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", arg0, arg1)
	ret0, _ := ret[0].(*[]models.Product)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductRepoMockRecorder) GetProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductRepo)(nil).GetProducts), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockProductRepo) UpdateProduct(arg0 *gorm.DB, arg1 int, arg2 *models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductRepoMockRecorder) UpdateProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductRepo)(nil).UpdateProduct), arg0, arg1, arg2)
}
