// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/catalog/internal (interfaces: ProductReader)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	requests "github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	reflect "reflect"
)

// MockProductReader is a mock of ProductReader interface
type MockProductReader struct {
	ctrl     *gomock.Controller
	recorder *MockProductReaderMockRecorder
}

// MockProductReaderMockRecorder is the mock recorder for MockProductReader
type MockProductReaderMockRecorder struct {
	mock *MockProductReader
}

// NewMockProductReader creates a new mock instance
func NewMockProductReader(ctrl *gomock.Controller) *MockProductReader {
	mock := &MockProductReader{ctrl: ctrl}
	mock.recorder = &MockProductReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductReader) EXPECT() *MockProductReaderMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockProductReader) GetByID(arg0 int32) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockProductReaderMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProductReader)(nil).GetByID), arg0)
}

// GetBySlug mocks base method
func (m *MockProductReader) GetBySlug(arg0 string) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", arg0)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug
func (mr *MockProductReaderMockRecorder) GetBySlug(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockProductReader)(nil).GetBySlug), arg0)
}

// ListAll mocks base method
func (m *MockProductReader) ListAll(arg0 requests.Pagination) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockProductReaderMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockProductReader)(nil).ListAll), arg0)
}

// ListByShopID mocks base method
func (m *MockProductReader) ListByShopID(arg0 int64, arg1 requests.Pagination) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByShopID", arg0, arg1)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShopID indicates an expected call of ListByShopID
func (mr *MockProductReaderMockRecorder) ListByShopID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShopID", reflect.TypeOf((*MockProductReader)(nil).ListByShopID), arg0, arg1)
}
