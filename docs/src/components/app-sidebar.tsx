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
import { GrMapLocation } from 'react-icons/gr';
import { MdOutlineTouchApp } from 'react-icons/md';
import { SlBookOpen } from 'react-icons/sl';

const navMenus: NavMenuGroup[] = [
  {
    title: 'Getting Started',
    url: '#',
    icon: MdOutlineTouchApp,
    isActive: true,
    items: [
      { title: 'Overview', url: '/' },
      { title: 'Tentang API', url: '/tentang-api' },
      { title: 'Panduan Instalasi', url: '/paduan-instalasi' },
    ],
  },
  {
    title: 'API Documentation',
    url: '#',
    icon: SlBookOpen,
    isActive: true,
    items: [
      { title: 'GET Provinsi', url: '/api-documentation/get-provinsi' },
      {
        title: 'GET Kabupaten/Kota',
        url: '/api-documentation/get-kabupaten-kota',
      },
      { title: 'GET Kecamatan', url: '/api-documentation/get-kecamatan' },
      {
        title: 'GET Desa/Kelurahan',
        url: '/api-documentation/get-desa-kelurahan',
      },
      {
        title: 'Search / Autocomplete',
        url: '/api-documentation/search',
      },
      {
        title: 'Reverse Geocoding',
        url: '/api-documentation/reverse-geocoding',
      },
    ],
  },
  {
    title: 'Data Wilayah',
    url: '#',
    icon: GrMapLocation,
    isActive: true,
    items: [
      { title: 'Provinsi', url: '/data-wilayah/provinsi' },
      { title: 'Kabupaten/Kota', url: '/data-wilayah/kabupaten-kota' },
      { title: 'Kecamatan', url: '/data-wilayah/kecamatan' },
      { title: 'Desa/Kelurahan', url: '/data-wilayah/desa-kelurahan' },
    ],
  },
];

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar
      collapsible="icon"
      {...props}
    >
      <SidebarHeader>
        <SidebarLogo />
      </SidebarHeader>
      <SidebarContent>
        <NavMenu items={navMenus} />
      </SidebarContent>
      <SidebarFooter>
        <NavSourceCode />
      </SidebarFooter>
    </Sidebar>
  );
}
