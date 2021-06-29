import React, { ReactElement } from 'react';
import { connect } from 'react-redux';
import { ArrowLeft, ArrowRight } from 'react-feather';
import { createStructuredSelector } from 'reselect';
import { formValueSelector } from 'redux-form';

import useFeatureFlagEnabled from 'hooks/useFeatureFlagEnabled';
import { knownBackendFlags } from 'utils/featureFlags';
import { actions as wizardActions } from 'reducers/policies/wizard';
import wizardStages from 'Containers/Policies/Wizard/wizardStages';
import PanelButton from 'Components/PanelButton';
import SavePolicyButton from '../SavePolicyButton';

type PreviewButtonsProps = {
    hasAuditLogEventSource: boolean;
    setWizardStage: (string) => void;
};

function PreviewButtons({
    hasAuditLogEventSource,
    setWizardStage,
}: PreviewButtonsProps): ReactElement {
    const auditLogEnabled = useFeatureFlagEnabled(knownBackendFlags.ROX_K8S_AUDIT_LOG_DETECTION);

    function goBackToEditBPL() {
        setWizardStage(wizardStages.editBPL);
    }

    function goToEnforcement() {
        setWizardStage(wizardStages.enforcement);
    }

    const skipEnforcement = auditLogEnabled && hasAuditLogEventSource;

    return (
        <>
            <PanelButton
                icon={<ArrowLeft className="h-4 w-4" />}
                className="btn btn-base mr-2"
                onClick={goBackToEditBPL}
                tooltip="Back to previous step"
            >
                Previous
            </PanelButton>
            {skipEnforcement ? (
                <SavePolicyButton />
            ) : (
                <PanelButton
                    icon={<ArrowRight className="h-4 w-4" />}
                    className="btn btn-base mr-2"
                    onClick={goToEnforcement}
                    tooltip="Go to next step"
                >
                    Next
                </PanelButton>
            )}
        </>
    );
}

const mapStateToProps = createStructuredSelector({
    hasAuditLogEventSource: (state) => {
        const eventSourceValue = formValueSelector('policyCreationForm')(state, 'eventSource');
        return eventSourceValue === 'AUDIT_LOG_EVENT';
    },
});

const mapDispatchToProps = {
    setWizardStage: wizardActions.setWizardStage,
};

export default connect(mapStateToProps, mapDispatchToProps)(PreviewButtons);