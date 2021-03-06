import React from 'react';
import PropTypes from 'prop-types';

import ReduxSelectField from 'Components/forms/ReduxSelectField';
import ReduxTextField from 'Components/forms/ReduxTextField';
import ReduxTextAreaField from 'Components/forms/ReduxTextAreaField';
import ReduxCheckboxField from 'Components/forms/ReduxCheckboxField';
import ReduxMultiSelectField from 'Components/forms/ReduxMultiSelectField';
import ReduxMultiSelectCreatableField from 'Components/forms/ReduxMultiSelectCreatableField';
import ReduxNumericInputField from 'Components/forms/ReduxNumericInputField';
import ReduxToggleField from 'Components/forms/ReduxToggleField';
import ReduxRadioButtonGroupField from 'Components/forms/ReduxRadioButtonGroupField';

export default function Field({ field, name, readOnly }) {
    if (field === undefined) {
        return null;
    }
    // this is to accomodate for recursive Fields (when type is 'group')
    const path = field.subpath ? name : `${name}.value`;

    switch (field.type) {
        case 'text':
            return (
                <label className="flex flex-col flex-1">
                    {field.label?.length > 0 && !field.hideInnerLabel && (
                        <span className="font-600 pl-1 mb-1">{field.label}</span>
                    )}
                    <ReduxTextField
                        key={path}
                        name={path}
                        disabled={readOnly || field.disabled}
                        placeholder={field.placeholder}
                    />
                </label>
            );
        case 'checkbox':
            return <ReduxCheckboxField name={path} disabled={readOnly || field.disabled} />;
        case 'toggle':
            return (
                <ReduxToggleField
                    name={path}
                    key={path}
                    disabled={readOnly || field.disabled}
                    reverse={field.reverse}
                    className="self-center"
                />
            );
        case 'radioGroup':
            return (
                <ReduxRadioButtonGroupField
                    name={path}
                    key={path}
                    buttons={field.radioButtons}
                    groupClassName="w-full"
                    disabled={readOnly}
                />
            );
        case 'select':
            return (
                <label className="flex flex-col flex-1">
                    {field.label?.length > 0 && !field.hideInnerLabel && (
                        <span className="font-600 pl-1 mb-1">{field.label}</span>
                    )}
                    <ReduxSelectField
                        key={path}
                        name={path}
                        options={field.options}
                        placeholder={field.placeholder}
                        disabled={readOnly || field.disabled}
                        value={field.default}
                    />
                </label>
            );
        case 'multiselect':
            return (
                <ReduxMultiSelectField name={path} options={field.options} disabled={readOnly} />
            );
        case 'multiselect-creatable':
            return (
                <ReduxMultiSelectCreatableField
                    name={path}
                    options={field.options}
                    disabled={readOnly}
                />
            );
        case 'textarea':
            return (
                <ReduxTextAreaField
                    name={path}
                    key={path}
                    disabled={readOnly || field.disabled}
                    placeholder={field.placeholder}
                />
            );
        case 'number':
            return (
                <ReduxNumericInputField
                    key={path}
                    name={path}
                    min={field.min}
                    max={field.max}
                    step={field.step}
                    placeholder={field.placeholder}
                    disabled={readOnly}
                />
            );
        case 'group':
            return field.subComponents.map((subField) => {
                const subFieldName = `${name}.${subField.subpath}`;
                return (
                    <Field
                        key={subFieldName}
                        name={subFieldName}
                        field={subField}
                        readOnly={readOnly}
                    />
                );
            });
        default:
            throw new Error(`Unknown field type: ${field.type}`);
    }
}

Field.propsTypes = {
    field: PropTypes.shape({
        type: PropTypes.string.isRequired,
    }).isRequired,
    name: PropTypes.string,
};

Field.defaultProps = {
    name: '',
};
