package resources

import (
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/docker/types"
	"github.com/stackrox/rox/pkg/registries"
	dockerFactory "github.com/stackrox/rox/pkg/registries/docker"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/tlscheck"
)

// RegistryStore stores cluster-internal registries by namespace.
// It is assumed all the registries are Docker registries.
type RegistryStore struct {
	factory registries.Factory
	// store maps a namespace to the names of registries accessible from within the namespace.
	// The registry maps to its credentials.
	store map[string]registries.Set

	mutex sync.RWMutex
}

// newRegistryStore creates a new registryStore.
func newRegistryStore() *RegistryStore {
	return &RegistryStore{
		factory: registries.NewFactory(registries.WithRegistryCreators(dockerFactory.Creator)),
		store:   make(map[string]registries.Set),
	}
}

func (rs *RegistryStore) getRegistries(namespace string) registries.Set {
	rs.mutex.Lock()
	defer rs.mutex.Unlock()

	regs := rs.store[namespace]
	if regs == nil {
		regs = registries.NewSet(rs.factory)
		rs.store[namespace] = regs
	}

	return regs
}

// upsertRegistry upserts the given registry with the given credentials in the given namespace into the store.
func (rs *RegistryStore) upsertRegistry(namespace, registry string, dce types.DockerConfigEntry) error {
	regs := rs.getRegistries(namespace)

	secure, err := tlscheck.CheckTLS(registry)
	if err != nil {
		return errors.Wrapf(err, "unable to check TLS for registry %q", registry)
	}

	err = regs.UpdateImageIntegration(&storage.ImageIntegration{
		Name:                 registry,
		Type:                 "docker",
		Categories:           []storage.ImageIntegrationCategory{storage.ImageIntegrationCategory_REGISTRY},
		IntegrationConfig:    &storage.ImageIntegration_Docker{
			Docker: &storage.DockerConfig{
				Endpoint:             registry,
				Username:             dce.Username,
				Password:             dce.Password,
				Insecure:             !secure,
			},
		},
	})
	if err != nil {
		return errors.Wrapf(err, "updating registry store with registry %q", registry)
	}

	return nil
}

// getAllInNamespace returns all the registries+credentials within a given namespace.
func (rs *RegistryStore) getAllInNamespace(namespace string) registries.Set {
	rs.mutex.RLock()
	defer rs.mutex.RUnlock()

	return rs.store[namespace]
}