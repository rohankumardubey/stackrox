package datastore

import (
	"github.com/stackrox/rox/central/authprovider/datastore/internal/store"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/pkg/auth/authproviders"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once         sync.Once
	soleInstance authproviders.Store
)

// Singleton returns the sole instance of the DataStore service.
func Singleton() authproviders.Store {
	once.Do(func() {
		soleInstance = New(store.New(globaldb.GetGlobalDB()))
	})
	return soleInstance
}
