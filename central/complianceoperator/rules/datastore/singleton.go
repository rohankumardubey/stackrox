package datastore

import (
	store "github.com/stackrox/rox/central/complianceoperator/rules/store"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	once sync.Once
	ds   DataStore
)

// Singleton returns the singleton datastore
func Singleton() DataStore {
	once.Do(func() {
		store, err := store.New(globaldb.GetRocksDB())
		utils.CrashOnError(err)

		ds, err = NewDatastore(store)
		utils.CrashOnError(err)
	})
	return ds
}
