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
