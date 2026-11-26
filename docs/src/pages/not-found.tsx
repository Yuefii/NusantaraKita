import { Button } from '@/components/ui/button';
import { FileCode, FileQuestion, Home } from 'lucide-react';
import { Link, useLocation } from 'react-router';

export const NotFound = () => {
  const location = useLocation();

  return (
    <section className="flex min-h-[65vh] w-full items-center justify-center py-8">
      <div className="border-border/80 bg-card/40 relative flex w-full max-w-xl flex-col items-center overflow-hidden rounded-xl border px-6 py-14 text-center shadow-2xs sm:px-10 sm:py-16">
        <div
          aria-hidden="true"
          className="bg-grid-pattern text-foreground pointer-events-none absolute inset-0 opacity-[0.03] dark:opacity-[0.06]"
        />

        <div className="relative z-10 flex flex-col items-center space-y-4">
          <div className="border-primary/20 bg-primary/10 text-primary inline-flex items-center gap-2 rounded-md border px-2.5 py-1 font-mono text-xs font-semibold">
            <FileQuestion className="h-3.5 w-3.5" />
            <span>404 • Not Found</span>
          </div>

          <p className="text-foreground font-mono text-5xl font-bold tracking-tighter sm:text-6xl">
            404
          </p>

          <h1 className="text-foreground text-2xl font-bold tracking-tight sm:text-3xl">
            Halaman Tidak Ditemukan
          </h1>

          <p className="text-muted-foreground max-w-md text-sm leading-relaxed sm:text-base">
            Maaf, halaman <code className="break-all">{location.pathname}</code>{' '}
            yang Anda tuju tidak tersedia atau telah dipindahkan.
          </p>

          <div className="flex flex-wrap items-center justify-center gap-3 pt-3">
            <Button asChild>
              <Link to="/">
                <Home className="h-4 w-4" />
                <span>Kembali ke Halaman Utama</span>
              </Link>
            </Button>

            <Button
              asChild
              variant="outline"
            >
              <Link to="/api-documentation/get-provinsi">
                <FileCode className="h-4 w-4" />
                <span>Lihat Dokumentasi API</span>
              </Link>
            </Button>
          </div>
        </div>
      </div>
    </section>
  );
};

export default NotFound;
