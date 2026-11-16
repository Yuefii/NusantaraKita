'use client';

import { useSidebar } from '@/components/ui/sidebar';
import { cn } from '@/lib/utils';
import { type LucideIcon } from 'lucide-react';
import type { IconType } from 'react-icons/lib';
import { Link, useLocation } from 'react-router';

export interface NavMenuItem {
  title: string;
  url: string;
  method?: string;
  badge?: string;
}

export interface NavMenuGroup {
  title: string;
  url: string;
  icon?: LucideIcon | IconType;
  isActive?: boolean;
  badge?: string;
  items?: NavMenuItem[];
}

export interface NavMenuProps {
  items: NavMenuGroup[];
}

export function NavMenu({ items }: NavMenuProps) {
  const { isMobile, setOpenMobile } = useSidebar();
  const location = useLocation();

  const handleSubItemClick = () => {
    if (isMobile) {
      setOpenMobile(false);
    }
  };

  return (
    <div className="flex flex-col space-y-4 px-2 py-2">
      {/* Nav Menu Groups */}
      {items.map((group) => {
        return (
          <div
            key={group.title}
            className="flex flex-col space-y-1"
          >
            {/* Section Header */}
            <div className="px-2 pt-1 pb-1">
              <h4 className="text-muted-foreground/80 text-[11px] font-semibold tracking-wider uppercase">
                {group.title}
              </h4>
            </div>

            {/* Tree Navigation Links */}
            <ul className="border-border/70 ml-2.5 space-y-0.5 border-l pl-2">
              {group.items?.map((subItem) => {
                const isActive = location.pathname === subItem.url;
                const hasMethod =
                  subItem.method ||
                  subItem.url.startsWith('/api-documentation');

                return (
                  <li key={subItem.title}>
                    <Link
                      to={subItem.url}
                      onClick={handleSubItemClick}
                      className={cn(
                        'group flex items-center justify-between gap-2 rounded-md px-2.5 py-1.5 text-xs font-medium transition-all sm:text-[13px]',
                        isActive
                          ? 'border-foreground bg-accent/70 text-foreground -ml-[9px] border-l-2 pl-[7px] font-semibold shadow-2xs'
                          : 'text-muted-foreground hover:bg-muted/40 hover:text-foreground font-normal',
                      )}
                    >
                      <span className="truncate">{subItem.title}</span>

                      {hasMethod && (
                        <span
                          className={cn(
                            'py-0.2 shrink-0 rounded px-1.5 font-mono text-[9px] font-bold tracking-wider transition-colors',
                            isActive
                              ? 'border border-emerald-500/35 bg-emerald-500/20 text-emerald-600 dark:text-emerald-300'
                              : 'border border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',
                          )}
                        >
                          {subItem.method || 'GET'}
                        </span>
                      )}

                      {subItem.badge && !hasMethod && (
                        <span className="border-border/60 bg-muted/60 text-muted-foreground py-0.2 shrink-0 rounded border px-1.5 font-mono text-[9px]">
                          {subItem.badge}
                        </span>
                      )}
                    </Link>
                  </li>
                );
              })}
            </ul>
          </div>
        );
      })}
    </div>
  );
}
