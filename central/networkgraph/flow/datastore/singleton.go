package datastore

import (
	"github.com/stackrox/rox/central/deployment/cache"
	graphConfigDS "github.com/stackrox/rox/central/networkgraph/config/datastore"
	networkEntityDS "github.com/stackrox/rox/central/networkgraph/entity/datastore"
	"github.com/stackrox/rox/central/networkgraph/flow/datastore/internal/store/singleton"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once     sync.Once
	instance ClusterDataStore
)

// Singleton provides the instance of ClusterDataStore to use.
func Singleton() ClusterDataStore {
	once.Do(func() {
		instance = NewClusterDataStore(singleton.Singleton(), graphConfigDS.Singleton(), networkEntityDS.Singleton(), cache.DeletedDeploymentCacheSingleton())
	})
	return instance
}
