'use client';

import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from '@/components/ui/sidebar';
import { ExternalLink } from 'lucide-react';
import { FaGithub } from 'react-icons/fa';

export function NavSourceCode() {
  const { open } = useSidebar();

  if (!open) {
    return (
      <SidebarMenu className="group-data-[collapsible=icon]:items-center">
        <SidebarMenuItem className="group-data-[collapsible=icon]:flex group-data-[collapsible=icon]:justify-center">
          <SidebarMenuButton
            size="default"
            asChild
            tooltip="GitHub Repository"
            className="border-sidebar-border bg-background/50 hover:bg-sidebar-accent text-sidebar-foreground mx-auto flex h-9 w-9 items-center justify-center rounded-md border transition-colors"
          >
            <a
              href="https://github.com/Yuefii/nusantarakita"
              target="_blank"
              rel="noopener noreferrer"
              aria-label="GitHub Repository"
            >
              <FaGithub className="h-4 w-4 shrink-0" />
            </a>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
    );
  }

  return (
    <a
      href="https://github.com/Yuefii/nusantarakita"
      target="_blank"
      rel="noopener noreferrer"
      className="group border-sidebar-border bg-background/50 hover:bg-sidebar-accent hover:text-sidebar-accent-foreground text-sidebar-foreground flex w-full items-center justify-between gap-2.5 rounded-md border px-3 py-2 text-xs font-medium shadow-2xs transition-all"
    >
      <div className="flex items-center gap-2.5">
        <FaGithub className="h-4 w-4 shrink-0" />
        <span>GitHub Repository</span>
      </div>
      <ExternalLink className="text-muted-foreground group-hover:text-sidebar-accent-foreground h-3.5 w-3.5 shrink-0 transition-colors" />
    </a>
  );
}
