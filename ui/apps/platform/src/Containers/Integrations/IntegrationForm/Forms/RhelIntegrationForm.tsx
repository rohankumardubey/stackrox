import React, { ReactElement } from 'react';
import { TextInput, PageSection, Form, Switch } from '@patternfly/react-core';
import * as yup from 'yup';

import usePageState from 'Containers/Integrations/hooks/usePageState';
import useIntegrationForm from '../useIntegrationForm';

import IntegrationFormActions from '../IntegrationFormActions';
import FormCancelButton from '../FormCancelButton';
import FormTestButton from '../FormTestButton';
import FormSaveButton from '../FormSaveButton';
import FormMessageBanner from '../FormMessageBanner';
import FormLabelGroup from '../FormLabelGroup';

export type RhelIntegration = {
    id?: string;
    name: string;
    categories: 'REGISTRY'[];
    docker: {
        endpoint: string;
        username: string;
        password: string;
    };
    skipTestIntegration: boolean;
    type: 'rhel';
    enabled: boolean;
    clusterIds: string[];
};

export type RhelIntegrationFormValues = {
    config: RhelIntegration;
    updatePassword: boolean;
};

export type RhelIntegrationFormProps = {
    initialValues: RhelIntegration | null;
    isEdittable?: boolean;
};

export const validationSchema = yup.object().shape({
    config: yup.object().shape({
        name: yup.string().required('Required'),
        categories: yup
            .array()
            .of(yup.string().oneOf(['REGISTRY']))
            .min(1, 'Must have at least one type selected')
            .required('Required'),
        docker: yup.object().shape({
            endpoint: yup.string().required('Required'),
            username: yup.string(),
            password: yup.string(),
        }),
        skipTestIntegration: yup.bool(),
        type: yup.string().matches(/rhel/),
        enabled: yup.bool(),
        clusterIds: yup.array().of(yup.string()),
    }),
    updatePassword: yup.bool(),
});

export const defaultValues: RhelIntegrationFormValues = {
    config: {
        name: '',
        categories: ['REGISTRY'],
        docker: {
            endpoint: '',
            username: '',
            password: '',
        },
        skipTestIntegration: false,
        type: 'rhel',
        enabled: true,
        clusterIds: [],
    },
    updatePassword: true,
};

function RhelIntegrationForm({
    initialValues = null,
    isEdittable = false,
}: RhelIntegrationFormProps): ReactElement {
    const formInitialValues = defaultValues;
    if (initialValues) {
        formInitialValues.config = { ...formInitialValues.config, ...initialValues };
        // We want to clear the password because backend returns '******' to represent that there
        // are currently stored credentials
        formInitialValues.config.docker.password = '';
    }
    const {
        values,
        errors,
        setFieldValue,
        isSubmitting,
        isTesting,
        onSave,
        onTest,
        onCancel,
        message,
    } = useIntegrationForm<RhelIntegrationFormValues, typeof validationSchema>({
        initialValues: formInitialValues,
        validationSchema,
    });
    const { isCreating } = usePageState();

    function onChange(value, event) {
        return setFieldValue(event.target.id, value, false);
    }

    return (
        <>
            {message && <FormMessageBanner message={message} />}
            <PageSection variant="light" isFilled hasOverflowScroll>
                <Form isWidthLimited>
                    <FormLabelGroup label="Name" isRequired fieldId="config.name" errors={errors}>
                        <TextInput
                            isRequired
                            type="text"
                            id="config.name"
                            name="config.name"
                            value={values.config.name}
                            placeholder="(ex. Red Hat Registry)"
                            onChange={onChange}
                            isDisabled={!isEdittable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Endpoint"
                        isRequired
                        fieldId="config.docker.endpoint"
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="config.docker.endpoint"
                            name="config.docker.endpoint"
                            value={values.config.docker.endpoint}
                            placeholder="(ex. registry.access.redhat.com)"
                            onChange={onChange}
                            isDisabled={!isEdittable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Username"
                        fieldId="config.docker.username"
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="config.docker.username"
                            name="config.docker.username"
                            value={values.config.docker.username}
                            onChange={onChange}
                            isDisabled={!isEdittable}
                        />
                    </FormLabelGroup>
                    {!isCreating && (
                        <FormLabelGroup
                            label="Update Password"
                            fieldId="updatePassword"
                            helperText="Setting this to false will use the currently stored credentials, if they exist."
                            errors={errors}
                        >
                            <Switch
                                id="updatePassword"
                                name="updatePassword"
                                aria-label="update password"
                                isChecked={values.updatePassword}
                                onChange={onChange}
                                isDisabled={!isEdittable}
                            />
                        </FormLabelGroup>
                    )}
                    {values.updatePassword && (
                        <FormLabelGroup
                            label="Password"
                            fieldId="config.docker.password"
                            errors={errors}
                        >
                            <TextInput
                                isRequired
                                type="password"
                                id="config.docker.password"
                                name="config.docker.password"
                                value={values.config.docker.password}
                                onChange={onChange}
                                isDisabled={!isEdittable}
                            />
                        </FormLabelGroup>
                    )}
                    <FormLabelGroup
                        label="Create Integration Without Testing"
                        fieldId="config.skipTestIntegration"
                        errors={errors}
                    >
                        <Switch
                            id="config.skipTestIntegration"
                            name="config.skipTestIntegration"
                            aria-label="skip test integration"
                            isChecked={values.config.skipTestIntegration}
                            onChange={onChange}
                            isDisabled={!isEdittable}
                        />
                    </FormLabelGroup>
                </Form>
            </PageSection>
            {isEdittable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                    >
                        Save
                    </FormSaveButton>
                    <FormTestButton
                        onTest={onTest}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default RhelIntegrationForm;
