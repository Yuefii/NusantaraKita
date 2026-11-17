import Code from '@/components/ui/code';
import EndpointUrl from '@/components/ui/endpoint-url';
import { Notes, type NotesProps } from '@/components/ui/notes';
import ResponseSuccess from '@/components/ui/response-success';
import Table from '@/components/ui/table';
import {
  errorHandlingTableRows,
  getKecamatanResponseStructureData,
  kecamatanResponse,
} from '@/constant/get-kecamatan.constant';
import {
  errorHandlingTableHeaders,
  queryTableHeaders,
  queryTableRows,
} from '@/constant/get-provinsi.constant';
import { Layers } from 'lucide-react';

const notesItems: NotesProps['items'] = [
  <>
    Jika <code>pagination=false</code>, field <code>pagination</code> tidak akan
    muncul di response.
  </>,
  'Parameter query dapat digabungkan sesuai kebutuhan filter data.',
];

export const GetKecamatan = () => {
  return (
    <article className="text-foreground w-full space-y-10 transition-all">
      {/* Header */}
      <div className="space-y-3">
        <div className="inline-flex items-center gap-2 rounded-md border border-emerald-500/20 bg-emerald-500/10 px-2.5 py-1 text-xs font-semibold text-emerald-600 dark:text-emerald-400">
          <Layers className="h-3.5 w-3.5" />
          <span>API Reference • GET /v2/kecamatan</span>
        </div>
        <h1 className="text-foreground text-3xl font-bold tracking-tight sm:text-4xl">
          GET Kecamatan
        </h1>
        <p className="text-muted-foreground text-base leading-relaxed sm:text-lg">
          Mengambil daftar kecamatan di Indonesia. Mendukung pencarian seluruh
          kecamatan atau filter spesifik berdasarkan kabupaten/kota induknya.
        </p>
      </div>

      <hr className="border-border/60" />

      {/* 1. Endpoints Section */}
      <section
        id="endpoint"
        className="scroll-mt-20 space-y-5"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Endpoint URL</span>
          <a
            href="#endpoint"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>

        <div className="space-y-5">
          <div className="space-y-2">
            <h3 className="text-foreground text-sm font-semibold">
              1. Semua Kecamatan (Global)
            </h3>
            <EndpointUrl
              method="GET"
              url="https://api.nusantarakita.web.id/v2/kecamatan"
            />
          </div>

          <div className="space-y-2">
            <h3 className="text-foreground text-sm font-semibold">
              2. Kecamatan Berdasarkan Kode Kabupaten / Kota
            </h3>
            <EndpointUrl
              method="GET"
              url="https://api.nusantarakita.web.id/v2/{kodeKabKota}/kecamatan"
            />
            <span className="text-muted-foreground mt-2 block text-xs">
              Ganti <code>{'{kodeKabKota}'}</code> dengan kode kabupaten/kota
              (contoh: <code>32.73</code> untuk Kota Bandung).
            </span>
          </div>

          <div className="space-y-2">
            <h3 className="text-foreground text-sm font-semibold">
              3. Detail Kecamatan by Kode
            </h3>
            <EndpointUrl
              method="GET"
              url="https://api.nusantarakita.web.id/v2/kecamatan/{kode}"
            />
            <span className="text-muted-foreground mt-2 block text-xs">
              Ganti <code>{'{kode}'}</code> dengan kode 6-digit titik kecamatan
              (contoh: <code>32.73.01</code> untuk Sukasari).
            </span>
          </div>
        </div>
      </section>

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
            Semua Kecamatan (Dengan Paginasi)
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/kecamatan"
          />
        </div>

        <div className="space-y-3">
          <div className="text-muted-foreground text-xs font-semibold tracking-wider uppercase">
            Kecamatan di Kabupaten Aceh Selatan (11.01)
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/11.01/kecamatan"
          />
        </div>

        <div className="space-y-3">
          <div className="text-muted-foreground text-xs font-semibold tracking-wider uppercase">
            Custom Limit Per Halaman
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/kecamatan?limit=20"
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
        <ResponseSuccess items={getKecamatanResponseStructureData} />
      </section>

      {/* 4. Response Examples Section */}
      <section
        id="contoh-response"
        className="scroll-mt-20 space-y-6"
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

        <div className="space-y-6">
          <div className="space-y-2">
            <h3 className="text-foreground text-sm font-semibold">
              Contoh Response Default:
            </h3>
            <Code
              content={JSON.stringify(kecamatanResponse, null, 2)}
              showCopyButton
            />
          </div>

          <div className="space-y-2">
            <h3 className="text-foreground text-sm font-semibold">
              Contoh Response Tanpa Paginasi:
            </h3>
            <Code
              content={JSON.stringify(
                { data: kecamatanResponse.data },
                null,
                2,
              )}
              showCopyButton
            />
          </div>
        </div>
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
