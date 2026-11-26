import AppLayout from '@/components/layouts/app-layout';
import AboutApi from '@/pages/about-api';
import { DesaKelurahan } from '@/pages/desa-kelurahan';
import { GetDesaKelurahan } from '@/pages/get-desa-kelurahan';
import { GetKabupatenKota } from '@/pages/get-kabupaten-kota';
import { GetKecamatan } from '@/pages/get-kecamatan';
import { GetProvinsi } from '@/pages/get-provinsi';
import { InstalationGuide } from '@/pages/instalation-guide';
import { KabupatenKota } from '@/pages/kabupaten-kota';
import { Kecamatan } from '@/pages/kecamatan';
import { NotFound } from '@/pages/not-found';
import Overview from '@/pages/overview';
import { Provinsi } from '@/pages/provinsi';
import { ReverseGeocodingApiDoc } from '@/pages/reverse-geocoding';
import { SearchApiDoc } from '@/pages/search';
import { createBrowserRouter } from 'react-router';

export const routers = createBrowserRouter([
  {
    path: '/',
    Component: AppLayout,
    // loader: loadRootData,
    children: [
      { path: '/', Component: Overview },
      { path: '/tentang-api', Component: AboutApi },
      { path: '/paduan-instalasi', Component: InstalationGuide },
      {
        path: '/api-documentation',
        children: [
          { path: 'get-provinsi', Component: GetProvinsi },
          { path: 'get-kabupaten-kota', Component: GetKabupatenKota },
          { path: 'get-kecamatan', Component: GetKecamatan },
          { path: 'get-desa-kelurahan', Component: GetDesaKelurahan },
          { path: 'search', Component: SearchApiDoc },
          { path: 'reverse-geocoding', Component: ReverseGeocodingApiDoc },
        ],
      },
      {
        path: '/data-wilayah',
        children: [
          { path: 'provinsi', Component: Provinsi },
          { path: 'kabupaten-kota', Component: KabupatenKota },
          { path: 'kecamatan', Component: Kecamatan },
          { path: 'desa-kelurahan', Component: DesaKelurahan },
        ],
      },
      { path: '*', Component: NotFound },
    ],
  },
]);
