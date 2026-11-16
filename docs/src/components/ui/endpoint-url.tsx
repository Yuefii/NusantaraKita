import { ExternalLink } from 'lucide-react';
import React, { useMemo, useState } from 'react';
import { TbCheck, TbCopy } from 'react-icons/tb';

interface EndpointUrlProps {
  method: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH';
  url: string;
  showCopyButton?: boolean;
}

const EndpointUrl: React.FC<EndpointUrlProps> = ({
  method,
  url,
  showCopyButton = true,
}) => {
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(url);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error('Failed to copy URL: ', err);
    }
  };

  const methodBadgeStyle = useMemo(() => {
    switch (method) {
      case 'GET':
        return 'bg-emerald-500/15 text-emerald-400 border-emerald-500/30 ring-1 ring-emerald-500/20';
      case 'POST':
        return 'bg-sky-500/15 text-sky-400 border-sky-500/30 ring-1 ring-sky-500/20';
      case 'PUT':
        return 'bg-amber-500/15 text-amber-400 border-amber-500/30 ring-1 ring-amber-500/20';
      case 'DELETE':
        return 'bg-rose-500/15 text-rose-400 border-rose-500/30 ring-1 ring-rose-500/20';
      case 'PATCH':
        return 'bg-orange-500/15 text-orange-400 border-orange-500/30 ring-1 ring-orange-500/20';
      default:
        return 'bg-zinc-500/15 text-zinc-400 border-zinc-500/30';
    }
  }, [method]);

  // Check if URL is fully formed (no placeholders like {kode}) to offer live test
  const canDirectOpen =
    method === 'GET' && !url.includes('{') && url.startsWith('http');

  // Format url highlighting parameter tags like {kode}
  const renderFormattedUrl = () => {
    const parts = url.split(/(\{.*?\})/g);
    return parts.map((part, i) => {
      if (part.startsWith('{') && part.endsWith('}')) {
        return (
          <span
            key={i}
            className="rounded border border-amber-500/30 bg-amber-500/15 px-1 py-0.5 font-semibold text-amber-400"
          >
            {part}
          </span>
        );
      }
      return <span key={i}>{part}</span>;
    });
  };

  return (
    <div className="group relative flex flex-wrap items-center justify-between gap-3 overflow-hidden rounded-lg border border-zinc-800 bg-zinc-950 p-2.5 shadow-md transition-all hover:border-zinc-700 sm:flex-nowrap sm:px-4 sm:py-3">
      <div className="no-scrollbar flex max-w-full min-w-0 items-center gap-3 overflow-x-auto">
        {/* Method Badge */}
        <span
          className={`shrink-0 rounded-md border px-2.5 py-1 text-xs font-bold tracking-wider uppercase ${methodBadgeStyle}`}
        >
          {method}
        </span>

        {/* URL Text */}
        <div className="no-scrollbar overflow-x-auto font-mono text-xs whitespace-nowrap text-zinc-200 select-all sm:text-sm">
          {renderFormattedUrl()}
        </div>
      </div>

      {/* Action Buttons */}
      <div className="ml-auto flex shrink-0 items-center gap-1.5">
        {canDirectOpen && (
          <a
            href={url}
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex h-8 cursor-pointer items-center gap-1.5 rounded-md border border-zinc-800 bg-zinc-900/80 px-2.5 text-xs font-medium text-zinc-300 transition-all hover:bg-zinc-800 hover:text-white"
            title="Buka endpoint langsung di tab baru"
          >
            <ExternalLink className="h-3.5 w-3.5" />
            <span className="hidden sm:inline">Test API</span>
          </a>
        )}

        {showCopyButton && (
          <button
            onClick={handleCopy}
            className="inline-flex h-8 cursor-pointer items-center gap-1.5 rounded-md border border-zinc-800 bg-zinc-900/80 px-2.5 text-xs font-medium text-zinc-300 transition-all hover:bg-zinc-800 hover:text-white"
            aria-label="Salin URL endpoint"
            title="Salin URL"
          >
            {copied ? (
              <>
                <TbCheck className="h-4 w-4 text-emerald-400" />
                <span className="font-medium text-emerald-400">Tersalin</span>
              </>
            ) : (
              <>
                <TbCopy className="h-3.5 w-3.5" />
                <span className="hidden sm:inline">Salin</span>
              </>
            )}
          </button>
        )}
      </div>
    </div>
  );
};

export default EndpointUrl;
