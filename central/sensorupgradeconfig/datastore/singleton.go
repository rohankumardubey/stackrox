package datastore

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/sensorupgradeconfig/datastore/internal/store"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	once      sync.Once
	singleton DataStore
)

var (
	defaultConfig = &storage.SensorUpgradeConfig{
		EnableAutoUpgrade: true,
	}
)

func addDefaultConfigIfEmpty(d DataStore) error {
	ctx := sac.WithAllAccess(context.Background())
	currentConfig, err := d.GetSensorUpgradeConfig(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to check initial sensor upgrade config")
	}
	if currentConfig != nil {
		return nil
	}
	return d.UpsertSensorUpgradeConfig(ctx, defaultConfig)
}

func initialize() {
	var err error
	singleton, err = New(store.New(globaldb.GetGlobalDB()))
	utils.CrashOnError(err)
	utils.Must(addDefaultConfigIfEmpty(singleton))
}

// Singleton returns the datastore instance.
func Singleton() DataStore {
	once.Do(initialize)
	return singleton
}
