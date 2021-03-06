import React, { ReactElement } from 'react';
import pluralize from 'pluralize';

import { NavigateToEntityHook } from 'Containers/Network/SidePanel/useNavigateToEntity';
import { filterModes } from 'constants/networkFilterModes';
import { Edge } from 'Containers/Network/networkTypes';
import { networkFlowStatus } from 'constants/networkGraph';
import useFetchNetworkBaselines from './useFetchNetworkBaselines';

import NetworkBaselines from '../NetworkBaselines';

export type NetworkFlowsProps = {
    deploymentId: string;
    edges: Edge[];
    filterState: number;
    onNavigateToEntity: NavigateToEntityHook;
    lastUpdatedTimestamp: string;
};

function getPanelHeaderText(numBaselineFlows: number, filterState): string {
    switch (filterState) {
        case filterModes.active:
            return `${numBaselineFlows} active ${pluralize('flow', numBaselineFlows)}`;
        case filterModes.allowed:
            return `${numBaselineFlows} allowed ${pluralize('flow', numBaselineFlows)}`;
        default:
            return `${numBaselineFlows} ${pluralize('flow', numBaselineFlows)}`;
    }
}

function NetworkFlows({
    deploymentId,
    edges,
    filterState,
    onNavigateToEntity,
    lastUpdatedTimestamp,
}: NetworkFlowsProps): ReactElement {
    const { data: networkBaselines, isLoading } = useFetchNetworkBaselines({
        deploymentId,
        edges,
        filterState,
        lastUpdatedTimestamp,
    });

    const header = getPanelHeaderText(networkBaselines.length, filterState);

    return (
        <NetworkBaselines
            header={header}
            headerComponents={null}
            isLoading={isLoading}
            networkBaselines={networkBaselines}
            deploymentId={deploymentId}
            filterState={filterModes}
            onNavigateToEntity={onNavigateToEntity}
            includedBaselineStatuses={[networkFlowStatus.ANOMALOUS, networkFlowStatus.BASELINE]}
            excludedSearchCategories={[]}
        />
    );
}

export default NetworkFlows;
