import OverviewMap from '@/components/overview/overview-map';
import OverViewSelects from '@/components/overview/overview-selects';
import EndpointUrl from '@/components/ui/endpoint-url';
import { useOverview } from '@/context/overview-provider';
import 'leaflet/dist/leaflet.css';
import { ArrowRight, Compass, FileCode, Sparkles } from 'lucide-react';
import { Link } from 'react-router';

const STATS = [
  {
    label: 'Provinsi',
    count: '38',
    description: 'Wilayah tingkat I',
  },
  {
    label: 'Kabupaten & Kota',
    count: '514',
    description: '416 Kab, 98 Kota',
  },
  {
    label: 'Kecamatan / Distrik',
    count: '7.277',
    description: 'Wilayah tingkat III',
  },
  {
    label: 'Desa & Kelurahan',
    count: '83.763',
    description: 'Termasuk kodepos',
  },
];

const Overview = () => {
  const { state } = useOverview();
  const { province, kabKota, kecamatan, desaKel } = state.selected;

  const endpoints = [
    {
      label: 'Provinsi',
      condition: true,
      url: 'https://api.nusantarakita.web.id/v2/provinsi?pagination=true&limit=15',
    },
    {
      label: 'Kab/Kota',
      condition: province && kabKota,
      url: province
        ? `https://api.nusantarakita.web.id/v2/${province.kode}/kab-kota?pagination=true&limit=15`
        : '',
    },
    {
      label: 'Kecamatan',
      condition: province && kabKota && kecamatan,
      url: kabKota
        ? `https://api.nusantarakita.web.id/v2/${kabKota.kode}/kecamatan?pagination=true&limit=15`
        : '',
    },
    {
      label: 'Desa/Kelurahan',
      condition: province && kabKota && kecamatan && desaKel,
      url: kecamatan
        ? `https://api.nusantarakita.web.id/v2/${kecamatan.kode}/desa-kel?pagination=true&limit=15`
        : '',
    },
  ];

  return (
    <article className="text-foreground w-full space-y-10 transition-all">
      {/* Header & Introduction */}
      <div className="space-y-4">
        <div className="border-primary/20 bg-primary/10 text-primary inline-flex items-center gap-2 rounded-md border px-2.5 py-1 text-xs font-semibold">
          <Sparkles className="h-3.5 w-3.5" />
          <span>Dokumentasi Resmi Nusantara Kita</span>
        </div>

        <h1 className="text-foreground text-3xl font-bold tracking-tight sm:text-4xl">
          API Data Wilayah & Geospasial Indonesia
        </h1>

        <p className="text-muted-foreground max-w-3xl text-base leading-relaxed sm:text-lg">
          Akses data wilayah administratif seluruh Indonesia secara cepat,
          akurat, dan terstruktur. Dilengkapi dengan koordinat spasial, GeoJSON
          batas poligon, kodepos, dan fitur pencarian autocomplete.
        </p>

        <div className="flex flex-wrap items-center gap-3 pt-2">
          <Link
            to="/api-documentation/get-provinsi"
            className="bg-primary text-primary-foreground inline-flex cursor-pointer items-center gap-2 rounded-md px-4 py-2.5 text-sm font-semibold shadow-2xs transition-all hover:opacity-90"
          >
            <FileCode className="h-4 w-4" />
            <span>Mulai Jelajahi Endpoint</span>
            <ArrowRight className="h-4 w-4" />
          </Link>

          <Link
            to="/paduan-instalasi"
            className="border-border/80 bg-background hover:bg-muted text-foreground inline-flex cursor-pointer items-center gap-2 rounded-md border px-4 py-2.5 text-sm font-medium transition-colors"
          >
            <Compass className="h-4 w-4" />
            <span>Panduan Instalasi</span>
          </Link>
        </div>
      </div>

      <hr className="border-border/60" />

      {/* 1. Quick Statistics Section */}
      <section
        id="statistik"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Cakupan Data Wilayah</span>
          <a
            href="#statistik"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Dataset wilayah administratif terlengkap yang disusun mengacu pada
          Kepmendagri 2025, mencakup 4 tingkatan pemerintahan di seluruh
          Indonesia:
        </p>

        <div className="grid grid-cols-2 gap-4 pt-2 sm:grid-cols-4">
          {STATS.map((stat) => (
            <div
              key={stat.label}
              className="border-foreground/25 border-l-2 py-1 pl-3.5"
            >
              <div className="text-foreground text-2xl font-bold tracking-tight sm:text-3xl">
                {stat.count}
              </div>
              <div className="text-muted-foreground mt-1 text-xs font-semibold tracking-wider uppercase">
                {stat.label}
              </div>
              <div className="text-muted-foreground/80 mt-0.5 text-xs">
                {stat.description}
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* 2. Interactive Map & Selector Section */}
      <section
        id="map-preview"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Peta Interaktif Geospasial</span>
          <a
            href="#map-preview"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Pilih tingkatan wilayah secara bertahap untuk melihat letak koordinat
          dan batas poligon (GeoJSON) secara realtime di peta interaktif.
        </p>

        {/* Dynamic Selectors Bar */}
        <div className="border-border/80 bg-muted/30 rounded-lg border p-3 sm:p-4">
          <OverViewSelects />
        </div>

        {/* Map Container */}
        <div className="border-border/80 overflow-hidden rounded-lg border">
          <OverviewMap />
        </div>
      </section>

      {/* 3. Endpoint Live Preview Section */}
      {state.isShowEndpoints && (
        <section
          id="generated-endpoints"
          className="scroll-mt-20 space-y-4"
        >
          <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
            <span>Generated Endpoint URL</span>
            <a
              href="#generated-endpoints"
              className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
            >
              #
            </a>
          </h2>
          <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
            Endpoint yang dihasilkan secara otomatis sesuai dengan wilayah yang
            sedang Anda pilih di atas:
          </p>

          <div className="flex flex-col gap-3 pt-1">
            {endpoints
              .filter((endpoint) => endpoint.condition)
              .map((endpoint) => (
                <div
                  key={endpoint.label}
                  className="space-y-1.5"
                >
                  <div className="text-muted-foreground flex items-center gap-2 text-xs font-semibold tracking-wider uppercase">
                    <span>Tingkat:</span>
                    <span className="text-foreground">{endpoint.label}</span>
                  </div>
                  <EndpointUrl
                    method="GET"
                    url={endpoint.url}
                  />
                </div>
              ))}
          </div>
        </section>
      )}
    </article>
  );
};

export default Overview;
