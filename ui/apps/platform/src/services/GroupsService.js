import axios from './instance';

const url = '/v1/groups';
const updateUrl = '/v1/groupsbatch';

/**
 * Fetches list of groups of rules
 *
 * @returns {Promise<Object, Error>} fulfilled with array of groups
 */
export function fetchGroups() {
    return axios.get(url).then((response) => ({
        response: response.data,
    }));
}

/**
 * Update/Add a group rule.
 *
 * @returns {Promise<Object, Error>}
 */
export function updateOrAddGroup({ oldGroups, newGroups }) {
    return axios.post(updateUrl, { previous_groups: oldGroups, required_groups: newGroups });
}

/**
 * Deletes a group rule.
 *
 * @returns {Promise<Object, Error>}
 */
export function deleteRuleGroup(data) {
    const { key, authProviderId, value } = data.props;
    return axios.delete(`${url}?authProviderId=${authProviderId}&key=${key}&value=${value}`, {});
}
