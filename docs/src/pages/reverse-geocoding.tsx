import Code from '@/components/ui/code';
import EndpointUrl from '@/components/ui/endpoint-url';
import { Notes, type NotesProps } from '@/components/ui/notes';
import ResponseSuccess from '@/components/ui/response-success';
import Table from '@/components/ui/table';
import {
  errorHandlingTableHeaders,
  errorHandlingTableRows,
  queryTableHeaders,
  queryTableRows,
  responseExampleReverse,
  reverseResponseStructureData,
} from '@/constant/reverse-geocoding.constant';
import { MapPin } from 'lucide-react';

const notesItems: NotesProps['items'] = [
  'Parameter lat dan lng wajib diisi dengan angka desimal floating-point yang valid.',
  'Titik koordinat harus berada di dalam batas wilayah kepulauan Indonesia (garis batas latitude -11 s/d 6 dan longitude 95 s/d 141).',
  'Field jarak_meter menghitung jarak Euclidean/Haversine dari koordinat input ke titik pusat administratif desa/kelurahan terdekat.',
  'Setiap tingkatan wilayah di dalam data response dilengkapi dengan URL CDN file GeoJSON batas wilayah.',
];

export const ReverseGeocodingApiDoc = () => {
  return (
    <article className="text-foreground w-full space-y-10 transition-all">
      {/* Header */}
      <div className="space-y-3">
        <div className="inline-flex items-center gap-2 rounded-md border border-rose-500/20 bg-rose-500/10 px-2.5 py-1 text-xs font-semibold text-rose-600 dark:text-rose-400">
          <MapPin className="h-3.5 w-3.5" />
          <span>API Reference • GET /v2/reverse</span>
        </div>
        <h1 className="text-foreground text-3xl font-bold tracking-tight sm:text-4xl">
          Reverse Geocoding
        </h1>
        <p className="text-muted-foreground text-base leading-relaxed sm:text-lg">
          Mengonversi titik koordinat GPS (latitude dan longitude) menjadi
          informasi hierarki administratif lengkap wilayah Indonesia (Provinsi,
          Kab/Kota, Kecamatan, Desa, dan Kodepos).
        </p>
      </div>

      <hr className="border-border/60" />

      {/* 1. Endpoints Section */}
      <section
        id="endpoint"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Endpoint Utama</span>
          <a
            href="#endpoint"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <EndpointUrl
          method="GET"
          url="https://api.nusantarakita.web.id/v2/reverse?lat={latitude}&lng={longitude}"
        />
        <span className="text-muted-foreground block text-xs">
          Ganti parameter <code>{'{latitude}'}</code> dan{' '}
          <code>{'{longitude}'}</code> dengan nilai koordinat GPS floating
          point.
        </span>
      </section>

      {/* Description Subsection */}
      <div className="space-y-2">
        <h3 className="text-foreground text-base font-semibold">
          Deskripsi Fungsi
        </h3>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Endpoint Reverse Geocoding memungkinkan Anda mengonversi koordinat
          geografis GPS (latitude & longitude) menjadi informasi hierarki
          administratif wilayah Indonesia yang lengkap: mulai dari nama
          Provinsi, Kabupaten/Kota, Kecamatan, hingga Desa/Kelurahan beserta
          kode pos dan estimasi jarak ke titik pusat administratif.
        </p>
      </div>

      {/* 2. Query Parameters Section */}
      <section
        id="parameter-query"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Parameter Query</span>
          <a
            href="#parameter-query"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <Table
          columnHeaders={queryTableHeaders}
          rowData={queryTableRows}
        />
      </section>

      {/* Example Requests Subsection */}
      <div className="space-y-4 pt-2">
        <h3 className="text-foreground text-base font-semibold">
          Contoh Penggunaan Request
        </h3>

        <div className="space-y-3">
          <div className="text-muted-foreground text-xs font-semibold tracking-wider uppercase">
            Contoh Pencarian Titik Kota Bandung
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/reverse?lat=-6.874112&lng=107.589123"
          />
        </div>

        <div className="space-y-3">
          <div className="text-muted-foreground text-xs font-semibold tracking-wider uppercase">
            Contoh Pencarian Titik Monas (DKI Jakarta)
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/reverse?lat=-6.175392&lng=106.827153"
          />
        </div>
      </div>

      {/* 3. Response Structure Section */}
      <section
        id="struktur-response"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Struktur Data Response</span>
          <a
            href="#struktur-response"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <ResponseSuccess items={reverseResponseStructureData} />
      </section>

      {/* 4. Response Examples Section */}
      <section
        id="contoh-response"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Contoh Response JSON</span>
          <a
            href="#contoh-response"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <Code
          content={JSON.stringify(responseExampleReverse, null, 2)}
          showCopyButton
        />
      </section>

      {/* 5. Error Handling Section */}
      <section
        id="error-handling"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Error Handling</span>
          <a
            href="#error-handling"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <Table
          columnHeaders={errorHandlingTableHeaders}
          rowData={errorHandlingTableRows}
        />
      </section>

      {/* Notes / Callout */}
      <Notes items={notesItems} />
    </article>
  );
};
