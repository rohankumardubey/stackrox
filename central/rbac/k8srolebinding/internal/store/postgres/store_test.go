// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	storage "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type RolebindingsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestRolebindingsStore(t *testing.T) {
	suite.Run(t, new(RolebindingsStoreSuite))
}

func (s *RolebindingsStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *RolebindingsStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *RolebindingsStoreSuite) TestStore() {
	ctx := context.Background()

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	k8SRoleBinding := &storage.K8SRoleBinding{}
	s.NoError(testutils.FullInit(k8SRoleBinding, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundK8SRoleBinding, exists, err := store.Get(ctx, k8SRoleBinding.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundK8SRoleBinding)

	s.NoError(store.Upsert(ctx, k8SRoleBinding))
	foundK8SRoleBinding, exists, err = store.Get(ctx, k8SRoleBinding.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(k8SRoleBinding, foundK8SRoleBinding)

	k8SRoleBindingCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(k8SRoleBindingCount, 1)

	k8SRoleBindingExists, err := store.Exists(ctx, k8SRoleBinding.GetId())
	s.NoError(err)
	s.True(k8SRoleBindingExists)
	s.NoError(store.Upsert(ctx, k8SRoleBinding))

	foundK8SRoleBinding, exists, err = store.Get(ctx, k8SRoleBinding.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(k8SRoleBinding, foundK8SRoleBinding)

	s.NoError(store.Delete(ctx, k8SRoleBinding.GetId()))
	foundK8SRoleBinding, exists, err = store.Get(ctx, k8SRoleBinding.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundK8SRoleBinding)

	var k8SRoleBindings []*storage.K8SRoleBinding
	for i := 0; i < 200; i++ {
		k8SRoleBinding := &storage.K8SRoleBinding{}
		s.NoError(testutils.FullInit(k8SRoleBinding, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		k8SRoleBindings = append(k8SRoleBindings, k8SRoleBinding)
	}

	s.NoError(store.UpsertMany(ctx, k8SRoleBindings))

	k8SRoleBindingCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(k8SRoleBindingCount, 200)
}
