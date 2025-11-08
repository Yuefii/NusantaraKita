import Code from '@/components/ui/code';
import EndpointUrl from '@/components/ui/endpoint-url';
import { Notes, type NotesProps } from '@/components/ui/notes';
import TableWithTitle from '@/components/ui/query-table-parameters';
import ResponseSuccess from '@/components/ui/response-success';
import {
  errorHandlingTableHeaders,
  errorHandlingTableRows,
  getProvinsiResponseStructureData,
  queryTableHeaders,
  queryTableRows,
  responseExampleDefault,
  responseExampleWithPagination,
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
];

export const GetProvinsi = () => {
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
        API Documentation - GET Provinsi
      </h1>

      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Endpoint API
        </h2>
        <EndpointUrl
          method="GET"
          url="https://api-nusakita.vercel.app/v2/provinsi"
        />
      </div>

      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Description
        </h2>
        <p className="leading-relaxed text-gray-600">
          API ini digunakan untuk mendapatkan daftar provinsi beserta informasi
          geografisnya. Response akan menampilkan data provinsi dengan paginasi
          default, tetapi dapat dimodifikasi melalui parameter query.
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
            Default (dengan paginasi)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Tanpa paginasi
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi?pagination=false"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Custom limit per halaman
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi?limit=20"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Menuju halaman tertentu
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi?halaman=2"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Gabungan parameter
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi?pagination=true&limit=5&halaman=3"
          />
        </div>
      </div>

      <div className="mb-10">
        <ResponseSuccess items={getProvinsiResponseStructureData} />
      </div>

      <div className="mb-10">
        <h3 className="mb-4 text-base font-semibold text-gray-700">
          Example Response (Default):
        </h3>
        <Code
          content={JSON.stringify(responseExampleDefault, null, 2)}
          showCopyButton
        />
      </div>

      <div className="mb-10">
        <h3 className="mb-4 text-base font-semibold text-gray-700">
          Example Response (Without Pagination):
        </h3>
        <Code
          content={JSON.stringify(responseExampleWithPagination, null, 2)}
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
