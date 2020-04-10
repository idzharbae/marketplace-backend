// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/resources/internal (interfaces: OwnershipReader)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	reflect "reflect"
)

// MockOwnershipReader is a mock of OwnershipReader interface.
type MockOwnershipReader struct {
	ctrl     *gomock.Controller
	recorder *MockOwnershipReaderMockRecorder
}

// MockOwnershipReaderMockRecorder is the mock recorder for MockOwnershipReader.
type MockOwnershipReaderMockRecorder struct {
	mock *MockOwnershipReader
}

// NewMockOwnershipReader creates a new mock instance.
func NewMockOwnershipReader(ctrl *gomock.Controller) *MockOwnershipReader {
	mock := &MockOwnershipReader{ctrl: ctrl}
	mock.recorder = &MockOwnershipReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOwnershipReader) EXPECT() *MockOwnershipReaderMockRecorder {
	return m.recorder
}

// GetByURL mocks base method.
func (m *MockOwnershipReader) GetByURL(arg0 string) (entity.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByURL", arg0)
	ret0, _ := ret[0].(entity.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByURL indicates an expected call of GetByURL.
func (mr *MockOwnershipReaderMockRecorder) GetByURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByURL", reflect.TypeOf((*MockOwnershipReader)(nil).GetByURL), arg0)
}