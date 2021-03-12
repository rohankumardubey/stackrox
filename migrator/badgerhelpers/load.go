package badgerhelpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/migrator/option"
)

const (
	badgerDBDirName = `badgerdb`
)

// New returns an instance of the persistent BadgerDB store
func New(path string) (*badger.DB, error) {
	if stat, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0700)
		if err != nil {
			return nil, errors.Wrapf(err, "error creating badger path %s", path)
		}
	} else if err != nil {
		return nil, err
	} else if !stat.IsDir() {
		return nil, fmt.Errorf("badger path %s is not a directory", path)
	}

	options := badger.DefaultOptions(path).
		WithDir(path).
		WithTruncate(true).
		WithNumLevelZeroTables(2).
		WithNumLevelZeroTablesStall(5)
	return badger.Open(options)
}

// NewWithDefaults returns an instance of the persistent BadgerDB store instantiated at the default filesystem location.
func NewWithDefaults() (*badger.DB, error) {
	return New(filepath.Join(option.MigratorOptions.DBPathBase, badgerDBDirName))
}

// NewTemp creates a new DB, but places it in the host temporary directory.
func NewTemp(dbName string) (*badger.DB, error) {
	tmpDir, err := ioutil.TempDir("", fmt.Sprintf("badger-%s", strings.Replace(dbName, "/", "_", -1)))
	if err != nil {
		return nil, err
	}
	return New(tmpDir)
}
