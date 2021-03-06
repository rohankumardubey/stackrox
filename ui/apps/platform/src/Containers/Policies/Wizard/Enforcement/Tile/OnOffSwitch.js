import React, { Component } from 'react';
import PropTypes from 'prop-types';

class OnOffSwitch extends Component {
    static propTypes = {
        enabled: PropTypes.bool.isRequired,
        applied: PropTypes.bool.isRequired,
        onAction: PropTypes.func.isRequired,
        offAction: PropTypes.func.isRequired,
    };

    renderSwitch = () => {
        const onColor = 'bg-success-600 border-success-600 text-base-100';
        const offColor = 'bg-alert-600 border-alert-600 text-base-100';
        const neutralColor = 'border-primary-300 text-primary-600';

        const onSwitchColor = this.props.enabled && this.props.applied ? onColor : neutralColor;
        const offSwitchColor = this.props.enabled && !this.props.applied ? offColor : neutralColor;

        const onSwitchClass = `px-2 py-1 border-2 bg-base-100 ${onSwitchColor} font-700 rounded-sm text-xs uppercase`;
        const offSwitchClass = `px-2 py-1 border-2 bg-base-100 ${offSwitchColor} font-700 rounded-sm text-xs uppercase`;

        return (
            <div
                className="flex py-2 w-full justify-center"
                data-testid="policy-enforcement-on-off"
            >
                <button
                    type="button"
                    className={onSwitchClass}
                    onClick={this.props.onAction}
                    disabled={!this.props.enabled}
                >
                    On
                </button>
                <button
                    type="button"
                    className={offSwitchClass}
                    onClick={this.props.offAction}
                    disabled={!this.props.enabled}
                >
                    Off
                </button>
            </div>
        );
    };

    render() {
        return <div className="flex bottom-0"> {this.renderSwitch()} </div>;
    }
}

export default OnOffSwitch;
