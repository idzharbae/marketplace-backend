// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/transaction/internal (interfaces: CartUC)

// Package ucmock is a generated GoMock package.
package ucmock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	request "github.com/idzharbae/marketplace-backend/svc/transaction/internal/request"
	reflect "reflect"
)

// MockCartUC is a mock of CartUC interface.
type MockCartUC struct {
	ctrl     *gomock.Controller
	recorder *MockCartUCMockRecorder
}

// MockCartUCMockRecorder is the mock recorder for MockCartUC.
type MockCartUCMockRecorder struct {
	mock *MockCartUC
}

// NewMockCartUC creates a new mock instance.
func NewMockCartUC(ctrl *gomock.Controller) *MockCartUC {
	mock := &MockCartUC{ctrl: ctrl}
	mock.recorder = &MockCartUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartUC) EXPECT() *MockCartUCMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCartUC) Add(arg0 entity.Cart) (entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCartUCMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCartUC)(nil).Add), arg0)
}

// List mocks base method.
func (m *MockCartUC) List(arg0 request.ListCartReq) ([]entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCartUCMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCartUC)(nil).List), arg0)
}

// Remove mocks base method.
func (m *MockCartUC) Remove(arg0, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockCartUCMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCartUC)(nil).Remove), arg0, arg1)
}

// Update mocks base method.
func (m *MockCartUC) Update(arg0 entity.Cart) (entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCartUCMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCartUC)(nil).Update), arg0)
}
