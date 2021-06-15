package sac

import (
	"github.com/stackrox/rox/central/dackbox"
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search/filtered"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	nsSAC = sac.ForResource(resources.Namespace)

	nsSACFilter filtered.Filter
	once        sync.Once
)

// GetSACFilter returns the sac filter for image ids.
func GetSACFilter() filtered.Filter {
	once.Do(func() {
		var err error
		nsSACFilter, err = filtered.NewSACFilter(
			filtered.WithResourceHelper(nsSAC),
			filtered.WithScopeTransform(dackbox.NamespaceSACTransform),
			filtered.WithReadAccess(),
		)
		utils.CrashOnError(err)
	})
	return nsSACFilter
}
