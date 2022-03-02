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

// MockClusterStore is a mock of ClusterStore interface.
type MockClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStoreMockRecorder
}

// MockClusterStoreMockRecorder is the mock recorder for MockClusterStore.
type MockClusterStoreMockRecorder struct {
	mock *MockClusterStore
}

// NewMockClusterStore creates a new mock instance.
func NewMockClusterStore(ctrl *gomock.Controller) *MockClusterStore {
	mock := &MockClusterStore{ctrl: ctrl}
	mock.recorder = &MockClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterStore) EXPECT() *MockClusterStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockClusterStore) Count(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockClusterStoreMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockClusterStore)(nil).Count), ctx)
}

// Delete mocks base method.
func (m *MockClusterStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterStoreMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterStore)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockClusterStore) Get(ctx context.Context, id string) (*storage.Cluster, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.Cluster)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockClusterStoreMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterStore)(nil).Get), ctx, id)
}

// GetMany mocks base method.
func (m *MockClusterStore) GetMany(ctx context.Context, ids []string) ([]*storage.Cluster, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ctx, ids)
	ret0, _ := ret[0].([]*storage.Cluster)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMany indicates an expected call of GetMany.
func (mr *MockClusterStoreMockRecorder) GetMany(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockClusterStore)(nil).GetMany), ctx, ids)
}

// Upsert mocks base method.
func (m *MockClusterStore) Upsert(ctx context.Context, cluster *storage.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockClusterStoreMockRecorder) Upsert(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockClusterStore)(nil).Upsert), ctx, cluster)
}

// Walk mocks base method.
func (m *MockClusterStore) Walk(ctx context.Context, fn func(*storage.Cluster) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockClusterStoreMockRecorder) Walk(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockClusterStore)(nil).Walk), ctx, fn)
}

// MockClusterHealthStore is a mock of ClusterHealthStore interface.
type MockClusterHealthStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterHealthStoreMockRecorder
}

// MockClusterHealthStoreMockRecorder is the mock recorder for MockClusterHealthStore.
type MockClusterHealthStoreMockRecorder struct {
	mock *MockClusterHealthStore
}

// NewMockClusterHealthStore creates a new mock instance.
func NewMockClusterHealthStore(ctrl *gomock.Controller) *MockClusterHealthStore {
	mock := &MockClusterHealthStore{ctrl: ctrl}
	mock.recorder = &MockClusterHealthStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterHealthStore) EXPECT() *MockClusterHealthStoreMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockClusterHealthStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterHealthStoreMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterHealthStore)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockClusterHealthStore) Get(ctx context.Context, id string) (*storage.ClusterHealthStatus, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.ClusterHealthStatus)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockClusterHealthStoreMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterHealthStore)(nil).Get), ctx, id)
}

// GetMany mocks base method.
func (m *MockClusterHealthStore) GetMany(ctx context.Context, ids []string) ([]*storage.ClusterHealthStatus, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ctx, ids)
	ret0, _ := ret[0].([]*storage.ClusterHealthStatus)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMany indicates an expected call of GetMany.
func (mr *MockClusterHealthStoreMockRecorder) GetMany(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockClusterHealthStore)(nil).GetMany), ctx, ids)
}

// UpsertManyWithIDs mocks base method.
func (m *MockClusterHealthStore) UpsertManyWithIDs(ctx context.Context, ids []string, objs []*storage.ClusterHealthStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertManyWithIDs", ctx, ids, objs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertManyWithIDs indicates an expected call of UpsertManyWithIDs.
func (mr *MockClusterHealthStoreMockRecorder) UpsertManyWithIDs(ctx, ids, objs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertManyWithIDs", reflect.TypeOf((*MockClusterHealthStore)(nil).UpsertManyWithIDs), ctx, ids, objs)
}

// UpsertWithID mocks base method.
func (m *MockClusterHealthStore) UpsertWithID(ctx context.Context, id string, obj *storage.ClusterHealthStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWithID", ctx, id, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWithID indicates an expected call of UpsertWithID.
func (mr *MockClusterHealthStoreMockRecorder) UpsertWithID(ctx, id, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWithID", reflect.TypeOf((*MockClusterHealthStore)(nil).UpsertWithID), ctx, id, obj)
}

// WalkAllWithID mocks base method.
func (m *MockClusterHealthStore) WalkAllWithID(ctx context.Context, fn func(string, *storage.ClusterHealthStatus) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkAllWithID", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkAllWithID indicates an expected call of WalkAllWithID.
func (mr *MockClusterHealthStoreMockRecorder) WalkAllWithID(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkAllWithID", reflect.TypeOf((*MockClusterHealthStore)(nil).WalkAllWithID), ctx, fn)
}
