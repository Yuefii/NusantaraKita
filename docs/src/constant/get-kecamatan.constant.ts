import type { ResponseJsonListItem } from '@/components/ui/response-json-structure';

const getKecamatanResponseStructureData: ResponseJsonListItem[] = [
  {
    name: 'pagination',
    details: '(object) - Hanya muncul jika pagination=true :',
    children: [
      { name: 'total_item', details: '(integer) - Total seluruh kecamatan' },
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
    details: '(array) - Daftar kecamatan:',
    children: [
      { name: 'kode', details: '(string) - Kode unik kecamatan' },
      { name: 'nama', details: '(string) - Nama kecamatan' },
      {
        name: 'kode_kabupaten_kota',
        details: '(string) - Kode kabupaten/kota induk',
      },
      { name: 'lat', details: '(double) - Koordinat latitude' },
      { name: 'lng', details: '(double) - Koordinat longitude' },
      { name: 'geojson_url', details: '(string) - URL GeoJSON kecamatan' },
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
      kode: '11.01.01',
      nama: 'Bakongan',
      lat: 2.960325743420683,
      lng: 97.46087307098534,
      kode_kabupaten_kota: '11.01',
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kecamatan/11.01.01.geojson',
    },
  ],
};

const errorHandlingTableRows = [
  {
    Code: '404',
    Description: 'Data tidak ditemukan',
    Example: JSON.stringify(
      { detail: 'tidak ditemukan data' },
      null,
      2,
    ),
  },
  {
    Code: '404',
    Description: 'Halaman melebihi batas',
    Example: JSON.stringify(
      { detail: 'nomor halaman melebihi total halaman. Halaman maksimum adalah 7' },
      null,
      2,
    ),
  },
  {
    Code: '500',
    Description: 'Server Error',
    Example: JSON.stringify(
      { detail: 'Terjadi kesalahan pada server' },
      null,
      2,
    ),
  },
];

export {
  errorHandlingTableRows,
  getKecamatanResponseStructureData,
  kecamatanResponse,
};
