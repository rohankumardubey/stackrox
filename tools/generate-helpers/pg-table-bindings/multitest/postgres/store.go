// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
)

const (
	countStmt  = "SELECT COUNT(*) FROM multikey"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM multikey WHERE Key1 = $1 AND Key2 = $2)"

	getStmt    = "SELECT serialized FROM multikey WHERE Key1 = $1 AND Key2 = $2"
	deleteStmt = "DELETE FROM multikey WHERE Key1 = $1 AND Key2 = $2"
	walkStmt   = "SELECT serialized FROM multikey"
)

var (
	log = logging.LoggerForModule()

	table = "multikey"
)

type Store interface {
	Count() (int, error)
	Exists(key1 string, key2 string) (bool, error)
	Get(key1 string, key2 string) (*storage.TestMultiKeyStruct, bool, error)
	Upsert(obj *storage.TestMultiKeyStruct) error
	UpsertMany(objs []*storage.TestMultiKeyStruct) error
	Delete(key1 string, key2 string) error

	Walk(fn func(obj *storage.TestMultiKeyStruct) error) error

	AckKeysIndexed(keys ...string) error
	GetKeysToIndex() ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableMultikey(db *pgxpool.Pool) {
	table := `
create table if not exists multikey (
    Key1 varchar,
    Key2 varchar,
    StringSlice text[],
    Bool bool,
    Uint64 numeric,
    Int64 numeric,
    Float numeric,
    Labels jsonb,
    Timestamp timestamp,
    Enum integer,
    Enums int[],
    Embedded_Embedded varchar,
    Oneofstring varchar,
    Oneofnested_Nested varchar,
    serialized bytea,
    PRIMARY KEY(Key1, Key2)
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

	createTableMultikeyNested(db)
}

func createTableMultikeyNested(db *pgxpool.Pool) {
	table := `
create table if not exists multikey_Nested (
    parent_Key1 varchar,
    parent_Key2 varchar,
    idx numeric,
    Nested varchar,
    Nested2_Nested2 varchar,
    PRIMARY KEY(parent_Key1, parent_Key2, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (parent_Key1, parent_Key2) REFERENCES multikey(Key1, Key2) ON DELETE CASCADE
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists multikeyNested_idx on multikey_Nested using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

}

func insertIntoMultikey(tx pgx.Tx, obj *storage.TestMultiKeyStruct) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start

		obj.GetKey1(),

		obj.GetKey2(),

		obj.GetStringSlice(),

		obj.GetBool(),

		obj.GetUint64(),

		obj.GetInt64(),

		obj.GetFloat(),

		obj.GetLabels(),

		pgutils.NilOrStringTimestamp(obj.GetTimestamp()),

		obj.GetEnum(),

		obj.GetEnums(),

		obj.GetEmbedded().GetEmbedded(),

		obj.GetOneofstring(),

		obj.GetOneofnested().GetNested(),

		serialized,
	}

	finalStr := "INSERT INTO multikey (Key1, Key2, StringSlice, Bool, Uint64, Int64, Float, Labels, Timestamp, Enum, Enums, Embedded_Embedded, Oneofstring, Oneofnested_Nested, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT(Key1, Key2) DO UPDATE SET Key1 = EXCLUDED.Key1, Key2 = EXCLUDED.Key2, StringSlice = EXCLUDED.StringSlice, Bool = EXCLUDED.Bool, Uint64 = EXCLUDED.Uint64, Int64 = EXCLUDED.Int64, Float = EXCLUDED.Float, Labels = EXCLUDED.Labels, Timestamp = EXCLUDED.Timestamp, Enum = EXCLUDED.Enum, Enums = EXCLUDED.Enums, Embedded_Embedded = EXCLUDED.Embedded_Embedded, Oneofstring = EXCLUDED.Oneofstring, Oneofnested_Nested = EXCLUDED.Oneofnested_Nested, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetNested() {
		if err := insertIntoMultikeyNested(tx, child, obj.GetKey1(), obj.GetKey2(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from multikey_Nested where parent_Key1 = $1 AND parent_Key2 = $2 AND idx >= $3"
	_, err = tx.Exec(context.Background(), query, obj.GetKey1(), obj.GetKey2(), len(obj.GetNested()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoMultikeyNested(tx pgx.Tx, obj *storage.TestMultiKeyStruct_Nested, parent_Key1 string, parent_Key2 string, idx int) error {

	values := []interface{}{
		// parent primary keys start

		parent_Key1,

		parent_Key2,

		idx,

		obj.GetNested(),

		obj.GetNested2().GetNested2(),
	}

	finalStr := "INSERT INTO multikey_Nested (parent_Key1, parent_Key2, idx, Nested, Nested2_Nested2) VALUES($1, $2, $3, $4, $5) ON CONFLICT(parent_Key1, parent_Key2, idx) DO UPDATE SET parent_Key1 = EXCLUDED.parent_Key1, parent_Key2 = EXCLUDED.parent_Key2, idx = EXCLUDED.idx, Nested = EXCLUDED.Nested, Nested2_Nested2 = EXCLUDED.Nested2_Nested2"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new Store instance using the provided sql instance.
func New(db *pgxpool.Pool) Store {
	globaldb.RegisterTable(table, "TestMultiKeyStruct")

	createTableMultikey(db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) upsert(objs ...*storage.TestMultiKeyStruct) error {
	conn, release := s.acquireConn(ops.Get, "TestMultiKeyStruct")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(context.Background())
		if err != nil {
			return err
		}

		if err := insertIntoMultikey(tx, obj); err != nil {
			if err := tx.Rollback(context.Background()); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(context.Background()); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(obj *storage.TestMultiKeyStruct) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "TestMultiKeyStruct")

	return s.upsert(obj)
}

func (s *storeImpl) UpsertMany(objs []*storage.TestMultiKeyStruct) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "TestMultiKeyStruct")

	return s.upsert(objs...)
}

// Count returns the number of objects in the store
func (s *storeImpl) Count() (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "TestMultiKeyStruct")

	row := s.db.QueryRow(context.Background(), countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(key1 string, key2 string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "TestMultiKeyStruct")

	row := s.db.QueryRow(context.Background(), existsStmt, key1, key2)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(key1 string, key2 string) (*storage.TestMultiKeyStruct, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "TestMultiKeyStruct")

	conn, release := s.acquireConn(ops.Get, "TestMultiKeyStruct")
	defer release()

	row := conn.QueryRow(context.Background(), getStmt, key1, key2)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.TestMultiKeyStruct
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(key1 string, key2 string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "TestMultiKeyStruct")

	conn, release := s.acquireConn(ops.Remove, "TestMultiKeyStruct")
	defer release()

	if _, err := conn.Exec(context.Background(), deleteStmt, key1, key2); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(fn func(obj *storage.TestMultiKeyStruct) error) error {
	rows, err := s.db.Query(context.Background(), walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.TestMultiKeyStruct
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableMultikey(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE multikey CASCADE")
	dropTableMultikeyNested(db)

}

func dropTableMultikeyNested(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE multikey_Nested CASCADE")

}

func Destroy(db *pgxpool.Pool) {
	dropTableMultikey(db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex() ([]string, error) {
	return nil, nil
}
