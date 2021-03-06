import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { selectors } from 'reducers';
import { createStructuredSelector } from 'reselect';

import fieldsMap from 'Containers/Policies/Wizard/Details/descriptors';

class Fields extends Component {
    static propTypes = {
        clustersById: PropTypes.shape({}).isRequired,
        policy: PropTypes.shape({}).isRequired,
        notifiers: PropTypes.arrayOf(
            PropTypes.shape({
                name: PropTypes.string.isRequired,
            })
        ).isRequired,
    };

    render() {
        const policy = { ...this.props.policy };
        const fields = Object.keys(policy);
        if (!fields) {
            return '';
        }

        return (
            <div className="px-3 pt-5">
                <div className="bg-base-100 shadow" data-testid="policy-details">
                    <div className="p-3 pb-2 border-b border-base-300 text-base-600 font-700 text-lg leading-normal">
                        Policy Details
                    </div>
                    <div className="h-full p-3 pb-1">
                        {fields.map((field) => {
                            if (!fieldsMap[field]) {
                                return '';
                            }
                            if (policy[field] === undefined) {
                                return '';
                            }
                            const { label } = fieldsMap[field];
                            const value = fieldsMap[field].formatValue(policy[field], {
                                clustersById: this.props.clustersById,
                                notifiers: this.props.notifiers,
                            });
                            if (!value) {
                                return '';
                            }
                            if (Array.isArray(value)) {
                                return (
                                    <div className="mb-4" key={field}>
                                        <div className="text-base-600 font-700">{label}:</div>
                                        {value.map((v) => (
                                            <div key={v} className="flex pt-1 leading-normal">
                                                {v}
                                            </div>
                                        ))}
                                    </div>
                                );
                            }
                            if (typeof value === 'object') {
                                return (
                                    <div key={field}>
                                        {Object.keys(value).map((key) => (
                                            <div className="mb-4" key={field}>
                                                <div className="text-base-600 font-700">{key}:</div>
                                                {value[key].map((v, i) => (
                                                    <div
                                                        // eslint-disable-next-line react/no-array-index-key
                                                        key={i}
                                                        className="flex pt-1 leading-normal"
                                                    >
                                                        {v}
                                                    </div>
                                                ))}
                                            </div>
                                        ))}
                                    </div>
                                );
                            }
                            return (
                                <div className="mb-4" key={field}>
                                    <div className="text-base-600 font-700">{label}:</div>
                                    <div className="flex pt-1 leading-normal">{value}</div>
                                </div>
                            );
                        })}
                    </div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = createStructuredSelector({
    clustersById: selectors.getClustersById,
    notifiers: selectors.getNotifiers,
});

export default connect(mapStateToProps)(Fields);
