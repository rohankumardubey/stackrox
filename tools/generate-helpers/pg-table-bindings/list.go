package main

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
)

var typeRegistry = make(map[string]string)

func init() {
	for s, r := range map[proto.Message]permissions.ResourceHandle{
		&storage.ClusterHealthStatus{}: resources.Cluster,
		&storage.ImageComponentEdge{}:  resources.ImageComponent,
		&storage.InitBundleMeta{}:      resources.InitBundle,
		&storage.K8SRole{}:             resources.K8sRole,
		&storage.K8SRoleBinding{}:      resources.K8sRoleBinding,
		&storage.NamespaceMetadata{}:   resources.Namespace,
		&storage.ProcessIndicator{}:    resources.Indicator,
		&storage.TokenMetadata{}:       resources.APIToken,
		// Tests
		&storage.TestMultiKeyStruct{}:  resources.Namespace,
		&storage.TestSingleKeyStruct{}: resources.Namespace,
	} {
		typeRegistry[fmt.Sprintf("%T", s)] = string(r.GetResource())
	}
}

func storageToResource(t string) string {
	if !strings.HasPrefix(t, "*") {
		t = "*" + t
	}
	s, ok := typeRegistry[t]
	if ok {
		return s
	}
	return strings.TrimPrefix(t, "*storage.")
}
