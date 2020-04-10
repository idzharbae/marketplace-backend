// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/idzharbae/marketplace-backend/svc/resources/internal (interfaces: FileUC)

// Package ucmock is a generated GoMock package.
package ucmock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/marketplace-backend/svc/resources/internal/entity"
	reflect "reflect"
)

// MockFileUC is a mock of FileUC interface.
type MockFileUC struct {
	ctrl     *gomock.Controller
	recorder *MockFileUCMockRecorder
}

// MockFileUCMockRecorder is the mock recorder for MockFileUC.
type MockFileUCMockRecorder struct {
	mock *MockFileUC
}

// NewMockFileUC creates a new mock instance.
func NewMockFileUC(ctrl *gomock.Controller) *MockFileUC {
	mock := &MockFileUC{ctrl: ctrl}
	mock.recorder = &MockFileUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileUC) EXPECT() *MockFileUCMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockFileUC) DeleteFile(arg0 entity.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockFileUCMockRecorder) DeleteFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileUC)(nil).DeleteFile), arg0)
}

// UploadFile mocks base method.
func (m *MockFileUC) UploadFile(arg0 entity.File) (entity.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0)
	ret0, _ := ret[0].(entity.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockFileUCMockRecorder) UploadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileUC)(nil).UploadFile), arg0)
}
