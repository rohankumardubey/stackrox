import React, { useContext } from 'react';
import entityTypes from 'constants/entityTypes';
import { gql, useQuery } from '@apollo/client';

import workflowStateContext from 'Containers/workflowStateContext';
import EntityTileLink from 'Components/EntityTileLink';
import queryService from 'utils/queryService';

const POLICIES_COUNT_QUERY = gql`
    query policiesCount($query: String) {
        policies(query: $query) {
            id
            alertCount
        }
    }
`;

const PoliciesCountTile = () => {
    const { loading, data = {} } = useQuery(POLICIES_COUNT_QUERY, {
        variables: {
            query: queryService.objectToWhereClause({
                Category: 'Vulnerability Management',
            }),
        },
    });

    const { policies = [] } = data;

    const policyCount = policies.length;
    const failingPoliciesCount = policies.reduce((sum, policy) => {
        return policy.alertCount ? sum + 1 : sum;
    }, 0);
    const failingPoliciesCountText = `(${failingPoliciesCount} failing)`;

    const workflowState = useContext(workflowStateContext);
    const url = workflowState.clear().pushList(entityTypes.POLICY).toUrl();

    return (
        <EntityTileLink
            count={policyCount}
            entityType={entityTypes.POLICY}
            position="first"
            short
            subText={failingPoliciesCountText}
            loading={loading}
            isError={!!failingPoliciesCount}
            url={url}
        />
    );
};

export default PoliciesCountTile;
