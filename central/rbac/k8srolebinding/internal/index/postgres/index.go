// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	mappings "github.com/stackrox/rox/central/rbac/k8srolebinding/mappings"
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

const table = "K8SRoleBinding"

func init() {
	mapping.RegisterCategoryToTable(v1.SearchCategory_ROLEBINDINGS, table)
}

func NewIndexer(db *pgxpool.Pool) *indexerImpl {
	return &indexerImpl {
		db: db,
	}
}

type indexerImpl struct {
	db *pgxpool.Pool
}

func (b *indexerImpl) AddK8sRoleBinding(deployment *storage.K8SRoleBinding) error {
	// Added as a part of normal DB op
	return nil
}

func (b *indexerImpl) AddK8sRoleBindings(_ []*storage.K8SRoleBinding) error {
	// Added as a part of normal DB op
	return nil
}

func (b *indexerImpl) DeleteK8sRoleBinding(id string) error {
	// Removed as a part of normal DB op
	return nil
}

func (b *indexerImpl) DeleteK8sRoleBindings(_ []string) error {
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
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "K8SRoleBinding")
	return postgres.RunCountRequest(v1.SearchCategory_ROLEBINDINGS, q, b.db, mappings.OptionsMap)
}

func (b *indexerImpl) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "K8SRoleBinding")
	return postgres.RunSearchRequest(v1.SearchCategory_ROLEBINDINGS, q, b.db, mappings.OptionsMap)
}
