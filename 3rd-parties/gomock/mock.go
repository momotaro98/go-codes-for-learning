// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package app is a generated GoMock package.
package app

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUserRepository is a mock of IUserRepository interface
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// InsertAUser mocks base method
func (m *MockIUserRepository) InsertAUser(user *UserModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAUser indicates an expected call of InsertAUser
func (mr *MockIUserRepositoryMockRecorder) InsertAUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAUser", reflect.TypeOf((*MockIUserRepository)(nil).InsertAUser), user)
}
