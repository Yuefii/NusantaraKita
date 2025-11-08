import React from 'react';

export interface TableProps {
  columnHeaders: string[];
  rowData: Array<Record<string, React.ReactNode>>;
}

const Table: React.FC<TableProps> = ({ columnHeaders, rowData }) => {
  const headerStyle = 'bg-[#4A4A4A] text-white text-left';
  const thStyle = 'py-3 px-4 font-semibold';
  const tdStyle =
    'py-3 px-4 border-b border-gray-200 text-gray-700 max-w-[250px] truncate';

  return (
    <div className="overflow-x-auto rounded-lg shadow-md">
      <table className="min-w-full bg-white">
        <thead className={headerStyle}>
          <tr>
            {columnHeaders.map((header, index) => (
              <th
                key={index}
                className={`${thStyle} max-w-[250px] truncate`}
              >
                {header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rowData.map((row, rowIndex) => (
            <tr
              key={rowIndex}
              className="hover:bg-gray-50"
            >
              {columnHeaders.map((headerKey, cellIndex) => (
                <td
                  key={cellIndex}
                  className={tdStyle}
                  title={String(row[headerKey] ?? '')}
                >
                  {row[headerKey]}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;
