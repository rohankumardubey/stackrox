import React, { ReactElement } from 'react';
import { TextInput, PageSection, Form, Checkbox, SelectOption } from '@patternfly/react-core';
import * as yup from 'yup';

import FormMultiSelect from 'Components/FormMultiSelect';
import usePageState from 'Containers/Integrations/hooks/usePageState';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormCancelButton from '../FormCancelButton';
import FormTestButton from '../FormTestButton';
import FormSaveButton from '../FormSaveButton';
import FormMessage from '../FormMessage';
import FormLabelGroup from '../FormLabelGroup';

export type DtrIntegration = {
    id?: string;
    name: string;
    categories: ('REGISTRY' | 'SCANNER')[];
    dtr: {
        endpoint: string;
        username: string;
        password: string;
        insecure: boolean;
    };
    skipTestIntegration: boolean;
    type: 'dtr';
    enabled: boolean;
    clusterIds: string[];
};

export type DtrIntegrationFormValues = {
    config: DtrIntegration;
    updatePassword: boolean;
};

export const validationSchema = yup.object().shape({
    config: yup.object().shape({
        name: yup.string().trim().required('An integration name is required'),
        categories: yup
            .array()
            .of(yup.string().trim().oneOf(['REGISTRY', 'SCANNER']))
            .min(1, 'Must have at least one type selected')
            .required('Required'),
        dtr: yup.object().shape({
            endpoint: yup.string().trim().required('An endpoint is required').min(1),
            username: yup.string(),
            password: yup
                .string()
                .test(
                    'password-test',
                    'A password is required',
                    (value, context: yup.TestContext) => {
                        const requirePasswordField =
                            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                            // @ts-ignore
                            context?.from[2]?.value?.updatePassword || false;

                        if (!requirePasswordField) {
                            return true;
                        }

                        const trimmedValue = value?.trim();
                        return !!trimmedValue;
                    }
                ),
            insecure: yup.bool(),
        }),
        skipTestIntegration: yup.bool(),
        type: yup.string().matches(/dtr/),
        enabled: yup.bool(),
        clusterIds: yup.array().of(yup.string()),
    }),
    updatePassword: yup.bool(),
});

export const defaultValues: DtrIntegrationFormValues = {
    config: {
        name: '',
        categories: [],
        dtr: {
            endpoint: '',
            username: '',
            password: '',
            insecure: false,
        },
        skipTestIntegration: false,
        type: 'dtr',
        enabled: true,
        clusterIds: [],
    },
    updatePassword: true,
};

function DtrIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<DtrIntegration>): ReactElement {
    const formInitialValues = defaultValues;
    if (initialValues) {
        formInitialValues.config = { ...formInitialValues.config, ...initialValues };
        // We want to clear the password because backend returns '******' to represent that there
        // are currently stored credentials
        formInitialValues.config.dtr.password = '';
    }
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
    } = useIntegrationForm<DtrIntegrationFormValues, typeof validationSchema>({
        initialValues: formInitialValues,
        validationSchema,
    });
    const { isCreating } = usePageState();

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    function onCustomChange(id, value) {
        return setFieldValue(id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                {message && <FormMessage message={message} />}
                <Form isWidthLimited>
                    <FormLabelGroup
                        label="Integration name"
                        isRequired
                        fieldId="config.name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="config.name"
                            placeholder="(ex. Prod Docker Trusted Registry)"
                            value={values.config.name}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Type"
                        isRequired
                        fieldId="config.categories"
                        touched={touched}
                        errors={errors}
                    >
                        <FormMultiSelect
                            id="config.categories"
                            values={values.config.categories}
                            onChange={onCustomChange}
                            isDisabled={!isEditable}
                        >
                            <SelectOption key={0} value="REGISTRY">
                                Registry
                            </SelectOption>
                            <SelectOption key={1} value="SCANNER">
                                Scanner
                            </SelectOption>
                        </FormMultiSelect>
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Endpoint"
                        isRequired
                        fieldId="config.dtr.endpoint"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="config.dtr.endpoint"
                            placeholder="(ex. dtr.example.com)"
                            value={values.config.dtr.endpoint}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Username"
                        fieldId="config.dtr.username"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="config.dtr.username"
                            value={values.config.dtr.username}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    {!isCreating && (
                        <FormLabelGroup
                            fieldId="updatePassword"
                            helperText="Leave this off to use the currently stored credentials."
                            errors={errors}
                        >
                            <Checkbox
                                label="Update stored credentials"
                                id="updatePassword"
                                isChecked={values.updatePassword}
                                onChange={onChange}
                                onBlur={handleBlur}
                                isDisabled={!isEditable}
                            />
                        </FormLabelGroup>
                    )}
                    <FormLabelGroup
                        isRequired={values.updatePassword}
                        label="Password"
                        fieldId="config.dtr.password"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired={values.updatePassword}
                            type="password"
                            id="config.dtr.password"
                            value={values.config.dtr.password}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable || !values.updatePassword}
                            placeholder={
                                values.updatePassword
                                    ? ''
                                    : 'Currently-stored password will be used.'
                            }
                        />
                    </FormLabelGroup>
                    <FormLabelGroup fieldId="config.dtr.insecure" touched={touched} errors={errors}>
                        <Checkbox
                            label="Disable TLS certificate validation (insecure)"
                            id="config.dtr.insecure"
                            isChecked={values.config.dtr.insecure}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        fieldId="config.skipTestIntegration"
                        touched={touched}
                        errors={errors}
                    >
                        <Checkbox
                            label="Create integration without testing"
                            id="config.skipTestIntegration"
                            isChecked={values.config.skipTestIntegration}
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

export default DtrIntegrationForm;
