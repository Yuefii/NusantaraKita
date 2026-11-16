import type { ResponseJsonListItem } from './response-json-structure';
import ResponseJsonStructure from './response-json-structure';

interface ResponseSuccessProps {
  items: ResponseJsonListItem[];
}

const ResponseSuccess: React.FC<ResponseSuccessProps> = ({ items }) => {
  return (
    <div className="space-y-3 py-2">
      <div className="flex items-center justify-between">
        <h3 className="text-foreground text-xl font-semibold tracking-tight">
          Struktur Response (Success 200)
        </h3>
        <span className="rounded-md border border-emerald-500/30 bg-emerald-500/10 px-2 py-0.5 font-mono text-xs font-semibold text-emerald-600 dark:text-emerald-400">
          JSON
        </span>
      </div>
      <ResponseJsonStructure items={items} />
    </div>
  );
};

export default ResponseSuccess;
