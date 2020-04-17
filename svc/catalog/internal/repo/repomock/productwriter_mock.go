// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/catalog/internal (interfaces: ProductWriter)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	reflect "reflect"
)

// MockProductWriter is a mock of ProductWriter interface.
type MockProductWriter struct {
	ctrl     *gomock.Controller
	recorder *MockProductWriterMockRecorder
}

// MockProductWriterMockRecorder is the mock recorder for MockProductWriter.
type MockProductWriterMockRecorder struct {
	mock *MockProductWriter
}

// NewMockProductWriter creates a new mock instance.
func NewMockProductWriter(ctrl *gomock.Controller) *MockProductWriter {
	mock := &MockProductWriter{ctrl: ctrl}
	mock.recorder = &MockProductWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductWriter) EXPECT() *MockProductWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductWriter) Create(arg0 entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductWriterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductWriter)(nil).Create), arg0)
}

// DeleteByID mocks base method.
func (m *MockProductWriter) DeleteByID(arg0 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockProductWriterMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockProductWriter)(nil).DeleteByID), arg0)
}

// DeleteBySlug mocks base method.
func (m *MockProductWriter) DeleteBySlug(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBySlug", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBySlug indicates an expected call of DeleteBySlug.
func (mr *MockProductWriterMockRecorder) DeleteBySlug(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBySlug", reflect.TypeOf((*MockProductWriter)(nil).DeleteBySlug), arg0)
}

// Update mocks base method.
func (m *MockProductWriter) Update(arg0 entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductWriterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductWriter)(nil).Update), arg0)
}
