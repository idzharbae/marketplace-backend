// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/resources/internal (interfaces: OwnershipWriter)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	reflect "reflect"
)

// MockOwnershipWriter is a mock of OwnershipWriter interface.
type MockOwnershipWriter struct {
	ctrl     *gomock.Controller
	recorder *MockOwnershipWriterMockRecorder
}

// MockOwnershipWriterMockRecorder is the mock recorder for MockOwnershipWriter.
type MockOwnershipWriterMockRecorder struct {
	mock *MockOwnershipWriter
}

// NewMockOwnershipWriter creates a new mock instance.
func NewMockOwnershipWriter(ctrl *gomock.Controller) *MockOwnershipWriter {
	mock := &MockOwnershipWriter{ctrl: ctrl}
	mock.recorder = &MockOwnershipWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOwnershipWriter) EXPECT() *MockOwnershipWriterMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockOwnershipWriter) DeleteByID(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockOwnershipWriterMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockOwnershipWriter)(nil).DeleteByID), arg0)
}

// Save mocks base method.
func (m *MockOwnershipWriter) Save(arg0 entity.File) (entity.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(entity.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockOwnershipWriterMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOwnershipWriter)(nil).Save), arg0)
}