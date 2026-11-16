import { Info } from 'lucide-react';
import type React from 'react';

export interface NotesProps {
  items: (string | React.JSX.Element)[];
  title?: string;
}

export const Notes: React.FC<NotesProps> = ({
  items,
  title = 'Catatan Penggunaan',
}) => {
  return (
    <div className="text-foreground my-6 rounded-lg border border-sky-500/25 bg-sky-500/5 p-4 dark:border-sky-500/20 dark:bg-sky-500/10">
      <div className="flex items-start gap-3">
        <Info className="mt-0.5 h-4 w-4 shrink-0 text-sky-500" />
        <div className="flex-1 space-y-1.5">
          <h4 className="text-foreground text-sm font-semibold tracking-tight">
            {title}
          </h4>
          <ul className="text-muted-foreground list-disc space-y-1 pl-4 text-xs leading-relaxed sm:text-sm">
            {items.map((item, index) => (
              <li key={`note-${index}`}>{item}</li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};
