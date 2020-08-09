// Code generated by MockGen. DO NOT EDIT.
// Source: com.capturetweet/pkg/tweet (interfaces: Repository)

// Package tweet is a generated GoMock package.
package tweet

import (
	context "context"
	reflect "reflect"

	anaconda "github.com/ChimeraCoder/anaconda"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Exist mocks base method.
func (m *MockRepository) Exist(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist.
func (mr *MockRepositoryMockRecorder) Exist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockRepository)(nil).Exist), arg0, arg1)
}

// FindAllOrderByUpdated mocks base method.
func (m *MockRepository) FindAllOrderByUpdated(arg0 context.Context, arg1 int) ([]Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllOrderByUpdated", arg0, arg1)
	ret0, _ := ret[0].([]Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllOrderByUpdated indicates an expected call of FindAllOrderByUpdated.
func (mr *MockRepositoryMockRecorder) FindAllOrderByUpdated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllOrderByUpdated", reflect.TypeOf((*MockRepository)(nil).FindAllOrderByUpdated), arg0, arg1)
}

// FindById mocks base method.
func (m *MockRepository) FindById(arg0 context.Context, arg1 string) (*Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockRepositoryMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepository)(nil).FindById), arg0, arg1)
}

// FindByIds mocks base method.
func (m *MockRepository) FindByIds(arg0 context.Context, arg1 []string) ([]Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIds", arg0, arg1)
	ret0, _ := ret[0].([]Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIds indicates an expected call of FindByIds.
func (mr *MockRepositoryMockRecorder) FindByIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIds", reflect.TypeOf((*MockRepository)(nil).FindByIds), arg0, arg1)
}

// FindByUser mocks base method.
func (m *MockRepository) FindByUser(arg0 context.Context, arg1 string) ([]Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0, arg1)
	ret0, _ := ret[0].([]Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser.
func (mr *MockRepositoryMockRecorder) FindByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockRepository)(nil).FindByUser), arg0, arg1)
}

// Store mocks base method.
func (m *MockRepository) Store(arg0 context.Context, arg1 *anaconda.Tweet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockRepositoryMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRepository)(nil).Store), arg0, arg1)
}

// UpdateLargeImage mocks base method.
func (m *MockRepository) UpdateLargeImage(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLargeImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLargeImage indicates an expected call of UpdateLargeImage.
func (mr *MockRepositoryMockRecorder) UpdateLargeImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLargeImage", reflect.TypeOf((*MockRepository)(nil).UpdateLargeImage), arg0, arg1, arg2)
}

// UpdateThumbImage mocks base method.
func (m *MockRepository) UpdateThumbImage(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThumbImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateThumbImage indicates an expected call of UpdateThumbImage.
func (mr *MockRepositoryMockRecorder) UpdateThumbImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThumbImage", reflect.TypeOf((*MockRepository)(nil).UpdateThumbImage), arg0, arg1, arg2)
}
