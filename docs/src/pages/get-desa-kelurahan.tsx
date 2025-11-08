import Code from '@/components/ui/code';
import EndpointUrl from '@/components/ui/endpoint-url';
import { Notes, type NotesProps } from '@/components/ui/notes';
import TableWithTitle from '@/components/ui/query-table-parameters';
import ResponseSuccess from '@/components/ui/response-success';
import {
  errorHandlingTableRows,
  getDesaKelurahanExampleResponse,
  getDesaKelurahanResponseStructureData,
} from '@/constant/get-desa-kelurahan.constant';
import {
  errorHandlingTableHeaders,
  queryTableHeaders,
  queryTableRows,
} from '@/constant/get-provinsi.constant';
import { useIsMobile } from '@/hooks/use-mobile';
import { useIsTablet } from '@/hooks/use-tablet';
import { cn } from '@/lib/utils';

const notesItems: NotesProps['items'] = [
  <>
    Jika <code className="rounded bg-gray-300 px-1">pagination=false</code>,
    field <code className="rounded bg-gray-100 px-1">pagination</code> tidak
    akan muncul di response
  </>,
  'Parameter dapat digabungkan sesuai kebutuhan',
  'Untuk endpoint berdasarkan Kecamatan, pastikan kode Kecamatan valid',
];

export const GetDesaKelurahan = () => {
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
        API Documentation - GET Desa/Kelurahan
      </h1>

      <div className="mb-10">
        <div className="mb-6">
          <h2 className="mb-4 text-xl font-semibold text-gray-600">
            Endpoint API (Semua Desa/Kelurahan)
          </h2>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/desa-kel"
          />
        </div>

        <div className="mb-6">
          <h2 className="mb-4 text-xl font-semibold text-gray-600">
            Endpoint API (Desa/Keluarahan berdasarkan kode Kecamatan)
          </h2>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/{kodeKecamatan}/desa-kel"
          />
          <span className="mt-1 block text-sm">
            Ganti {'{'}kodeKecamatan{'}'} dengan kode Kecamatan yang valid
          </span>
        </div>
      </div>

      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Description
        </h2>
        <p className="leading-relaxed text-gray-600">
          API ini digunakan untuk mendapatkan daftar Desa/Keluarahan beserta
          informasi geografisnya. Dapat menampilkan semua Desa/Kelurahan atau
          berdasarkan Kecamatan tertentu. Response akan menampilkan data dengan
          paginasi default, tetapi dapat dimodifikasi melalui parameter query.
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
            Semua Desa/Kelurahan (dengan paginasi)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/desa-kel"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Semua Desa/Kelurahan (tanpa paginasi)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/desa-kel?pagination=false"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Custom limit per halaman
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/desa-kel?limit=20"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Menuju halaman tertentu
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/desa-kel?page=2"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Desa/Kelurahan berdasarkan kode Kecamatan
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/11.01/desa-kel"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Gabungan parameter
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/desa-kel?pagination=true&limit=5&page=3"
          />
        </div>
      </div>

      <div className="mb-10">
        <ResponseSuccess items={getDesaKelurahanResponseStructureData} />
      </div>

      <div className="mb-10">
        <h3 className="mb-4 text-base font-semibold text-gray-700">
          Example Response (Default):
        </h3>
        <Code
          content={JSON.stringify(getDesaKelurahanExampleResponse, null, 2)}
          showCopyButton
        />
      </div>

      <div className="mb-10">
        <h3 className="mb-4 text-base font-semibold text-gray-700">
          Example Response (Without Pagination):
        </h3>
        <Code
          content={JSON.stringify(
            { data: getDesaKelurahanExampleResponse.data },
            null,
            2,
          )}
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
