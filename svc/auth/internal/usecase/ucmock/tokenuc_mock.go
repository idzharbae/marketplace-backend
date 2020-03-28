// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/auth/internal (interfaces: TokenUC)

// Package ucmock is a generated GoMock package.
package ucmock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	request "github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	reflect "reflect"
)

// MockTokenUC is a mock of TokenUC interface
type MockTokenUC struct {
	ctrl     *gomock.Controller
	recorder *MockTokenUCMockRecorder
}

// MockTokenUCMockRecorder is the mock recorder for MockTokenUC
type MockTokenUCMockRecorder struct {
	mock *MockTokenUC
}

// NewMockTokenUC creates a new mock instance
func NewMockTokenUC(ctrl *gomock.Controller) *MockTokenUC {
	mock := &MockTokenUC{ctrl: ctrl}
	mock.recorder = &MockTokenUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenUC) EXPECT() *MockTokenUCMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTokenUC) Get(arg0 request.GetToken) (entity.AuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(entity.AuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTokenUCMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTokenUC)(nil).Get), arg0)
}

// Refresh mocks base method
func (m *MockTokenUC) Refresh(arg0 request.RefreshToken) (entity.AuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0)
	ret0, _ := ret[0].(entity.AuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh
func (mr *MockTokenUCMockRecorder) Refresh(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockTokenUC)(nil).Refresh), arg0)
}
