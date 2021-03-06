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
	bucketName = []byte("installationInfo")
)

type Store interface {
	GetInstallationInfo() (*storage.InstallationInfo, error)
	AddInstallationInfo(installationinfo *storage.InstallationInfo) error
}

func New(db *bbolt.DB) Store {
	return &store{underlying: singletonstore.New(db, bucketName, func() proto.Message {
		return new(storage.InstallationInfo)
	}, "InstallationInfo")}
}

type store struct {
	underlying singletonstore.SingletonStore
}

func (s *store) GetInstallationInfo() (*storage.InstallationInfo, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Get, "InstallationInfo")
	msg, err := s.underlying.Get()
	if err != nil {
		return nil, err
	}
	if msg == nil {
		return nil, nil
	}
	return msg.(*storage.InstallationInfo), nil
}

func (s *store) AddInstallationInfo(installationinfo *storage.InstallationInfo) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Add, "InstallationInfo")
	return s.underlying.Create(installationinfo)
}
