package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

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
	// API V2 Base List Endpoints
	mux.HandleFunc("GET /v2/provinsi", h.handleProvinsi)
	mux.HandleFunc("GET /v2/kab-kota", h.handleKabKota)
	mux.HandleFunc("GET /v2/kecamatan", h.handleKecamatan)
	mux.HandleFunc("GET /v2/desa-kel", h.handleDesaKel)

	// Search & Reverse Geocoding
	mux.HandleFunc("GET /v2/search", h.handleSearch)
	mux.HandleFunc("GET /v2/reverse", h.handleReverseGeocode)

	// Two-segment dispatcher:
	// - Detail: /v2/provinsi/{kode}, /v2/kab-kota/{kode}, /v2/kecamatan/{kode}, /v2/desa-kel/{kode}
	// - Backwards-compatible parent-child: /v2/{kode_provinsi}/kab-kota, /v2/{kode_kab}/kecamatan, /v2/{kode_kec}/desa-kel
	mux.HandleFunc("GET /v2/{part1}/{part2}", h.handleTwoSegmentRoute)

	// Hierarchical RESTful parent-child routes:
	// - /v2/provinsi/{kode}/kab-kota
	// - /v2/kab-kota/{kode}/kecamatan
	// - /v2/kecamatan/{kode}/desa-kel
	mux.HandleFunc("GET /v2/provinsi/{kode}/kab-kota", h.handleKabKotaByProvinsi)
	mux.HandleFunc("GET /v2/kab-kota/{kode}/kecamatan", h.handleKecamatanByKabKota)
	mux.HandleFunc("GET /v2/kecamatan/{kode}/desa-kel", h.handleDesaKelByKecamatan)
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

// Error codes returned in the "error.code" field.
const (
	codeNotFound         = "NOT_FOUND"
	codeBadRequest       = "BAD_REQUEST"
	codeMethodNotAllowed = "METHOD_NOT_ALLOWED"
	codePageOutOfRange   = "PAGE_OUT_OF_RANGE"
	codeInternalError    = "INTERNAL_ERROR"
)

var errMethodNotAllowed = errors.New("metode HTTP tidak diizinkan")

type statusProbeWriter struct {
	header http.Header
	status int
}

func (p *statusProbeWriter) Header() http.Header {
	if p.header == nil {
		p.header = make(http.Header)
	}
	return p.header
}

func (p *statusProbeWriter) WriteHeader(code int) {
	if p.status == 0 {
		p.status = code
	}
}

func (p *statusProbeWriter) Write(b []byte) (int, error) {
	if p.status == 0 {
		p.status = http.StatusOK
	}
	return len(b), nil
}

// WrapMux wraps an *http.ServeMux so unmatched routes (404) and unsupported
// HTTP methods (405) return structured JSON error responses instead of Go's
// default plain-text bodies.
func WrapMux(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h, pattern := mux.Handler(r)
		if pattern != "" {
			mux.ServeHTTP(w, r)
			return
		}

		probe := &statusProbeWriter{}
		h.ServeHTTP(probe, r)

		switch probe.status {
		case http.StatusMethodNotAllowed:
			if allow := probe.Header().Get("Allow"); allow != "" {
				w.Header().Set("Allow", allow)
			}
			writeError(w, http.StatusMethodNotAllowed, errMethodNotAllowed, codeMethodNotAllowed, "Metode HTTP tidak diizinkan")
		case http.StatusNotFound:
			writeError(w, http.StatusNotFound, service.ErrNotFound, codeNotFound, "Endpoint tidak ditemukan")
		default:
			for k, vals := range probe.Header() {
				for _, v := range vals {
					w.Header().Add(k, v)
				}
			}
			if probe.status != 0 {
				w.WriteHeader(probe.status)
			}
		}
	})
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type errorResponse struct {
	Error errorBody `json:"error"`
}

func writeError(w http.ResponseWriter, status int, err error, code, message string) {
	slog.Error("API request failed", "error", err, "status", status, "code", code, "message", message)
	writeJSON(w, status, errorResponse{Error: errorBody{Code: code, Message: message}})
}

// handleWithService is a DRY helper to parse params, call service, and handle standard errors/responses
func (h *Handler) handleWithService(w http.ResponseWriter, r *http.Request, svcCall func(ctx context.Context, p queryParams) (any, error)) {
	params := parseQueryParams(r)

	resp, err := svcCall(r.Context(), params)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}

		var pageErr service.ErrPageExceeded
		if errors.As(err, &pageErr) {
			writeError(w, http.StatusNotFound, err, codePageOutOfRange, err.Error())
			return
		}

		writeError(w, http.StatusInternalServerError, err, codeInternalError, "Terjadi kesalahan pada server")
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Dispatcher for Two-Segment Routes ----------------
func (h *Handler) handleTwoSegmentRoute(w http.ResponseWriter, r *http.Request) {
	part1 := r.PathValue("part1")
	part2 := r.PathValue("part2")

	// Check if part1 is a known resource type (Detail request: /v2/{entity}/{kode})
	switch part1 {
	case "provinsi":
		h.handleProvinsiByKode(w, r, part2)
		return
	case "kab-kota":
		h.handleKabKotaByKode(w, r, part2)
		return
	case "kecamatan":
		h.handleKecamatanByKode(w, r, part2)
		return
	case "desa-kel":
		h.handleDesaKelByKode(w, r, part2)
		return
	}

	// Check if part2 is a child resource type (Parent-child filter: /v2/{parent_kode}/{child_entity})
	switch part2 {
	case "kab-kota":
		h.handleKabKotaByProvinsiWithKode(w, r, part1)
		return
	case "kecamatan":
		h.handleKecamatanByKabKotaWithKode(w, r, part1)
		return
	case "desa-kel":
		h.handleDesaKelByKecamatanWithKode(w, r, part1)
		return
	}

	writeError(w, http.StatusNotFound, service.ErrNotFound, codeNotFound, "Endpoint tidak ditemukan")
}

// ---------------- Provinsi ----------------
func (h *Handler) handleProvinsi(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetProvinsi(ctx, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleProvinsiByKode(w http.ResponseWriter, r *http.Request, kode string) {
	resp, err := h.svc.GetProvinsiByKode(r.Context(), kode)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, err, codeInternalError, "Terjadi kesalahan pada server")
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Kabupaten / Kota ----------------
func (h *Handler) handleKabKota(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKabKota(ctx, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleKabKotaByKode(w http.ResponseWriter, r *http.Request, kode string) {
	resp, err := h.svc.GetKabKotaByKode(r.Context(), kode)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, err, codeInternalError, "Terjadi kesalahan pada server")
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) handleKabKotaByProvinsiWithKode(w http.ResponseWriter, r *http.Request, kodeProv string) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKabKotaByProvinsi(ctx, kodeProv, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleKabKotaByProvinsi(w http.ResponseWriter, r *http.Request) {
	kodeProv := r.PathValue("kode")
	if kodeProv == "" {
		kodeProv = r.PathValue("kode_provinsi")
	}
	h.handleKabKotaByProvinsiWithKode(w, r, kodeProv)
}

// ---------------- Kecamatan ----------------
func (h *Handler) handleKecamatan(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKecamatan(ctx, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleKecamatanByKode(w http.ResponseWriter, r *http.Request, kode string) {
	resp, err := h.svc.GetKecamatanByKode(r.Context(), kode)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, err, codeInternalError, "Terjadi kesalahan pada server")
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) handleKecamatanByKabKotaWithKode(w http.ResponseWriter, r *http.Request, kodeKab string) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetKecamatanByKabKota(ctx, kodeKab, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleKecamatanByKabKota(w http.ResponseWriter, r *http.Request) {
	kodeKab := r.PathValue("kode")
	if kodeKab == "" {
		kodeKab = r.PathValue("kode_kabupaten_kota")
	}
	h.handleKecamatanByKabKotaWithKode(w, r, kodeKab)
}

// ---------------- Desa / Kelurahan ----------------
func (h *Handler) handleDesaKel(w http.ResponseWriter, r *http.Request) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetDesaKelurahan(ctx, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleDesaKelByKode(w http.ResponseWriter, r *http.Request, kode string) {
	resp, err := h.svc.GetDesaKelurahanByKode(r.Context(), kode)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, err, codeInternalError, "Terjadi kesalahan pada server")
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) handleDesaKelByKecamatanWithKode(w http.ResponseWriter, r *http.Request, kodeKec string) {
	h.handleWithService(w, r, func(ctx context.Context, p queryParams) (any, error) {
		return h.svc.GetDesaKelurahanByKecamatan(ctx, kodeKec, p.limit, p.halaman, p.pagination)
	})
}

func (h *Handler) handleDesaKelByKecamatan(w http.ResponseWriter, r *http.Request) {
	kodeKec := r.PathValue("kode")
	if kodeKec == "" {
		kodeKec = r.PathValue("kode_kecamatan")
	}
	h.handleDesaKelByKecamatanWithKode(w, r, kodeKec)
}

// ---------------- Search / Autocomplete ----------------
func (h *Handler) handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	tipe := r.URL.Query().Get("tipe")
	params := parseQueryParams(r)

	resp, err := h.svc.SearchWilayah(r.Context(), q, tipe, params.limit, params.halaman, params.pagination)
	if err != nil {
		if errors.Is(err, service.ErrQueryTooShort) {
			writeError(w, http.StatusBadRequest, err, codeBadRequest, err.Error())
			return
		}

		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}

		var pageErr service.ErrPageExceeded
		if errors.As(err, &pageErr) {
			writeError(w, http.StatusNotFound, err, codePageOutOfRange, err.Error())
			return
		}

		writeError(w, http.StatusInternalServerError, err, codeInternalError, "Terjadi kesalahan pada server")
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// ---------------- Reverse Geocoding ----------------
func (h *Handler) handleReverseGeocode(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")

	if latStr == "" || lngStr == "" {
		writeError(w, http.StatusBadRequest, service.ErrBadRequest, codeBadRequest, "parameter 'lat' dan 'lng' wajib diisi")
		return
	}

	lat, errLat := strconv.ParseFloat(latStr, 64)
	lng, errLng := strconv.ParseFloat(lngStr, 64)
	if errLat != nil || errLng != nil || lat < -90 || lat > 90 || lng < -180 || lng > 180 {
		writeError(w, http.StatusBadRequest, service.ErrBadRequest, codeBadRequest, "koordinat 'lat' atau 'lng' tidak valid")
		return
	}

	resp, err := h.svc.ReverseGeocode(r.Context(), lat, lng)
	if err != nil {
		if errors.Is(err, service.ErrOutsideIndonesia) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, err, codeNotFound, err.Error())
			return
		}

		writeError(w, http.StatusInternalServerError, err, codeInternalError, "Terjadi kesalahan pada server")
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
