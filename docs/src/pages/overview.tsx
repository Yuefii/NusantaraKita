import OverviewDescription from "@/components/overview/overview-description";
import OverviewMap from "@/components/overview/overview-map";
import OverViewSelects from "@/components/overview/overview-selects";
import EndpointUrl from "@/components/ui/endpoint-url";
import { useOverview } from "@/context/overview-provider";
import { useIsMobile } from "@/hooks/use-mobile";
import { useIsTablet } from "@/hooks/use-tablet";
import { cn } from "@/lib/utils";
import "leaflet/dist/leaflet.css";

const Overview = () => {
  const { state } = useOverview();
  const isMobile = useIsMobile();
  const isTablet = useIsTablet();
  const isSmallScreen = isMobile || isTablet;

  return (
    <section
      className={cn("bg-white text-gray-800 max-w-5xl", {
        "p-5": isSmallScreen,
        "p-8": !isSmallScreen,
      })}
    >
      <h1 className="text-3xl font-bold mb-10 text-gray-700">Overview</h1>

      <div className="mb-10">
        <OverviewDescription
          title="Map View"
          description="Menampilkan Posisi Wilayah Yang dipilih"
        />
        <OverViewSelects />
        <OverviewMap />
      </div>

      {state.isShowEndpoints && (
        <div className="mb-10">
          <OverviewDescription
            title="Endpoint"
            description="Menampilkan Endpoint berdasarkan data yang dipilih"
          />
          <div className="flex flex-col gap-5">
            {/* Provinsi */}
            <EndpointUrl
              method="GET"
              url="https://api-nusakita.vercel.app/v2/provinsi?pagination=true&limit=15"
            />

            {/* Kab/Kota */}
            {state.selected.province && (
              <EndpointUrl
                method="GET"
                url={`https://api-nusakita.vercel.app/v2/${state.selected.province.kode}/kab-kota?pagination=true&limit=15`}
              />
            )}

            {/* Kecamatan */}
            {state.selected.kabKota && (
              <EndpointUrl
                method="GET"
                url={`https://api-nusakita.vercel.app/v2/${state.selected.kabKota.kode}/kecamatan?pagination=true&limit=15`}
              />
            )}

            {/* Desa / Kelurahan */}
            {state.selected.kecamatan && (
              <EndpointUrl
                method="GET"
                url={`https://api-nusakita.vercel.app/v2/${state.selected.kecamatan.kode}/desa-kel?pagination=true&limit=15`}
              />
            )}
          </div>
        </div>
      )}
    </section>
  );
};

export default Overview;
