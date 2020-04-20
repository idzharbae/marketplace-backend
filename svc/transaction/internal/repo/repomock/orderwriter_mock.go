// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/transaction/internal (interfaces: OrderWriter)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/transaction/internal/entity"
	reflect "reflect"
)

// MockOrderWriter is a mock of OrderWriter interface.
type MockOrderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockOrderWriterMockRecorder
}

// MockOrderWriterMockRecorder is the mock recorder for MockOrderWriter.
type MockOrderWriterMockRecorder struct {
	mock *MockOrderWriter
}

// NewMockOrderWriter creates a new mock instance.
func NewMockOrderWriter(ctrl *gomock.Controller) *MockOrderWriter {
	mock := &MockOrderWriter{ctrl: ctrl}
	mock.recorder = &MockOrderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderWriter) EXPECT() *MockOrderWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderWriter) Create(arg0 entity.Order) (entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderWriterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderWriter)(nil).Create), arg0)
}

// DeleteByID mocks base method.
func (m *MockOrderWriter) DeleteByID(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockOrderWriterMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockOrderWriter)(nil).DeleteByID), arg0)
}

// Update mocks base method.
func (m *MockOrderWriter) Update(arg0 entity.Order) (entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrderWriterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderWriter)(nil).Update), arg0)
}
