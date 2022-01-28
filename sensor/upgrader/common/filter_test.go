package common

import (
	"testing"

	"github.com/stackrox/rox/pkg/k8sutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	serviceAccountYAML = `apiVersion: v1
kind: ServiceAccount
metadata:
name: sensor
namespace: stackrox
labels:
app.kubernetes.io/name: stackrox
auto-upgrade.stackrox.io/component: "sensor"
imagePullSecrets:
- name: stackrox
`

	sensorTLSSecretYAML = `apiVersion: v1
kind: Secret
data:
  ca.pem: |
    LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUIwekNDQVhxZ0F3SUJBZ0lVRjBwTFZBcUVlc09qSW5FSnRQNUg0YlpZaTRjd0NnWUlLb1pJemowRUF3SXcKU0RFbk1DVUdBMVVFQXhNZVUzUmhZMnRTYjNnZ1EyVnlkR2xtYVdOaGRHVWdRWFYwYUc5eWFYUjVNUjB3R3dZRApWUVFGRXhReE1UUTJNRGMxT1RVME5qYzNOVGMwTXpZek5UQWVGdzB5TURBM01qa3hORFU0TURCYUZ3MHlOVEEzCk1qZ3hORFU0TURCYU1FZ3hKekFsQmdOVkJBTVRIbE4wWVdOclVtOTRJRU5sY25ScFptbGpZWFJsSUVGMWRHaHYKY21sMGVURWRNQnNHQTFVRUJSTVVNVEUwTmpBM05UazFORFkzTnpVM05ETTJNelV3V1RBVEJnY3Foa2pPUFFJQgpCZ2dxaGtqT1BRTUJCd05DQUFUY1oySHgvK2RqQ3NYL2VzSVo1RTFrOFRNell2cnJOYXhPQ00vWEp5L3F0eHdsCnZvUXJ4cWY2ck5COXBaWnBMVVFEYjBlUlFXM2YrekltSHdQSkRMTVpvMEl3UURBT0JnTlZIUThCQWY4RUJBTUMKQVFZd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVXdTZVd2bStDeE5TRHRzaWZLMGh4Z3VzdApVa2N3Q2dZSUtvWkl6ajBFQXdJRFJ3QXdSQUlnUVF5amVpZkRWQkpIc1NDRXAvZDkwUHFDWHViN1U2b3EvZW0rCndicDJKd29DSUJ4Y1UvL0YvSDlnbmdzbnI4elVob2JIbWZsRzdvZjRBTTlLS2dmZWt1VDAKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  sensor-cert.pem: |
    LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNqakNDQWpXZ0F3SUJBZ0lVUXFFT2NWcUtWQXVVYW56UWZhejlBNjM4OU5Nd0NnWUlLb1pJemowRUF3SXcKU0RFbk1DVUdBMVVFQXhNZVUzUmhZMnRTYjNnZ1EyVnlkR2xtYVdOaGRHVWdRWFYwYUc5eWFYUjVNUjB3R3dZRApWUVFGRXhReE1UUTJNRGMxT1RVME5qYzNOVGMwTXpZek5UQWVGdzB5TURBM01qa3hOREEwTURCYUZ3MHlNVEEzCk1qa3hOVEEwTURCYU1IWXhGekFWQmdOVkJBc01EbE5GVGxOUFVsOVRSVkpXU1VORk1UMHdPd1lEVlFRREREUlQKUlU1VFQxSmZVMFZTVmtsRFJUb2daamt6T1dJNE5EY3RNemc0WmkwME5UVmhMVGt4WmpRdE16QXlaVEUyTjJRegpaamsxTVJ3d0dnWURWUVFGRXhNeU5qa3dNVFF6TkRrMU1UQTBOVEE0TnpFNE1Ga3dFd1lIS29aSXpqMENBUVlJCktvWkl6ajBEQVFjRFFnQUUraldoZUJuN05BVDFHWEpxaHNyVm0wajh5QTdzcm9PSEg2dmxaaDUvQ1N5OTVmTW8KLzlTS0JON2VzbThhaTFNMzhjL3JGMjFRMUZUOXJOQS9YN1JtaTZPQnpqQ0J5ekFPQmdOVkhROEJBZjhFQkFNQwpCYUF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01Bd0dBMVVkRXdFQi93UUNNQUF3CkhRWURWUjBPQkJZRUZLWUhJMFh3b0ZnMFU2bzZiaEYyZ3NrdkNTWGFNQjhHQTFVZEl3UVlNQmFBRk1FbmxyNXYKZ3NUVWc3YklueXRJY1lMckxWSkhNRXdHQTFVZEVRUkZNRU9DRDNObGJuTnZjaTV6ZEdGamEzSnZlSUlUYzJWdQpjMjl5TG5OMFlXTnJjbTk0TG5OMlk0SWJjMlZ1YzI5eUxYZGxZbWh2YjJzdWMzUmhZMnR5YjNndWMzWmpNQW9HCkNDcUdTTTQ5QkFNQ0EwY0FNRVFDSUJoRGN4WVFOUzZ6Y1lVQzltUDUwRDE5enVidkhIa0RSQ3JCUnQrM0pMRkoKQWlCUFROZDUzR1BIWFVjTmR1am5aV29iM1c1eTZvM25BaTl6dDdidzB6UHFSUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
  sensor-key.pem: |
    LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUtOL3d2cm5PSXY5Y1ZPT2U3YXA1U0V4VWg4ZVFMS2p3b2I3b1VqcnVoSXpvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFK2pXaGVCbjdOQVQxR1hKcWhzclZtMGo4eUE3c3JvT0hINnZsWmg1L0NTeTk1Zk1vLzlTSwpCTjdlc204YWkxTTM4Yy9yRjIxUTFGVDlyTkEvWDdSbWl3PT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo=
metadata:
  labels:
    auto-upgrade.stackrox.io/component: sensor
  name: sensor-tls
  namespace: stackrox
type: Opaque
`
	centralTLSSecretYAML = `apiVersion: v1
kind: Secret
type: Opaque
metadata:
  name: central-tls
  namespace: stackrox
  labels:
    app.kubernetes.io/name: stackrox

  annotations:
    "helm.sh/hook": "pre-install"

data:
  ca-key.pem: LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSVBiR3FLUFJiQVFXbmxLaENnMndsb2ZKRjlkTk15NlJGK3MyWlprVmYvM09vQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFM0dkaDhmL25Zd3JGLzNyQ0dlUk5aUEV6TTJMNjZ6V3NUZ2pQMXljdjZyY2NKYjZFSzhhbgorcXpRZmFXV2FTMUVBMjlIa1VGdDMvc3lKaDhEeVF5ekdRPT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo=
  ca.pem: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUIwekNDQVhxZ0F3SUJBZ0lVRjBwTFZBcUVlc09qSW5FSnRQNUg0YlpZaTRjd0NnWUlLb1pJemowRUF3SXcKU0RFbk1DVUdBMVVFQXhNZVUzUmhZMnRTYjNnZ1EyVnlkR2xtYVdOaGRHVWdRWFYwYUc5eWFYUjVNUjB3R3dZRApWUVFGRXhReE1UUTJNRGMxT1RVME5qYzNOVGMwTXpZek5UQWVGdzB5TURBM01qa3hORFU0TURCYUZ3MHlOVEEzCk1qZ3hORFU0TURCYU1FZ3hKekFsQmdOVkJBTVRIbE4wWVdOclVtOTRJRU5sY25ScFptbGpZWFJsSUVGMWRHaHYKY21sMGVURWRNQnNHQTFVRUJSTVVNVEUwTmpBM05UazFORFkzTnpVM05ETTJNelV3V1RBVEJnY3Foa2pPUFFJQgpCZ2dxaGtqT1BRTUJCd05DQUFUY1oySHgvK2RqQ3NYL2VzSVo1RTFrOFRNell2cnJOYXhPQ00vWEp5L3F0eHdsCnZvUXJ4cWY2ck5COXBaWnBMVVFEYjBlUlFXM2YrekltSHdQSkRMTVpvMEl3UURBT0JnTlZIUThCQWY4RUJBTUMKQVFZd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVXdTZVd2bStDeE5TRHRzaWZLMGh4Z3VzdApVa2N3Q2dZSUtvWkl6ajBFQXdJRFJ3QXdSQUlnUVF5amVpZkRWQkpIc1NDRXAvZDkwUHFDWHViN1U2b3EvZW0rCndicDJKd29DSUJ4Y1UvL0YvSDlnbmdzbnI4elVob2JIbWZsRzdvZjRBTTlLS2dmZWt1VDAKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  cert.pem: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNXekNDQWdDZ0F3SUJBZ0lVQnNxaHFPb3pFdFdJeVBmNFcyRnlzYXN4MHBJd0NnWUlLb1pJemowRUF3SXcKU0RFbk1DVUdBMVVFQXhNZVUzUmhZMnRTYjNnZ1EyVnlkR2xtYVdOaGRHVWdRWFYwYUc5eWFYUjVNUjB3R3dZRApWUVFGRXhReE1UUTJNRGMxT1RVME5qYzNOVGMwTXpZek5UQWVGdzB5TURBM01qa3hOREF6TURCYUZ3MHlNVEEzCk1qa3hOVEF6TURCYU1Gd3hHREFXQmdOVkJBc01EME5GVGxSU1FVeGZVMFZTVmtsRFJURWhNQjhHQTFVRUF3d1kKUTBWT1ZGSkJURjlUUlZKV1NVTkZPaUJEWlc1MGNtRnNNUjB3R3dZRFZRUUZFeFF4TkRNMU9UZzFNalkxTlRnMQpNemc0TmpjNE1qQlpNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDlBd0VIQTBJQUJGd25SY2d4RkQrSFk5N2RaSE1DCnc5SFZicHRYKzExSHVkNWRiV3EyQ1NSWjd1MjNvRXRhNFJLUE5XR0Y4QmZJRFQ4M3Zwa3F0YldRa2lSOTZqOUkKcUZ1amdiTXdnYkF3RGdZRFZSMFBBUUgvQkFRREFnV2dNQjBHQTFVZEpRUVdNQlFHQ0NzR0FRVUZCd01CQmdncgpCZ0VGQlFjREFqQU1CZ05WSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJTK21iSUFTT3RaU2Rob2FqbHFUU2t3CkQyMCsvREFmQmdOVkhTTUVHREFXZ0JUQko1YStiNExFMUlPMnlKOHJTSEdDNnkxU1J6QXhCZ05WSFJFRUtqQW8KZ2hCalpXNTBjbUZzTG5OMFlXTnJjbTk0Z2hSalpXNTBjbUZzTG5OMFlXTnJjbTk0TG5OMll6QUtCZ2dxaGtqTwpQUVFEQWdOSkFEQkdBaUVBK3ZsNFRJRjU0a3lBN0dUNlkyVytqWkV6QmpoRG5vY2F4ak9kZC9UcWhvNENJUUR6CjRnUm9UcnAwWjk3Z0J4Z0RKbHBsdG9rWlltTk1leHB0VkhHZDVSUy9Vdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
  jwt-key.der: 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
  key.pem: LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSURGV2JWV0FDT0lFVHVoeGpud0NPMEVHaVk1cWoySWp0bWd6VXVsc0ZMdmlvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFWENkRnlERVVQNGRqM3Qxa2N3TEQwZFZ1bTFmN1hVZTUzbDF0YXJZSkpGbnU3YmVnUzFyaApFbzgxWVlYd0Y4Z05QemUrbVNxMXRaQ1NKSDNxUDBpb1d3PT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo=
`

	additionalCASensorYAML = `apiVersion: v1
data:
  default-central-ca.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNyakNDQVpZQ0NRQ09CWHVkbjM0b01EQU5CZ2txaGtpRzl3MEJBUXNGQURBWk1SY3dGUVlEVlFRRERBNVMKYjI5MElGTmxjblpsY2lCRFFUQWVGdzB5TURBNU1qVXdOakEwTlROYUZ3MHlNREV3TWpVd05qQTBOVE5hTUJreApGekFWQmdOVkJBTU1EbEp2YjNRZ1UyVnlkbVZ5SUVOQk1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBCk1JSUJDZ0tDQVFFQTFvaGoralI5UjQzMWwvMTNYSkxZN1VuNXpHK1gyazBzVlNrV1FBSTV1dEoySUFHMDIzc0cKTlBXNlN4bHMrL2hJOVNOZ2t6eFpQOUk1TVBXMUpOaUZka2ZhTmxnRlROTElRN1hRR1UzUmlKUkQyOGFlMEpENQp1YWFzbnE3LzZFRnRjaVB6Um5ldjd4Ly9ZOVZDQVVyS0FBblRqOXFLSWtCQmxTWkgwem1rcS9PMHlpdTlyZGh2CllrejQ1bEtjRjlRRmNneXhKTXdNMjJSQTdmN0ZxQ1FUTzdmOHRqeVBwMU8yQXUwcHFEOE1PTGpUcFlkZEFjWmUKblI4YWN6eGs2ZkoxeUtMTmR6UTZ1WUhoalpSQmRUTU5ndXBsWmRwZFBLRDU3czJOQ3MvbytsWmpzSmorQTNORApESktOSGU5YUtYVDBjd3JsR1JXdi8zSzdVT3Z3SGRHWmFRSURBUUFCTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCCkFRRFF1RTl4YTFiTkxpUjQzWWd6V1J4VHZHS0xKUUpadXpMcFJRYWxQWGIvSytNaUV0WStYcTBubWJGdXk2VkkKaUNWQmhHTENzSUJxazBvWVZ4TSt1UnpBdGR5SXFTSVZhMzNwRy9zRERseEFaSFJWdzJ6djBvNzNweklMYTlXbQp3ZjFvSS85RndYTXlGSTRzS2wyMDNrT2ZBdUhjZEl2S1lhM0g0M2dtVEg4WnAwV0FTd1F3ZG9hRXlxN2RDSE9TCjBXV0NtOWw2aXE4c0hqdnJRK3J2NFJ4K1J2TEt3Tk1nSVJWT2w5Sm9JWlc1TTNLR1BCYTdsYjJ2cFdqelFkSTIKRkJWTmxrL2dyYU9lVGdoc240K3pEYktWOUpFKzR0b2Nqb2ZDdjAyelBwZmRnNnV6MGdnSWdzcnlzdnpTcjZ6WgplUWU0L2NmZS82b2t6dVpZWjcyc1RWUkwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
kind: Secret
metadata:
  labels:
    app.kubernetes.io/name: stackrox
  name: additional-ca-sensor
  namespace: stackrox
type: Opaque`
)

func mustGetObjFromYAML(t *testing.T, yaml string) k8sutil.Object {
	obj, err := k8sutil.UnstructuredFromYAML(yaml)
	require.NoError(t, err)
	return obj
}

func TestFilterToOnlyCertObjects(t *testing.T) {
	serviceAccount := mustGetObjFromYAML(t, serviceAccountYAML)
	sensorTLS := mustGetObjFromYAML(t, sensorTLSSecretYAML)
	centralTLS := mustGetObjFromYAML(t, centralTLSSecretYAML)
	additionalCA := mustGetObjFromYAML(t, additionalCASensorYAML)
	filtered := []k8sutil.Object{serviceAccount, sensorTLS, centralTLS, additionalCA}
	Filter(&filtered, CertObjectPredicate)
	assert.Equal(t, []k8sutil.Object{sensorTLS}, filtered)
}

func TestFilterAdditionalCASecretObjects(t *testing.T) {
	serviceAccount := mustGetObjFromYAML(t, serviceAccountYAML)
	sensorTLS := mustGetObjFromYAML(t, sensorTLSSecretYAML)
	centralTLS := mustGetObjFromYAML(t, centralTLSSecretYAML)
	additionalCA := mustGetObjFromYAML(t, additionalCASensorYAML)
	filtered := []k8sutil.Object{serviceAccount, sensorTLS, centralTLS, additionalCA}
	Filter(&filtered, AdditionalCASecretPredicate)
	assert.Equal(t, []k8sutil.Object{additionalCA}, filtered)
}

func TestFilterNotAdditionalCASecretObjects(t *testing.T) {
	serviceAccount := mustGetObjFromYAML(t, serviceAccountYAML)
	sensorTLS := mustGetObjFromYAML(t, sensorTLSSecretYAML)
	centralTLS := mustGetObjFromYAML(t, centralTLSSecretYAML)
	additionalCA := mustGetObjFromYAML(t, additionalCASensorYAML)
	filtered := []k8sutil.Object{serviceAccount, sensorTLS, centralTLS, additionalCA}
	Filter(&filtered, Not(AdditionalCASecretPredicate))
	assert.Equal(t, []k8sutil.Object{serviceAccount, sensorTLS, centralTLS}, filtered)
}