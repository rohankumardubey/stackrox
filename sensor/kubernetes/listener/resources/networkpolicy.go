package resources

import (
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	networkPolicyConversion "github.com/stackrox/rox/pkg/protoconv/networkpolicy"
	"github.com/stackrox/rox/sensor/common/detector"
	networkingV1 "k8s.io/api/networking/v1"
)

type networkPolicyStore interface {
	Size() int
	All() map[string]*storage.NetworkPolicy
	Get(id string) *storage.NetworkPolicy
	Upsert(ns *storage.NetworkPolicy)
	Find(namespace string, labels map[string]string) map[string]*storage.NetworkPolicy
	Delete(ID, ns string)
}

// networkPolicyDispatcher handles network policy resource events.
type networkPolicyDispatcher struct {
	netpolStore     networkPolicyStore
	deploymentStore *DeploymentStore
	detector        detector.Detector
}

func newNetworkPolicyDispatcher(networkPolicyStore networkPolicyStore, deploymentStore *DeploymentStore, detector detector.Detector) *networkPolicyDispatcher {
	return &networkPolicyDispatcher{
		netpolStore:     networkPolicyStore,
		deploymentStore: deploymentStore,
		detector:        detector,
	}
}

// Process processes a network policy resource event, and returns the sensor events to generate.
func (h *networkPolicyDispatcher) ProcessEvent(obj, _ interface{}, action central.ResourceAction) []*central.SensorEvent {
	np := obj.(*networkingV1.NetworkPolicy)

	roxNetpol := networkPolicyConversion.KubernetesNetworkPolicyWrap{NetworkPolicy: np}.ToRoxNetworkPolicy()

	if features.NetworkPolicySystemPolicy.Enabled() {
		sel := h.getSelector(roxNetpol, action)
		if action == central.ResourceAction_REMOVE_RESOURCE {
			h.netpolStore.Delete(roxNetpol.GetId(), roxNetpol.GetNamespace())
		} else {
			h.netpolStore.Upsert(roxNetpol)
		}

		h.updateDeploymentsFromStore(roxNetpol, sel)

		return []*central.SensorEvent{
			{
				Id:     string(np.UID),
				Action: action,
				Resource: &central.SensorEvent_NetworkPolicy{
					NetworkPolicy: roxNetpol,
				},
			},
		}
	}

	return []*central.SensorEvent{
		{
			Id:     string(np.UID),
			Action: action,
			Resource: &central.SensorEvent_NetworkPolicy{
				NetworkPolicy: roxNetpol,
			},
		},
	}
}

func (h *networkPolicyDispatcher) getSelector(np *storage.NetworkPolicy, action central.ResourceAction) selector {
	var sel selector
	oldWrap := h.netpolStore.Get(np.GetId())
	if oldWrap != nil {
		sel = SelectorFromMap(oldWrap.GetSpec().GetPodSelector().GetMatchLabels())
	}

	if action == central.ResourceAction_UPDATE_RESOURCE {
		if sel != nil {
			sel = or(sel, SelectorFromMap(np.GetSpec().GetPodSelector().GetMatchLabels()))
		} else {
			sel = SelectorFromMap(np.GetSpec().GetPodSelector().GetMatchLabels())
		}
	} else if action == central.ResourceAction_CREATE_RESOURCE {
		sel = SelectorFromMap(np.GetSpec().GetPodSelector().GetMatchLabels())
	}
	return sel
}

func (h *networkPolicyDispatcher) updateDeploymentsFromStore(np *storage.NetworkPolicy, sel selector) {
	for _, deploymentWrap := range h.deploymentStore.getMatchingDeployments(np.GetNamespace(), sel) {
		h.detector.ProcessDeployment(deploymentWrap.GetDeployment(), central.ResourceAction_UPDATE_RESOURCE)
	}
}
