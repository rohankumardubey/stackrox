package datastore

import (
	"github.com/stackrox/rox/central/globaldb"
	graphConfigDS "github.com/stackrox/rox/central/networkgraph/config/datastore"
	"github.com/stackrox/rox/central/networkgraph/entity/datastore/internal/store/rocksdb"
	"github.com/stackrox/rox/central/networkgraph/entity/networktree"
	"github.com/stackrox/rox/central/sensor/service/connection"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	once sync.Once
	ds   EntityDataStore
)

// Singleton provides the instance of EntityDataStore to use.
func Singleton() EntityDataStore {
	storage, err := rocksdb.New(globaldb.GetRocksDB())
	utils.CrashOnError(err)

	once.Do(func() {
		ds = NewEntityDataStore(storage, graphConfigDS.Singleton(), networktree.Singleton(), connection.ManagerSingleton())
	})
	return ds
}