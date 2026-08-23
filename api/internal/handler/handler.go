package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"vercel-go-starter/internal/model"
)

const cdnBaseURL = "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson"

type Handler struct {
	db     *sql.DB
}

func New(db *sql.DB) *Handler {
	return &Handler{
		db:     db,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// API V2
	mux.HandleFunc("GET /v2/provinsi", h.handleProvinsi)
	mux.HandleFunc("GET /v2/kab-kota", h.handleKabKota)
	mux.HandleFunc("GET /v2/{kode_provinsi}/kab-kota", h.handleKabKotaByProvinsi)
	mux.HandleFunc("GET /v2/kecamatan", h.handleKecamatan)
	mux.HandleFunc("GET /v2/{kode_kabupaten_kota}/kecamatan", h.handleKecamatanByKabKota)
	mux.HandleFunc("GET /v2/desa-kel", h.handleDesaKel)
	mux.HandleFunc("GET /v2/{kode_kecamatan}/desa-kel", h.handleDesaKelByKecamatan)
}

type queryParams struct {
	limit      int
	halaman    int
	pagination bool
}

func parseQueryParams(r *http.Request) (queryParams, error) {
	q := r.URL.Query()

	limit := 10
	if limitStr := q.Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l >= 1 {
			limit = l
		}
	}

	halaman := 1
	if halStr := q.Get("halaman"); halStr != "" {
		if h, err := strconv.Atoi(halStr); err == nil && h >= 1 {
			halaman = h
		}
	}

	pagination := true
	if pagStr := q.Get("pagination"); pagStr != "" {
		if p, err := strconv.ParseBool(pagStr); err == nil {
			pagination = p
		}
	}

	return queryParams{limit: limit, halaman: halaman, pagination: pagination}, nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"detail": message})
}

// replacePlaceholders replaces ? with $1, $2, etc for PostgreSQL
func replacePlaceholders(query string) string {
	n := 1
	for {
		idx := strings.Index(query, "?")
		if idx == -1 {
			break
		}
		query = query[:idx] + fmt.Sprintf("$%d", n) + query[idx+1:]
		n++
	}
	return query
}

// ---------------- Provinsi ----------------
func (h *Handler) handleProvinsi(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)

	if !params.pagination {
		rows, err := h.db.Query("SELECT kode, nama, lat, lng FROM nk_provinsi")
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var data []model.Provinsi
		for rows.Next() {
			var p model.Provinsi
			if err := rows.Scan(&p.Kode, &p.Nama, &p.Lat, &p.Lng); err != nil {
				writeError(w, http.StatusInternalServerError, err.Error())
				return
			}
			p.GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", cdnBaseURL, p.Kode)
			data = append(data, p)
		}

		if len(data) == 0 {
			writeError(w, http.StatusInternalServerError, "tidak ditemukan data")
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{"data": data})
		return
	}

	var totalItem int
	err := h.db.QueryRow("SELECT COUNT(*) FROM nk_provinsi").Scan(&totalItem)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	totalHalaman := (totalItem + params.limit - 1) / params.limit
	if params.halaman > totalHalaman && totalHalaman > 0 {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("nomor halaman melebihi total halaman. Halaman maksimum adalah %d", totalHalaman))
		return
	}

	offset := (params.halaman - 1) * params.limit
	
	query := replacePlaceholders("SELECT kode, nama, lat, lng FROM nk_provinsi LIMIT ? OFFSET ?")
	rows, err := h.db.Query(query, params.limit, offset)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var data []model.Provinsi
	for rows.Next() {
		var p model.Provinsi
		if err := rows.Scan(&p.Kode, &p.Nama, &p.Lat, &p.Lng); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		p.GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", cdnBaseURL, p.Kode)
		data = append(data, p)
	}

	if len(data) == 0 {
		writeError(w, http.StatusInternalServerError, "tidak ditemukan data untuk halaman yang diminta")
		return
	}

	resp := model.PaginatedProvinsiResponse{
		Pagination: &model.PaginationMeta{
			TotalItem:      totalItem,
			TotalHalaman:   totalHalaman,
			HalamanSaatIni: params.halaman,
			UkuranHalaman:  params.limit,
		},
		Data: data,
	}
	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Kabupaten / Kota ----------------
func (h *Handler) handleKabKota(w http.ResponseWriter, r *http.Request) {
	h.fetchKabKota(w, r, "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota", "SELECT COUNT(*) FROM nk_kabupaten_kota", nil)
}

func (h *Handler) handleKabKotaByProvinsi(w http.ResponseWriter, r *http.Request) {
	kodeProv := r.PathValue("kode_provinsi")
	h.fetchKabKota(w, r, "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode_provinsi = ?", "SELECT COUNT(*) FROM nk_kabupaten_kota WHERE kode_provinsi = ?", []interface{}{kodeProv})
}

func (h *Handler) fetchKabKota(w http.ResponseWriter, r *http.Request, query string, countQuery string, args []interface{}) {
	params, _ := parseQueryParams(r)

	if !params.pagination {
		rows, err := h.db.Query(replacePlaceholders(query), args...)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var data []model.KabupatenKota
		for rows.Next() {
			var k model.KabupatenKota
			if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi); err != nil {
				writeError(w, http.StatusInternalServerError, err.Error())
				return
			}
			k.GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", cdnBaseURL, k.Kode)
			data = append(data, k)
		}

		if len(data) == 0 {
			writeError(w, http.StatusInternalServerError, "tidak ditemukan data")
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{"data": data})
		return
	}

	var totalItem int
	err := h.db.QueryRow(replacePlaceholders(countQuery), args...).Scan(&totalItem)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	totalHalaman := (totalItem + params.limit - 1) / params.limit
	if params.halaman > totalHalaman && totalHalaman > 0 {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("nomor halaman melebihi total halaman. Halaman maksimum adalah %d", totalHalaman))
		return
	}

	offset := (params.halaman - 1) * params.limit
	
	finalArgs := append(args, params.limit, offset)
	rows, err := h.db.Query(replacePlaceholders(query+" LIMIT ? OFFSET ?"), finalArgs...)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var data []model.KabupatenKota
	for rows.Next() {
		var k model.KabupatenKota
		if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		k.GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", cdnBaseURL, k.Kode)
		data = append(data, k)
	}

	if len(data) == 0 {
		writeError(w, http.StatusInternalServerError, "tidak ditemukan data untuk halaman yang diminta")
		return
	}

	resp := model.PaginatedKabupatenKotaResponse{
		Pagination: &model.PaginationMeta{
			TotalItem:      totalItem,
			TotalHalaman:   totalHalaman,
			HalamanSaatIni: params.halaman,
			UkuranHalaman:  params.limit,
		},
		Data: data,
	}
	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Kecamatan ----------------
func (h *Handler) handleKecamatan(w http.ResponseWriter, r *http.Request) {
	h.fetchKecamatan(w, r, "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan", "SELECT COUNT(*) FROM nk_kecamatan", nil)
}

func (h *Handler) handleKecamatanByKabKota(w http.ResponseWriter, r *http.Request) {
	kodeKab := r.PathValue("kode_kabupaten_kota")
	h.fetchKecamatan(w, r, "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan WHERE kode_kabupaten_kota = ?", "SELECT COUNT(*) FROM nk_kecamatan WHERE kode_kabupaten_kota = ?", []interface{}{kodeKab})
}

func (h *Handler) fetchKecamatan(w http.ResponseWriter, r *http.Request, query string, countQuery string, args []interface{}) {
	params, _ := parseQueryParams(r)

	if !params.pagination {
		rows, err := h.db.Query(replacePlaceholders(query), args...)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var data []model.Kecamatan
		for rows.Next() {
			var k model.Kecamatan
			if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota); err != nil {
				writeError(w, http.StatusInternalServerError, err.Error())
				return
			}
			k.GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", cdnBaseURL, k.Kode)
			data = append(data, k)
		}

		if len(data) == 0 {
			writeError(w, http.StatusInternalServerError, "tidak ditemukan data")
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{"data": data})
		return
	}

	var totalItem int
	err := h.db.QueryRow(replacePlaceholders(countQuery), args...).Scan(&totalItem)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	totalHalaman := (totalItem + params.limit - 1) / params.limit
	if params.halaman > totalHalaman && totalHalaman > 0 {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("nomor halaman melebihi total halaman. Halaman maksimum adalah %d", totalHalaman))
		return
	}

	offset := (params.halaman - 1) * params.limit
	
	finalArgs := append(args, params.limit, offset)
	rows, err := h.db.Query(replacePlaceholders(query+" LIMIT ? OFFSET ?"), finalArgs...)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var data []model.Kecamatan
	for rows.Next() {
		var k model.Kecamatan
		if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		k.GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", cdnBaseURL, k.Kode)
		data = append(data, k)
	}

	if len(data) == 0 {
		writeError(w, http.StatusInternalServerError, "tidak ditemukan data untuk halaman yang diminta")
		return
	}

	resp := model.PaginatedKecamatanResponse{
		Pagination: &model.PaginationMeta{
			TotalItem:      totalItem,
			TotalHalaman:   totalHalaman,
			HalamanSaatIni: params.halaman,
			UkuranHalaman:  params.limit,
		},
		Data: data,
	}
	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Desa / Kelurahan ----------------
func (h *Handler) handleDesaKel(w http.ResponseWriter, r *http.Request) {
	h.fetchDesaKel(w, r, "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan", "SELECT COUNT(*) FROM nk_desa_kelurahan", nil)
}

func (h *Handler) handleDesaKelByKecamatan(w http.ResponseWriter, r *http.Request) {
	kodeKec := r.PathValue("kode_kecamatan")
	h.fetchDesaKel(w, r, "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan WHERE kode_kecamatan = ?", "SELECT COUNT(*) FROM nk_desa_kelurahan WHERE kode_kecamatan = ?", []interface{}{kodeKec})
}

func (h *Handler) fetchDesaKel(w http.ResponseWriter, r *http.Request, query string, countQuery string, args []interface{}) {
	params, _ := parseQueryParams(r)

	if !params.pagination {
		rows, err := h.db.Query(replacePlaceholders(query), args...)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var data []model.DesaKelurahan
		for rows.Next() {
			var d model.DesaKelurahan
			if err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos); err != nil {
				writeError(w, http.StatusInternalServerError, err.Error())
				return
			}
			d.GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", cdnBaseURL, d.Kode)
			data = append(data, d)
		}

		if len(data) == 0 {
			writeError(w, http.StatusInternalServerError, "tidak ditemukan data")
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{"data": data})
		return
	}

	var totalItem int
	err := h.db.QueryRow(replacePlaceholders(countQuery), args...).Scan(&totalItem)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	totalHalaman := (totalItem + params.limit - 1) / params.limit
	if params.halaman > totalHalaman && totalHalaman > 0 {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("nomor halaman melebihi total halaman. Halaman maksimum adalah %d", totalHalaman))
		return
	}

	offset := (params.halaman - 1) * params.limit
	
	finalArgs := append(args, params.limit, offset)
	rows, err := h.db.Query(replacePlaceholders(query+" LIMIT ? OFFSET ?"), finalArgs...)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var data []model.DesaKelurahan
	for rows.Next() {
		var d model.DesaKelurahan
		if err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		d.GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", cdnBaseURL, d.Kode)
		data = append(data, d)
	}

	if len(data) == 0 {
		writeError(w, http.StatusInternalServerError, "tidak ditemukan data untuk halaman yang diminta")
		return
	}

	resp := model.PaginatedDesaKelurahanResponse{
		Pagination: &model.PaginationMeta{
			TotalItem:      totalItem,
			TotalHalaman:   totalHalaman,
			HalamanSaatIni: params.halaman,
			UkuranHalaman:  params.limit,
		},
		Data: data,
	}
	writeJSON(w, http.StatusOK, resp)
}
