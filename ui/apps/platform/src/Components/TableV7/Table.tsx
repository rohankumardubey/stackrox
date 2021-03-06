import React, { ReactElement, ReactNode } from 'react';

export type TableProps = {
    children: ReactNode;
};

export function Table({ children }: TableProps): ReactElement {
    return <table className="w-full relative border-separate border-spacing-0">{children}</table>;
}
