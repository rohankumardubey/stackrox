// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockEntityStore is a mock of EntityStore interface.
type MockEntityStore struct {
	ctrl     *gomock.Controller
	recorder *MockEntityStoreMockRecorder
}

// MockEntityStoreMockRecorder is the mock recorder for MockEntityStore.
type MockEntityStoreMockRecorder struct {
	mock *MockEntityStore
}

// NewMockEntityStore creates a new mock instance.
func NewMockEntityStore(ctrl *gomock.Controller) *MockEntityStore {
	mock := &MockEntityStore{ctrl: ctrl}
	mock.recorder = &MockEntityStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityStore) EXPECT() *MockEntityStoreMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockEntityStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEntityStoreMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEntityStore)(nil).Delete), ctx, id)
}

// DeleteMany mocks base method.
func (m *MockEntityStore) DeleteMany(ctx context.Context, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMany", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMany indicates an expected call of DeleteMany.
func (mr *MockEntityStoreMockRecorder) DeleteMany(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMany", reflect.TypeOf((*MockEntityStore)(nil).DeleteMany), ctx, ids)
}

// Exists mocks base method.
func (m *MockEntityStore) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockEntityStoreMockRecorder) Exists(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockEntityStore)(nil).Exists), ctx, id)
}

// Get mocks base method.
func (m *MockEntityStore) Get(ctx context.Context, id string) (*storage.NetworkEntity, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.NetworkEntity)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockEntityStoreMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEntityStore)(nil).Get), ctx, id)
}

// GetIDs mocks base method.
func (m *MockEntityStore) GetIDs(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDs", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDs indicates an expected call of GetIDs.
func (mr *MockEntityStoreMockRecorder) GetIDs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDs", reflect.TypeOf((*MockEntityStore)(nil).GetIDs), ctx)
}

// Upsert mocks base method.
func (m *MockEntityStore) Upsert(ctx context.Context, entity *storage.NetworkEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockEntityStoreMockRecorder) Upsert(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockEntityStore)(nil).Upsert), ctx, entity)
}

// UpsertMany mocks base method.
func (m *MockEntityStore) UpsertMany(ctx context.Context, objs []*storage.NetworkEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMany", ctx, objs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMany indicates an expected call of UpsertMany.
func (mr *MockEntityStoreMockRecorder) UpsertMany(ctx, objs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMany", reflect.TypeOf((*MockEntityStore)(nil).UpsertMany), ctx, objs)
}

// Walk mocks base method.
func (m *MockEntityStore) Walk(ctx context.Context, fn func(*storage.NetworkEntity) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockEntityStoreMockRecorder) Walk(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockEntityStore)(nil).Walk), ctx, fn)
}
