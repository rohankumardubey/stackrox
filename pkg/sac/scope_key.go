package sac

import (
	"fmt"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
)

// ScopeKind identifies the kind of an access scope.
type ScopeKind int

const (
	// GlobalScopeKind identifies the global scope. This scope does not have a key.
	GlobalScopeKind ScopeKind = iota
	// AccessModeScopeKind identifies the access mode scope (read or read/write).
	AccessModeScopeKind
	// ResourceScopeKind identifies the resource scope.
	ResourceScopeKind
	// ClusterScopeKind identifies the cluster scope.
	ClusterScopeKind
	// NamespaceScopeKind identifies the namespace scope.
	NamespaceScopeKind
)

// ScopeKey is a common superinterface for all access scope keys.
// This interface can only be implemented by types in this package. The intention is to
// ensure strong typing for every kind of scope key.
type ScopeKey interface {
	fmt.Stringer
	ScopeKind() ScopeKind

	isScopeKey()
}

// AccessModeScopeKey is the scope key for the access mode scope.
type AccessModeScopeKey storage.Access

func (AccessModeScopeKey) isScopeKey() {}

// ScopeKind returns `AccessModeScopeKind`.
func (AccessModeScopeKey) ScopeKind() ScopeKind {
	return AccessModeScopeKind
}

// String returns a string representation for this access scope key.
func (k AccessModeScopeKey) String() string {
	return storage.Access(k).String()
}

// Verb returns a string version of this access scope suitable for sending to third party auth plugins
func (k AccessModeScopeKey) Verb() string {
	switch storage.Access(k) {
	case storage.Access_READ_ACCESS:
		return "view"
	case storage.Access_READ_WRITE_ACCESS:
		return "edit"
	default:
		return ""
	}
}

// AccessModeScopeKeys wraps the given access modes in a scope key slice.
func AccessModeScopeKeys(ams ...storage.Access) []ScopeKey {
	keys := make([]ScopeKey, len(ams))
	for i, am := range ams {
		keys[i] = AccessModeScopeKey(am)
	}
	return keys
}

// ResourceScopeKey is the scope key for the resource scope.
type ResourceScopeKey permissions.Resource

func (ResourceScopeKey) isScopeKey() {}

// ScopeKind returns `ResourceScopeKind`.
func (ResourceScopeKey) ScopeKind() ScopeKind {
	return ResourceScopeKind
}

// String returns a string representation for this access scope key.
func (k ResourceScopeKey) String() string {
	return string(k)
}

// ResourceScopeKeys wraps the given resources in a scope key slice.
func ResourceScopeKeys(resources ...permissions.ResourceHandle) []ScopeKey {
	keys := make([]ScopeKey, len(resources))
	for i, resource := range resources {
		keys[i] = ResourceScopeKey(resource.GetResource())
	}
	return keys
}

// ClusterScopeKey is the scope key for the cluster scope.
type ClusterScopeKey string

func (ClusterScopeKey) isScopeKey() {}

// ScopeKind returns `ClusterScopeKind`.
func (ClusterScopeKey) ScopeKind() ScopeKind {
	return ClusterScopeKind
}

// String returns a string representation for this access scope key.
func (k ClusterScopeKey) String() string {
	return string(k)
}

// ClusterScopeKeys wraps the given cluster IDs in a scope key slice.
func ClusterScopeKeys(clusterIDs ...string) []ScopeKey {
	keys := make([]ScopeKey, len(clusterIDs))
	for i, clusterID := range clusterIDs {
		keys[i] = ClusterScopeKey(clusterID)
	}
	return keys
}

// NamespaceScopeKey is the scope key for the namespace scope.
type NamespaceScopeKey string

func (NamespaceScopeKey) isScopeKey() {}

// ScopeKind returns `NamespaceScopeKind`.
func (NamespaceScopeKey) ScopeKind() ScopeKind {
	return NamespaceScopeKind
}

// String returns a string representation for this access scope key.
func (k NamespaceScopeKey) String() string {
	return string(k)
}

// NamespaceScopeKeys wraps the given namespaces in a scope key slice.
func NamespaceScopeKeys(namespaces ...string) []ScopeKey {
	keys := make([]ScopeKey, len(namespaces))
	for i, namespace := range namespaces {
		keys[i] = NamespaceScopeKey(namespace)
	}
	return keys
}

// ScopePredicate is a common interface for all objects that can be interpreted as an expression over scopes.
type ScopePredicate interface {
	TryAllowed(sc ScopeChecker) TryAllowedResult
}

// ScopeSuffix is a predicate that checks if the given scope suffix (relative to the checker) is allowed.
type ScopeSuffix []ScopeKey

// TryAllowed implements the ScopePredicate interface.
func (i ScopeSuffix) TryAllowed(sc ScopeChecker) TryAllowedResult {
	return sc.TryAllowed(i...)
}

// AnyScope is a scope predicate that evaluates to Allowed if any of the given scopes is allowed.
type AnyScope []ScopePredicate

// TryAllowed implements the ScopePredicate interface.
func (p AnyScope) TryAllowed(sc ScopeChecker) TryAllowedResult {
	res := Deny
	for _, pred := range p {
		if predRes := pred.TryAllowed(sc); predRes == Allow {
			return Allow
		} else if predRes == Unknown {
			res = Unknown
		}
	}
	return res
}

// AllScopes is a scope predicate that evaluates to Allowed if all of the given scopes are allowed.
type AllScopes []ScopePredicate

// TryAllowed implements the ScopePredicate interface.
func (p AllScopes) TryAllowed(sc ScopeChecker) TryAllowedResult {
	res := Allow
	for _, pred := range p {
		if predRes := pred.TryAllowed(sc); predRes == Deny {
			return Deny
		} else if predRes == Unknown {
			res = Unknown
		}
	}
	return res
}
