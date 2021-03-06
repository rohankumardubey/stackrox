import scopeSelectors from '../helpers/scopeSelectors';
import search from '../selectors/search';

export const url = '/main/network';

const networkPanels = {
    creatorPanel: '[data-testid="network-creator-panel"]',
    simulatorPanel: '[data-testid="network-simulator-panel"]',
    uploadPanel: '[data-testid="upload-yaml-panel"]',
    detailsPanel: '[data-testid="network-details-panel"]',
};

const networkEntityTabbedOverlay = '[data-testid="network-entity-tabbed-overlay"]';

export const selectors = {
    cytoscapeContainer: '#cytoscapeContainer',
    simulatorSuccessMessage: 'div[data-testid="message-body"]:contains("Policies processed")',
    panels: networkPanels,
    legend: {
        deployments: '[data-testid="deployment-legend"] div',
        namespaces: '[data-testid="namespace-legend"] div',
        connections: '[data-testid="connection-legend"] div',
    },
    namespaces: {
        all: 'g.container > rect',
        getNamespace: (namespace) => `g.container > rect.namespace-${namespace}`,
    },
    services: {
        all: 'g.namespace .node',
        getServicesForNamespace: (namespace) => `g.namespace-${namespace} .node`,
    },
    links: {
        all: '.link',
        bidirectional: '.link[marker-start="url(#start)"]',
        namespaces: '.link.namespace',
        services: '.link.service',
    },
    buttons: {
        viewActiveYamlButton: '[data-testid="view-active-yaml-button"]',
        simulatorButtonOff: '[data-testid="simulator-button-off"]',
        generateNetworkPolicies: 'button:contains("Generate and simulate network policies")',
        applyNetworkPolicies: 'button:contains("Apply Network Policies")',
        apply: 'div[aria-modal="true"] button:contains("Apply")',
        // Select buttons by data-testid attribute and contains text, because "allowed" and "all" are ambiguous:
        activeFilter: 'button[data-testid="network-connections-filter-active"]:contains("active")',
        allowedFilter:
            'button[data-testid="network-connections-filter-allowed"]:contains("allowed")',
        allFilter: 'button[data-testid="network-connections-filter-all"]:contains("all")',
        hideNsEdgesFilter: '[data-testid="namespace-flows-filter"] button:contains("Hide")',
        stopSimulation: '.simulator-mode button:contains("Stop")',
    },
    detailsPanel: scopeSelectors(networkPanels.detailsPanel, {
        header: '[data-testid="network-details-panel-header"]',
        search,
        table: {
            rows: '.rt-tbody .rt-tr',
        },
    }),
    networkEntityTabbedOverlay: scopeSelectors(networkEntityTabbedOverlay, {
        header: '[data-testid="network-entity-tabbed-overlay-header"]',
    }),
    toolbar: scopeSelectors('[data-testid="network-graph-toolbar"]', {
        namespaceSelect: '.namespace-select > button',
        filterSelect: search.multiSelectInput,
    }),
};
