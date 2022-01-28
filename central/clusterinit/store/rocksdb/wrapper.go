package rocksdb

import (
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/clusterinit/store"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/rocksdb"
	"github.com/stackrox/rox/pkg/sync"
)

type rocksDBStoreWrapper struct {
	store             Store
	uniqueIDMutex     sync.Mutex
	uniqueUpdateMutex sync.Mutex
}

// NewStore returns a new rocksdb-backed store.
func NewStore(db *rocksdb.RocksDB) (store.Store, error) {
	rawStore, err := New(db)
	if err != nil {
		return nil, err
	}
	return &rocksDBStoreWrapper{store: rawStore}, nil
}

func (w *rocksDBStoreWrapper) GetAll() ([]*storage.InitBundleMeta, error) {
	var result []*storage.InitBundleMeta
	if err := w.store.Walk(func(obj *storage.InitBundleMeta) error {
		if obj.GetIsRevoked() {
			return nil
		}
		result = append(result, obj)
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func (w *rocksDBStoreWrapper) Get(id string) (*storage.InitBundleMeta, error) {
	obj, exists, err := w.store.Get(id)
	if err != nil {
		return nil, err
	} else if !exists {
		return nil, store.ErrInitBundleNotFound
	}
	return obj, nil
}

func (w *rocksDBStoreWrapper) Add(meta *storage.InitBundleMeta) error {
	w.uniqueIDMutex.Lock()
	defer w.uniqueIDMutex.Unlock()

	if err := w.checkDuplicateName(meta); err != nil {
		return err
	}

	if exists, err := w.store.Exists(meta.GetId()); err != nil {
		return err
	} else if exists {
		return store.ErrInitBundleIDCollision
	}

	return w.store.Upsert(meta)
}

func (w *rocksDBStoreWrapper) checkDuplicateName(meta *storage.InitBundleMeta) error {
	metas, err := w.GetAll()
	if err != nil {
		return err
	}
	for _, m := range metas {
		if m.Name == meta.Name && !m.IsRevoked {
			return store.ErrInitBundleDuplicateName
		}
	}
	return nil
}

func (w *rocksDBStoreWrapper) Revoke(id string) error {
	w.uniqueUpdateMutex.Lock()
	defer w.uniqueUpdateMutex.Unlock()

	meta, err := w.Get(id)
	if err != nil {
		if errors.Is(err, store.ErrInitBundleNotFound) {
			return errors.Errorf("init bundle %q does not exist", meta.GetId())
		}
		return errors.Wrapf(err, "reading init bundle %q", id)
	}

	meta.IsRevoked = true
	if err := w.store.Upsert(meta); err != nil {
		return err
	}
	return nil
}