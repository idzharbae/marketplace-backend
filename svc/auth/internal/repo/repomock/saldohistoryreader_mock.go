// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/auth/internal (interfaces: SaldoHistoryReader)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	reflect "reflect"
)

// MockSaldoHistoryReader is a mock of SaldoHistoryReader interface.
type MockSaldoHistoryReader struct {
	ctrl     *gomock.Controller
	recorder *MockSaldoHistoryReaderMockRecorder
}

// MockSaldoHistoryReaderMockRecorder is the mock recorder for MockSaldoHistoryReader.
type MockSaldoHistoryReaderMockRecorder struct {
	mock *MockSaldoHistoryReader
}

// NewMockSaldoHistoryReader creates a new mock instance.
func NewMockSaldoHistoryReader(ctrl *gomock.Controller) *MockSaldoHistoryReader {
	mock := &MockSaldoHistoryReader{ctrl: ctrl}
	mock.recorder = &MockSaldoHistoryReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaldoHistoryReader) EXPECT() *MockSaldoHistoryReaderMockRecorder {
	return m.recorder
}

// ListByUserID mocks base method.
func (m *MockSaldoHistoryReader) ListByUserID(arg0 int64) ([]entity.SaldoHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", arg0)
	ret0, _ := ret[0].([]entity.SaldoHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockSaldoHistoryReaderMockRecorder) ListByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockSaldoHistoryReader)(nil).ListByUserID), arg0)
}
