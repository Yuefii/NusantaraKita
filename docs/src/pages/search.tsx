import Code from '@/components/ui/code';
import EndpointUrl from '@/components/ui/endpoint-url';
import { Notes, type NotesProps } from '@/components/ui/notes';
import TableWithTitle from '@/components/ui/query-table-parameters';
import ResponseSuccess from '@/components/ui/response-success';
import {
  errorHandlingTableHeaders,
  errorHandlingTableRows,
  queryTableHeaders,
  queryTableRows,
  responseExampleSearch,
  searchResponseStructureData,
} from '@/constant/search.constant';
import { useIsMobile } from '@/hooks/use-mobile';
import { useIsTablet } from '@/hooks/use-tablet';
import { cn } from '@/lib/utils';

const notesItems: NotesProps['items'] = [
  'Parameter q minimal terdiri dari 2 karakter.',
  'Pencarian mendukung nama provinsi, kabupaten/kota, kecamatan, desa/kelurahan, dan 5 digit kode pos.',
  'Hasil pencarian diurutkan berdasarkan tingkat relevansi tertinggi (Exact Match > Prefix Match > Substring Match).',
  'Field hierarki mempermudah integrasi fitur autocomplete atau dropdown pencarian alamat di frontend.',
];

export const SearchApiDoc = () => {
  const isMobile = useIsMobile();
  const isTablet = useIsTablet();
  const isSmallScreen = isMobile || isTablet;

  return (
    <section
      className={cn('max-w-5xl bg-white text-gray-800', {
        'p-5': isSmallScreen,
        'p-8': !isSmallScreen,
      })}
    >
      <h1 className="mb-10 text-3xl font-bold text-gray-700">
        API Documentation - Search / Autocomplete
      </h1>

      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Endpoint API
        </h2>
        <EndpointUrl
          method="GET"
          url="https://api.nusantarakita.web.id/v2/search?q={keyword}"
        />
      </div>

      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Description
        </h2>
        <p className="leading-relaxed text-gray-600">
          Endpoint ini dirancang khusus untuk kebutuhan pencarian cepat dan autocomplete nama wilayah di Indonesia.
          Anda dapat mencari lintas seluruh tingkat administratif (provinsi, kabupaten/kota, kecamatan, dan kelurahan/desa)
          atau membatasi pencarian pada jenis wilayah tertentu menggunakan parameter <code className="rounded bg-gray-200 px-1">tipe</code>.
        </p>
      </div>

      <div className="mb-10">
        <TableWithTitle
          titleText="Query Parameters"
          columnHeaders={queryTableHeaders}
          rowData={queryTableRows}
        />
      </div>

      <div className="mb-10">
        <h2 className="mb-6 text-2xl font-semibold text-gray-700">
          Example Requests
        </h2>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Pencarian umum (semua jenis wilayah)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/search?q=bandung"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Pencarian khusus tingkat Kabupaten/Kota
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/search?q=bandung&tipe=kab-kota"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Pencarian berdasarkan Kode Pos
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/search?q=40151"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Pencarian dengan paginasi kustom
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/search?q=sukajadi&limit=5&halaman=1"
          />
        </div>
      </div>

      <div className="mb-10">
        <ResponseSuccess items={searchResponseStructureData} />
      </div>

      <div className="mb-10">
        <h3 className="mb-4 text-base font-semibold text-gray-700">
          Example Response:
        </h3>
        <Code
          content={JSON.stringify(responseExampleSearch, null, 2)}
          showCopyButton
        />
      </div>

      <div className="mb-10">
        <TableWithTitle
          titleText="Error Handling"
          columnHeaders={errorHandlingTableHeaders}
          rowData={errorHandlingTableRows}
        />
      </div>

      <div className="mb-10">
        <Notes items={notesItems} />
      </div>
    </section>
  );
};
