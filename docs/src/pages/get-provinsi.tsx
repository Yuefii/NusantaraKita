import Code from "@/components/ui/code";
import EndpointUrl from "@/components/ui/endpoint-url";
import { Notes, type NotesProps } from "@/components/ui/notes";
import TableWithTitle from "@/components/ui/query-table-parameters";
import type { ResponseJsonListItem } from "@/components/ui/response-json-structure";
import ResponseSuccess from "@/components/ui/response-success";
import { useIsMobile } from "@/hooks/use-mobile";
import { useIsTablet } from "@/hooks/use-tablet";
import { cn } from "@/lib/utils";

export const queryTableHeaders = [
  "Parameter",
  "Type",
  "Description",
  "Default",
];
const errorHandlingTableHeaders = ["Code", "Description", "Example"];

export const queryTableRows = [
  {
    Parameter: "pagination",
    Type: "boolean",
    Description: "Menampilkan/menyembunyikan informasi paginasi (true/false)",
    Default: "true",
  },
  {
    Parameter: "limit",
    Type: "integer",
    Description: "Menentukan jumlah item per halaman",
    Default: "10",
  },
  {
    Parameter: "halaman",
    Type: "integer",
    Description: "Menentukan halaman yang ingin ditampilkan",
    Default: "1",
  },
];

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
    Example: JSON.stringify({ error: "tidak ditemukan data" }, null, 2),
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

const getProvinsiResponseStructureData: ResponseJsonListItem[] = [
  {
    name: "pagination",
    details: "(object) - Hanya muncul jika pagination=true :",
    children: [
      { name: "total_item", details: "(integer) - Total seluruh provinsi" },
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
    details: "(array) - Daftar provinsi:",
    children: [
      { name: "kode", details: "(string) - Kode unik provinsi" },
      { name: "nama", details: "(string) - Nama provinsi" },
      { name: "lat", details: "(double) - Koordinat latitude" },
      { name: "lng", details: "(double) - Koordinat longitude" },
    ],
  },
];

const responseExampleDefault = {
  pagination: {
    total_item: 38,
    total_halaman: 4,
    halaman_saat_ini: 1,
    ukuran_halaman: 10,
  },
  data: [
    {
      kode: "11",
      nama: "Aceh",
      lat: 4.225728583038235,
      lng: 96.91187408609952,
    },
  ],
};

const responseExampleWithPagination = {
  data: [
    {
      kode: "11",
      nama: "Aceh",
      lat: 4.225728583038235,
      lng: 96.91187408609952,
    },
  ],
};

const notesItems: NotesProps["items"] = [
  <>
    Jika <code className="bg-gray-300 px-1 rounded">pagination=false</code>,
    field <code className="bg-gray-100 px-1 rounded">pagination</code> tidak
    akan muncul di response
  </>,
  "Parameter dapat digabungkan sesuai kebutuhan",
];

export const GetProvinsi = () => {
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
        API Documentation - GET Provinsi
      </h1>

      <div className="mb-10">
        <h2 className="text-xl font-semibold mb-4 text-gray-600">
          Endpoint API
        </h2>
        <EndpointUrl
          method="GET"
          url="https://api-nusakita.vercel.app/v2/provinsi"
        />
      </div>

      <div className="mb-10">
        <h2 className="text-xl font-semibold mb-4 text-gray-600">
          Description
        </h2>
        <p className="text-gray-600 leading-relaxed">
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
        <h2 className="text-2xl font-semibold mb-6 text-gray-700">
          Example Requests
        </h2>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Default (dengan paginasi)
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Tanpa paginasi
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi?pagination=false"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Custom limit per halaman
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi?limit=20"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
            Menuju halaman tertentu
          </h3>
          <EndpointUrl
            method="GET"
            url="https://api-nusakita.vercel.app/v2/provinsi?halaman=2"
          />
        </div>

        <div className="mb-6">
          <h3 className="text-lg font-medium mb-2 text-gray-600">
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
        <h3 className="text-base font-semibold mb-4 text-gray-700">
          Example Response (Default):
        </h3>
        <Code
          content={JSON.stringify(responseExampleDefault, null, 2)}
          showCopyButton
        />
      </div>

      <div className="mb-10">
        <h3 className="text-base font-semibold mb-4 text-gray-700">
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
