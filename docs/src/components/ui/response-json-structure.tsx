import React from 'react';

export interface ResponseJsonListItem {
  name: string;
  details?: string;
  children?: ResponseJsonListItem[];
}

interface ResponseJsonStructureProps {
  items: ResponseJsonListItem[];
}

interface RecursiveListItemProps {
  item: ResponseJsonListItem;
}

const RecursiveListItem: React.FC<RecursiveListItemProps> = ({ item }) => {
  return (
    <li className="list-none space-y-1.5">
      <div className="flex flex-wrap items-center gap-2">
        <span className="border-border bg-muted/80 text-foreground inline-block rounded-md border px-2 py-0.5 font-mono text-xs font-semibold">
          {item.name}
        </span>
        {item.details && (
          <span className="text-muted-foreground text-xs sm:text-sm">
            {item.details}
          </span>
        )}
      </div>
      {item.children && item.children.length > 0 && (
        <ul className="border-border/70 mt-2 ml-2 space-y-2 border-l-2 pl-4">
          {item.children.map((child, index) => (
            <RecursiveListItem
              key={`${item.name}-child-${index}`}
              item={child}
            />
          ))}
        </ul>
      )}
    </li>
  );
};

const ResponseJsonStructure: React.FC<ResponseJsonStructureProps> = ({
  items,
}) => {
  if (!items || items.length === 0) return null;

  return (
    <ul className="border-border bg-card/60 space-y-2.5 rounded-lg border p-4">
      {items.map((item, index) => (
        <RecursiveListItem
          key={`root-${index}`}
          item={item}
        />
      ))}
    </ul>
  );
};

export default ResponseJsonStructure;
