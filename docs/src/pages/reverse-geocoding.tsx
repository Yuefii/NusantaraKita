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
  responseExampleReverse,
  reverseResponseStructureData,
} from '@/constant/reverse-geocoding.constant';
import { useIsMobile } from '@/hooks/use-mobile';
import { useIsTablet } from '@/hooks/use-tablet';
import { cn } from '@/lib/utils';

const notesItems: NotesProps['items'] = [
  'Parameter lat dan lng wajib diisi dengan angka desimal floating-point yang valid.',
  'Titik koordinat harus berada di dalam batas wilayah kepulauan Indonesia (garis batas latitude -11 s/d 6 dan longitude 95 s/d 141).',
  'Field jarak_meter menghitung jarak Euclidean/Haversine dari koordinat input ke titik pusat administratif desa/kelurahan terdekat.',
  'Setiap tingkatan wilayah di dalam data response dilengkapi dengan URL CDN file GeoJSON.',
];

export const ReverseGeocodingApiDoc = () => {
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
        API Documentation - Reverse Geocoding
      </h1>

      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Endpoint API
        </h2>
        <EndpointUrl
          method="GET"
          url="https://api.nusantarakita.web.id/v2/reverse?lat={latitude}&lng={longitude}"
        />
      </div>

      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Description
        </h2>
        <p className="leading-relaxed text-gray-600">
          Endpoint Reverse Geocoding memungkinkan Anda mengonversi koordinat
          geografis GPS (latitude & longitude) menjadi informasi hierarki
          administratif wilayah Indonesia yang lengkap: mulai dari nama
          Provinsi, Kabupaten/Kota, Kecamatan, hingga Desa/Kelurahan beserta
          kode pos dan perkiraan jarak ke pusat wilayah.
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
            Contoh Pencarian Koordinat Kota Bandung
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/reverse?lat=-6.874112&lng=107.589123"
          />
        </div>

        <div className="mb-6">
          <h3 className="mb-2 text-lg font-medium text-gray-600">
            Contoh Pencarian Koordinat Monas (Jakarta)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api.nusantarakita.web.id/v2/reverse?lat=-6.175392&lng=106.827153"
          />
        </div>
      </div>

      <div className="mb-10">
        <ResponseSuccess items={reverseResponseStructureData} />
      </div>

      <div className="mb-10">
        <h3 className="mb-4 text-base font-semibold text-gray-700">
          Example Response:
        </h3>
        <Code
          content={JSON.stringify(responseExampleReverse, null, 2)}
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
