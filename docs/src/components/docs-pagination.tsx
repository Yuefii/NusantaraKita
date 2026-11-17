import { ChevronLeft, ChevronRight } from 'lucide-react';
import { Link, useLocation } from 'react-router';

export const DOCS_PAGES = [
  { title: 'Overview', url: '/' },
  { title: 'Tentang API', url: '/tentang-api' },
  { title: 'Panduan Instalasi', url: '/paduan-instalasi' },
  { title: 'GET Provinsi', url: '/api-documentation/get-provinsi' },
  {
    title: 'GET Kabupaten / Kota',
    url: '/api-documentation/get-kabupaten-kota',
  },
  { title: 'GET Kecamatan', url: '/api-documentation/get-kecamatan' },
  {
    title: 'GET Desa / Kelurahan',
    url: '/api-documentation/get-desa-kelurahan',
  },
  { title: 'Search & Autocomplete', url: '/api-documentation/search' },
  { title: 'Reverse Geocoding', url: '/api-documentation/reverse-geocoding' },
  { title: 'Data Wilayah: Provinsi', url: '/data-wilayah/provinsi' },
  {
    title: 'Data Wilayah: Kabupaten / Kota',
    url: '/data-wilayah/kabupaten-kota',
  },
  { title: 'Data Wilayah: Kecamatan', url: '/data-wilayah/kecamatan' },
  {
    title: 'Data Wilayah: Desa / Kelurahan',
    url: '/data-wilayah/desa-kelurahan',
  },
];

export function DocsPagination() {
  const location = useLocation();
  const currentIndex = DOCS_PAGES.findIndex(
    (page) => page.url === location.pathname,
  );

  if (currentIndex === -1) {
    return null;
  }

  const prev = currentIndex > 0 ? DOCS_PAGES[currentIndex - 1] : null;
  const next =
    currentIndex < DOCS_PAGES.length - 1 ? DOCS_PAGES[currentIndex + 1] : null;

  return (
    <nav
      aria-label="Dokumentasi Navigasi Halaman"
      className="border-border/80 mt-14 flex flex-col gap-3 border-t pt-6 sm:flex-row sm:items-stretch sm:justify-between sm:gap-4"
    >
      {prev ? (
        <Link
          to={prev.url}
          className="group border-border/70 bg-card/60 hover:border-primary/50 hover:bg-muted/40 flex flex-1 flex-col items-start gap-1 rounded-lg border p-4 shadow-2xs transition-all"
        >
          <span className="text-muted-foreground group-hover:text-primary flex items-center gap-1 text-xs font-medium transition-colors">
            <ChevronLeft className="h-3.5 w-3.5 transition-transform group-hover:-translate-x-0.5" />
            <span>Sebelumnya</span>
          </span>
          <span className="text-foreground group-hover:text-primary text-sm font-semibold tracking-tight transition-colors">
            {prev.title}
          </span>
        </Link>
      ) : (
        <div className="hidden flex-1 sm:block" />
      )}

      {next ? (
        <Link
          to={next.url}
          className="group border-border/70 bg-card/60 hover:border-primary/50 hover:bg-muted/40 flex flex-1 flex-col items-end gap-1 rounded-lg border p-4 text-right shadow-2xs transition-all"
        >
          <span className="text-muted-foreground group-hover:text-primary flex items-center gap-1 text-xs font-medium transition-colors">
            <span>Selanjutnya</span>
            <ChevronRight className="h-3.5 w-3.5 transition-transform group-hover:translate-x-0.5" />
          </span>
          <span className="text-foreground group-hover:text-primary text-sm font-semibold tracking-tight transition-colors">
            {next.title}
          </span>
        </Link>
      ) : (
        <div className="hidden flex-1 sm:block" />
      )}
    </nav>
  );
}
