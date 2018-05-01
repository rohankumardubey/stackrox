import React from 'react';
import PropTypes from 'prop-types';
import ReactModal from 'react-modal';

const Modal = props => (
    <ReactModal
        isOpen={props.isOpen}
        contentLabel="Modal"
        ariaHideApp={false}
        overlayClassName="ReactModal__Overlay react-modal-overlay p-4 flex"
        className="ReactModal__Content dialog mx-auto my-0 flex flex-col self-center bg-primary-100 overflow-hidden max-h-full transition p-4"
    >
        <div className="py-4">{props.text}</div>
        <div className="flex flex-row justify-end">
            <button className="btn-base" onClick={props.onCancel}>
                Cancel
            </button>
            <button className="btn-success" onClick={props.onConfirm}>
                Confirm
            </button>
        </div>
    </ReactModal>
);

Modal.propTypes = {
    isOpen: PropTypes.bool.isRequired,
    text: PropTypes.string.isRequired,
    onCancel: PropTypes.func.isRequired,
    onConfirm: PropTypes.func.isRequired
};

export default Modal;
