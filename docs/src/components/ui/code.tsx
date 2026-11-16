import React, { useEffect, useMemo, useRef, useState } from 'react';
import { FiChevronDown, FiChevronUp } from 'react-icons/fi';
import { TbCheck, TbCopy } from 'react-icons/tb';

interface CodeProps {
  content: string;
  showCopyButton?: boolean;
  language?: string;
}

/**
 * Intelligent lightweight JSON syntax highlighter without heavy dependencies.
 */
function highlightJson(json: string) {
  // Regex to match JSON tokens:
  // 1: string (key or value)
  // 2: numbers
  // 3: booleans / null
  // 4: punctuation
  const tokenRegex =
    /("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+-]?\d+)?|[{}[],:])/g;

  const parts: React.ReactNode[] = [];
  let lastIndex = 0;
  let match: RegExpExecArray | null;

  while ((match = tokenRegex.exec(json)) !== null) {
    const textBefore = json.substring(lastIndex, match.index);
    if (textBefore) parts.push(textBefore);

    const token = match[0];
    if (token.endsWith(':')) {
      // Key
      const keyStr = token.slice(0, -1);
      parts.push(
        <span
          key={match.index}
          className="font-medium text-sky-300 dark:text-sky-400"
        >
          {keyStr}
        </span>,
      );
      parts.push(
        <span
          key={`${match.index}-colon`}
          className="text-zinc-500"
        >
          :
        </span>,
      );
    } else if (token.startsWith('"')) {
      // String value
      parts.push(
        <span
          key={match.index}
          className="text-emerald-300 dark:text-emerald-400"
        >
          {token}
        </span>,
      );
    } else if (/^(true|false)$/.test(token)) {
      // Boolean
      parts.push(
        <span
          key={match.index}
          className="font-medium text-purple-400 dark:text-purple-300"
        >
          {token}
        </span>,
      );
    } else if (token === 'null') {
      // Null
      parts.push(
        <span
          key={match.index}
          className="font-medium text-rose-400 dark:text-rose-300"
        >
          {token}
        </span>,
      );
    } else if (/^-?\d/.test(token)) {
      // Number
      parts.push(
        <span
          key={match.index}
          className="text-amber-300 dark:text-amber-400"
        >
          {token}
        </span>,
      );
    } else {
      // Brackets, commas
      parts.push(
        <span
          key={match.index}
          className="text-zinc-500"
        >
          {token}
        </span>,
      );
    }
    lastIndex = tokenRegex.lastIndex;
  }

  const trailingText = json.substring(lastIndex);
  if (trailingText) parts.push(trailingText);

  return parts;
}

const Code: React.FC<CodeProps> = ({
  content,
  showCopyButton = false,
  language,
}) => {
  const [copied, setCopied] = useState(false);
  const [isExpanded, setIsExpanded] = useState(false);
  const [showExpandButton, setShowExpandButton] = useState(false);
  const preRef = useRef<HTMLPreElement>(null);

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(content);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error('Failed to copy content: ', err);
    }
  };

  const detectedLanguage = useMemo(() => {
    if (language) return language.toUpperCase();
    const trimmed = content.trim();
    if (trimmed.startsWith('{') || trimmed.startsWith('[')) return 'JSON';
    if (
      trimmed.startsWith('git') ||
      trimmed.startsWith('npm') ||
      trimmed.startsWith('cd') ||
      trimmed.startsWith('go ') ||
      trimmed.startsWith('cp ')
    ) {
      return 'BASH';
    }
    return 'CODE';
  }, [content, language]);

  const renderedContent = useMemo(() => {
    if (detectedLanguage === 'JSON') {
      try {
        return highlightJson(content);
      } catch {
        return content;
      }
    }
    return content;
  }, [content, detectedLanguage]);

  useEffect(() => {
    const preEl = preRef.current;
    if (preEl && preEl.scrollHeight > 360) {
      setShowExpandButton(true);
    }
  }, [content]);

  return (
    <div className="relative overflow-hidden rounded-lg border border-zinc-800 bg-zinc-950 text-zinc-100 shadow-lg">
      {/* Code Header Bar */}
      <div className="flex items-center justify-between border-b border-zinc-800/80 bg-zinc-900/60 px-4 py-2 text-xs">
        <div className="flex items-center gap-2">
          {/* Terminal Dots */}
          <div className="flex items-center gap-1.5 opacity-60">
            <span className="h-2.5 w-2.5 rounded-full bg-rose-500/80" />
            <span className="h-2.5 w-2.5 rounded-full bg-amber-500/80" />
            <span className="h-2.5 w-2.5 rounded-full bg-emerald-500/80" />
          </div>
          <span className="ml-2 font-mono text-[11px] font-semibold tracking-wider text-zinc-400">
            {detectedLanguage}
          </span>
        </div>

        {showCopyButton && (
          <button
            onClick={handleCopy}
            className="flex cursor-pointer items-center gap-1.5 rounded-md px-2 py-1 text-xs text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-zinc-200"
            aria-label="Copy code to clipboard"
            title="Salin kode"
          >
            {copied ? (
              <>
                <TbCheck className="h-3.5 w-3.5 text-emerald-400" />
                <span className="font-medium text-emerald-400">Tersalin</span>
              </>
            ) : (
              <>
                <TbCopy className="h-3.5 w-3.5" />
                <span>Salin</span>
              </>
            )}
          </button>
        )}
      </div>

      {/* Code Body */}
      <div className="relative p-4">
        <pre
          ref={preRef}
          className={`no-scrollbar overflow-x-auto font-mono text-xs leading-relaxed whitespace-pre transition-all duration-300 sm:text-sm ${
            isExpanded ? 'max-h-none' : 'max-h-[360px] overflow-y-auto'
          }`}
        >
          <code>{renderedContent}</code>
        </pre>

        {showExpandButton && !isExpanded && (
          <div className="pointer-events-none absolute inset-x-0 bottom-0 flex h-20 items-end justify-center bg-gradient-to-t from-zinc-950 via-zinc-950/80 to-transparent pb-3" />
        )}
      </div>

      {/* Expand/Collapse Control */}
      {showExpandButton && (
        <div className="flex justify-center border-t border-zinc-800/80 bg-zinc-900/40 p-2">
          <button
            onClick={() => setIsExpanded((prev) => !prev)}
            className="flex cursor-pointer items-center gap-1.5 rounded-md px-3 py-1 text-xs text-zinc-400 transition-colors hover:bg-zinc-800 hover:text-zinc-200"
          >
            {isExpanded ? (
              <>
                <FiChevronUp className="h-3.5 w-3.5" />
                <span>Ciutkan</span>
              </>
            ) : (
              <>
                <FiChevronDown className="h-3.5 w-3.5" />
                <span>Tampilkan Selengkapnya</span>
              </>
            )}
          </button>
        </div>
      )}
    </div>
  );
};

export default Code;
