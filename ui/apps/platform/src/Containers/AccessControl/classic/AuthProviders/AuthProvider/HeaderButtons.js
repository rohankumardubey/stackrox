import React from 'react';
import PropTypes from 'prop-types';

import SaveButton from 'Components/SaveButton';

function HeaderButtons({ editText, isEditing, onEdit, onCancel, onTest, editDisabled }) {
    if (!isEditing) {
        return (
            <div className="flex mr-3">
                {!!onTest && (
                    <button
                        className="mr-2 border-2 bg-primary-200 border-primary-400 text-sm text-primary-700 hover:bg-primary-300 hover:border-primary-500 rounded-sm block px-3 py-2 uppercase"
                        type="button"
                        onClick={onTest}
                        disabled={editDisabled}
                    >
                        Test Login
                    </button>
                )}
                <button
                    className="border-2 bg-primary-200 border-primary-400 text-sm text-primary-700 hover:bg-primary-300 hover:border-primary-500 rounded-sm block px-3 py-2 uppercase"
                    type="button"
                    onClick={onEdit}
                    disabled={editDisabled}
                >
                    {editText}
                </button>
            </div>
        );
    }
    return (
        <div className="flex mr-3">
            <button className="btn btn-base mr-2" type="button" onClick={onCancel}>
                Cancel
            </button>
            <SaveButton formName="auth-provider-form" />
        </div>
    );
}

HeaderButtons.propTypes = {
    editText: PropTypes.string.isRequired,
    isEditing: PropTypes.bool.isRequired,
    onEdit: PropTypes.func.isRequired,
    onCancel: PropTypes.func.isRequired,
    onTest: PropTypes.func,
    editDisabled: PropTypes.bool.isRequired,
};

HeaderButtons.defaultProps = {
    onTest: null,
};

export default HeaderButtons;