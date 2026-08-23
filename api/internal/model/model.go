package model

type PaginationMeta struct {
	TotalItem      int `json:"total_item"`
	TotalHalaman   int `json:"total_halaman"`
	HalamanSaatIni int `json:"halaman_saat_ini"`
	UkuranHalaman  int `json:"ukuran_halaman"`
}

type Provinsi struct {
	Kode       string  `json:"kode"`
	Nama       string  `json:"nama"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	GeojsonURL string  `json:"geojson_url"`
}

type PaginatedProvinsiResponse struct {
	Pagination *PaginationMeta `json:"pagination,omitempty"`
	Data       []Provinsi      `json:"data"`
}

type KabupatenKota struct {
	Kode         string  `json:"kode"`
	Nama         string  `json:"nama"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	KodeProvinsi string  `json:"kode_provinsi"`
	GeojsonURL   string  `json:"geojson_url"`
}

type PaginatedKabupatenKotaResponse struct {
	Pagination *PaginationMeta `json:"pagination,omitempty"`
	Data       []KabupatenKota `json:"data"`
}

type Kecamatan struct {
	Kode              string  `json:"kode"`
	Nama              string  `json:"nama"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
	KodeKabupatenKota string  `json:"kode_kabupaten_kota"`
	GeojsonURL        string  `json:"geojson_url"`
}

type PaginatedKecamatanResponse struct {
	Pagination *PaginationMeta `json:"pagination,omitempty"`
	Data       []Kecamatan     `json:"data"`
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

type PaginatedDesaKelurahanResponse struct {
	Pagination *PaginationMeta `json:"pagination,omitempty"`
	Data       []DesaKelurahan `json:"data"`
}
