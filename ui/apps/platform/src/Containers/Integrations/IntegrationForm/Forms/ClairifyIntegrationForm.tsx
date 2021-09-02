import React, { ReactElement } from 'react';
import { TextInput, SelectOption, PageSection, Form } from '@patternfly/react-core';
import * as yup from 'yup';

import FormMultiSelect from 'Components/FormMultiSelect';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormCancelButton from '../FormCancelButton';
import FormTestButton from '../FormTestButton';
import FormSaveButton from '../FormSaveButton';
import FormMessage from '../FormMessage';
import FormLabelGroup from '../FormLabelGroup';

export type ClairifyIntegration = {
    id?: string;
    name: string;
    categories: ('NODE_SCANNER' | 'SCANNER')[];
    clairify: {
        endpoint: string;
        grpcEndpoint: string;
        numConcurrentScans: string;
    };
    type: 'clairify';
    enabled: boolean;
    clusterIds: string[];
};

export const validationSchema = yup.object().shape({
    name: yup.string().trim().required('An integration name is required'),
    categories: yup
        .array()
        .of(yup.string().trim().oneOf(['NODE_SCANNER', 'SCANNER']))
        .min(1, 'Must have at least one type selected')
        .required('Required'),
    clairify: yup.object().shape({
        endpoint: yup.string().trim().required('An endpoint is required'),
        grpcEndpoint: yup.string().trim(),
        numConcurrentScans: yup.string().trim(),
    }),
    type: yup.string().matches(/clairify/),
    enabled: yup.bool(),
    clusterIds: yup.array().of(yup.string()),
});

export const defaultValues: ClairifyIntegration = {
    name: '',
    categories: [],
    clairify: {
        endpoint: '',
        grpcEndpoint: '',
        numConcurrentScans: '0',
    },
    type: 'clairify',
    enabled: true,
    clusterIds: [],
};

function ClairifyIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<ClairifyIntegration>): ReactElement {
    const formInitialValues = initialValues
        ? { ...defaultValues, ...initialValues }
        : defaultValues;
    const {
        values,
        touched,
        errors,
        dirty,
        isValid,
        setFieldValue,
        handleBlur,
        isSubmitting,
        isTesting,
        onSave,
        onTest,
        onCancel,
        message,
    } = useIntegrationForm<ClairifyIntegration, typeof validationSchema>({
        initialValues: formInitialValues,
        validationSchema,
    });

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    function onCustomChange(id, value) {
        return setFieldValue(id, value, false);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                {message && <FormMessage message={message} />}
                <Form isWidthLimited>
                    <FormLabelGroup
                        label="Integration name"
                        isRequired
                        fieldId="name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="name"
                            value={values.name}
                            placeholder="(example, StackRox Scanner Integration)"
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Type"
                        isRequired
                        fieldId="categories"
                        touched={touched}
                        errors={errors}
                    >
                        <FormMultiSelect
                            id="categories"
                            values={values.categories}
                            onChange={onCustomChange}
                            isDisabled={!isEditable}
                        >
                            <SelectOption key={0} value="SCANNER">
                                Image Scanner
                            </SelectOption>
                            <SelectOption key={1} value="NODE_SCANNER">
                                Node Scanner
                            </SelectOption>
                        </FormMultiSelect>
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Endpoint"
                        isRequired
                        fieldId="clairify.endpoint"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="clairify.endpoint"
                            value={values.clairify.endpoint}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="GRPC endpoint"
                        fieldId="clairify.grpcEndpoint"
                        helperText="Used For Node Scanning"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="clairify.grpcEndpoint"
                            value={values.clairify.grpcEndpoint}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Max concurrent image scans"
                        fieldId="clairify.numConcurrentScans"
                        helperText="0 for default"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="number"
                            id="clairify.numConcurrentScans"
                            value={values.clairify.numConcurrentScans}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                </Form>
            </PageSection>
            {isEditable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!dirty || !isValid}
                    >
                        Save
                    </FormSaveButton>
                    <FormTestButton
                        onTest={onTest}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isValid={isValid}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default ClairifyIntegrationForm;
