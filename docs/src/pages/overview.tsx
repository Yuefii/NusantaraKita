import OverviewDescription from '@/components/overview/overview-description';
import OverviewMap from '@/components/overview/overview-map';
import OverViewSelects from '@/components/overview/overview-selects';
import EndpointUrl from '@/components/ui/endpoint-url';
import { useOverview } from '@/context/overview-provider';
import { useIsMobile } from '@/hooks/use-mobile';
import { useIsTablet } from '@/hooks/use-tablet';
import { cn } from '@/lib/utils';
import 'leaflet/dist/leaflet.css';

const Overview = () => {
  const isMobile = useIsMobile();
  const isTablet = useIsTablet();
  const { state } = useOverview();
  const isSmallScreen = isMobile || isTablet;

  const { province, kabKota, kecamatan, desaKel } = state.selected;

  const endpoints = [
    {
      label: 'Provinsi',
      condition: true,
      url: 'https://api-nusakita.vercel.app/v2/provinsi?pagination=true&limit=15',
    },
    {
      label: 'Kab/Kota',
      condition: province && kabKota,
      url: province
        ? `https://api-nusakita.vercel.app/v2/${province.kode}/kab-kota?pagination=true&limit=15`
        : '',
    },
    {
      label: 'Kecamatan',
      condition: province && kabKota && kecamatan,
      url: kabKota
        ? `https://api-nusakita.vercel.app/v2/${kabKota.kode}/kecamatan?pagination=true&limit=15`
        : '',
    },
    {
      label: 'Desa/Kelurahan',
      condition: province && kabKota && kecamatan && desaKel,
      url: kecamatan
        ? `https://api-nusakita.vercel.app/v2/${kecamatan.kode}/desa-kel?pagination=true&limit=15`
        : '',
    },
  ];

  return (
    <section
      className={cn('max-w-5xl bg-white text-gray-800', {
        'p-5': isSmallScreen,
        'p-8': !isSmallScreen,
      })}
    >
      <h1 className="mb-10 text-3xl font-bold text-gray-700">Overview</h1>

      {/* Map Section */}
      <div className="mb-10">
        <OverviewDescription
          title="Map View"
          description="Menampilkan posisi wilayah yang dipilih"
        />
        <OverViewSelects />
        <OverviewMap />
      </div>

      {/* Endpoint Section */}
      {state.isShowEndpoints && (
        <div className="mb-10">
          <OverviewDescription
            title="Endpoint"
            description="Menampilkan endpoint berdasarkan data yang dipilih"
          />
          <div className="flex flex-col gap-5">
            {endpoints
              .filter((endpoint) => endpoint.condition)
              .map((endpoint) => (
                <EndpointUrl
                  key={endpoint.label}
                  method="GET"
                  url={endpoint.url}
                />
              ))}
          </div>
        </div>
      )}
    </section>
  );
};

export default Overview;
