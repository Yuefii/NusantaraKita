import { KabupatenKodeDescription } from '@/components/data-wilayah/kabupaten-kota/kabupaten-kota-description';
import { KabupatenKotaDisplayData } from '@/components/data-wilayah/kabupaten-kota/kabupaten-kota-display-data';
import { KabupatenKotaPagination } from '@/components/data-wilayah/kabupaten-kota/kabupaten-kota-pagination';
import LimitPerPage from '@/components/ui/limit-per-page';
import ResponseTypeSwitch from '@/components/ui/response-type';
import useGetKabKota from '@/hooks/get-kabupaten-kota/use-get-kab-kota';
import { useIsMobile } from '@/hooks/use-mobile';
import useParams from '@/hooks/use-params';
import { useIsTablet } from '@/hooks/use-tablet';
import { cn } from '@/lib/utils';
import { Building2 } from 'lucide-react';

export const KabupatenKota = () => {
  const { params } = useParams();
  const { data, isError, isLoading, isSuccess, isPending } =
    useGetKabKota(params);

  const isMobile = useIsMobile();
  const isTablet = useIsTablet();
  const isSmallScreen = isMobile || isTablet;

  return (
    <section
      className={cn(
        'text-foreground mx-auto max-w-5xl space-y-8 transition-all',
        {
          'p-4 sm:p-6': isSmallScreen,
          'p-8 lg:p-10': !isSmallScreen,
        },
      )}
    >
      {/* Header */}
      <div className="space-y-3">
        <div className="inline-flex items-center gap-2 rounded-md border border-indigo-500/20 bg-indigo-500/10 px-2.5 py-1 text-xs font-semibold text-indigo-600 dark:text-indigo-400">
          <Building2 className="h-3.5 w-3.5" />
          <span>Data Explorer</span>
        </div>
        <h1 className="text-foreground text-3xl font-bold tracking-tight">
          Data Wilayah - Kabupaten / Kota
        </h1>
        <KabupatenKodeDescription />
      </div>

      {/* Controls Bar */}
      <div className="border-border/80 bg-card flex flex-wrap items-center justify-between gap-4 rounded-lg border p-4 shadow-xs">
        <ResponseTypeSwitch />
        <LimitPerPage />
      </div>

      {/* Data & Pagination */}
      <div className="flex w-full flex-col justify-between gap-6">
        <KabupatenKotaDisplayData
          data={data}
          isError={isError}
          isLoading={isLoading}
          isSuccess={isSuccess}
          isPending={isPending}
        />
        <div className="w-fit self-center">
          <KabupatenKotaPagination
            data={data}
            isLoading={isLoading}
            isPending={isPending}
          />
        </div>
      </div>
    </section>
  );
};
