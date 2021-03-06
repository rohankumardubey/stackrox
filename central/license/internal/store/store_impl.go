// Code generated by boltbindings generator. DO NOT EDIT.

package store

import (
	proto "github.com/gogo/protobuf/proto"
	metrics "github.com/stackrox/rox/central/metrics"
	storage "github.com/stackrox/rox/generated/storage"
	protoCrud "github.com/stackrox/rox/pkg/bolthelper/crud/proto"
	ops "github.com/stackrox/rox/pkg/metrics"
	bbolt "go.etcd.io/bbolt"
	"time"
)

var (
	bucketName = []byte("licenseKeys")
)

type store struct {
	crud protoCrud.MessageCrud
}

func key(msg proto.Message) []byte {
	return []byte(msg.(*storage.StoredLicenseKey).GetLicenseId())
}

func alloc() proto.Message {
	return new(storage.StoredLicenseKey)
}

func newStore(db *bbolt.DB) (*store, error) {
	newCrud, err := protoCrud.NewMessageCrud(db, bucketName, key, alloc)
	if err != nil {
		return nil, err
	}
	return &store{crud: newCrud}, nil
}

func (s *store) DeleteLicenseKey(id string) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Remove, "LicenseKey")
	_, _, err := s.crud.Delete(id)
	return err
}

func (s *store) ListLicenseKeys() ([]*storage.StoredLicenseKey, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.GetAll, "LicenseKey")
	msgs, err := s.crud.ReadAll()
	if err != nil {
		return nil, err
	}
	storedKeys := make([]*storage.StoredLicenseKey, len(msgs))
	for i, msg := range msgs {
		storedKeys[i] = msg.(*storage.StoredLicenseKey)
	}
	return storedKeys, nil
}

func (s *store) UpsertLicenseKeys(licensekeys []*storage.StoredLicenseKey) error {
	msgs := make([]proto.Message, len(licensekeys))
	for i, key := range licensekeys {
		msgs[i] = key
	}
	_, _, err := s.crud.UpsertBatch(msgs)
	return err
}
