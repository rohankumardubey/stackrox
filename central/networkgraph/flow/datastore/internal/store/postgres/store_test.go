// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"

)

type NetworkflowStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestNetworkflowStore(t *testing.T) {
	suite.Run(t, new(NetworkflowStoreSuite))
}

func (s *NetworkflowStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *NetworkflowStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *NetworkflowStoreSuite) TestStore() {
	ctx := context.Background()
	clusterId := "22"

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool, clusterId)

	deleteBefore := time.Now().Add(orphanWindow)
	log.Info(deleteBefore)

	networkFlow := &storage.NetworkFlow		{
		Props: &storage.NetworkFlowProperties{
			SrcEntity:  &storage.NetworkEntityInfo{Type: storage.NetworkEntityInfo_DEPLOYMENT, Id: "a"},
			DstEntity:  &storage.NetworkEntityInfo{Type: storage.NetworkEntityInfo_DEPLOYMENT, Id: "a"},
			DstPort:    1,
			L4Protocol: storage.L4Protocol_L4_PROTOCOL_TCP,
		},
		ClusterId: clusterId,
	}

	foundNetworkFlow, exists, err := store.Get(ctx, networkFlow.GetProps().GetSrcEntity().GetType(), networkFlow.GetProps().GetSrcEntity().GetId(), networkFlow.GetProps().GetDstEntity().GetType(), networkFlow.GetProps().GetDstEntity().GetId(), networkFlow.GetProps().GetDstPort(), networkFlow.GetProps().GetL4Protocol())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkFlow)

	// Todo: The time and how it gets stored is hosing up the equal checks.  figure out best way to deal with that.
	s.NoError(store.Upsert(ctx, networkFlow))
	foundNetworkFlow, exists, err = store.Get(ctx, networkFlow.GetProps().GetSrcEntity().GetType(), networkFlow.GetProps().GetSrcEntity().GetId(), networkFlow.GetProps().GetDstEntity().GetType(), networkFlow.GetProps().GetDstEntity().GetId(), networkFlow.GetProps().GetDstPort(), networkFlow.GetProps().GetL4Protocol())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkFlow, foundNetworkFlow)

	networkFlowCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(networkFlowCount, 1)

	networkFlowExists, err := store.Exists(ctx, networkFlow.GetProps().GetSrcEntity().GetType(), networkFlow.GetProps().GetSrcEntity().GetId(), networkFlow.GetProps().GetDstEntity().GetType(), networkFlow.GetProps().GetDstEntity().GetId(), networkFlow.GetProps().GetDstPort(), networkFlow.GetProps().GetL4Protocol())
	s.NoError(err)
	s.True(networkFlowExists)
	s.NoError(store.Upsert(ctx, networkFlow))

	foundNetworkFlow, exists, err = store.Get(ctx, networkFlow.GetProps().GetSrcEntity().GetType(), networkFlow.GetProps().GetSrcEntity().GetId(), networkFlow.GetProps().GetDstEntity().GetType(), networkFlow.GetProps().GetDstEntity().GetId(), networkFlow.GetProps().GetDstPort(), networkFlow.GetProps().GetL4Protocol())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkFlow, foundNetworkFlow)

	s.NoError(store.Delete(ctx, networkFlow.GetProps().GetSrcEntity().GetType(), networkFlow.GetProps().GetSrcEntity().GetId(), networkFlow.GetProps().GetDstEntity().GetType(), networkFlow.GetProps().GetDstEntity().GetId(), networkFlow.GetProps().GetDstPort(), networkFlow.GetProps().GetL4Protocol()))
	foundNetworkFlow, exists, err = store.Get(ctx, networkFlow.GetProps().GetSrcEntity().GetType(), networkFlow.GetProps().GetSrcEntity().GetId(), networkFlow.GetProps().GetDstEntity().GetType(), networkFlow.GetProps().GetDstEntity().GetId(), networkFlow.GetProps().GetDstPort(), networkFlow.GetProps().GetL4Protocol())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkFlow)

	var networkFlows []*storage.NetworkFlow
	flowCount := 100
	for i := 0; i < flowCount; i++ {
		networkFlow := &storage.NetworkFlow{}
		s.NoError(testutils.FullInit(networkFlow, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		networkFlows = append(networkFlows, networkFlow)
	}

	s.NoError(store.UpsertMany(ctx, networkFlows))

	networkFlowCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(networkFlowCount, flowCount)

	// Clean up
	Destroy(ctx, pool)
}
