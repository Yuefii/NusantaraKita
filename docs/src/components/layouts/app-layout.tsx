import { AppSidebar } from '@/components/app-sidebar';
import { DocsPagination } from '@/components/docs-pagination';
import { InPageBreadcrumbs } from '@/components/in-page-breadcrumb';
import { SearchDialog } from '@/components/search-dialog';
import { ThemeToggle } from '@/components/theme-toggle';
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from '@/components/ui/sidebar';
import { cn } from '@/lib/utils';
import { Search } from 'lucide-react';
import { useState } from 'react';
import { FaGithub } from 'react-icons/fa';
import { Link, NavLink, Outlet, useLocation } from 'react-router';

/** Application layout with modern full-width header, docs sidebar, in-page breadcrumb, and TOC ala nextjs.org/docs. */
export default function AppLayout() {
  const [searchOpen, setSearchOpen] = useState(false);
  const location = useLocation();

  return (
    <SidebarProvider defaultOpen={true}>
      <div className="bg-background text-foreground flex min-h-screen w-full flex-col">
        {/* Full-width sticky header ala nextjs.org/docs */}
        <header className="border-border/80 bg-background/80 sticky top-0 z-40 w-full border-b backdrop-blur-md">
          <div className="flex h-14 w-full items-center justify-between px-4 sm:h-16 sm:px-6 lg:px-8">
            {/* Left: Mobile Trigger, Brand Logo, Nav Links */}
            <div className="flex items-center gap-4 sm:gap-6">
              <SidebarTrigger className="text-muted-foreground hover:text-foreground -ml-1 size-8 cursor-pointer md:hidden" />

              <Link
                to="/"
                className="flex items-center gap-2.5 transition-opacity hover:opacity-90"
              >
                <img
                  src="/logo.png"
                  alt="Nusantara Kita"
                  className="h-7 w-7 object-contain sm:h-8 sm:w-8"
                />
                <span className="text-foreground text-sm font-bold tracking-tight sm:text-base">
                  Nusantara Kita
                </span>
              </Link>

              {/* Main Nav Links (Next.js docs header style) */}
              <nav className="hidden items-center gap-5 text-sm font-medium md:flex">
                <NavLink
                  to="/"
                  end
                  className={({ isActive }) =>
                    isActive
                      ? 'text-foreground font-semibold'
                      : 'text-muted-foreground hover:text-foreground transition-colors'
                  }
                >
                  Overview
                </NavLink>
                <NavLink
                  to="/api-documentation/get-provinsi"
                  className={({ isActive }) =>
                    isActive ||
                    location.pathname.startsWith('/api-documentation')
                      ? 'text-foreground font-semibold'
                      : 'text-muted-foreground hover:text-foreground transition-colors'
                  }
                >
                  API Reference
                </NavLink>
                <NavLink
                  to="/data-wilayah/provinsi"
                  className={({ isActive }) =>
                    isActive || location.pathname.startsWith('/data-wilayah')
                      ? 'text-foreground font-semibold'
                      : 'text-muted-foreground hover:text-foreground transition-colors'
                  }
                >
                  Data Wilayah
                </NavLink>
                <NavLink
                  to="/paduan-instalasi"
                  className={({ isActive }) =>
                    isActive
                      ? 'text-foreground font-semibold'
                      : 'text-muted-foreground hover:text-foreground transition-colors'
                  }
                >
                  Panduan
                </NavLink>
              </nav>
            </div>

            {/* Right: Search bar, Status, GitHub, Theme Toggle */}
            <div className="flex items-center gap-2 sm:gap-3">
              <button
                onClick={() => setSearchOpen(true)}
                className="border-border/80 bg-muted/40 text-muted-foreground hover:bg-muted hover:text-foreground flex h-9 cursor-pointer items-center gap-2 rounded-md border px-2.5 text-xs shadow-2xs transition-all sm:w-48 sm:px-3 md:w-56"
                aria-label="Cari dokumentasi"
              >
                <Search className="h-3.5 w-3.5 shrink-0" />
                <span className="hidden truncate sm:inline">
                  Cari dokumentasi...
                </span>
              </button>

              {/* Live API Status Badge (Desktop only) */}
              <div
                className="hidden items-center gap-1.5 rounded-md border border-emerald-500/25 bg-emerald-500/10 px-2.5 py-1 text-xs font-medium text-emerald-600 lg:flex dark:text-emerald-400"
                title="API Nusantara Kita Online & Siap Digunakan"
              >
                <span className="relative flex h-2 w-2">
                  <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75"></span>
                  <span className="relative inline-flex h-2 w-2 rounded-full bg-emerald-500"></span>
                </span>
                <span>Active</span>
              </div>

              {/* GitHub Link */}
              <a
                href="https://github.com/Yuefii/nusantarakita"
                target="_blank"
                rel="noopener noreferrer"
                className="border-border/70 bg-background/50 text-muted-foreground hover:bg-muted hover:text-foreground flex h-9 w-9 cursor-pointer items-center justify-center rounded-md border transition-colors"
                title="GitHub Repository"
                aria-label="GitHub Repository"
              >
                <FaGithub className="h-4 w-4" />
              </a>

              {/* Theme Toggle */}
              <ThemeToggle />
            </div>
          </div>
        </header>

        {/* Main Layout Body: Sidebar + Content */}
        <div className="flex w-full flex-1">
          <AppSidebar />

          <SidebarInset className="bg-background text-foreground flex min-w-0 flex-1 flex-col">
            <div className="w-full flex-1 overflow-x-hidden">
              <div
                className={cn(
                  'w-full px-6 py-8 sm:px-10 lg:px-14 xl:px-16',
                  location.pathname === '/'
                    ? 'max-w-6xl'
                    : location.pathname.startsWith('/data-wilayah')
                      ? 'max-w-5xl'
                      : 'max-w-4xl',
                )}
              >
                {location.pathname !== '/' && <InPageBreadcrumbs />}
                <Outlet />
                <div className="pt-4 pb-10">
                  <DocsPagination />
                </div>
              </div>
            </div>
          </SidebarInset>
        </div>

        {/* Global Search Dialog Modal */}
        <SearchDialog
          open={searchOpen}
          onOpenChange={setSearchOpen}
        />
      </div>
    </SidebarProvider>
  );
}
