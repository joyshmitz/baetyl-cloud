// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/v2/service (interfaces: ApplicationService)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	models "github.com/baetyl/baetyl-cloud/v2/models"
	v1 "github.com/baetyl/baetyl-go/v2/spec/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockApplicationService) Create(arg0 interface{}, arg1 string, arg2 *v1.Application) (*v1.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockApplicationServiceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApplicationService)(nil).Create), arg0, arg1, arg2)
}

// CreateWithBase mocks base method.
func (m *MockApplicationService) CreateWithBase(arg0 interface{}, arg1 string, arg2, arg3 *v1.Application) (*v1.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithBase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithBase indicates an expected call of CreateWithBase.
func (mr *MockApplicationServiceMockRecorder) CreateWithBase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithBase", reflect.TypeOf((*MockApplicationService)(nil).CreateWithBase), arg0, arg1, arg2, arg3)
}

// Delete mocks base method.
func (m *MockApplicationService) Delete(arg0 interface{}, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockApplicationServiceMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApplicationService)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Get mocks base method.
func (m *MockApplicationService) Get(arg0, arg1, arg2 string) (*v1.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockApplicationServiceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationService)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockApplicationService) List(arg0 string, arg1 *models.ListOptions) (*models.ApplicationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*models.ApplicationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockApplicationServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApplicationService)(nil).List), arg0, arg1)
}

// ListByNames mocks base method.
func (m *MockApplicationService) ListByNames(arg0 string, arg1 []string) ([]models.AppItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByNames", arg0, arg1)
	ret0, _ := ret[0].([]models.AppItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByNames indicates an expected call of ListByNames.
func (mr *MockApplicationServiceMockRecorder) ListByNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByNames", reflect.TypeOf((*MockApplicationService)(nil).ListByNames), arg0, arg1)
}

// Update mocks base method.
func (m *MockApplicationService) Update(arg0 interface{}, arg1 string, arg2 *v1.Application) (*v1.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockApplicationServiceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockApplicationService)(nil).Update), arg0, arg1, arg2)
}
