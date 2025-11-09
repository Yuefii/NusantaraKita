import type { IconType } from 'react-icons/lib';

export type CardItemData = {
  icon: IconType;
  title?: string;
  value: string;
};

export interface CardItemProps {
  items: CardItemData[];
}

export const CardItem: React.FC<CardItemProps> = ({ items }) => {
  return (
    <div className="relative flex flex-col gap-2 overflow-hidden rounded-xl border border-gray-300 bg-white p-4 pl-6 shadow-sm">
      <div className="bg-primary absolute top-0 left-0 h-full w-1.5" />
      {items.map((item, index) => (
        <div
          key={index}
          className="flex items-center gap-2 text-[#4A4A4A]"
        >
          <div className="flex h-6 w-6 items-center justify-center">
            <item.icon />
          </div>
          <div
            className="flex gap-1 overflow-hidden font-medium"
            title={item.value}
          >
            <span className="whitespace-nowrap">
              {item.title ? `${item.title}:` : ''}
            </span>
            <span className="truncate italic">{item.value}</span>
          </div>
        </div>
      ))}
    </div>
  );
};
