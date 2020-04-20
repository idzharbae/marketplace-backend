// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/transaction/internal (interfaces: AuthGateway)

// Package gatewaymock is a generated GoMock package.
package gatewaymock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	reflect "reflect"
)

// MockAuthGateway is a mock of AuthGateway interface.
type MockAuthGateway struct {
	ctrl     *gomock.Controller
	recorder *MockAuthGatewayMockRecorder
}

// MockAuthGatewayMockRecorder is the mock recorder for MockAuthGateway.
type MockAuthGatewayMockRecorder struct {
	mock *MockAuthGateway
}

// NewMockAuthGateway creates a new mock instance.
func NewMockAuthGateway(ctrl *gomock.Controller) *MockAuthGateway {
	mock := &MockAuthGateway{ctrl: ctrl}
	mock.recorder = &MockAuthGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthGateway) EXPECT() *MockAuthGatewayMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method.
func (m *MockAuthGateway) GetUserByID(arg0 int64) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockAuthGatewayMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAuthGateway)(nil).GetUserByID), arg0)
}

// UpdateUserSaldo mocks base method.
func (m *MockAuthGateway) UpdateUserSaldo(arg0, arg1 int64) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserSaldo", arg0, arg1)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserSaldo indicates an expected call of UpdateUserSaldo.
func (mr *MockAuthGatewayMockRecorder) UpdateUserSaldo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSaldo", reflect.TypeOf((*MockAuthGateway)(nil).UpdateUserSaldo), arg0, arg1)
}
