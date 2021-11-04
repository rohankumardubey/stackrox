package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/networkgraph/flow/datastore/internal/store"
)

const (
	createTableQuery = "create table if not exists networkflows (id varchar primary key, clusterid varchar, value jsonb)"
)

// NewClusterStore returns a new ClusterStore instance using the provided rocksdb instance.
func NewClusterStore(db *pgxpool.Pool) store.ClusterStore {
	if _, err := db.Exec(context.Background(), createTableQuery); err != nil {
		panic(err)
	}

	if _, err := db.Exec(context.Background(), `create index if not exists networkflows_dst_entity on networkflows using hash ((value->'dstEntity'->>'id'))`); err != nil {
		panic(err)
	}
	if _, err := db.Exec(context.Background(), `create index if not exists networkflows_src_entity on networkflows using hash ((value->'srcEntity'->>'id'))`); err != nil {
		panic(err)
	}

	return &clusterStoreImpl{
		db: db,
	}
}

type clusterStoreImpl struct {
	db *pgxpool.Pool
}

// GetFlowStore returns the FlowStore for the cluster ID, or nil if none exists.
func (s *clusterStoreImpl) GetFlowStore(clusterID string) store.FlowStore {
	return &flowStoreImpl{
		db:        s.db,
		clusterID: clusterID,
	}
}

// CreateFlowStore returns the FlowStore for the cluster ID, or creates one if none exists.
func (s *clusterStoreImpl) CreateFlowStore(clusterID string) (store.FlowStore, error) {
	fs := &flowStoreImpl{
		db:        s.db,
		clusterID: clusterID,
	}
	return fs, nil
}