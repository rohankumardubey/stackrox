import React, { useContext } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { useQuery } from 'react-apollo';
import gql from 'graphql-tag';
import sortBy from 'lodash/sortBy';
import { format } from 'date-fns';

import workflowStateContext from 'Containers/workflowStateContext';

import Button from 'Components/Button';
import Loader from 'Components/Loader';
import Widget from 'Components/Widget';
import LabeledBarGraph from 'Components/visuals/LabeledBarGraph';
import NoResultsMessage from 'Components/NoResultsMessage';
import dateTimeFormat from 'constants/dateTimeFormat';
import entityTypes from 'constants/entityTypes';
import { severityLabels } from 'messages/common';
import queryService from 'modules/queryService';

const FREQUENTLY_VIOLATED_POLICIES = gql`
    query frequentlyViolatedPolicies($query: String) {
        results: policies(query: $query) {
            id
            name
            enforcementActions
            severity
            alertCount
            categories
            description
            latestViolation
        }
    }
`;

const ViewAllButton = ({ url }) => {
    return (
        <Link to={url} className="no-underline">
            <Button className="btn-sm btn-base" type="button" text="View All" />
        </Link>
    );
};

const processData = (data, workflowState, limit) => {
    const results = sortBy(data.results, ['alertCount']).slice(-limit); // @TODO: Remove when we have pagination on Policies
    return results
        .filter(datum => datum.alertCount)
        .map(datum => {
            const {
                id,
                name,
                description,
                enforcementActions,
                severity,
                alertCount,
                latestViolation,
                categories
            } = datum;
            const url = workflowState.pushRelatedEntity(entityTypes.POLICY, id).toUrl();
            const isEnforced = enforcementActions.length ? 'Yes' : 'No';
            const categoriesStr = categories.join(', ');

            const tooltipBody = (
                <ul className="flex-1 list-reset border-base-300 overflow-hidden">
                    <li className="py-1" key="categories">
                        <span className="text-base-600 font-700 mr-2">Category:</span>
                        <span className="font-600">{categoriesStr}</span>
                    </li>
                    <li className="py-1" key="description">
                        <span className="text-base-600 font-700 mr-2">Description:</span>
                        <span className="font-600">{description}</span>
                    </li>
                    <li className="py-1" key="latestViolation">
                        <span className="text-base-600 font-700 mr-2">Last violated:</span>
                        <span className="font-600">{format(latestViolation, dateTimeFormat)}</span>
                    </li>
                </ul>
            );

            return {
                x: alertCount,
                y: `${name} / Enforced: ${isEnforced} / Severity: ${severityLabels[severity]}`,
                url,
                hint: { title: name, body: tooltipBody }
            };
        });
};

const FrequentlyViolatedPolicies = ({ entityContext, limit }) => {
    const { loading, data = {} } = useQuery(FREQUENTLY_VIOLATED_POLICIES, {
        variables: {
            query: `${queryService.entityContextToQueryString(entityContext)}+
            ${queryService.objectToWhereClause({ Category: 'Vulnerability Management' })}`
        }
    });

    let content = <Loader />;

    const workflowState = useContext(workflowStateContext);
    if (!loading) {
        const processedData = processData(data, workflowState, limit);

        if (!processedData || processedData.length === 0) {
            content = (
                <NoResultsMessage
                    message="No deployments with policy violations found"
                    className="p-6"
                    icon="info"
                />
            );
        } else {
            content = <LabeledBarGraph data={processedData} title="Failing Deployments" />;
        }
    }

    const viewAllURL = workflowState
        .pushList(entityTypes.POLICY)
        .setSort([{ id: 'policyStatus', desc: false }, { id: 'severity', desc: false }])
        .toUrl();

    return (
        <Widget
            className="h-full pdf-page"
            bodyClassName="px-2"
            header="Frequently Violated Policies"
            headerComponents={<ViewAllButton url={viewAllURL} />}
        >
            {content}
        </Widget>
    );
};

FrequentlyViolatedPolicies.propTypes = {
    entityContext: PropTypes.shape({}),
    limit: PropTypes.number
};

FrequentlyViolatedPolicies.defaultProps = {
    entityContext: {},
    limit: 9
};

export default FrequentlyViolatedPolicies;
