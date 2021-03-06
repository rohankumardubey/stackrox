import React, { ReactElement, useState } from 'react';
import { Modal, ModalVariant } from '@patternfly/react-core';

import { importPolicies } from 'services/PoliciesService';
import { ListPolicy } from 'types/policy.proto';
import {
    parsePolicyImportErrors,
    getResolvedPolicies,
    getErrorMessages,
    checkDupeOnlyErrors,
} from 'Containers/Policies/Table/PolicyImport.utils';
import ImportPolicyJSONSuccess from './ImportPolicyJSONSuccess';
import ImportPolicyJSONModalError from './ImportPolicyJSONModalError';
import ImportPolicyJSONUpload from './ImportPolicyJSONUpload';

const RESOLUTION = { resolution: '', newName: '' };

type DuplicateErrors = {
    type: string;
    incomingName: string;
    incomingId: string;
    duplicateName: string;
};

type ImportPolicyJSONModalType = 'upload' | 'success' | 'error';

type ImportPolicyJSONModalProps = {
    cancelModal: () => void;
    isOpen: boolean;
    fetchPoliciesWithQuery: () => void;
};

function ImportPolicyJSONModal({
    cancelModal,
    isOpen,
    fetchPoliciesWithQuery,
}: ImportPolicyJSONModalProps): ReactElement {
    const [policies, setPolicies] = useState<ListPolicy[]>([]);
    const [duplicateErrors, setDuplicateErrors] = useState<DuplicateErrors[]>([]);
    const [duplicateResolution, setDuplicateResolution] = useState(RESOLUTION);
    const [modalType, setModalType] = useState<ImportPolicyJSONModalType>('upload');
    const [errorMessages, setErrorMessages] = useState<string[]>([]);

    function startImportPolicies() {
        // Note: this only resolves errors on one policy for MVP,
        //   see decision in comment on Jira story, https://stack-rox.atlassian.net/browse/ROX-4409
        const [policiesToImport, metadata] = getResolvedPolicies(
            policies,
            duplicateErrors,
            duplicateResolution
        );
        importPolicies(policiesToImport, metadata)
            .then((response) => {
                if (response.allSucceeded) {
                    setModalType('success');
                    // TODO: multiple policies import will be handled in
                    // https://stack-rox.atlassian.net/browse/ROX-8613
                    setPolicies([response.responses[0].policy]);
                    setTimeout(() => {
                        handleCancelModal();
                        fetchPoliciesWithQuery();
                    }, 3000);
                } else {
                    const errors = parsePolicyImportErrors(response?.responses);
                    const onlyHasDupeErrors = checkDupeOnlyErrors(errors);
                    if (onlyHasDupeErrors) {
                        setDuplicateErrors(errors[0]);
                    }
                    const errorMessageArray = getErrorMessages(errors[0]).map(
                        ({ msg }) => msg as string
                    );
                    setErrorMessages(errorMessageArray);
                    setModalType('error');
                }
            })
            .catch((err) => {
                setErrorMessages([`A network error occurred: ${err.message as string}`]);
                setModalType('error');
            });
    }

    function handleCancelModal() {
        setPolicies([]);
        setModalType('upload');
        setErrorMessages([]);
        cancelModal();
    }

    return (
        <Modal
            title="Import policy JSON"
            isOpen={isOpen}
            variant={ModalVariant.small}
            onClose={handleCancelModal}
            data-testid="import-policy-modal"
            aria-label="Import policy"
            hasNoBodyWrapper
        >
            {modalType === 'upload' && (
                <ImportPolicyJSONUpload
                    cancelModal={handleCancelModal}
                    startImportPolicies={startImportPolicies}
                    setPolicies={setPolicies}
                    policies={policies}
                />
            )}
            {modalType === 'error' && (
                <ImportPolicyJSONModalError
                    handleCancelModal={handleCancelModal}
                    policies={policies}
                    startImportPolicies={startImportPolicies}
                    duplicateErrors={duplicateErrors}
                    errorMessages={errorMessages}
                    duplicateResolution={duplicateResolution}
                    setDuplicateResolution={setDuplicateResolution}
                />
            )}
            {modalType === 'success' && (
                <ImportPolicyJSONSuccess policies={policies} handleCloseModal={handleCancelModal} />
            )}
        </Modal>
    );
}

export default ImportPolicyJSONModal;
