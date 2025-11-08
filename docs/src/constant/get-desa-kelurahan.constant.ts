import type { ResponseJsonListItem } from '@/components/ui/response-json-structure';

const getDesaKelurahanResponseStructureData: ResponseJsonListItem[] = [
  {
    name: 'pagination',
    details: '(object) - Hanya muncul jika pagination=true:',
    children: [
      {
        name: 'total_item',
        details: '(integer) - Total seluruh Desa/Kelurahan',
      },
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
    details: '(array) - Daftar Desa/Kelurahan:',
    children: [
      { name: 'kode', details: '(string) - Kode unik Desa/Kelurahan' },
      { name: 'nama', details: '(string) - Nama Desa/Kelurahan' },
      { name: 'lat', details: '(double) - Koordinat latitude' },
      { name: 'lng', details: '(double) - Koordinat longitude' },
      { name: 'kode_kecamatan', details: '(string) - Kode Kecamatan induk' },
      { name: 'geojson_url', details: '(string) - URL GeoJSON Desa/Kelurahan' },
    ],
  },
];

const getDesaKelurahanExampleResponse = {
  pagination: {
    total_item: 516,
    total_halaman: 52,
    halaman_saat_ini: 1,
    ukuran_halaman: 10,
  },
  data: [
    {
      kode: '11.01.01.2001',
      nama: 'Keude Bakongan',
      lat: 2.931094803160483,
      lng: 97.48458404258515,
      kode_kecamatan: '11.01.01',
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/desa_kelurahan/11.01.01.2001.geojson',
    },
  ],
};

const errorHandlingTableRows = [
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
    Example: JSON.stringify(
      { error: 'tidak ditemukan data desa/kelurahan' },
      null,
      2,
    ),
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

export {
  errorHandlingTableRows,
  getDesaKelurahanExampleResponse,
  getDesaKelurahanResponseStructureData,
};
