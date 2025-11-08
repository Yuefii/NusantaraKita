import React, { useMemo, useState } from 'react';
import { TbCopy, TbCopyCheck, TbCopyCheckFilled } from 'react-icons/tb';
import { Button } from './button';

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
  const [copyButtonIcon, setCopyButtonIcon] = useState<
    'COPY' | 'COPIED' | 'FAILED'
  >('COPY');

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(url);
      setCopyButtonIcon('COPIED');
      setTimeout(() => setCopyButtonIcon('COPY'), 2000);
    } catch (err) {
      console.error('Failed to copy URL: ', err);
      setCopyButtonIcon('FAILED');
      setTimeout(() => setCopyButtonIcon('COPY'), 2000);
    }
  };

  const renderCopyButtonIcon = useMemo(() => {
    switch (copyButtonIcon) {
      case 'COPY':
        return <TbCopy />;
      case 'COPIED':
        return <TbCopyCheck />;
      case 'FAILED':
        return <TbCopyCheckFilled />;
      default:
        return <TbCopy />;
    }
  }, [copyButtonIcon]);

  const getMethodTextColorClass = (httpMethod: string): string => {
    switch (httpMethod) {
      case 'GET':
        return 'text-[#28a745]';
      case 'POST':
        return 'text-[#007bff]';
      case 'PUT':
        return 'text-[#ffc107]';
      case 'DELETE':
        return 'text-[#dc3545]';
      case 'PATCH':
        return 'text-[#fd7e14]';
      default:
        return 'text-[#6c757d]';
    }
  };

  const methodClasses = `font-bold mr-2.5 ${getMethodTextColorClass(method)}`;

  const componentClasses = `
    bg-[#4A4A4A]
    px-4 py-2.5
    rounded-lg
    flex items-center
    font-mono text-sm
    text-[#E0E0E0]
  `;

  return (
    <div className={componentClasses.trim().replace(/\s+/g, ' ')}>
      <span className={methodClasses}>{method}</span>
      <div className="w-full overflow-x-auto">
        <code className="mr-2 block min-w-fit whitespace-nowrap">{url}</code>
      </div>

      {showCopyButton && (
        <Button
          onClick={handleCopy}
          className="ml-auto flex cursor-pointer items-center justify-center p-1"
          aria-label="Copy endpoint URL"
        >
          {renderCopyButtonIcon}
        </Button>
      )}
    </div>
  );
};

export default EndpointUrl;
