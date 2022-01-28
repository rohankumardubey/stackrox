// Code generated by MockGen. DO NOT EDIT.
// Source: policy_compiler.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	detection "github.com/stackrox/rox/pkg/detection"
	reflect "reflect"
)

// MockPolicyCompiler is a mock of PolicyCompiler interface
type MockPolicyCompiler struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyCompilerMockRecorder
}

// MockPolicyCompilerMockRecorder is the mock recorder for MockPolicyCompiler
type MockPolicyCompilerMockRecorder struct {
	mock *MockPolicyCompiler
}

// NewMockPolicyCompiler creates a new mock instance
func NewMockPolicyCompiler(ctrl *gomock.Controller) *MockPolicyCompiler {
	mock := &MockPolicyCompiler{ctrl: ctrl}
	mock.recorder = &MockPolicyCompilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPolicyCompiler) EXPECT() *MockPolicyCompilerMockRecorder {
	return m.recorder
}

// CompilePolicy mocks base method
func (m *MockPolicyCompiler) CompilePolicy(policy *storage.Policy) (detection.CompiledPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompilePolicy", policy)
	ret0, _ := ret[0].(detection.CompiledPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompilePolicy indicates an expected call of CompilePolicy
func (mr *MockPolicyCompilerMockRecorder) CompilePolicy(policy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompilePolicy", reflect.TypeOf((*MockPolicyCompiler)(nil).CompilePolicy), policy)
}