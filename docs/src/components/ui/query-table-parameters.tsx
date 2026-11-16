import React from 'react';
import type { TableProps } from './table';
import Table from './table';

interface TableWithTitleProps extends TableProps {
  titleText?: string;
  description?: string;
}

const TableWithTitle: React.FC<TableWithTitleProps> = ({
  columnHeaders,
  rowData,
  titleText = 'Query Parameters',
  description,
}) => {
  return (
    <section className="py-2">
      <div className="mb-4">
        <h3 className="text-foreground text-xl font-semibold tracking-tight">
          {titleText}
        </h3>
        {description && (
          <p className="text-muted-foreground mt-1 text-sm">{description}</p>
        )}
      </div>
      <Table
        columnHeaders={columnHeaders}
        rowData={rowData}
      />
    </section>
  );
};

export default TableWithTitle;
