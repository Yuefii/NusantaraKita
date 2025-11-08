import type { ResponseJsonListItem } from '@/components/ui/response-json-structure';

const getKabKotaResponseStructureData: ResponseJsonListItem[] = [
  {
    name: 'pagination',
    details: '(object) - Hanya muncul jika pagination=true :',
    children: [
      {
        name: 'total_item',
        details: '(integer) - Total seluruh kabupaten/kota',
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
    details: '(array) - Daftar kabupaten/kota:',
    children: [
      { name: 'kode', details: '(string) - Kode unik kabupaten/kota' },
      { name: 'nama', details: '(string) - Nama kabupaten/kota' },
      { name: 'kode_provinsi', details: '(string) - Kode provinsi induk' },
      { name: 'lat', details: '(double) - Koordinat latitude' },
      { name: 'lng', details: '(double) - Koordinat longitude' },
      { name: 'geojson_url', details: '(string) - URL GeoJSON kabupaten/kota' },
    ],
  },
];

const kabKotaResponse = {
  pagination: {
    total_item: 516,
    total_halaman: 52,
    halaman_saat_ini: 1,
    ukuran_halaman: 10,
  },
  data: [
    {
      kode: '11.01',
      nama: 'Aceh Selatan',
      lat: 3.161853840894135,
      lng: 97.43651771865193,
      kode_provinsi: '11',
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kabupaten_kota/11.01.geojson',
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
      { error: 'tidak ditemukan data kabupaten/kota' },
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
  getKabKotaResponseStructureData,
  kabKotaResponse,
};
