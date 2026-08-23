package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"vercel-go-starter/internal/service"
)

type Handler struct {
	svc *service.DaerahService
}

func New(svc *service.DaerahService) *Handler {
	return &Handler{
		svc: svc,
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

func writeError(w http.ResponseWriter, status int, err error, message string) {
	log.Printf("[ERROR] %v", err)
	writeJSON(w, status, map[string]string{"detail": message})
}

// ---------------- Provinsi ----------------
func (h *Handler) handleProvinsi(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)

	resp, err := h.svc.GetProvinsi(r.Context(), params.limit, params.halaman, params.pagination)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "nomor halaman melebihi") {
			writeError(w, http.StatusNotFound, err, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err, "Terjadi kesalahan pada server")
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Kabupaten / Kota ----------------
func (h *Handler) handleKabKota(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)

	resp, err := h.svc.GetKabKota(r.Context(), params.limit, params.halaman, params.pagination)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "nomor halaman melebihi") {
			writeError(w, http.StatusNotFound, err, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err, "Terjadi kesalahan pada server")
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) handleKabKotaByProvinsi(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)
	kodeProv := r.PathValue("kode_provinsi")

	resp, err := h.svc.GetKabKotaByProvinsi(r.Context(), kodeProv, params.limit, params.halaman, params.pagination)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "nomor halaman melebihi") {
			writeError(w, http.StatusNotFound, err, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err, "Terjadi kesalahan pada server")
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Kecamatan ----------------
func (h *Handler) handleKecamatan(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)

	resp, err := h.svc.GetKecamatan(r.Context(), params.limit, params.halaman, params.pagination)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "nomor halaman melebihi") {
			writeError(w, http.StatusNotFound, err, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err, "Terjadi kesalahan pada server")
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) handleKecamatanByKabKota(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)
	kodeKab := r.PathValue("kode_kabupaten_kota")

	resp, err := h.svc.GetKecamatanByKabKota(r.Context(), kodeKab, params.limit, params.halaman, params.pagination)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "nomor halaman melebihi") {
			writeError(w, http.StatusNotFound, err, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err, "Terjadi kesalahan pada server")
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Desa / Kelurahan ----------------
func (h *Handler) handleDesaKel(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)

	resp, err := h.svc.GetDesaKelurahan(r.Context(), params.limit, params.halaman, params.pagination)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "nomor halaman melebihi") {
			writeError(w, http.StatusNotFound, err, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err, "Terjadi kesalahan pada server")
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) handleDesaKelByKecamatan(w http.ResponseWriter, r *http.Request) {
	params, _ := parseQueryParams(r)
	kodeKec := r.PathValue("kode_kecamatan")

	resp, err := h.svc.GetDesaKelurahanByKecamatan(r.Context(), kodeKec, params.limit, params.halaman, params.pagination)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") || strings.Contains(err.Error(), "nomor halaman melebihi") {
			writeError(w, http.StatusNotFound, err, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err, "Terjadi kesalahan pada server")
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
