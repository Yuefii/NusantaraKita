import { formatBreadcrumb } from '@/utils/formate-breadcrumb';
import { ChevronRight } from 'lucide-react';
import React from 'react';
import { Link, useLocation } from 'react-router';

export function InPageBreadcrumbs() {
  const location = useLocation();
  const paths = location.pathname.split('/').filter(Boolean);

  if (paths.length === 0) {
    return null;
  }

  return (
    <nav
      aria-label="Breadcrumb"
      className="text-muted-foreground mb-4 flex flex-wrap items-center gap-1.5 text-xs font-medium"
    >
      <Link
        to="/"
        className="hover:text-foreground transition-colors"
      >
        Docs
      </Link>
      {paths.map((path, index) => {
        const isLast = index === paths.length - 1;
        const formatted = formatBreadcrumb(path);
        const url = '/' + paths.slice(0, index + 1).join('/');

        return (
          <React.Fragment key={path}>
            <ChevronRight className="text-muted-foreground/60 h-3 w-3" />
            {isLast ? (
              <span className="text-foreground font-semibold">{formatted}</span>
            ) : (
              <Link
                to={url}
                className="hover:text-foreground transition-colors"
              >
                {formatted}
              </Link>
            )}
          </React.Fragment>
        );
      })}
    </nav>
  );
}
