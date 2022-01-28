package store

import (
	"github.com/stackrox/rox/generated/storage"
)

// EntityStore stores network graph entities.
//go:generate mockgen-wrapper
type EntityStore interface {
	Exists(id string) (bool, error)

	GetIDs() ([]string, error)
	Get(id string) (*storage.NetworkEntity, bool, error)

	Upsert(entity *storage.NetworkEntity) error
	UpsertMany(objs []*storage.NetworkEntity) error
	Delete(id string) error
	DeleteMany(ids []string) error

	Walk(fn func(obj *storage.NetworkEntity) error) error
}