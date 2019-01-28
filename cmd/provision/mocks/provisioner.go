// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/provision (interfaces: Provisioner)

// Package mocks is a generated GoMock package.
package mocks

import (
	provision "code.cloudfoundry.org/cfdev/provision"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProvisioner is a mock of Provisioner interface
type MockProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockProvisionerMockRecorder
}

// MockProvisionerMockRecorder is the mock recorder for MockProvisioner
type MockProvisionerMockRecorder struct {
	mock *MockProvisioner
}

// NewMockProvisioner creates a new mock instance
func NewMockProvisioner(ctrl *gomock.Controller) *MockProvisioner {
	mock := &MockProvisioner{ctrl: ctrl}
	mock.recorder = &MockProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvisioner) EXPECT() *MockProvisionerMockRecorder {
	return m.recorder
}

// DeployBosh mocks base method
func (m *MockProvisioner) DeployBosh() error {
	ret := m.ctrl.Call(m, "DeployBosh")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployBosh indicates an expected call of DeployBosh
func (mr *MockProvisionerMockRecorder) DeployBosh() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployBosh", reflect.TypeOf((*MockProvisioner)(nil).DeployBosh))
}

// DeployServices mocks base method
func (m *MockProvisioner) DeployServices(arg0 provision.UI, arg1 []provision.Service, arg2 []string) error {
	ret := m.ctrl.Call(m, "DeployServices", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployServices indicates an expected call of DeployServices
func (mr *MockProvisionerMockRecorder) DeployServices(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployServices", reflect.TypeOf((*MockProvisioner)(nil).DeployServices), arg0, arg1, arg2)
}

// Ping mocks base method
func (m *MockProvisioner) Ping() error {
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockProvisionerMockRecorder) Ping() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockProvisioner)(nil).Ping))
}

// WhiteListServices mocks base method
func (m *MockProvisioner) WhiteListServices(arg0 string, arg1 []provision.Service) ([]provision.Service, error) {
	ret := m.ctrl.Call(m, "WhiteListServices", arg0, arg1)
	ret0, _ := ret[0].([]provision.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhiteListServices indicates an expected call of WhiteListServices
func (mr *MockProvisionerMockRecorder) WhiteListServices(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhiteListServices", reflect.TypeOf((*MockProvisioner)(nil).WhiteListServices), arg0, arg1)
}
