package fixtures

import (
	ptypes "github.com/gogo/protobuf/types"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/images/types"
)

// GetAlert returns a Mock Alert
func GetAlert() *storage.Alert {
	return &storage.Alert{
		Id: "Alert1",
		Violations: []*storage.Alert_Violation{
			{
				Message: "Deployment is affected by 'CVE-2017-15804'",
			},
			{
				Message: "Deployment is affected by 'CVE-2017-15670'",
			},
			{
				Message: "This is a kube event violation",
				MessageAttributes: &storage.Alert_Violation_KeyValueAttrs_{
					KeyValueAttrs: &storage.Alert_Violation_KeyValueAttrs{
						Attrs: []*storage.Alert_Violation_KeyValueAttrs_KeyValueAttr{
							{Key: "pod", Value: "nginx"},
							{Key: "container", Value: "nginx"},
						},
					},
				},
			},
		},
		ProcessViolation: &storage.Alert_ProcessViolation{
			Message: "This is a process violation",
		},
		Time:   ptypes.TimestampNow(),
		Policy: GetPolicy(),
		Entity: &storage.Alert_Deployment_{
			Deployment: &storage.Alert_Deployment{
				Name:        "nginx_server",
				Id:          "s79mdvmb6dsl",
				ClusterId:   "prod cluster",
				ClusterName: "prod cluster",
				Namespace:   "stackrox",
				Labels: map[string]string{
					"com.docker.stack.namespace":    "prevent",
					"com.docker.swarm.service.name": "prevent_sensor",
					"email":                         "vv@stackrox.com",
					"owner":                         "stackrox",
				},
				Containers: []*storage.Alert_Deployment_Container{
					{
						Name:  "nginx110container",
						Image: types.ToContainerImage(LightweightDeploymentImage()),
					},
				},
			},
		},
	}
}

// GetAlertWithMitre returns a mock Alert with MITRE ATT&CK
func GetAlertWithMitre() *storage.Alert {
	alert := GetAlert()
	alert.Policy = GetPolicyWithMitre()
	return alert
}

// GetResourceAlert returns a Mock Alert with a resource entity
func GetResourceAlert() *storage.Alert {
	return &storage.Alert{
		Id: "some-resource-alert-on-secret",
		Violations: []*storage.Alert_Violation{
			{
				Message: "Access to secret \"my-secret\" in \"cluster-id / stackrox\"",
				Type:    storage.Alert_Violation_K8S_EVENT,
				MessageAttributes: &storage.Alert_Violation_KeyValueAttrs_{
					KeyValueAttrs: &storage.Alert_Violation_KeyValueAttrs{
						Attrs: []*storage.Alert_Violation_KeyValueAttrs_KeyValueAttr{
							{Key: "Kubernetes API Verb", Value: "CREATE"},
							{Key: "username", Value: "test-user"},
							{Key: "user groups", Value: "groupA, groupB"},
							{Key: "resource", Value: "/api/v1/namespace/stackrox/secrets/my-secret"},
							{Key: "user agent", Value: "oc/4.7.0 (darwin/amd64) kubernetes/c66c03f"},
							{Key: "IP address", Value: "192.168.0.1, 127.0.0.1"},
							{Key: "impersonated username", Value: "central-service-account"},
							{Key: "impersonated user groups", Value: "service-accounts, groupB"},
						},
					},
				},
			},
		},
		Time:   ptypes.TimestampNow(),
		Policy: GetAuditLogEventSourcePolicy(),
		Entity: &storage.Alert_Resource_{
			Resource: &storage.Alert_Resource{
				ResourceType: storage.Alert_Resource_SECRETS,
				Name:         "my-secret",
				ClusterId:    "cluster-id",
				ClusterName:  "prod cluster",
				Namespace:    "stackrox",
				NamespaceId:  "aaaa-bbbb-cccc-dddd",
			},
		},
		LifecycleStage: storage.LifecycleStage_RUNTIME,
	}
}

// GetImageAlert returns a Mock alert with an image for entity
func GetImageAlert() *storage.Alert {
	imageAlert := GetAlert()
	imageAlert.Entity = &storage.Alert_Image{Image: types.ToContainerImage(GetImage())}

	return imageAlert
}

// GetAlertWithID returns a mock alert with the specified id.
func GetAlertWithID(id string) *storage.Alert {
	alert := GetAlert()
	alert.Id = id
	return alert
}
