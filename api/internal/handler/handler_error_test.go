package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"nusantarakita/internal/repository"
	"nusantarakita/internal/service"
)

const (
	provCountQuery = "^SELECT COUNT\\(\\*\\) FROM nk_provinsi$"
	provDataQuery  = "^SELECT kode, nama, lat, lng FROM nk_provinsi LIMIT \\$1 OFFSET \\$2$"
	provAllQuery   = "^SELECT kode, nama, lat, lng FROM nk_provinsi$"
	kabCountQuery  = "^SELECT COUNT\\(\\*\\) FROM nk_kabupaten_kota WHERE kode_provinsi = \\$1$"
	kabDataQuery   = "^SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode_provinsi = \\$1 LIMIT \\$2 OFFSET \\$3$"
)

func newTestMux(t *testing.T) (*http.ServeMux, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	t.Cleanup(func() { db.Close() })
	mock.MatchExpectationsInOrder(false)

	mux := http.NewServeMux()
	New(service.NewDaerahService(repository.NewDaerahRepository(db))).RegisterRoutes(mux)
	return mux, mock
}

func provRows(n int) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"})
	for i := 0; i < n; i++ {
		rows.AddRow("3"+string(rune('0'+i)), "PROVINSI", -6.0, 106.0)
	}
	return rows
}

func expectProv(mock sqlmock.Sqlmock, total, limit, offset, returned int) {
	mock.ExpectQuery(provCountQuery).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	mock.ExpectQuery(provDataQuery).WithArgs(limit, offset).WillReturnRows(provRows(returned))
}

func serve(mux *http.ServeMux, target string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, target, nil))
	return w
}

// assertError checks the status and that the body is exactly {"error":{"code","message"}}.
func assertError(t *testing.T, w *httptest.ResponseRecorder, status int, code, messageContains string) {
	t.Helper()
	if w.Code != status {
		t.Fatalf("expected status %d, got %d (body %s)", status, w.Code, w.Body.String())
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(w.Body.Bytes(), &raw); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if len(raw) != 1 {
		t.Errorf("expected only the \"error\" key, got %s", w.Body.String())
	}
	if _, ok := raw["detail"]; ok {
		t.Errorf("old \"detail\" key still present: %s", w.Body.String())
	}

	var body struct {
		Error map[string]string `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid error body: %v", err)
	}
	if len(body.Error) != 2 {
		t.Errorf("expected error to hold only code and message, got %v", body.Error)
	}
	if body.Error["code"] != code {
		t.Errorf("expected code %q, got %q", code, body.Error["code"])
	}
	if !strings.Contains(body.Error["message"], messageContains) {
		t.Errorf("expected message to contain %q, got %q", messageContains, body.Error["message"])
	}
}

func TestErrorResponses(t *testing.T) {
	t.Run("page past the last page returns PAGE_OUT_OF_RANGE", func(t *testing.T) {
		mux, mock := newTestMux(t)
		expectProv(mock, 38, 10, 980, 0)
		assertError(t, serve(mux, "/v2/provinsi?halaman=99"), http.StatusNotFound,
			"PAGE_OUT_OF_RANGE", "Halaman maksimum adalah 4")
	})

	t.Run("first page past the end when total is an exact multiple", func(t *testing.T) {
		mux, mock := newTestMux(t)
		expectProv(mock, 20, 10, 20, 0)
		assertError(t, serve(mux, "/v2/provinsi?halaman=3"), http.StatusNotFound,
			"PAGE_OUT_OF_RANGE", "Halaman maksimum adalah 2")
	})

	t.Run("last page itself is still OK", func(t *testing.T) {
		mux, mock := newTestMux(t)
		expectProv(mock, 21, 10, 20, 1)
		if w := serve(mux, "/v2/provinsi?halaman=3"); w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d (body %s)", w.Code, w.Body.String())
		}
	})

	t.Run("limit larger than the total gives max page 1", func(t *testing.T) {
		mux, mock := newTestMux(t)
		expectProv(mock, 2, 1000, 1000, 0)
		assertError(t, serve(mux, "/v2/provinsi?limit=1000&halaman=2"), http.StatusNotFound,
			"PAGE_OUT_OF_RANGE", "Halaman maksimum adalah 1")
	})

	t.Run("limit=1 gives one page per item", func(t *testing.T) {
		mux, mock := newTestMux(t)
		expectProv(mock, 5, 1, 5, 0)
		assertError(t, serve(mux, "/v2/provinsi?limit=1&halaman=6"), http.StatusNotFound,
			"PAGE_OUT_OF_RANGE", "Halaman maksimum adalah 5")
	})

	t.Run("empty table returns NOT_FOUND, not PAGE_OUT_OF_RANGE", func(t *testing.T) {
		mux, mock := newTestMux(t)
		expectProv(mock, 0, 10, 40, 0)
		assertError(t, serve(mux, "/v2/provinsi?halaman=5"), http.StatusNotFound,
			"NOT_FOUND", "tidak ditemukan data")
	})

	t.Run("pagination=false ignores halaman", func(t *testing.T) {
		mux, mock := newTestMux(t)
		mock.ExpectQuery(provAllQuery).WillReturnRows(provRows(2))
		if w := serve(mux, "/v2/provinsi?pagination=false&halaman=99"); w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d (body %s)", w.Code, w.Body.String())
		}
	})

	t.Run("pagination=false on an empty table returns NOT_FOUND", func(t *testing.T) {
		mux, mock := newTestMux(t)
		mock.ExpectQuery(provAllQuery).WillReturnRows(provRows(0))
		assertError(t, serve(mux, "/v2/provinsi?pagination=false"), http.StatusNotFound,
			"NOT_FOUND", "tidak ditemukan data")
	})

	t.Run("database error returns INTERNAL_ERROR without leaking the cause", func(t *testing.T) {
		mux, mock := newTestMux(t)
		mock.ExpectQuery(provCountQuery).WillReturnError(errors.New("pq: password authentication failed"))
		mock.ExpectQuery(provDataQuery).WithArgs(10, 0).WillReturnError(errors.New("pq: password authentication failed"))
		w := serve(mux, "/v2/provinsi")
		assertError(t, w, http.StatusInternalServerError, "INTERNAL_ERROR", "Terjadi kesalahan pada server")
		if strings.Contains(w.Body.String(), "pq:") {
			t.Errorf("database error leaked into the response: %s", w.Body.String())
		}
	})

	t.Run("unknown kode_provinsi on a nested route returns NOT_FOUND", func(t *testing.T) {
		mux, mock := newTestMux(t)
		mock.ExpectQuery(kabCountQuery).WithArgs("99").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
		mock.ExpectQuery(kabDataQuery).WithArgs("99", 10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"kode", "nama", "lat", "lng", "kode_provinsi"}))
		assertError(t, serve(mux, "/v2/99/kab-kota"), http.StatusNotFound, "NOT_FOUND", "tidak ditemukan data")
	})

	t.Run("nested route past the last page returns PAGE_OUT_OF_RANGE", func(t *testing.T) {
		mux, mock := newTestMux(t)
		mock.ExpectQuery(kabCountQuery).WithArgs("32").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(27))
		mock.ExpectQuery(kabDataQuery).WithArgs("32", 10, 90).
			WillReturnRows(sqlmock.NewRows([]string{"kode", "nama", "lat", "lng", "kode_provinsi"}))
		assertError(t, serve(mux, "/v2/32/kab-kota?halaman=10"), http.StatusNotFound,
			"PAGE_OUT_OF_RANGE", "Halaman maksimum adalah 3")
	})

	t.Run("invalid halaman values fall back to page 1", func(t *testing.T) {
		for _, h := range []string{"0", "-1", "abc", "99999999999999999999", "1.5"} {
			t.Run(h, func(t *testing.T) {
				mux, mock := newTestMux(t)
				expectProv(mock, 38, 10, 0, 10)
				if w := serve(mux, "/v2/provinsi?halaman="+h); w.Code != http.StatusOK {
					t.Fatalf("halaman=%s: expected 200, got %d (body %s)", h, w.Code, w.Body.String())
				}
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("halaman=%s did not query page 1: %v", h, err)
				}
			})
		}
	})
}
