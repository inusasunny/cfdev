// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/start (interfaces: Env)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEnv is a mock of Env interface
type MockEnv struct {
	ctrl     *gomock.Controller
	recorder *MockEnvMockRecorder
}

// MockEnvMockRecorder is the mock recorder for MockEnv
type MockEnvMockRecorder struct {
	mock *MockEnv
}

// NewMockEnv creates a new mock instance
func NewMockEnv(ctrl *gomock.Controller) *MockEnv {
	mock := &MockEnv{ctrl: ctrl}
	mock.recorder = &MockEnvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnv) EXPECT() *MockEnvMockRecorder {
	return m.recorder
}

// CreateDirs mocks base method
func (m *MockEnv) CreateDirs() error {
	ret := m.ctrl.Call(m, "CreateDirs")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDirs indicates an expected call of CreateDirs
func (mr *MockEnvMockRecorder) CreateDirs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDirs", reflect.TypeOf((*MockEnv)(nil).CreateDirs))
}

// SetupState mocks base method
func (m *MockEnv) SetupState(arg0 string) error {
	ret := m.ctrl.Call(m, "SetupState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupState indicates an expected call of SetupState
func (mr *MockEnvMockRecorder) SetupState(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupState", reflect.TypeOf((*MockEnv)(nil).SetupState), arg0)
}
