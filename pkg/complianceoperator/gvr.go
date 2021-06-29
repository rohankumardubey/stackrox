package complianceoperator

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	complianceGroup   = "compliance.openshift.io"
	complianceVersion = "v1alpha1"

	// ComplianceGroupVersion is the group version for compliance operator objects
	ComplianceGroupVersion = complianceGroup + "/" + complianceVersion
)

// GroupVersionResources for compliance operator resources
var (
	CheckResultGVR = schema.GroupVersionResource{
		Group:    complianceGroup,
		Version:  complianceVersion,
		Resource: "compliancecheckresults",
	}

	ProfileGVR = schema.GroupVersionResource{
		Group:    complianceGroup,
		Version:  complianceVersion,
		Resource: "profiles",
	}

	ScanSettingBindingGVR = schema.GroupVersionResource{
		Group:    complianceGroup,
		Version:  complianceVersion,
		Resource: "scansettingbindings",
	}

	RuleGVR = schema.GroupVersionResource{
		Group:    complianceGroup,
		Version:  complianceVersion,
		Resource: "rules",
	}
)