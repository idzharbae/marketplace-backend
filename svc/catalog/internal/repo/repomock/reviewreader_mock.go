// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/catalog/internal (interfaces: ReviewReader)

// Package repomock is a generated GoMock package.
package repomock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	requests "github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	reflect "reflect"
)

// MockReviewReader is a mock of ReviewReader interface.
type MockReviewReader struct {
	ctrl     *gomock.Controller
	recorder *MockReviewReaderMockRecorder
}

// MockReviewReaderMockRecorder is the mock recorder for MockReviewReader.
type MockReviewReaderMockRecorder struct {
	mock *MockReviewReader
}

// NewMockReviewReader creates a new mock instance.
func NewMockReviewReader(ctrl *gomock.Controller) *MockReviewReader {
	mock := &MockReviewReader{ctrl: ctrl}
	mock.recorder = &MockReviewReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewReader) EXPECT() *MockReviewReaderMockRecorder {
	return m.recorder
}

// GetByCustomerIDAndProductID mocks base method.
func (m *MockReviewReader) GetByCustomerIDAndProductID(arg0, arg1 int64) (entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCustomerIDAndProductID", arg0, arg1)
	ret0, _ := ret[0].(entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCustomerIDAndProductID indicates an expected call of GetByCustomerIDAndProductID.
func (mr *MockReviewReaderMockRecorder) GetByCustomerIDAndProductID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCustomerIDAndProductID", reflect.TypeOf((*MockReviewReader)(nil).GetByCustomerIDAndProductID), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockReviewReader) GetByID(arg0 int64) (entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockReviewReaderMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockReviewReader)(nil).GetByID), arg0)
}

// ListByProductID mocks base method.
func (m *MockReviewReader) ListByProductID(arg0 int64, arg1 requests.Pagination) ([]entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByProductID", arg0, arg1)
	ret0, _ := ret[0].([]entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByProductID indicates an expected call of ListByProductID.
func (mr *MockReviewReaderMockRecorder) ListByProductID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByProductID", reflect.TypeOf((*MockReviewReader)(nil).ListByProductID), arg0, arg1)
}

// ListByShopID mocks base method.
func (m *MockReviewReader) ListByShopID(arg0 int64, arg1 requests.Pagination) ([]entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByShopID", arg0, arg1)
	ret0, _ := ret[0].([]entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShopID indicates an expected call of ListByShopID.
func (mr *MockReviewReaderMockRecorder) ListByShopID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShopID", reflect.TypeOf((*MockReviewReader)(nil).ListByShopID), arg0, arg1)
}
