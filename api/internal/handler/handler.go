package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"nusantarakita/internal/service"
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

func parseQueryParams(r *http.Request) queryParams {
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

	return queryParams{limit: limit, halaman: halaman, pagination: pagination}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, err error, message string) {
	slog.Error("API request failed", "error", err, "status", status, "message", message)
	writeJSON(w, status, map[string]string{"detail": message})
}

// handleWithService is a DRY helper to parse params, call service, and handle standard errors/responses
func (h *Handler) handleWithService(w http.ResponseWriter, r *http.Request, svcCall func(ctx context.Context, p queryParams) (any, error)) {
	params := parseQueryParams(r)

	resp, err := svcCall(r.Context(), params)
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

// ---------------- Provinsi ----------------
func (h *Handler) handleProvinsi(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetProvinsi(ctx, p.limit, p.halaman, p.pagination)
	})
}

// ---------------- Kabupaten / Kota ----------------
func (h *Handler) handleKabKota(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKabKota(ctx, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleKabKotaByProvinsi(w http.ResponseWriter, r *http.Request) {
	kodeProv := r.PathValue("kode_provinsi")
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKabKotaByProvinsi(ctx, kodeProv, p.limit, p.halaman, p.pagination)
	})
}

// ---------------- Kecamatan ----------------
func (h *Handler) handleKecamatan(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKecamatan(ctx, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleKecamatanByKabKota(w http.ResponseWriter, r *http.Request) {
	kodeKab := r.PathValue("kode_kabupaten_kota")
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKecamatanByKabKota(ctx, kodeKab, p.limit, p.halaman, p.pagination)
	})
}

// ---------------- Desa / Kelurahan ----------------
func (h *Handler) handleDesaKel(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetDesaKelurahan(ctx, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleDesaKelByKecamatan(w http.ResponseWriter, r *http.Request) {
	kodeKec := r.PathValue("kode_kecamatan")
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetDesaKelurahanByKecamatan(ctx, kodeKec, p.limit, p.halaman, p.pagination)
	})
}
