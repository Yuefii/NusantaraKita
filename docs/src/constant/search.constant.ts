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
    Parameter: 'q',
    Type: 'string',
    Description: 'Kata kunci pencarian nama wilayah atau 5 digit kode pos (minimal 2 karakter)',
    Default: '-',
  },
  {
    Parameter: 'tipe',
    Type: 'string',
    Description: 'Filter jenis wilayah: provinsi, kab-kota, kecamatan, desa-kel (opsional)',
    Default: 'all',
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
  {
    Parameter: 'pagination',
    Type: 'boolean',
    Description: 'Menampilkan/menyembunyikan informasi paginasi (true/false)',
    Default: 'true',
  },
];

export const errorHandlingTableRows = [
  {
    Code: '400',
    Description: 'Parameter q kurang dari 2 karakter',
    Example: JSON.stringify(
      { error: { code: 'BAD_REQUEST', message: "parameter 'q' minimal 2 karakter" } },
      null,
      2,
    ),
  },
  {
    Code: '404',
    Description: 'Data tidak ditemukan',
    Example: JSON.stringify(
      { error: { code: 'NOT_FOUND', message: 'tidak ditemukan data' } },
      null,
      2,
    ),
  },
  {
    Code: '500',
    Description: 'Server Error',
    Example: JSON.stringify(
      { error: { code: 'INTERNAL_ERROR', message: 'Terjadi kesalahan pada server' } },
      null,
      2,
    ),
  },
];

export const searchResponseStructureData: ResponseJsonListItem[] = [
  {
    name: 'pagination',
    details: '(object) - Hanya muncul jika pagination=true :',
    children: [
      { name: 'total_item', details: '(integer) - Total item yang cocok' },
      { name: 'total_halaman', details: '(integer) - Total halaman' },
      { name: 'halaman_saat_ini', details: '(integer) - Halaman aktif' },
      { name: 'ukuran_halaman', details: '(integer) - Jumlah item per halaman' },
    ],
  },
  {
    name: 'data',
    details: '(array) - Hasil pencarian wilayah:',
    children: [
      { name: 'kode', details: '(string) - Kode wilayah' },
      { name: 'nama', details: '(string) - Nama wilayah' },
      { name: 'tipe', details: '(string) - Tipe wilayah (provinsi, kabupaten_kota, kecamatan, desa_kelurahan)' },
      { name: 'lat', details: '(double) - Koordinat latitude' },
      { name: 'lng', details: '(double) - Koordinat longitude' },
      { name: 'kode_pos', details: '(string, opsional) - Kode pos (untuk desa/kelurahan)' },
      { name: 'hierarki', details: '(string) - Jalur hierarki wilayah (breadcrumb)' },
      { name: 'geojson_url', details: '(string) - URL CDN file GeoJSON' },
    ],
  },
];

export const responseExampleSearch = {
  pagination: {
    total_item: 12,
    total_halaman: 2,
    halaman_saat_ini: 1,
    ukuran_halaman: 10,
  },
  data: [
    {
      kode: '32.73',
      nama: 'Kota Bandung',
      tipe: 'kabupaten_kota',
      lat: -6.917464,
      lng: 107.619123,
      hierarki: 'Kota Bandung, Jawa Barat',
      geojson_url: 'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kabupaten_kota/32.73.geojson',
    },
    {
      kode: '32.04',
      nama: 'Kabupaten Bandung',
      tipe: 'kabupaten_kota',
      lat: -7.025253,
      lng: 107.519760,
      hierarki: 'Kabupaten Bandung, Jawa Barat',
      geojson_url: 'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kabupaten_kota/32.04.geojson',
    },
    {
      kode: '32.73.08',
      nama: 'Bandung Wetan',
      tipe: 'kecamatan',
      lat: -6.903891,
      lng: 107.618645,
      hierarki: 'Bandung Wetan, Kota Bandung, Jawa Barat',
      geojson_url: 'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kecamatan/32.73.08.geojson',
    },
  ],
};
