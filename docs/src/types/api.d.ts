type ParamsApi = {
  pagination?: boolean; // Menampilkan/menyembunyikan informasi paginasi (true/false)
  limit?: number; // Menentukan jumlah item per halaman
  halaman?: number; // Menentukan halaman yang ingin ditampilkan
};

type PaginationApi = {
  total_item: number; // (integer) - Total seluruh kecamatan
  total_halaman: number; // (integer) - Total halaman berdasarkan limit
  halaman_saat_ini: number; // (integer) - Halaman yang sedang diakses
  ukuran_halaman: number; // (integer) - Jumlah item per halaman
};

type ProvinsiApi = {
  kode: string; // (string) - Kode unik provinsi
  nama: string; // (string) - Nama provinsi
  lat: number; // (double) - Koordinat latitude
  lng: number; // (double) - Koordinat longitude
  geojson_url: string; // (string) - CDN URL file GeoJSON provinsi
};

type KabKotaApi = {
  kode: string; // Kode unik kabupaten/kota
  nama: string; // Nama kabupaten/kota
  kode_provinsi: string; // Kode provinsi induk
  lat: number; // Koordinat latitude
  lng: number; // Koordinat longitude
  geojson_url: string; // CDN URL file GeoJSON kabupaten/kota
};

type KecamatanApi = {
  kode: string; // Kode unik kecamatan
  nama: string; // Nama kecamatan
  kode_kabupaten_kota: string; // Kode kabupaten/kota induk
  lat: number; // Koordinat latitude
  lng: number; // Koordinat longitude
  geojson_url: string; // CDN URL file GeoJSON kecamatan
};

type DesaKelApi = {
  kode: string; // Kode unik Desa/Kelurahan
  nama: string; // Nama Desa/Kelurahan
  lat: number; // Koordinat latitude
  lng: number; // Koordinat longitude
  kode_kecamatan: string; // Kode Kecamatan induk
  kode_pos: string; // Kode pos Desa/Kelurahan
  geojson_url: string; // CDN URL file GeoJSON Desa/Kelurahan
};

type ApiResponse<T> = {
  pagination?: PaginationApi;
  data: T[];
};

type SingleApiResponse<T> = {
  data: T;
};

type SearchItemApi = {
  kode: string;
  nama: string;
  tipe: 'provinsi' | 'kabupaten_kota' | 'kecamatan' | 'desa_kelurahan';
  lat: number;
  lng: number;
  kode_pos?: string;
  hierarki?: string;
  geojson_url: string;
};

type RegionInfoApi = {
  kode: string;
  nama: string;
  geojson_url?: string;
};

type DesaKelurahanInfoApi = {
  kode: string;
  nama: string;
  kode_pos?: string;
  geojson_url?: string;
};

type CoordinateInfoApi = {
  lat: number;
  lng: number;
};

type ReverseGeocodeApi = {
  provinsi: RegionInfoApi;
  kabupaten_kota: RegionInfoApi;
  kecamatan: RegionInfoApi;
  desa_kelurahan: DesaKelurahanInfoApi;
  koordinat: CoordinateInfoApi;
  jarak_meter: number;
};

type KabKotaApiRes = ApiResponse<KabKotaApi>;
type KecamatanApiRes = ApiResponse<KecamatanApi>;
type DesaKelApiRes = ApiResponse<DesaKelApi>;
type ProvinsiApiRes = ApiResponse<ProvinsiApi>;
type SearchApiRes = ApiResponse<SearchItemApi>;
type ReverseGeocodeApiRes = SingleApiResponse<ReverseGeocodeApi>;
