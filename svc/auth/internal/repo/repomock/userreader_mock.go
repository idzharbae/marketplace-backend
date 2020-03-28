// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/auth/internal (interfaces: UserReader)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	request "github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	reflect "reflect"
)

// MockUserReader is a mock of UserReader interface
type MockUserReader struct {
	ctrl     *gomock.Controller
	recorder *MockUserReaderMockRecorder
}

// MockUserReaderMockRecorder is the mock recorder for MockUserReader
type MockUserReaderMockRecorder struct {
	mock *MockUserReader
}

// NewMockUserReader creates a new mock instance
func NewMockUserReader(ctrl *gomock.Controller) *MockUserReader {
	mock := &MockUserReader{ctrl: ctrl}
	mock.recorder = &MockUserReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserReader) EXPECT() *MockUserReaderMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockUserReader) GetByID(arg0 int64) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUserReaderMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserReader)(nil).GetByID), arg0)
}

// ListAll mocks base method
func (m *MockUserReader) ListAll(arg0 request.ListUser) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockUserReaderMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockUserReader)(nil).ListAll), arg0)
}
