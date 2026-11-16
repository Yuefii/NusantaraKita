import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from '@/components/ui/sidebar';
import { cn } from '@/lib/utils';
import { Link } from 'react-router';

export function SidebarLogo() {
  const { open } = useSidebar();

  return (
    <SidebarMenu className="group-data-[collapsible=icon]:items-center">
      <SidebarMenuItem className="group-data-[collapsible=icon]:flex group-data-[collapsible=icon]:justify-center">
        <SidebarMenuButton
          size="lg"
          asChild
          tooltip="Nusantara Kita"
          className={cn(
            'hover:bg-sidebar-accent/70 group h-auto rounded-md transition-colors',
            open
              ? 'p-1.5'
              : 'flex items-center justify-center group-data-[collapsible=icon]:size-9! group-data-[collapsible=icon]:p-0!',
          )}
        >
          <Link
            to="/"
            className={cn(
              'flex items-center transition-all',
              open ? 'w-full gap-3' : 'size-full justify-center',
            )}
          >
            {/* Plain logo without borders, centered when collapsed */}
            <div
              className={cn(
                'flex aspect-square shrink-0 items-center justify-center transition-all',
                open ? 'size-11' : 'size-7',
              )}
            >
              <img
                src="/logo.png"
                alt="Nusantara Kita Logo"
                className="h-full w-full object-contain"
              />
            </div>

            {open && (
              <div className="flex min-w-0 flex-1 flex-col leading-tight">
                <div className="flex items-center justify-between gap-1">
                  <span className="text-sidebar-foreground truncate text-sm font-bold tracking-tight">
                    Nusantara Kita
                  </span>
                  <span className="py-0.2 shrink-0 rounded border border-emerald-500/25 bg-emerald-500/10 px-1.5 font-mono text-[10px] font-semibold text-emerald-600 dark:text-emerald-400">
                    v2.2
                  </span>
                </div>
                <span className="text-muted-foreground mt-0.5 flex items-center gap-1.5 truncate text-[11px] font-medium">
                  <span className="relative flex h-1.5 w-1.5">
                    <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75"></span>
                    <span className="relative inline-flex h-1.5 w-1.5 rounded-full bg-emerald-500"></span>
                  </span>
                  <span>Data Wilayah & Spasial</span>
                </span>
              </div>
            )}
          </Link>
        </SidebarMenuButton>
      </SidebarMenuItem>
    </SidebarMenu>
  );
}
