import React from 'react';

export interface TableProps {
  columnHeaders: string[];
  rowData: Array<Record<string, React.ReactNode>>;
}

const Table: React.FC<TableProps> = ({ columnHeaders, rowData }) => {
  return (
    <div className="border-border bg-card overflow-x-auto rounded-lg border shadow-xs">
      <table className="divide-border min-w-full divide-y text-left">
        <thead className="bg-muted/50 text-muted-foreground text-xs font-semibold tracking-wider uppercase">
          <tr>
            {columnHeaders.map((header, index) => (
              <th
                key={index}
                className="px-4.5 py-3 font-medium"
              >
                {header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="divide-border/60 divide-y text-sm">
          {rowData.map((row, rowIndex) => (
            <tr
              key={rowIndex}
              className="hover:bg-muted/30 transition-colors"
            >
              {columnHeaders.map((headerKey, cellIndex) => {
                const cellValue = row[headerKey];
                return (
                  <td
                    key={cellIndex}
                    className="text-foreground/90 px-4.5 py-3 align-top"
                    title={
                      typeof cellValue === 'string' ? cellValue : undefined
                    }
                  >
                    {cellValue}
                  </td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;
