import { Dialog, DialogContent, DialogTitle } from '@/components/ui/dialog';
import { cn } from '@/lib/utils';
import {
  BookOpen,
  Compass,
  FileCode,
  MapPin,
  Search,
  Terminal,
} from 'lucide-react';
import React, { useEffect, useMemo, useState } from 'react';
import { useNavigate } from 'react-router';

interface SearchItem {
  id: string;
  title: string;
  subtitle?: string;
  category: 'Getting Started' | 'API Documentation' | 'Data Wilayah';
  url: string;
  badge?: string;
  icon: React.ReactNode;
}

const SEARCH_ITEMS: SearchItem[] = [
  // Getting Started
  {
    id: 'overview',
    title: 'Overview & Peta Interaktif',
    subtitle: 'Pengenalan dan demo interaktif data wilayah Indonesia',
    category: 'Getting Started',
    url: '/',
    badge: 'Docs',
    icon: <Compass className="h-4 w-4 text-emerald-500" />,
  },
  {
    id: 'tentang-api',
    title: 'Tentang API & Sumber Data',
    subtitle: 'Informasi dataset Kepmendagri dan kontributor',
    category: 'Getting Started',
    url: '/tentang-api',
    badge: 'About',
    icon: <BookOpen className="h-4 w-4 text-blue-500" />,
  },
  {
    id: 'paduan-instalasi',
    title: 'Panduan Instalasi & Self-Host',
    subtitle: 'Cara setup, clone, dan deploy API dengan Go atau Docker',
    category: 'Getting Started',
    url: '/paduan-instalasi',
    badge: 'Guide',
    icon: <Terminal className="h-4 w-4 text-purple-500" />,
  },

  // API Documentation
  {
    id: 'get-provinsi',
    title: 'GET Provinsi',
    subtitle: 'Endpoint daftar provinsi dan detail by kode',
    category: 'API Documentation',
    url: '/api-documentation/get-provinsi',
    badge: 'GET',
    icon: <FileCode className="h-4 w-4 text-emerald-500" />,
  },
  {
    id: 'get-kabupaten-kota',
    title: 'GET Kabupaten / Kota',
    subtitle: 'Endpoint daftar kabupaten/kota per provinsi',
    category: 'API Documentation',
    url: '/api-documentation/get-kabupaten-kota',
    badge: 'GET',
    icon: <FileCode className="h-4 w-4 text-emerald-500" />,
  },
  {
    id: 'get-kecamatan',
    title: 'GET Kecamatan',
    subtitle: 'Endpoint daftar kecamatan per kabupaten/kota',
    category: 'API Documentation',
    url: '/api-documentation/get-kecamatan',
    badge: 'GET',
    icon: <FileCode className="h-4 w-4 text-emerald-500" />,
  },
  {
    id: 'get-desa-kelurahan',
    title: 'GET Desa / Kelurahan',
    subtitle: 'Endpoint daftar desa dan kelurahan beserta kodepos',
    category: 'API Documentation',
    url: '/api-documentation/get-desa-kelurahan',
    badge: 'GET',
    icon: <FileCode className="h-4 w-4 text-emerald-500" />,
  },
  {
    id: 'search-api',
    title: 'Search & Autocomplete Wilayah',
    subtitle: 'Pencarian nama wilayah cepat secara global',
    category: 'API Documentation',
    url: '/api-documentation/search',
    badge: 'GET',
    icon: <Search className="h-4 w-4 text-amber-500" />,
  },
  {
    id: 'reverse-geocoding',
    title: 'Reverse Geocoding',
    subtitle:
      'Cari wilayah administratif berdasarkan koordinat latitude & longitude',
    category: 'API Documentation',
    url: '/api-documentation/reverse-geocoding',
    badge: 'GET',
    icon: <MapPin className="h-4 w-4 text-rose-500" />,
  },

  // Data Wilayah
  {
    id: 'data-provinsi',
    title: 'Data Explorer: Provinsi',
    subtitle: 'Tabel & Card explorer 38 provinsi di Indonesia',
    category: 'Data Wilayah',
    url: '/data-wilayah/provinsi',
    badge: 'Data',
    icon: <MapPin className="h-4 w-4 text-indigo-500" />,
  },
  {
    id: 'data-kab-kota',
    title: 'Data Explorer: Kabupaten / Kota',
    subtitle: 'Tabel & Card explorer 514 kabupaten dan kota',
    category: 'Data Wilayah',
    url: '/data-wilayah/kabupaten-kota',
    badge: 'Data',
    icon: <MapPin className="h-4 w-4 text-indigo-500" />,
  },
  {
    id: 'data-kecamatan',
    title: 'Data Explorer: Kecamatan',
    subtitle: 'Tabel & Card explorer seluruh kecamatan',
    category: 'Data Wilayah',
    url: '/data-wilayah/kecamatan',
    badge: 'Data',
    icon: <MapPin className="h-4 w-4 text-indigo-500" />,
  },
  {
    id: 'data-desa-kelurahan',
    title: 'Data Explorer: Desa / Kelurahan',
    subtitle: 'Tabel & Card explorer seluruh desa/kelurahan',
    category: 'Data Wilayah',
    url: '/data-wilayah/desa-kelurahan',
    badge: 'Data',
    icon: <MapPin className="h-4 w-4 text-indigo-500" />,
  },
];

interface SearchDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function SearchDialog({ open, onOpenChange }: SearchDialogProps) {
  const [query, setQuery] = useState('');
  const [selectedIndex, setSelectedIndex] = useState(0);
  const navigate = useNavigate();

  // Keyboard shortcut Cmd+K / Ctrl+K & custom open event
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
        e.preventDefault();
        onOpenChange(!open);
      }
    };
    const handleOpenEvent = () => {
      onOpenChange(true);
    };
    window.addEventListener('keydown', handleKeyDown);
    window.addEventListener('open-search-dialog', handleOpenEvent);
    return () => {
      window.removeEventListener('keydown', handleKeyDown);
      window.removeEventListener('open-search-dialog', handleOpenEvent);
    };
  }, [open, onOpenChange]);

  const filteredItems = useMemo(() => {
    if (!query.trim()) return SEARCH_ITEMS;
    const lower = query.toLowerCase();
    return SEARCH_ITEMS.filter(
      (item) =>
        item.title.toLowerCase().includes(lower) ||
        item.subtitle?.toLowerCase().includes(lower) ||
        item.category.toLowerCase().includes(lower) ||
        item.url.toLowerCase().includes(lower),
    );
  }, [query]);

  // Reset index when query changes
  useEffect(() => {
    setSelectedIndex(0);
  }, [query]);

  const handleSelect = (item: SearchItem) => {
    navigate(item.url);
    onOpenChange(false);
    setQuery('');
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'ArrowDown') {
      e.preventDefault();
      setSelectedIndex((prev) =>
        prev < filteredItems.length - 1 ? prev + 1 : 0,
      );
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      setSelectedIndex((prev) =>
        prev > 0 ? prev - 1 : filteredItems.length - 1,
      );
    } else if (e.key === 'Enter') {
      e.preventDefault();
      if (filteredItems[selectedIndex]) {
        handleSelect(filteredItems[selectedIndex]);
      }
    }
  };

  return (
    <Dialog
      open={open}
      onOpenChange={onOpenChange}
    >
      <DialogContent className="border-border/80 bg-popover/95 max-w-xl gap-0 overflow-hidden border p-0 shadow-2xl backdrop-blur-xl">
        <DialogTitle className="sr-only">Search Documentation</DialogTitle>
        {/* Search Input Bar */}
        <div className="border-border/60 flex items-center gap-3 border-b px-4 py-3.5">
          <Search className="text-muted-foreground h-5 w-5 shrink-0" />
          <input
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={handleKeyDown}
            placeholder="Cari endpoint, panduan, atau data wilayah..."
            className="text-foreground placeholder:text-muted-foreground/70 flex-1 bg-transparent text-sm focus:outline-hidden"
            autoFocus
          />
        </div>

        {/* Results List */}
        <div className="max-h-[360px] overflow-y-auto p-2">
          {filteredItems.length === 0 ? (
            <div className="text-muted-foreground py-12 text-center text-sm">
              Tidak ada hasil untuk &ldquo;
              <span className="text-foreground font-medium">{query}</span>
              &rdquo;
            </div>
          ) : (
            <div className="space-y-1">
              {filteredItems.map((item, idx) => {
                const isSelected = idx === selectedIndex;
                return (
                  <button
                    key={item.id}
                    onClick={() => handleSelect(item)}
                    onMouseEnter={() => setSelectedIndex(idx)}
                    className={cn(
                      'flex w-full cursor-pointer items-center justify-between rounded-lg px-3 py-2.5 text-left text-sm transition-all',
                      isSelected
                        ? 'bg-accent text-accent-foreground shadow-xs'
                        : 'text-muted-foreground hover:bg-muted/50 hover:text-foreground',
                    )}
                  >
                    <div className="flex min-w-0 items-center gap-3">
                      <div
                        className={cn(
                          'flex h-8 w-8 shrink-0 items-center justify-center rounded-md border',
                          isSelected
                            ? 'border-border/80 bg-background/80'
                            : 'border-border/40 bg-muted/40',
                        )}
                      >
                        {item.icon}
                      </div>
                      <div className="min-w-0 truncate">
                        <div className="text-foreground truncate font-medium">
                          {item.title}
                        </div>
                        {item.subtitle && (
                          <div className="text-muted-foreground truncate text-xs">
                            {item.subtitle}
                          </div>
                        )}
                      </div>
                    </div>
                    {item.badge && (
                      <span
                        className={cn(
                          'ml-2 shrink-0 rounded px-1.5 py-0.5 text-[10px] font-semibold tracking-wider uppercase',
                          item.badge === 'GET'
                            ? 'border border-emerald-500/20 bg-emerald-500/15 text-emerald-600 dark:text-emerald-400'
                            : 'bg-muted text-muted-foreground border-border/60 border',
                        )}
                      >
                        {item.badge}
                      </span>
                    )}
                  </button>
                );
              })}
            </div>
          )}
        </div>

        {/* Footer */}
        <div className="border-border/60 text-muted-foreground bg-muted/20 flex items-center justify-between border-t px-4 py-2.5 text-xs">
          <span>Pilih hasil untuk membuka dokumentasi</span>
          <span>Nusantara Kita</span>
        </div>
      </DialogContent>
    </Dialog>
  );
}
