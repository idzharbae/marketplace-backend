// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/catalog/internal (interfaces: ProductUC)

// Package ucmock is a generated GoMock package.
package ucmock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/catalog/internal/entity"
	requests "github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
	reflect "reflect"
)

// MockProductUC is a mock of ProductUC interface.
type MockProductUC struct {
	ctrl     *gomock.Controller
	recorder *MockProductUCMockRecorder
}

// MockProductUCMockRecorder is the mock recorder for MockProductUC.
type MockProductUCMockRecorder struct {
	mock *MockProductUC
}

// NewMockProductUC creates a new mock instance.
func NewMockProductUC(ctrl *gomock.Controller) *MockProductUC {
	mock := &MockProductUC{ctrl: ctrl}
	mock.recorder = &MockProductUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUC) EXPECT() *MockProductUCMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductUC) Create(arg0 entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductUCMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductUC)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockProductUC) Delete(arg0 entity.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductUCMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductUC)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockProductUC) Get(arg0 entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductUCMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductUC)(nil).Get), arg0)
}

// GetTotal mocks base method.
func (m *MockProductUC) GetTotal(arg0 int32) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotal", arg0)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotal indicates an expected call of GetTotal.
func (mr *MockProductUCMockRecorder) GetTotal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotal", reflect.TypeOf((*MockProductUC)(nil).GetTotal), arg0)
}

// List mocks base method.
func (m *MockProductUC) List(arg0 requests.ListProduct) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductUCMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductUC)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockProductUC) Update(arg0 entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductUCMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductUC)(nil).Update), arg0)
}
