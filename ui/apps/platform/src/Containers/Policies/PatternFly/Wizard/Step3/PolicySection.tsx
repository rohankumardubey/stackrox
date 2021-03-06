import React from 'react';
import { useFormikContext } from 'formik';
import {
    Card,
    CardHeader,
    CardTitle,
    CardActions,
    CardBody,
    Button,
    Divider,
    Flex,
    FlexItem,
    TextInput,
} from '@patternfly/react-core';
import { PencilAltIcon, TrashIcon, CheckIcon } from '@patternfly/react-icons';

import { Policy } from 'types/policy.proto';
import { Descriptor } from 'Containers/Policies/Wizard/Form/descriptors';

import PolicyGroupCard from './PolicyGroupCard';
import PolicySectionDropTarget from './PolicySectionDropTarget';

import './PolicySection.css';

type PolicySectionProps = {
    sectionIndex: number;
    descriptors: Descriptor[];
    readOnly?: boolean;
};

function PolicySection({ sectionIndex, descriptors, readOnly = false }: PolicySectionProps) {
    const [isEditingName, setIsEditingName] = React.useState(false);
    const { values, setFieldValue, handleChange } = useFormikContext<Policy>();
    const { sectionName, policyGroups } = values.policySections[sectionIndex];

    function onEditSectionName(_, e) {
        handleChange(e);
    }

    function onDeleteSection() {
        setFieldValue(
            'policySections',
            values.policySections.filter((_, i) => i !== sectionIndex)
        );
    }

    return (
        <Card isFlat isCompact className={!readOnly ? 'policy-section-card' : ''}>
            <CardHeader className="policy-section-card-header pf-u-p-0">
                <CardTitle className="pf-u-display-flex pf-u-align-self-stretch">
                    <Flex
                        alignItems={{ default: 'alignItemsCenter' }}
                        flexWrap={{ default: 'nowrap' }}
                    >
                        <FlexItem className="pf-u-pl-md">{sectionIndex + 1}</FlexItem>
                        <Divider component="div" isVertical />
                        <FlexItem>
                            {isEditingName ? (
                                <TextInput
                                    id={`policySections[${sectionIndex}].sectionName`}
                                    name={`policySections[${sectionIndex}].sectionName`}
                                    value={values.policySections[sectionIndex].sectionName}
                                    onChange={onEditSectionName}
                                />
                            ) : (
                                <div className="pf-u-py-sm" data-testid="policy-section-name">
                                    {sectionName}
                                </div>
                            )}
                        </FlexItem>
                    </Flex>
                </CardTitle>
                {!readOnly && (
                    <CardActions hasNoOffset>
                        <Button
                            variant="plain"
                            className="pf-u-px-sm"
                            onClick={() => setIsEditingName(!isEditingName)}
                        >
                            {isEditingName ? (
                                <CheckIcon data-testid="save-section-name-btn" />
                            ) : (
                                <PencilAltIcon data-testid="edit-section-name-btn" />
                            )}
                        </Button>
                        <Divider component="div" isVertical />
                        <Button
                            variant="plain"
                            className="pf-u-mr-xs pf-u-px-sm pf-u-py-md"
                            data-testid="delete-section-btn"
                            onClick={onDeleteSection}
                        >
                            <TrashIcon />
                        </Button>
                    </CardActions>
                )}
            </CardHeader>
            <CardBody className="policy-section-card-body">
                {policyGroups.map((group, groupIndex) => {
                    const descriptor = descriptors.find(
                        (descriptorField) =>
                            group.fieldName === descriptorField.name ||
                            group.fieldName === descriptorField.label
                    );
                    return (
                        descriptor && (
                            <PolicyGroupCard
                                key={descriptor.name}
                                descriptor={descriptor}
                                groupIndex={groupIndex}
                                sectionIndex={sectionIndex}
                                readOnly={readOnly}
                            />
                        )
                    );
                })}
                {!readOnly && (
                    <PolicySectionDropTarget
                        sectionIndex={sectionIndex}
                        descriptors={descriptors}
                    />
                )}
            </CardBody>
        </Card>
    );
}

export default PolicySection;
