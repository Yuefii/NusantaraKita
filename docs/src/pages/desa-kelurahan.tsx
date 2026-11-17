import DesaKelDescription from '@/components/data-wilayah/desa-kelurahan/desa-kel-description';
import { DesaKelDisplayData } from '@/components/data-wilayah/desa-kelurahan/desa-kel-display-data';
import { DesaKelPagination } from '@/components/data-wilayah/desa-kelurahan/desa-kel-pagination';
import LimitPerPage from '@/components/ui/limit-per-page';
import ResponseTypeSwitch from '@/components/ui/response-type';
import useGetDesaKelurahan from '@/hooks/get-desa-kelurahan/use-get-desa-kelurahan';
import useParams from '@/hooks/use-params';
import { Home } from 'lucide-react';

export const DesaKelurahan = () => {
  const { params } = useParams();
  const { data, isError, isLoading, isSuccess, isPending } =
    useGetDesaKelurahan(params);

  return (
    <article className="text-foreground w-full space-y-8 transition-all">
      {/* Header */}
      <div className="space-y-3">
        <div className="inline-flex items-center gap-2 rounded-md border border-indigo-500/20 bg-indigo-500/10 px-2.5 py-1 text-xs font-semibold text-indigo-600 dark:text-indigo-400">
          <Home className="h-3.5 w-3.5" />
          <span>Data Explorer</span>
        </div>
        <h1 className="text-foreground text-3xl font-bold tracking-tight">
          Data Wilayah - Desa / Kelurahan
        </h1>
        <DesaKelDescription />
      </div>

      {/* Controls Bar */}
      <div className="border-border/80 bg-card flex flex-wrap items-center justify-between gap-4 rounded-lg border p-4 shadow-xs">
        <ResponseTypeSwitch />
        <LimitPerPage />
      </div>

      {/* Data & Pagination */}
      <div className="flex w-full flex-col justify-between gap-6">
        <DesaKelDisplayData
          data={data}
          isError={isError}
          isLoading={isLoading}
          isSuccess={isSuccess}
          isPending={isPending}
        />
        <div className="w-fit self-center">
          <DesaKelPagination
            data={data}
            isLoading={isLoading}
            isPending={isPending}
          />
        </div>
      </div>
    </article>
  );
};
