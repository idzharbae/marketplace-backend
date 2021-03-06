// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/auth/internal (interfaces: SaldoHistoryUC)

// Package ucmock is a generated GoMock package.
package ucmock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	request "github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	reflect "reflect"
)

// MockSaldoHistoryUC is a mock of SaldoHistoryUC interface.
type MockSaldoHistoryUC struct {
	ctrl     *gomock.Controller
	recorder *MockSaldoHistoryUCMockRecorder
}

// MockSaldoHistoryUCMockRecorder is the mock recorder for MockSaldoHistoryUC.
type MockSaldoHistoryUCMockRecorder struct {
	mock *MockSaldoHistoryUC
}

// NewMockSaldoHistoryUC creates a new mock instance.
func NewMockSaldoHistoryUC(ctrl *gomock.Controller) *MockSaldoHistoryUC {
	mock := &MockSaldoHistoryUC{ctrl: ctrl}
	mock.recorder = &MockSaldoHistoryUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaldoHistoryUC) EXPECT() *MockSaldoHistoryUCMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSaldoHistoryUC) Create(arg0 entity.SaldoHistory) (entity.SaldoHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.SaldoHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSaldoHistoryUCMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSaldoHistoryUC)(nil).Create), arg0)
}

// List mocks base method.
func (m *MockSaldoHistoryUC) List(arg0 request.ListSaldoHistory) ([]entity.SaldoHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]entity.SaldoHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSaldoHistoryUCMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSaldoHistoryUC)(nil).List), arg0)
}
