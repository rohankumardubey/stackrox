/* eslint-disable no-nested-ternary */
/* eslint-disable react/no-array-index-key */
import React, { ReactElement } from 'react';
import { Bullseye, PageSection, PageSectionVariants, Spinner } from '@patternfly/react-core';

import usePagination from 'hooks/patternfly/usePagination';
import ACSEmptyState from 'Components/ACSEmptyState';
import DeferredCVEsTable from './DeferredCVEsTable';
import useImageVulnerabilities from '../useImageVulnerabilities';

type DeferredCVEsProps = {
    imageId: string;
};

function DeferredCVEs({ imageId }: DeferredCVEsProps): ReactElement {
    const { page, perPage, onSetPage, onPerPageSelect } = usePagination();
    const { isLoading, data, refetchQuery } = useImageVulnerabilities({
        imageId,
        vulnsQuery: 'Vulnerability State:DEFERRED',
        pagination: {
            limit: perPage,
            offset: (page - 1) * perPage,
            sortOption: {
                field: 'Severity',
                reversed: true,
            },
        },
    });

    if (isLoading) {
        return (
            <Bullseye>
                <Spinner isSVG size="sm" />
            </Bullseye>
        );
    }

    const itemCount = data?.image?.vulnCount || 0;
    const rows = data?.image?.vulns || [];

    if (!isLoading && rows && rows.length === 0) {
        return (
            <PageSection variant={PageSectionVariants.light} isFilled>
                <ACSEmptyState title="No deferral requests were approved." />
            </PageSection>
        );
    }

    return (
        <DeferredCVEsTable
            rows={rows}
            isLoading={isLoading}
            itemCount={itemCount}
            page={page}
            perPage={perPage}
            onSetPage={onSetPage}
            onPerPageSelect={onPerPageSelect}
            updateTable={refetchQuery}
        />
    );
}

export default DeferredCVEs;
