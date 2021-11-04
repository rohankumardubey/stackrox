// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/batcher"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/set"
)

const (
		countStmt = "select count(*) from networkgraphconfig"
		existsStmt = "select exists(select 1 from networkgraphconfig where id = $1)"
		getIDsStmt = "select id from networkgraphconfig"
		getStmt = "select value from networkgraphconfig where id = $1"
		getManyStmt = "select value from networkgraphconfig where id = ANY($1::text[])"
		upsertStmt = "insert into networkgraphconfig (id, value) values($1, $2) on conflict(id) do update set value = EXCLUDED.value"
		deleteStmt = "delete from networkgraphconfig where id = $1"
		deleteManyStmt = "delete from networkgraphconfig where id = ANY($1::text[])"
		walkStmt = "select value from networkgraphconfig"
		walkWithIDStmt = "select id, value from networkgraphconfig"
)

var (
	log = logging.LoggerForModule()

	table = "networkgraphconfig"

	marshaler = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true}
)

type Store interface {
	Count() (int, error)
	Exists(id string) (bool, error)
	GetIDs() ([]string, error)
	Get(id string) (*storage.NetworkGraphConfig, bool, error)
	GetMany(ids []string) ([]*storage.NetworkGraphConfig, []int, error)
	UpsertWithID(id string, obj *storage.NetworkGraphConfig) error
	UpsertManyWithIDs(ids []string, objs []*storage.NetworkGraphConfig) error
	Delete(id string) error
	DeleteMany(ids []string) error
	WalkAllWithID(fn func(id string, obj *storage.NetworkGraphConfig) error) error
	AckKeysIndexed(keys ...string) error
	GetKeysToIndex() ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func alloc() proto.Message {
	return &storage.NetworkGraphConfig{}
}

const (
	createTableQuery = "create table if not exists networkgraphconfig (id varchar primary key, value jsonb)"
	createIDIndexQuery = "create index if not exists networkgraphconfig_id on networkgraphconfig using hash ((id))"

	batchInsertTemplate = "insert into networkgraphconfig (id, value) values %s on conflict(id) do update set value = EXCLUDED.value"
)

// New returns a new Store instance using the provided sql instance.
func New(db *pgxpool.Pool) Store {
	globaldb.RegisterTable(table, "NetworkGraphConfig")

	_, err := db.Exec(context.Background(), createTableQuery)
	if err != nil {
		panic("error creating table")
	}

	_, err = db.Exec(context.Background(), createIDIndexQuery)
	if err != nil {
		panic("error creating index")
	}

//
	return &storeImpl{
		db: db,
	}
//
}

// Count returns the number of objects in the store
func (s *storeImpl) Count() (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "NetworkGraphConfig")

	row := s.db.QueryRow(context.Background(), countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "NetworkGraphConfig")

	row := s.db.QueryRow(context.Background(), existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, nilNoRows(err)
	}
	return exists, nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs() ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "NetworkGraphConfigIDs")

	rows, err := s.db.Query(context.Background(), getIDsStmt)
	if err != nil {
		return nil, nilNoRows(err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func nilNoRows(err error) error {
	if err == pgx.ErrNoRows {
		return nil
	}
	return err
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(id string) (*storage.NetworkGraphConfig, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "NetworkGraphConfig")

	conn, release := s.acquireConn(ops.Get, "NetworkGraphConfig")
	defer release()

	row := conn.QueryRow(context.Background(), getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, nilNoRows(err)
	}

	msg := alloc()
	buf := bytes.NewBuffer(data)
	defer metrics.SetJSONPBOperationDurationTime(time.Now(), "Unmarshal", "NetworkGraphConfig")
	if err := jsonpb.Unmarshal(buf, msg); err != nil {
		return nil, false, err
	}
	return msg.(*storage.NetworkGraphConfig), true, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice 
func (s *storeImpl) GetMany(ids []string) ([]*storage.NetworkGraphConfig, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "NetworkGraphConfig")

	conn, release := s.acquireConn(ops.GetMany, "NetworkGraphConfig")
	defer release()

	rows, err := conn.Query(context.Background(), getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	elems := make([]*storage.NetworkGraphConfig, 0, len(ids))
	foundSet := set.NewStringSet()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := alloc()
		buf := bytes.NewBuffer(data)
		t := time.Now()
		if err := jsonpb.Unmarshal(buf, msg); err != nil {
			return nil, nil, err
		}
		metrics.SetJSONPBOperationDurationTime(t, "Unmarshal", "NetworkGraphConfig")
		elem := msg.(*storage.NetworkGraphConfig)
		foundSet.Add(elem.GetId())
		elems = append(elems, elem)
	}
	missingIndices := make([]int, 0, len(ids)-len(foundSet))
	for i, id := range ids {
		if !foundSet.Contains(id) {
			missingIndices = append(missingIndices, i)
		}
	}
	return elems, missingIndices, nil
}
// UpsertWithID inserts the object into the DB
func (s *storeImpl) UpsertWithID(id string, obj *storage.NetworkGraphConfig) error {
	return upsert(id, obj)
}

// UpsertManyWithIDs batches objects into the DB
func (s *storeImpl) UpsertManyWithIDs(ids []string, objs []*storage.NetworkGraphConfig) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.AddMany, "NetworkGraphConfig")

	// txn? or partial? what is the impact of one not being upserted
	for i, id := range ids {
		if err := s.upsert(id, objs(i)); err != nil {
			return err
		}
	}
	return nil
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "NetworkGraphConfig")

	conn, release := s.acquireConn(ops.Remove, "NetworkGraphConfig")
	defer release()

	if _, err := conn.Exec(context.Background(), deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// Delete removes the specified IDs from the store
func (s *storeImpl) DeleteMany(ids []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "NetworkGraphConfig")

	conn, release := s.acquireConn(ops.RemoveMany, "NetworkGraphConfig")
	defer release()
	if _, err := conn.Exec(context.Background(), deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}
// WalkAllWithID iterates over all of the objects in the store and applies the closure
func (s *storeImpl) WalkAllWithID(fn func(id string, obj *storage.NetworkGraphConfig) error) error {

	panic("unimplemented")	
//return b.crud.WalkAllWithID(func(id []byte, msg proto.Message) error {
	rows, err := s.db.Query(context.Background(), walkStmt)
	if err != nil {
		return nilNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var data []byte
		if err := rows.Scan(&id, &data); err != nil {
			return err
		}
		msg := alloc()
		buf := bytes.NewBuffer(data)
		if err := jsonpb.Unmarshal(buf, msg); err != nil {
			return err
		}
		return fn(id, msg.(*storage.NetworkGraphConfig))
	}
	return nil
}

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex() ([]string, error) {
	return nil, nil
}