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
  geojson_url: string; // CDN URL file GeoJSON Desa/Kelurahan
};

type ApiResponse<T> = {
  pagination?: PaginationApi;
  data: T[];
};

type KabKotaApiRes = ApiResponse<KabKotaApi>;
type KecamatanApiRes = ApiResponse<KecamatanApi>;
type DesaKelApiRes = ApiResponse<DesaKelApi>;
type ProvinsiApiRes = ApiResponse<ProvinsiApi>;
