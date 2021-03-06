import React, { ReactElement } from 'react';
import { PlusCircle } from 'react-feather';

export type AddTacticButtonProps = {
    onClick: () => void;
};

function AddTacticButton({ onClick }: AddTacticButtonProps): ReactElement {
    return (
        <button
            type="button"
            className="flex flex-1 justify-center p-3 w-full border-dashed border border-base-500 hover:bg-primary-100"
            onClick={onClick}
        >
            <PlusCircle className="h-4 w-4 text-base-500 mr-4" />
            Add tactic
        </button>
    );
}
export default AddTacticButton;
