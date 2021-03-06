import React from 'react';
import PropTypes from 'prop-types';
import { Field } from 'redux-form';

const ReduxTextAreaField = ({ name, disabled, placeholder, maxLength }) => (
    <Field
        key={name}
        name={name}
        component="textarea"
        className="bg-base-100 border rounded-l py-3 px-2 border-base-300 text-base-600 w-full font-600 leading-normal"
        disabled={disabled}
        rows={4}
        placeholder={placeholder}
        maxLength={maxLength}
    />
);

ReduxTextAreaField.propTypes = {
    name: PropTypes.string.isRequired,
    disabled: PropTypes.bool,
    placeholder: PropTypes.string.isRequired,
    maxLength: PropTypes.string,
};

ReduxTextAreaField.defaultProps = {
    disabled: false,
    maxLength: null,
};

export default ReduxTextAreaField;
