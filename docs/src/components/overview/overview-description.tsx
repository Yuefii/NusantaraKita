import React from 'react';

interface OverviewDescriptionProps {
  title: string;
  description: string;
}

const OverviewDescription: React.FC<OverviewDescriptionProps> = ({
  description,
  title,
}) => {
  return (
    <div className="mb-4">
      <h2 className="text-foreground text-xl font-bold tracking-tight">
        {title}
      </h2>
      <p className="text-muted-foreground mt-1 text-sm leading-relaxed">
        {description}
      </p>
    </div>
  );
};

export default OverviewDescription;
