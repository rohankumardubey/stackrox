import React from 'react';
import { useFormikContext } from 'formik';
import {
    Flex,
    FlexItem,
    Title,
    Button,
    Divider,
    Grid,
    GridItem,
    Select,
    SelectVariant,
    SelectOption,
    Form,
    FormGroup,
} from '@patternfly/react-core';

import { Policy } from 'types/policy.proto';
import { Image } from 'types/image.proto';
import { ListDeployment } from 'types/deployment.proto';
import { fetchImages } from 'services/ImagesService';
import { fetchDeployments } from 'services/DeploymentsService';
import PolicyScopeCard from './PolicyScopeCard';

const MAX_INCLUSION_SCOPES = 10;
const MAX_EXCLUSION_SCOPES = 10;

function PolicyScopeForm({ clusters }) {
    const [isExcludeImagesOpen, setIsExcludeImagesOpen] = React.useState(false);
    const [images, setImages] = React.useState<Image[]>([]);
    const [deployments, setDeployments] = React.useState<ListDeployment[]>([]);
    const { values, setFieldValue } = useFormikContext<Policy>();
    const { scope, excludedDeploymentScopes, excludedImageNames } = values;

    const hasAuditLogEventSource = values.eventSource === 'AUDIT_LOG_EVENT';
    const hasBuildLifecycle = values.lifecycleStages.includes('BUILD');

    function addNewInclusionScope() {
        if (scope.length < MAX_INCLUSION_SCOPES) {
            setFieldValue('scope', [...scope, {}]);
        }
    }

    function deleteInclusionScope(index) {
        const newScope = scope.filter((_, i) => i !== index);
        setFieldValue('scope', newScope);
    }

    function addNewExclusionDeploymentScope() {
        if (excludedDeploymentScopes.length < MAX_EXCLUSION_SCOPES) {
            setFieldValue('excludedDeploymentScopes', [...excludedDeploymentScopes, {}]);
        }
    }

    function deleteExclusionDeploymentScope(index) {
        const newScope = scope.filter((_, i) => i !== index);
        setFieldValue('excludedDeploymentScopes', newScope);
    }

    function handleChangeMultiSelect(e, selectedImage) {
        setIsExcludeImagesOpen(false);
        if (excludedImageNames.includes(selectedImage)) {
            const newExclusions = excludedImageNames.filter((image) => image !== selectedImage);
            setFieldValue('excludedImageNames', newExclusions);
        } else {
            setFieldValue('excludedImageNames', [...excludedImageNames, selectedImage]);
        }
    }

    React.useEffect(() => {
        fetchImages()
            .then((response) => {
                setImages(response);
            })
            .catch(() => {
                // TODO
            });

        fetchDeployments()
            .then((response) => {
                const deploymentList = response.map((item) => item.deployment as ListDeployment);
                setDeployments(deploymentList);
            })
            .catch(() => {
                // TODO
            });
    }, []);

    return (
        <Flex direction={{ default: 'column' }}>
            <FlexItem flex={{ default: 'flex_1' }}>
                <Title headingLevel="h2">Policy scope</Title>
                <div className="pf-u-mt-sm">
                    Create scopes to restrict or exclude your policy from entities within your
                    environment.
                </div>
            </FlexItem>
            <Divider component="div" />
            <Flex direction={{ default: 'column' }}>
                <Flex>
                    <FlexItem flex={{ default: 'flex_1' }}>
                        <Title headingLevel="h3">Restrict by scope</Title>
                        <div className="pf-u-mt-sm">
                            Use Restrict by scope to enable this policy only for a specific cluster,
                            namespace, or label. You can add multiple scope and also use regular
                            expressions (RE2 syntax) for namespaces and labels.
                        </div>
                    </FlexItem>
                    <FlexItem className="pf-u-pr-md" alignSelf={{ default: 'alignSelfCenter' }}>
                        <Button variant="secondary" onClick={addNewInclusionScope}>
                            Add inclusion scope
                        </Button>
                    </FlexItem>
                </Flex>
                <FlexItem>
                    <Grid hasGutter>
                        {scope.map((_, index) => (
                            // eslint-disable-next-line react/no-array-index-key
                            <GridItem span={4} rowSpan={4} key={index}>
                                <PolicyScopeCard
                                    type="inclusion"
                                    name={`scope[${index}]`}
                                    index={index}
                                    clusters={clusters}
                                    onDelete={() => deleteInclusionScope(index)}
                                    hasAuditLogEventSource={hasAuditLogEventSource}
                                />
                            </GridItem>
                        ))}
                    </Grid>
                </FlexItem>
            </Flex>
            <Divider component="div" />
            <Flex direction={{ default: 'column' }}>
                <Flex>
                    <FlexItem flex={{ default: 'flex_1' }}>
                        <Title headingLevel="h3">Exclude by scope</Title>
                        <div className="pf-u-mt-sm">
                            Use Exclude by scope to exclude entities from your policy. You can add
                            multiple scopes and also use regular expressions (RE2 syntax) for
                            namespaces and labels. However, you can&apos;t use regular expressions
                            for selecting deployments.
                        </div>
                    </FlexItem>
                    <FlexItem className="pf-u-pr-md" alignSelf={{ default: 'alignSelfCenter' }}>
                        <Button variant="secondary" onClick={addNewExclusionDeploymentScope}>
                            Add exclusion scope
                        </Button>
                    </FlexItem>
                </Flex>
                <FlexItem>
                    <Grid hasGutter>
                        {excludedDeploymentScopes.map((_, index) => (
                            // eslint-disable-next-line react/no-array-index-key
                            <GridItem span={4} rowSpan={4} key={index}>
                                <PolicyScopeCard
                                    type="exclusion"
                                    name={`excludedDeploymentScopes[${index}]`}
                                    index={index}
                                    clusters={clusters}
                                    deployments={deployments}
                                    onDelete={() => deleteExclusionDeploymentScope(index)}
                                    hasAuditLogEventSource={hasAuditLogEventSource}
                                />
                            </GridItem>
                        ))}
                    </Grid>
                </FlexItem>
            </Flex>
            <Divider component="div" />
            <Flex direction={{ default: 'column' }}>
                <FlexItem flex={{ default: 'flex_1' }}>
                    <Title headingLevel="h3">Exclude images</Title>
                    <div className="pf-u-mt-sm">
                        The exclude images setting only applies when you check images in a
                        continuous integration system (the Build lifecycle stage). It won&apos;t
                        have any effect if you use this policy to check running deployments (the
                        Deploy lifecycle stage) or runtime activities (the Run lifecycle stage).
                    </div>
                </FlexItem>
                <FlexItem>
                    <Form>
                        <FormGroup
                            label="Exclude images (Build lifecycle only)"
                            fieldId="exclude-images"
                            helperText="Select all images from the list for which you don't want to trigger a violation for the policy."
                        >
                            <Select
                                onToggle={() => setIsExcludeImagesOpen(!isExcludeImagesOpen)}
                                isOpen={isExcludeImagesOpen}
                                variant={SelectVariant.typeaheadMulti}
                                selections={excludedImageNames}
                                onSelect={handleChangeMultiSelect}
                                isDisabled={hasAuditLogEventSource || !hasBuildLifecycle}
                                onClear={() => setFieldValue('excludedImageNames', [])}
                                placeholderText="Select images to exclude"
                            >
                                {images.map((image) => (
                                    <SelectOption key={image.name} value={image.name}>
                                        {image.name}
                                    </SelectOption>
                                ))}
                            </Select>
                        </FormGroup>
                    </Form>
                </FlexItem>
            </Flex>
        </Flex>
    );
}

export default PolicyScopeForm;