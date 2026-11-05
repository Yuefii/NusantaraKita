'use client';

import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from '@/components/ui/collapsible';
import {
  SidebarGroup,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
  useSidebar,
} from '@/components/ui/sidebar';
import { ChevronRight, type LucideIcon } from 'lucide-react';
import type { IconType } from 'react-icons/lib';
import { Link, useLocation } from 'react-router';

export interface NavMenuItem {
  title: string;
  url: string;
}

export interface NavMenuGroup {
  title: string;
  url: string;
  icon?: LucideIcon | IconType;
  isActive?: boolean;
  items?: NavMenuItem[];
}

export interface NavMenuProps {
  items: NavMenuGroup[];
}

export function NavMenu({ items }: NavMenuProps) {
  const { setOpen, open, isMobile, setOpenMobile } = useSidebar();
  const location = useLocation();

  const handleSidebarItemClick = () => {
    if (!open) setOpen(true);
  };

  const handleSubItemClick = () => {
    if (isMobile) {
      setOpenMobile(false);
    }
  };

  return (
    <SidebarGroup>
      <SidebarMenu>
        {items.map((item) => (
          <Collapsible
            key={item.title}
            asChild
            defaultOpen={item.isActive}
            className="group/collapsible"
          >
            <SidebarMenuItem>
              <CollapsibleTrigger asChild>
                <SidebarMenuButton
                  tooltip={item.title}
                  onClick={handleSidebarItemClick}
                >
                  {item.icon && <item.icon />}
                  <span>{item.title}</span>
                  <ChevronRight className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
                </SidebarMenuButton>
              </CollapsibleTrigger>
              <CollapsibleContent>
                <SidebarMenuSub>
                  {item.items?.map((subItem) => (
                    <SidebarMenuSubItem key={subItem.title}>
                      <SidebarMenuSubButton
                        asChild
                        isActive={location.pathname === subItem.url}
                      >
                        <Link
                          to={subItem.url}
                          onClick={handleSubItemClick}
                        >
                          <span>{subItem.title}</span>
                        </Link>
                      </SidebarMenuSubButton>
                    </SidebarMenuSubItem>
                  ))}
                </SidebarMenuSub>
              </CollapsibleContent>
            </SidebarMenuItem>
          </Collapsible>
        ))}
      </SidebarMenu>
    </SidebarGroup>
  );
}
