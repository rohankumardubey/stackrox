package datastore

import (
	"context"

	roleStore "github.com/stackrox/rox/central/role/datastore/internal/store"
	rocksDBStore "github.com/stackrox/rox/central/role/store"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
)

// DataStore is the datastore for roles.
//go:generate mockgen-wrapper
type DataStore interface {
	GetRole(ctx context.Context, name string) (*storage.Role, error)
	GetAllRoles(ctx context.Context) ([]*storage.Role, error)
	AddRole(ctx context.Context, role *storage.Role) error
	UpdateRole(ctx context.Context, role *storage.Role) error
	RemoveRole(ctx context.Context, name string) error

	GetPermissionSet(ctx context.Context, id string) (*storage.PermissionSet, bool, error)
	GetAllPermissionSets(ctx context.Context) ([]*storage.PermissionSet, error)
	AddPermissionSet(ctx context.Context, permissionSet *storage.PermissionSet) error
	UpdatePermissionSet(ctx context.Context, permissionSet *storage.PermissionSet) error
	RemovePermissionSet(ctx context.Context, id string) error

	GetAccessScope(ctx context.Context, id string) (*storage.SimpleAccessScope, bool, error)
	GetAllAccessScopes(ctx context.Context) ([]*storage.SimpleAccessScope, error)
	AddAccessScope(ctx context.Context, scope *storage.SimpleAccessScope) error
	UpdateAccessScope(ctx context.Context, scope *storage.SimpleAccessScope) error
	RemoveAccessScope(ctx context.Context, id string) error

	ResolveRoles(ctx context.Context, roles []*storage.Role) ([]*permissions.ResolvedRole, error)
	GetAndResolveRole(ctx context.Context, name string) (*permissions.ResolvedRole, error)
}

// New returns a new DataStore instance.
func New(roleStorage roleStore.Store, permissionSetStore rocksDBStore.PermissionSetStore, accessScopeStore rocksDBStore.SimpleAccessScopeStore, sacV2Enabled bool) DataStore {
	return &dataStoreImpl{
		roleStorage:          roleStorage,
		permissionSetStorage: permissionSetStore,
		accessScopeStorage:   accessScopeStore,
		sacV2Enabled:         sacV2Enabled,
	}
}
