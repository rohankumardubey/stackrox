// Code generated by singletonstore generator. DO NOT EDIT.

package store

import (
	proto "github.com/gogo/protobuf/proto"
	metrics "github.com/stackrox/rox/central/metrics"
	storage "github.com/stackrox/rox/generated/storage"
	singletonstore "github.com/stackrox/rox/pkg/bolthelper/singletonstore"
	ops "github.com/stackrox/rox/pkg/metrics"
	bbolt "go.etcd.io/bbolt"
	"time"
)

var (
	bucketName = []byte("config")
)

type Store interface {
	GetConfig() (*storage.Config, error)
	UpsertConfig(config *storage.Config) error
}

func New(db *bbolt.DB) Store {
	return &store{underlying: singletonstore.New(db, bucketName, func() proto.Message {
		return new(storage.Config)
	}, "Config")}
}

type store struct {
	underlying singletonstore.SingletonStore
}

func (s *store) GetConfig() (*storage.Config, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Get, "Config")
	msg, err := s.underlying.Get()
	if err != nil {
		return nil, err
	}
	if msg == nil {
		return nil, nil
	}
	return msg.(*storage.Config), nil
}

func (s *store) UpsertConfig(config *storage.Config) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Upsert, "Config")
	return s.underlying.Upsert(config)
}