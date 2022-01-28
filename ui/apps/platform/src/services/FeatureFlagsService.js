import axios from './instance';

const url = '/v1/featureflags';

/**
 * Fetches the list of feature flags and their current values from the backend.
 */
// eslint-disable-next-line import/prefer-default-export
export function fetchFeatureFlags() {
    return axios.get(url).then((response) => ({
        response: response.data,
    }));
}