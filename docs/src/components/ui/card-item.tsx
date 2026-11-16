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
    <div className="group border-border/70 bg-card hover:border-primary/40 relative flex flex-col gap-2.5 overflow-hidden rounded-lg border p-4 pl-5 shadow-xs transition-all duration-200 hover:shadow-md">
      <div className="from-primary via-primary/70 to-primary/30 absolute top-0 left-0 h-full w-1 bg-gradient-to-b" />
      {items.map((item, index) => (
        <div
          key={index}
          className="text-foreground/80 flex items-center gap-2.5 text-sm"
        >
          <div className="text-muted-foreground group-hover:text-primary flex h-6 w-6 shrink-0 items-center justify-center transition-colors">
            <item.icon />
          </div>
          <div
            className="flex gap-1.5 overflow-hidden text-sm"
            title={item.value}
          >
            {item.title && (
              <span className="text-muted-foreground font-medium whitespace-nowrap">
                {item.title}:
              </span>
            )}
            <span className="text-foreground truncate font-medium">
              {item.value}
            </span>
          </div>
        </div>
      ))}
    </div>
  );
};
