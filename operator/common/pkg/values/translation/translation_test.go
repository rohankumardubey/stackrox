package translation

import (
	"context"
	"testing"

	common "github.com/stackrox/rox/operator/common/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/chartutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetCustomize(t *testing.T) {
	tests := map[string]struct {
		customizeSpec *common.CustomizeSpec
		values        chartutil.Values
		wantValues    chartutil.Values
	}{
		"nil": {
			customizeSpec: nil,
			wantValues:    chartutil.Values{},
		},
		"empty": {
			customizeSpec: &common.CustomizeSpec{},
			wantValues:    chartutil.Values{},
		},
		"all-data": {
			customizeSpec: &common.CustomizeSpec{
				Labels:         map[string]string{"label1": "value2"},
				Annotations:    map[string]string{"annotation1": "value3"},
				PodLabels:      map[string]string{"pod-label1": "value4"},
				PodAnnotations: map[string]string{"pod-annotation1": "value5"},
				EnvVars:        map[string]string{"ENV_VAR1": "value6"},
			},
			wantValues: chartutil.Values{
				"labels":         map[string]string{"label1": "value2"},
				"annotations":    map[string]string{"annotation1": "value3"},
				"podLabels":      map[string]string{"pod-label1": "value4"},
				"podAnnotations": map[string]string{"pod-annotation1": "value5"},
				"envVars":        map[string]string{"ENV_VAR1": "value6"},
			},
		},
		"partial-data": {
			customizeSpec: &common.CustomizeSpec{
				Labels: map[string]string{"value2": "should-apply"},
			},
			wantValues: chartutil.Values{
				"labels": map[string]string{"value2": "should-apply"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			values, err := GetCustomize(tt.customizeSpec).Build()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantValues, values)
		})
	}
}

func TestGetResources(t *testing.T) {
	tests := map[string]struct {
		resources  *common.Resources
		wantValues chartutil.Values
	}{
		"nil": {
			resources:  nil,
			wantValues: chartutil.Values{},
		},
		"nil-override": {
			resources:  &common.Resources{},
			wantValues: chartutil.Values{},
		},
		"data-full": {
			resources: &common.Resources{
				Override: &corev1.ResourceRequirements{
					Limits: corev1.ResourceList{
						corev1.ResourceCPU:              resource.Quantity{Format: "1"},
						corev1.ResourceEphemeralStorage: resource.Quantity{Format: "3"},
					},
					Requests: corev1.ResourceList{
						corev1.ResourceMemory: resource.Quantity{Format: "2"},
					},
				},
			},
			wantValues: chartutil.Values{
				"limits": corev1.ResourceList{
					"cpu":               resource.Quantity{Format: "1"},
					"ephemeral-storage": resource.Quantity{Format: "3"},
				},
				"requests": corev1.ResourceList{
					"memory": resource.Quantity{Format: "2"},
				},
			},
		},
		"data-no-limits": {
			resources: &common.Resources{
				Override: &corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceMemory: resource.Quantity{Format: "5"},
					},
				},
			},
			wantValues: chartutil.Values{
				"requests": corev1.ResourceList{
					"memory": resource.Quantity{Format: "5"},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			values, err := GetResources(tt.resources).Build()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantValues, values)
		})
	}
}

func TestGetServiceTLS(t *testing.T) {
	type args struct {
		clientSet  kubernetes.Interface
		serviceTLS *corev1.LocalObjectReference
	}
	tests := map[string]struct {
		args       args
		wantErr    string
		wantValues chartutil.Values
	}{
		"nil": {
			args: args{
				clientSet:  nil,
				serviceTLS: nil,
			},
			wantValues: chartutil.Values{},
		},
		"success": {
			args: args{
				clientSet: fake.NewSimpleClientset(&corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{Name: "secret-name", Namespace: "nsname"},
					StringData: map[string]string{
						"key":  "mock-key",
						"cert": "mock-cert",
					},
				}),
				serviceTLS: &corev1.LocalObjectReference{Name: "secret-name"},
			},
			wantValues: chartutil.Values{
				"serviceTLS": chartutil.Values{
					"cert": "mock-cert",
					"key":  "mock-key",
				},
			},
		},
		"get-fail": {
			args: args{
				clientSet:  fake.NewSimpleClientset(),
				serviceTLS: &corev1.LocalObjectReference{Name: "secret-name"},
			},
			wantErr: "failed to fetch.* secret.* secrets \"secret-name\" not found",
		},
		"key-fail": {
			args: args{
				clientSet: fake.NewSimpleClientset(&corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{Name: "secret-name", Namespace: "nsname"},
					StringData: map[string]string{
						"not-cert": "something else",
					},
				}),
				serviceTLS: &corev1.LocalObjectReference{Name: "secret-name"},
			},
			wantErr: "secret \"secret-name\" in namespace \"nsname\" does not contain member \"key\"",
		},
		"data-fail": {
			args: args{
				clientSet: fake.NewSimpleClientset(&corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{Name: "secret-name", Namespace: "nsname"},
					StringData: map[string]string{
						"key":      "mock-key",
						"not-cert": "something else",
					},
				}),
				serviceTLS: &corev1.LocalObjectReference{Name: "secret-name"},
			},
			wantErr: "secret \"secret-name\" in namespace \"nsname\" does not contain member \"cert\"",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			values, err := GetServiceTLS(context.Background(), tt.args.clientSet, "nsname", tt.args.serviceTLS).Build()
			if tt.wantErr != "" {
				assert.Regexp(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantValues, values)
			}
		})
	}
}