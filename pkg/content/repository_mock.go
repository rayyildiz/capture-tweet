// Code generated by MockGen. DO NOT EDIT.
// Source: capturetweet.com/pkg/content (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen -package=content -self_package=capturetweet.com/pkg/content -destination=repository_mock.go . Repository
//

// Package content is a generated GoMock package.
package content

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
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

// ContactUs mocks base method.
func (m *MockRepository) ContactUs(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContactUs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContactUs indicates an expected call of ContactUs.
func (mr *MockRepositoryMockRecorder) ContactUs(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContactUs", reflect.TypeOf((*MockRepository)(nil).ContactUs), arg0, arg1, arg2, arg3)
}
