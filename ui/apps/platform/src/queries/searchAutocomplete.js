import { gql } from '@apollo/client';

const SEARCH_AUTOCOMPLETE_QUERY = gql`
    query autocomplete($query: String!, $categories: [SearchCategory!]) {
        searchAutocomplete(query: $query, categories: $categories)
    }
`;

export default SEARCH_AUTOCOMPLETE_QUERY;
