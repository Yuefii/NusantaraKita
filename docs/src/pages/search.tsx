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
  responseExampleSearch,
  searchResponseStructureData,
} from '@/constant/search.constant';
import { Search } from 'lucide-react';

const notesItems: NotesProps['items'] = [
  'Parameter q minimal terdiri dari 2 karakter.',
  'Pencarian mendukung nama provinsi, kabupaten/kota, kecamatan, desa/kelurahan, dan 5 digit kode pos.',
  'Hasil pencarian diurutkan berdasarkan tingkat relevansi tertinggi (Exact Match > Prefix Match > Substring Match).',
  'Field hierarki mempermudah integrasi fitur autocomplete atau dropdown pencarian alamat di antarmuka web dan mobile.',
];

export const SearchApiDoc = () => {
  return (
    <article className="text-foreground w-full space-y-10 transition-all">
      {/* Header */}
      <div className="space-y-3">
        <div className="inline-flex items-center gap-2 rounded-md border border-amber-500/20 bg-amber-500/10 px-2.5 py-1 text-xs font-semibold text-amber-600 dark:text-amber-400">
          <Search className="h-3.5 w-3.5" />
          <span>API Reference • GET /v2/search</span>
        </div>
        <h1 className="text-foreground text-3xl font-bold tracking-tight sm:text-4xl">
          Search & Autocomplete
        </h1>
        <p className="text-muted-foreground text-base leading-relaxed sm:text-lg">
          Pencarian cepat nama wilayah secara global. Sangat optimal untuk input
          autocomplete, dropdown formulir pemesanan, filter logistik, atau
          pencarian kodepos.
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
          url="https://api.nusantarakita.web.id/v2/search?q={keyword}"
        />
        <span className="text-muted-foreground block text-xs">
          Ganti <code>{'{keyword}'}</code> dengan kata kunci nama wilayah atau
          kodepos (contoh: <code>bandung</code> atau <code>40151</code>).
        </span>
      </section>

      {/* Description Subsection */}
      <div className="space-y-2">
        <h3 className="text-foreground text-base font-semibold">
          Deskripsi Fungsi
        </h3>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Endpoint ini dirancang khusus untuk kebutuhan pencarian cepat dan
          autocomplete nama wilayah di Indonesia. Anda dapat mencari lintas
          seluruh tingkat administratif (provinsi, kabupaten/kota, kecamatan,
          dan kelurahan/desa) atau membatasi cakupan pencarian menggunakan
          parameter <code>tipe</code>.
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
            1. Pencarian Umum (Semua Tingkat Wilayah)
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/search?q=bandung"
          />
        </div>

        <div className="space-y-3">
          <div className="text-muted-foreground text-xs font-semibold tracking-wider uppercase">
            2. Pencarian Khusus Tingkat Kabupaten / Kota
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/search?q=bandung&tipe=kab-kota"
          />
        </div>

        <div className="space-y-3">
          <div className="text-muted-foreground text-xs font-semibold tracking-wider uppercase">
            3. Pencarian Berdasarkan Kode Pos 5 Digit
          </div>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/search?q=40151"
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
        <ResponseSuccess items={searchResponseStructureData} />
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
          content={JSON.stringify(responseExampleSearch, null, 2)}
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
