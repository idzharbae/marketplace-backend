// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/auth/internal (interfaces: UserWriter)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	authproto "github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	entity "github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	request "github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	reflect "reflect"
)

// MockUserWriter is a mock of UserWriter interface.
type MockUserWriter struct {
	ctrl     *gomock.Controller
	recorder *MockUserWriterMockRecorder
}

// MockUserWriterMockRecorder is the mock recorder for MockUserWriter.
type MockUserWriterMockRecorder struct {
	mock *MockUserWriter
}

// NewMockUserWriter creates a new mock instance.
func NewMockUserWriter(ctrl *gomock.Controller) *MockUserWriter {
	mock := &MockUserWriter{ctrl: ctrl}
	mock.recorder = &MockUserWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserWriter) EXPECT() *MockUserWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserWriter) Create(arg0 entity.User) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserWriterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserWriter)(nil).Create), arg0)
}

// DeleteByID mocks base method.
func (m *MockUserWriter) DeleteByID(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockUserWriterMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockUserWriter)(nil).DeleteByID), arg0)
}

// TransferSaldo mocks base method.
func (m *MockUserWriter) TransferSaldo(arg0 request.Transfer) (authproto.TransferSaldoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferSaldo", arg0)
	ret0, _ := ret[0].(authproto.TransferSaldoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferSaldo indicates an expected call of TransferSaldo.
func (mr *MockUserWriterMockRecorder) TransferSaldo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferSaldo", reflect.TypeOf((*MockUserWriter)(nil).TransferSaldo), arg0)
}

// Update mocks base method.
func (m *MockUserWriter) Update(arg0 entity.User) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserWriterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserWriter)(nil).Update), arg0)
}

// UpdateSaldo mocks base method.
func (m *MockUserWriter) UpdateSaldo(arg0 request.TopUp) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSaldo", arg0)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSaldo indicates an expected call of UpdateSaldo.
func (mr *MockUserWriterMockRecorder) UpdateSaldo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSaldo", reflect.TypeOf((*MockUserWriter)(nil).UpdateSaldo), arg0)
}
