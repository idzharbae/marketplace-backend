// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/transaction/internal (interfaces: OrderReader)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	request "github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	reflect "reflect"
)

// MockOrderReader is a mock of OrderReader interface.
type MockOrderReader struct {
	ctrl     *gomock.Controller
	recorder *MockOrderReaderMockRecorder
}

// MockOrderReaderMockRecorder is the mock recorder for MockOrderReader.
type MockOrderReaderMockRecorder struct {
	mock *MockOrderReader
}

// NewMockOrderReader creates a new mock instance.
func NewMockOrderReader(ctrl *gomock.Controller) *MockOrderReader {
	mock := &MockOrderReader{ctrl: ctrl}
	mock.recorder = &MockOrderReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderReader) EXPECT() *MockOrderReaderMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockOrderReader) GetByID(arg0 int64) (entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockOrderReaderMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockOrderReader)(nil).GetByID), arg0)
}

// ListByShopID mocks base method.
func (m *MockOrderReader) ListByShopID(arg0 int64, arg1 int32, arg2 request.Pagination) ([]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByShopID", arg0, arg1, arg2)
	ret0, _ := ret[0].([]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShopID indicates an expected call of ListByShopID.
func (mr *MockOrderReaderMockRecorder) ListByShopID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShopID", reflect.TypeOf((*MockOrderReader)(nil).ListByShopID), arg0, arg1, arg2)
}

// ListByUserID mocks base method.
func (m *MockOrderReader) ListByUserID(arg0 int64, arg1 int32, arg2 request.Pagination) ([]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", arg0, arg1, arg2)
	ret0, _ := ret[0].([]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockOrderReaderMockRecorder) ListByUserID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockOrderReader)(nil).ListByUserID), arg0, arg1, arg2)
}
