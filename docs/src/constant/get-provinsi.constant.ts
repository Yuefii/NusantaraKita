import type { ResponseJsonListItem } from '@/components/ui/response-json-structure';

export const queryTableHeaders = [
  'Parameter',
  'Type',
  'Description',
  'Default',
];
export const errorHandlingTableHeaders = ['Code', 'Description', 'Example'];

export const queryTableRows = [
  {
    Parameter: 'pagination',
    Type: 'boolean',
    Description: 'Menampilkan/menyembunyikan informasi paginasi (true/false)',
    Default: 'true',
  },
  {
    Parameter: 'limit',
    Type: 'integer',
    Description: 'Menentukan jumlah item per halaman',
    Default: '10',
  },
  {
    Parameter: 'halaman',
    Type: 'integer',
    Description: 'Menentukan halaman yang ingin ditampilkan',
    Default: '1',
  },
];

export const errorHandlingTableRows = [
  {
    Code: '400',
    Description: 'Parameter tidak valid',
    Example: JSON.stringify(
      { error: 'nomor halaman tidak valid, halaman harus lebih besar dari 0' },
      null,
      2,
    ),
  },
  {
    Code: '404',
    Description: 'Data tidak ditemukan',
    Example: JSON.stringify({ error: 'tidak ditemukan data' }, null, 2),
  },
  {
    Code: '500',
    Description: 'Server Error',
    Example: JSON.stringify(
      { error: 'Gagal mengambil data: [error detail]' },
      null,
      2,
    ),
  },
];

export const getProvinsiResponseStructureData: ResponseJsonListItem[] = [
  {
    name: 'pagination',
    details: '(object) - Hanya muncul jika pagination=true :',
    children: [
      { name: 'total_item', details: '(integer) - Total seluruh provinsi' },
      {
        name: 'total_halaman',
        details: '(integer) - Total halaman berdasarkan limit',
      },
      {
        name: 'halaman_saat_ini',
        details: '(integer) - Halaman yang sedang diakses',
      },
      {
        name: 'ukuran_halaman',
        details: '(integer) - Jumlah item per halaman',
      },
    ],
  },
  {
    name: 'data',
    details: '(array) - Daftar provinsi:',
    children: [
      { name: 'kode', details: '(string) - Kode unik provinsi' },
      { name: 'nama', details: '(string) - Nama provinsi' },
      { name: 'lat', details: '(double) - Koordinat latitude' },
      { name: 'lng', details: '(double) - Koordinat longitude' },
      { name: 'geojson_url', details: '(string) - URL GeoJSON provinsi' },
    ],
  },
];

export const responseExampleDefault = {
  pagination: {
    total_item: 38,
    total_halaman: 4,
    halaman_saat_ini: 1,
    ukuran_halaman: 10,
  },
  data: [
    {
      kode: '11',
      nama: 'Aceh',
      lat: 4.225728583038235,
      lng: 96.91187408609952,
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/provinsi/11.geojson',
    },
  ],
};

export const responseExampleWithPagination = {
  data: [
    {
      kode: '11',
      nama: 'Aceh',
      lat: 4.225728583038235,
      lng: 96.91187408609952,
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/provinsi/11.geojson',
    },
  ],
};
