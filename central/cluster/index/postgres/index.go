// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	mappings "github.com/stackrox/rox/central/cluster/index/mappings"
	metrics "github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	ops "github.com/stackrox/rox/pkg/metrics"
	search "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
	"github.com/stackrox/rox/pkg/search/postgres"
	"time"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/pkg/search/blevesearch"
	"github.com/stackrox/rox/pkg/logging"
)

var log = logging.LoggerForModule()

const table = "Cluster"

func init() {
	mapping.RegisterCategoryToTable(v1.SearchCategory_CLUSTERS, table)
}

func NewIndexer(db *pgxpool.Pool) *indexerImpl {
	return &indexerImpl {
		db: db,
	}
}

type indexerImpl struct {
	db *pgxpool.Pool
}

func (b *indexerImpl) AddCluster(deployment *storage.Cluster) error {
	// Added as a part of normal DB op
	return nil
}

func (b *indexerImpl) AddClusters(_ []*storage.Cluster) error {
	// Added as a part of normal DB op
	return nil
}

func (b *indexerImpl) DeleteCluster(id string) error {
	// Removed as a part of normal DB op
	return nil
}

func (b *indexerImpl) DeleteClusters(_ []string) error {
	// Added as a part of normal DB op
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}

func (b *indexerImpl) Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "Cluster")
	return postgres.RunCountRequest(v1.SearchCategory_CLUSTERS, q, b.db, mappings.OptionsMap)
}

func (b *indexerImpl) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "Cluster")
	return postgres.RunSearchRequest(v1.SearchCategory_CLUSTERS, q, b.db, mappings.OptionsMap)
}
