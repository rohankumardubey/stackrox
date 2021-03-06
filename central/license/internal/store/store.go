// Code generated by boltbindings generator. DO NOT EDIT.

package store

import (
	storage "github.com/stackrox/rox/generated/storage"
	bbolt "go.etcd.io/bbolt"
)

type Store interface {
	DeleteLicenseKey(id string) error
	ListLicenseKeys() ([]*storage.StoredLicenseKey, error)
	UpsertLicenseKeys(licensekeys []*storage.StoredLicenseKey) error
}

func New(db *bbolt.DB) (Store, error) {
	return newStore(db)
}
