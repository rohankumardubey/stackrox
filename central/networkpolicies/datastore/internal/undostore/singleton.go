package undostore

import (
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	undoStoreInstance     UndoStore
	undoStoreInstanceInit sync.Once
)

// Singleton returns the singleton instance of the undo store.
func Singleton() UndoStore {
	undoStoreInstanceInit.Do(func() {
		undoStoreInstance = New(globaldb.GetGlobalDB())
	})
	return undoStoreInstance
}
