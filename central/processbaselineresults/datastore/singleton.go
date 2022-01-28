package datastore

import (
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/processbaselineresults/datastore/internal/store/rocksdb"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	singleton DataStore

	log = logging.LoggerForModule()
)

func initialize() {
	storage := rocksdb.New(globaldb.GetRocksDB())
	singleton = New(storage)
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return singleton
}