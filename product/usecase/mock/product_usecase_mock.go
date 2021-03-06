// Code generated by MockGen. DO NOT EDIT.
// Source: product_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	product "github.com/goweb3/product"
	reflect "reflect"
)

// MockProductUsecase is a mock of ProductUsecase interface
type MockProductUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUsecaseMockRecorder
}

// MockProductUsecaseMockRecorder is the mock recorder for MockProductUsecase
type MockProductUsecaseMockRecorder struct {
	mock *MockProductUsecase
}

// NewMockProductUsecase creates a new mock instance
func NewMockProductUsecase(ctrl *gomock.Controller) *MockProductUsecase {
	mock := &MockProductUsecase{ctrl: ctrl}
	mock.recorder = &MockProductUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductUsecase) EXPECT() *MockProductUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockProductUsecase) Create(title, description string, userID int64) (*product.Product, error) {
	ret := m.ctrl.Call(m, "Create", title, description, userID)
	ret0, _ := ret[0].(*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProductUsecaseMockRecorder) Create(title, description, userID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductUsecase)(nil).Create), title, description, userID)
}

// GetByID mocks base method
func (m *MockProductUsecase) GetByID(id int64) (*product.Product, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockProductUsecaseMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProductUsecase)(nil).GetByID), id)
}

// GetByTitle mocks base method
func (m *MockProductUsecase) GetByTitle(title string) (*product.Product, error) {
	ret := m.ctrl.Call(m, "GetByTitle", title)
	ret0, _ := ret[0].(*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTitle indicates an expected call of GetByTitle
func (mr *MockProductUsecaseMockRecorder) GetByTitle(title interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTitle", reflect.TypeOf((*MockProductUsecase)(nil).GetByTitle), title)
}

// Update mocks base method
func (m *MockProductUsecase) Update(productID int64, title, description string) (*product.Product, error) {
	ret := m.ctrl.Call(m, "Update", productID, title, description)
	ret0, _ := ret[0].(*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockProductUsecaseMockRecorder) Update(productID, title, description interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductUsecase)(nil).Update), productID, title, description)
}

// Delete mocks base method
func (m *MockProductUsecase) Delete(id int64) error {
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProductUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductUsecase)(nil).Delete), id)
}

// Fetch mocks base method
func (m *MockProductUsecase) Fetch(offset, limit int64) ([]*product.Product, error) {
	ret := m.ctrl.Call(m, "Fetch", offset, limit)
	ret0, _ := ret[0].([]*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockProductUsecaseMockRecorder) Fetch(offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProductUsecase)(nil).Fetch), offset, limit)
}
