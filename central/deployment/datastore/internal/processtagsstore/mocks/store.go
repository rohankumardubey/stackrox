// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	analystnotes "github.com/stackrox/rox/central/analystnotes"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetTagsForProcessKey mocks base method.
func (m *MockStore) GetTagsForProcessKey(key *analystnotes.ProcessNoteKey) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsForProcessKey", key)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsForProcessKey indicates an expected call of GetTagsForProcessKey.
func (mr *MockStoreMockRecorder) GetTagsForProcessKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsForProcessKey", reflect.TypeOf((*MockStore)(nil).GetTagsForProcessKey), key)
}

// RemoveProcessTags mocks base method.
func (m *MockStore) RemoveProcessTags(key *analystnotes.ProcessNoteKey, tags []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessTags", key, tags)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessTags indicates an expected call of RemoveProcessTags.
func (mr *MockStoreMockRecorder) RemoveProcessTags(key, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessTags", reflect.TypeOf((*MockStore)(nil).RemoveProcessTags), key, tags)
}

// UpsertProcessTags mocks base method.
func (m *MockStore) UpsertProcessTags(key *analystnotes.ProcessNoteKey, tags []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertProcessTags", key, tags)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertProcessTags indicates an expected call of UpsertProcessTags.
func (mr *MockStoreMockRecorder) UpsertProcessTags(key, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertProcessTags", reflect.TypeOf((*MockStore)(nil).UpsertProcessTags), key, tags)
}

// WalkTagsForDeployment mocks base method.
func (m *MockStore) WalkTagsForDeployment(deploymentID string, f func(string) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkTagsForDeployment", deploymentID, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkTagsForDeployment indicates an expected call of WalkTagsForDeployment.
func (mr *MockStoreMockRecorder) WalkTagsForDeployment(deploymentID, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkTagsForDeployment", reflect.TypeOf((*MockStore)(nil).WalkTagsForDeployment), deploymentID, f)
}
