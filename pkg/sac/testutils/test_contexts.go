package testutils

import (
	"context"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/sac"
)

const (
	Cluster1   = "cluster1"
	Cluster2   = "cluster2"
	Cluster3   = "cluster3"
	NamespaceA = "namespaceA"
	NamespaceB = "namespaceB"
	NamespaceC = "namespaceC"

	UnrestrictedReadCtx              = "UnrestrictedReadCtx"
	UnrestrictedReadWriteCtx         = "UnrestrictedReadWriteCtx"
	Cluster1ReadWriteCtx             = "Cluster1ReadWriteCtx"
	Cluster1NamespaceAReadWriteCtx   = "Cluster1NamespaceAReadWriteCtx"
	Cluster1NamespaceBReadWriteCtx   = "Cluster1NamespaceBReadWriteCtx"
	Cluster1NamespaceCReadWriteCtx   = "Cluster1NamespaceCReadWriteCtx"
	Cluster1NamespacesABReadWriteCtx = "Cluster1NamespacesABReadWriteCtx"
	Cluster1NamespacesACReadWriteCtx = "Cluster1NamespacesACReadWriteCtx"
	Cluster1NamespacesBCReadWriteCtx = "Cluster1NamespacesBCReadWriteCtx"
	Cluster2ReadWriteCtx             = "Cluster2ReadWriteCtx"
	Cluster2NamespaceAReadWriteCtx   = "Cluster2NamespaceAReadWriteCtx"
	Cluster2NamespaceBReadWriteCtx   = "Cluster2NamespaceBReadWriteCtx"
	Cluster2NamespaceCReadWriteCtx   = "Cluster2NamespaceCReadWriteCtx"
	Cluster2NamespacesABReadWriteCtx = "Cluster2NamespacesABReadWriteCtx"
	Cluster2NamespacesACReadWriteCtx = "Cluster2NamespacesACReadWriteCtx"
	Cluster2NamespacesBCReadWriteCtx = "Cluster2NamespacesBCReadWriteCtx"
	Cluster3ReadWriteCtx             = "Cluster3ReadWriteCtx"
	MixedClusterAndNamespaceReadCtx  = "MixedClusterAndNamespaceReadCtx"
)

func GetNamespaceScopedTestContexts(ctx context.Context, resource permissions.Resource) map[string]context.Context {
	contextMap := make(map[string]context.Context, 0)

	contextMap[UnrestrictedReadCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS),
				sac.ResourceScopeKeys(resource)))

	contextMap[UnrestrictedReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource)))

	contextMap[Cluster1ReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster1)))

	contextMap[Cluster1NamespaceAReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster1),
				sac.NamespaceScopeKeys(NamespaceA)))

	contextMap[Cluster1NamespaceBReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster1),
				sac.NamespaceScopeKeys(NamespaceB)))

	contextMap[Cluster1NamespaceCReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster1),
				sac.NamespaceScopeKeys(NamespaceC)))

	contextMap[Cluster1NamespacesABReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster1),
				sac.NamespaceScopeKeys(NamespaceA, NamespaceB)))

	contextMap[Cluster1NamespacesACReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster1),
				sac.NamespaceScopeKeys(NamespaceA, NamespaceC)))

	contextMap[Cluster1NamespacesBCReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster1),
				sac.NamespaceScopeKeys(NamespaceB, NamespaceC)))

	contextMap[Cluster2ReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster2)))

	contextMap[Cluster2NamespaceAReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster2),
				sac.NamespaceScopeKeys(NamespaceA)))

	contextMap[Cluster2NamespaceBReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster2),
				sac.NamespaceScopeKeys(NamespaceB)))

	contextMap[Cluster2NamespaceCReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster2),
				sac.NamespaceScopeKeys(NamespaceC)))

	contextMap[Cluster2NamespacesABReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster2),
				sac.NamespaceScopeKeys(NamespaceA, NamespaceB)))

	contextMap[Cluster2NamespacesACReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster2),
				sac.NamespaceScopeKeys(NamespaceA, NamespaceC)))

	contextMap[Cluster2NamespacesBCReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster2),
				sac.NamespaceScopeKeys(NamespaceB, NamespaceC)))

	contextMap[Cluster3ReadWriteCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resource),
				sac.ClusterScopeKeys(Cluster3)))

	contextMap[MixedClusterAndNamespaceReadCtx] =
		sac.WithGlobalAccessScopeChecker(ctx,
			sac.OneStepSCC{
				sac.AccessModeScopeKey(storage.Access_READ_ACCESS): sac.OneStepSCC{
					sac.ResourceScopeKey(resource): sac.OneStepSCC{
						sac.ClusterScopeKey(Cluster1): sac.AllowFixedScopes(sac.NamespaceScopeKeys(NamespaceA)),
						sac.ClusterScopeKey(Cluster2): sac.AllowAllAccessScopeChecker(),
						sac.ClusterScopeKey(Cluster3): sac.AllowFixedScopes(sac.NamespaceScopeKeys(NamespaceC)),
					},
				},
			})

	return contextMap
}
