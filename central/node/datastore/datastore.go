package datastore

import (
	"github.com/stackrox/rox/central/node/store"
)

//go:generate mockgen-wrapper

// DataStore is a wrapper around a store that provides search functionality
type DataStore interface {
	store.Store
}
