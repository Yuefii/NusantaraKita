package model

type PaginationMeta struct {
	TotalItem      int `json:"total_item"`
	TotalHalaman   int `json:"total_halaman"`
	HalamanSaatIni int `json:"halaman_saat_ini"`
	UkuranHalaman  int `json:"ukuran_halaman"`
}

// PaginatedResponse is a generic struct for all paginated endpoints
type PaginatedResponse[T any] struct {
	Pagination *PaginationMeta `json:"pagination,omitempty"`
	Data       []T             `json:"data"`
}

// SingleResponse is a generic struct for single item endpoints
type SingleResponse[T any] struct {
	Data T `json:"data"`
}

type Provinsi struct {
	Kode       string  `json:"kode"`
	Nama       string  `json:"nama"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	GeojsonURL string  `json:"geojson_url"`
}

type KabupatenKota struct {
	Kode         string  `json:"kode"`
	Nama         string  `json:"nama"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	KodeProvinsi string  `json:"kode_provinsi"`
	GeojsonURL   string  `json:"geojson_url"`
}

type Kecamatan struct {
	Kode              string  `json:"kode"`
	Nama              string  `json:"nama"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
	KodeKabupatenKota string  `json:"kode_kabupaten_kota"`
	GeojsonURL        string  `json:"geojson_url"`
}

type DesaKelurahan struct {
	Kode          string  `json:"kode"`
	Nama          string  `json:"nama"`
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	KodeKecamatan string  `json:"kode_kecamatan"`
	KodePos       string  `json:"kode_pos"`
	GeojsonURL    string  `json:"geojson_url"`
}

type SearchItem struct {
	Kode       string  `json:"kode"`
	Nama       string  `json:"nama"`
	Tipe       string  `json:"tipe"` // "provinsi", "kabupaten_kota", "kecamatan", "desa_kelurahan"
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	KodePos    string  `json:"kode_pos,omitempty"`
	Hierarki   string  `json:"hierarki,omitempty"`
	GeojsonURL string  `json:"geojson_url"`
}

type RegionInfo struct {
	Kode       string `json:"kode"`
	Nama       string `json:"nama"`
	GeojsonURL string `json:"geojson_url,omitempty"`
}

type DesaKelurahanInfo struct {
	Kode       string `json:"kode"`
	Nama       string `json:"nama"`
	KodePos    string `json:"kode_pos,omitempty"`
	GeojsonURL string `json:"geojson_url,omitempty"`
}

type CoordinateInfo struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ReverseGeocodeResult struct {
	Provinsi      RegionInfo        `json:"provinsi"`
	KabupatenKota RegionInfo        `json:"kabupaten_kota"`
	Kecamatan     RegionInfo        `json:"kecamatan"`
	DesaKelurahan DesaKelurahanInfo `json:"desa_kelurahan"`
	Koordinat     CoordinateInfo    `json:"koordinat"`
	JarakMeter    float64           `json:"jarak_meter"`
}
