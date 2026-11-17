'use client';

import { NavMenu, type NavMenuGroup } from '@/components/nav-menu';
import { NavSourceCode } from '@/components/nav-source-code';
import { SidebarLogo } from '@/components/sidebar-logo';
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
} from '@/components/ui/sidebar';
import * as React from 'react';

const navMenus: NavMenuGroup[] = [
  {
    title: 'Getting Started',
    url: '#',
    isActive: true,
    badge: '3',
    items: [
      { title: 'Overview', url: '/' },
      { title: 'Tentang API', url: '/tentang-api' },
      { title: 'Panduan Instalasi', url: '/paduan-instalasi' },
    ],
  },
  {
    title: 'API Documentation',
    url: '#',
    isActive: true,
    badge: '6',
    items: [
      {
        title: 'Provinsi',
        url: '/api-documentation/get-provinsi',
        method: 'GET',
      },
      {
        title: 'Kabupaten / Kota',
        url: '/api-documentation/get-kabupaten-kota',
        method: 'GET',
      },
      {
        title: 'Kecamatan',
        url: '/api-documentation/get-kecamatan',
        method: 'GET',
      },
      {
        title: 'Desa / Kelurahan',
        url: '/api-documentation/get-desa-kelurahan',
        method: 'GET',
      },
      {
        title: 'Search & Autocomplete',
        url: '/api-documentation/search',
        method: 'GET',
      },
      {
        title: 'Reverse Geocoding',
        url: '/api-documentation/reverse-geocoding',
        method: 'GET',
      },
    ],
  },
  {
    title: 'Data Wilayah',
    url: '#',
    isActive: true,
    badge: '4',
    items: [
      { title: 'Provinsi', url: '/data-wilayah/provinsi' },
      { title: 'Kabupaten / Kota', url: '/data-wilayah/kabupaten-kota' },
      { title: 'Kecamatan', url: '/data-wilayah/kecamatan' },
      { title: 'Desa / Kelurahan', url: '/data-wilayah/desa-kelurahan' },
    ],
  },
];

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar
      collapsible="icon"
      className="border-border bg-background border-r"
      {...props}
    >
      <SidebarHeader className="border-border border-b p-3 md:hidden">
        <SidebarLogo />
      </SidebarHeader>
      <SidebarContent className="no-scrollbar bg-background px-1 py-2">
        <NavMenu items={navMenus} />
      </SidebarContent>
      <SidebarFooter className="border-border bg-background border-t p-2.5 group-data-[collapsible=icon]:p-1.5">
        <NavSourceCode />
      </SidebarFooter>
    </Sidebar>
  );
}
