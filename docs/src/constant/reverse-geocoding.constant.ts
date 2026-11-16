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
    Parameter: 'lat',
    Type: 'number',
    Description:
      'Koordinat latitude (-90 s/d 90, wilayah Indonesia sekitar -11 s/d 6)',
    Default: 'wajib',
  },
  {
    Parameter: 'lng',
    Type: 'number',
    Description:
      'Koordinat longitude (-180 s/d 180, wilayah Indonesia sekitar 95 s/d 141)',
    Default: 'wajib',
  },
];

export const errorHandlingTableRows = [
  {
    Code: '400',
    Description: 'Parameter lat atau lng tidak valid / tidak diisi',
    Example: JSON.stringify(
      {
        error: {
          code: 'BAD_REQUEST',
          message: "parameter 'lat' dan 'lng' wajib diisi",
        },
      },
      null,
      2,
    ),
  },
  {
    Code: '404',
    Description: 'Koordinat berada di luar wilayah Indonesia',
    Example: JSON.stringify(
      {
        error: {
          code: 'NOT_FOUND',
          message: 'koordinat berada di luar wilayah Indonesia',
        },
      },
      null,
      2,
    ),
  },
  {
    Code: '500',
    Description: 'Server Error',
    Example: JSON.stringify(
      {
        error: {
          code: 'INTERNAL_ERROR',
          message: 'Terjadi kesalahan pada server',
        },
      },
      null,
      2,
    ),
  },
];

export const reverseResponseStructureData: ResponseJsonListItem[] = [
  {
    name: 'data',
    details: '(object) - Hasil reverse geocoding:',
    children: [
      {
        name: 'provinsi',
        details: '(object) - Informasi provinsi:',
        children: [
          { name: 'kode', details: '(string) - Kode provinsi' },
          { name: 'nama', details: '(string) - Nama provinsi' },
          {
            name: 'geojson_url',
            details: '(string) - URL CDN GeoJSON provinsi',
          },
        ],
      },
      {
        name: 'kabupaten_kota',
        details: '(object) - Informasi kabupaten/kota:',
        children: [
          { name: 'kode', details: '(string) - Kode kabupaten/kota' },
          { name: 'nama', details: '(string) - Nama kabupaten/kota' },
          {
            name: 'geojson_url',
            details: '(string) - URL CDN GeoJSON kabupaten/kota',
          },
        ],
      },
      {
        name: 'kecamatan',
        details: '(object) - Informasi kecamatan:',
        children: [
          { name: 'kode', details: '(string) - Kode kecamatan' },
          { name: 'nama', details: '(string) - Nama kecamatan' },
          {
            name: 'geojson_url',
            details: '(string) - URL CDN GeoJSON kecamatan',
          },
        ],
      },
      {
        name: 'desa_kelurahan',
        details: '(object) - Informasi desa/kelurahan terdekat:',
        children: [
          { name: 'kode', details: '(string) - Kode desa/kelurahan' },
          { name: 'nama', details: '(string) - Nama desa/kelurahan' },
          { name: 'kode_pos', details: '(string) - Kode pos' },
          {
            name: 'geojson_url',
            details: '(string) - URL CDN GeoJSON desa/kelurahan',
          },
        ],
      },
      {
        name: 'koordinat',
        details: '(object) - Koordinat titik input pengguna:',
        children: [
          { name: 'lat', details: '(double) - Latitude' },
          { name: 'lng', details: '(double) - Longitude' },
        ],
      },
      {
        name: 'jarak_meter',
        details:
          '(double) - Jarak titik input ke titik pusat desa/kelurahan (dalam meter)',
      },
    ],
  },
];

export const responseExampleReverse = {
  data: {
    provinsi: {
      kode: '32',
      nama: 'Jawa Barat',
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/provinsi/32.geojson',
    },
    kabupaten_kota: {
      kode: '32.73',
      nama: 'Kota Bandung',
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kabupaten_kota/32.73.geojson',
    },
    kecamatan: {
      kode: '32.73.01',
      nama: 'Sukasari',
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/kecamatan/32.73.01.geojson',
    },
    desa_kelurahan: {
      kode: '32.73.01.1001',
      nama: 'Sarijadi',
      kode_pos: '40151',
      geojson_url:
        'https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson/desa_kelurahan/32.73.01.1001.geojson',
    },
    koordinat: {
      lat: -6.874112,
      lng: 107.589123,
    },
    jarak_meter: 124.52,
  },
};
