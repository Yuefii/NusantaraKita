import EndpointUrl from "@/components/ui/endpoint-url";
import TableWithTitle from "@/components/ui/query-table-parameters";
import { useIsMobile } from "@/hooks/use-mobile";
import { useIsTablet } from "@/hooks/use-tablet";
import { cn } from "@/lib/utils";
import { queryTableHeaders, queryTableRows } from "./get-provinsi";
import ResponseSuccess from "@/components/ui/response-success";
import type { ResponseJsonListItem } from "@/components/ui/response-json-structure";
import Code from "@/components/ui/code";
import { Notes, type NotesProps } from "@/components/ui/notes";

const getKecamatanResponseStructureData: ResponseJsonListItem[] = [
  {
    name: "pagination",
    details: "(object) - Hanya muncul jika pagination=true :",
    children: [
      { name: "total_item", details: "(integer) - Total seluruh kecamatan" },
      {
        name: "total_halaman",
        details: "(integer) - Total halaman berdasarkan limit",
      },
      {
        name: "halaman_saat_ini",
        details: "(integer) - Halaman yang sedang diakses",
      },
      {
        name: "ukuran_halaman",
        details: "(integer) - Jumlah item per halaman",
      },
    ],
  },
  {
    name: "data",
    details: "(array) - Daftar kecamatan:",
    children: [
      { name: "kode", details: "(string) - Kode unik kecamatan" },
      { name: "nama", details: "(string) - Nama kecamatan" },
      {
        name: "kode_kabupaten_kota",
        details: "(string) - Kode kabupaten/kota induk",
      },
      { name: "lat", details: "(double) - Koordinat latitude" },
      { name: "lng", details: "(double) - Koordinat longitude" },
    ],
  },
];

const kecamatanResponse = {
  pagination: {
    total_item: 516,
    total_halaman: 52,
    halaman_saat_ini: 1,
    ukuran_halaman: 10,
  },
  data: [
    {
      kode: "11.01.01",
      nama: "Bakongan",
      lat: 2.960325743420683,
      lng: 97.46087307098534,
      kode_kabupaten_kota: "11.01",
    },
  ],
};

const errorHandlingTableHeaders = ["Code", "Description", "Example"];
const errorHandlingTableRows = [
  {
    Code: "400",
    Description: "Parameter tidak valid",
    Example: JSON.stringify(
      { error: "nomor halaman tidak valid, halaman harus lebih besar dari 0" },
      null,
      2
    ),
  },
  {
    Code: "404",
    Description: "Data tidak ditemukan",
    Example: JSON.stringify(
      { error: "tidak ditemukan data kecamatan" },
      null,
      2
    ),
  },
  {
    Code: "500",
    Description: "Server Error",
    Example: JSON.stringify(
      { error: "Gagal mengambil data: [error detail]" },
      null,
      2
    ),
  },
];

const notesItems: NotesProps["items"] = [
  <>
    Jika <code className="bg-gray-300 px-1 rounded">pagination=false</code>,
    field <code className="bg-gray-100 px-1 rounded">pagination</code> tidak
    akan muncul di response
  </>,
  "Parameter dapat digabungkan sesuai kebutuhan",
  "Untuk endpoint berdasarkan kabupaten/kota, pastikan kode kabupaten/kota valid",
];

export const GetKecamatan = () => {
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
      <h1 className="text-3xl font-bold mb-10 text-gray-700">
        API Documentation - GET Kecamatan
      </h1>

      <div className="mb-10">
        <div className="mb-6">
          <h2 className="text-xl font-semibold mb-4 text-gray-600">
            Endpoint API (Semua Kecamatan)
          </h2>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/kecamatan"
          />
        </div>

        <div className="mb-6">
          <h2 className="text-xl font-semibold mb-4 text-gray-600">
            Endpoint API (Kecamatan berdasarkan kode Kabupaten/Kota)
          </h2>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/{kodeKabKota}/kecamatan"
          />
          <span className="text-sm mt-1 block">
            Ganti {"{"}kodeKabKota{"}"} dengan kode kabupaten/kota yang valid
          </span>
        </div>
      </div>

      <div className="mb-10">
        <h2 className="text-xl font-semibold mb-4 text-gray-600">
          Description
        </h2>
        <p className="text-gray-600 leading-relaxed">
          API ini digunakan untuk mendapatkan daftar kecamatan beserta informasi
          geografisnya. Dapat menampilkan semua kecamatan atau berdasarkan
          kabupaten/kota tertentu. Response akan menampilkan data dengan
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
        <h2 className="text-2xl font-semibold mb-6 text-gray-700">
          Example Requests
        </h2>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Semua Kecamatan (dengan paginasi)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/kecamatan"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Semua Kecamatan (tanpa paginasi)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/kecamatan?pagination=false"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Custom limit per halaman
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/kecamatan?limit=20"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Menuju halaman tertentu
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/kecamatan?page=2"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Kecamatan berdasarkan kode kabupaten/kota
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/11.01/kecamatan"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Gabungan parameter
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/kecamatan?pagination=true&limit=5&page=3"
          />
        </div>
      </div>

      <div className="mb-10">
        <ResponseSuccess items={getKecamatanResponseStructureData} />
      </div>

      <div className="mb-10">
        <h3 className="text-base font-semibold mb-4 text-gray-700">
          Example Response (Default):
        </h3>
        <Code
          content={JSON.stringify(kecamatanResponse, null, 2)}
          showCopyButton
        />
      </div>

      <div className="mb-10">
        <h3 className="text-base font-semibold mb-4 text-gray-700">
          Example Response (Without Pagination):
        </h3>
        <Code
          content={JSON.stringify({ data: kecamatanResponse.data }, null, 2)}
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
