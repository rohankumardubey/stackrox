package store

import (
	"github.com/stackrox/rox/central/processindicator"
	"github.com/stackrox/rox/generated/storage"
)

// Store provides storage functionality for alerts.
//go:generate mockgen-wrapper
type Store interface {
	GetProcessIndicator(id string) (*storage.ProcessIndicator, bool, error)
	GetProcessIndicators() ([]*storage.ProcessIndicator, error)
	GetBatchProcessIndicators(ids []string) ([]*storage.ProcessIndicator, []int, error)
	GetProcessInfoToArgs() (map[processindicator.ProcessWithContainerInfo][]processindicator.IDAndArgs, error)

	AddProcessIndicators(...*storage.ProcessIndicator) error
	RemoveProcessIndicators(id []string) error

	AckKeysIndexed(keys ...string) error
	GetKeysToIndex() ([]string, error)

	WalkAll(func(pi *storage.ProcessIndicator) error) error
}
