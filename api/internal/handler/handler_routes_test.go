package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"nusantarakita/internal/repository"
	"nusantarakita/internal/service"
)

func TestHandleProvinsi(t *testing.T) {
	t.Run("GET /v2/provinsi returns paginated JSON", func(t *testing.T) {
		// 1. Setup Mock DB
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to open sqlmock database: %s", err)
		}
		defer db.Close()
		mock.MatchExpectationsInOrder(false)

		// 2. Setup Dependency Injection
		repo := repository.NewDaerahRepository(db)
		svc := service.NewDaerahService(repo)
		h := New(svc)

		mux := http.NewServeMux()
		h.RegisterRoutes(mux)
		// Mock the count and data query
		countRows := sqlmock.NewRows([]string{"count"}).AddRow(2)
		mock.ExpectQuery("^SELECT COUNT\\(\\*\\) FROM nk_provinsi$").WillReturnRows(countRows)

		dataRows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("31", "DKI JAKARTA", -6.175, 106.827).
			AddRow("32", "JAWA BARAT", -6.920, 107.604)
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi LIMIT \\$1 OFFSET \\$2$").
			WithArgs(10, 0).
			WillReturnRows(dataRows)

		// 3. Create HTTP Request
		req := httptest.NewRequest(http.MethodGet, "/v2/provinsi?limit=10&halaman=1", nil)
		w := httptest.NewRecorder()

		// 4. Serve the request
		mux.ServeHTTP(w, req)

		// 5. Assertions
		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}

		// Check pagination meta
		if meta, ok := response["pagination"].(map[string]interface{}); ok {
			if int(meta["total_item"].(float64)) != 2 {
				t.Errorf("expected 2 total items, got %v", meta["total_item"])
			}
		} else {
			t.Error("expected pagination meta in response")
		}

		// Check data
		if data, ok := response["data"].([]interface{}); ok {
			if len(data) != 2 {
				t.Errorf("expected 2 data items, got %d", len(data))
			}
			firstItem := data[0].(map[string]interface{})
			if firstItem["nama"] != "DKI JAKARTA" {
				t.Errorf("expected DKI JAKARTA, got %s", firstItem["nama"])
			}
		} else {
			t.Error("expected data array in response")
		}
	})

	t.Run("GET /v2/provinsi handling empty data (404)", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to open sqlmock database: %s", err)
		}
		defer db.Close()
		mock.MatchExpectationsInOrder(false)

		repo := repository.NewDaerahRepository(db)
		svc := service.NewDaerahService(repo)
		h := New(svc)

		mux := http.NewServeMux()
		h.RegisterRoutes(mux)

		// Return 0 count and empty rows
		countRows := sqlmock.NewRows([]string{"count"}).AddRow(0)
		mock.ExpectQuery("^SELECT COUNT\\(\\*\\) FROM nk_provinsi$").WillReturnRows(countRows)

		dataRows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}) // Empty
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi LIMIT \\$1 OFFSET \\$2$").
			WithArgs(10, 0).
			WillReturnRows(dataRows)

		req := httptest.NewRequest(http.MethodGet, "/v2/provinsi", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		// Should return 404 because len(data) == 0 causes ErrNotFound
		if w.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestHandleDetailRoutes(t *testing.T) {
	t.Run("GET /v2/provinsi/{kode} returns single JSON", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to open sqlmock: %v", err)
		}
		defer db.Close()

		repo := repository.NewDaerahRepository(db)
		svc := service.NewDaerahService(repo)
		h := New(svc)

		mux := http.NewServeMux()
		h.RegisterRoutes(mux)

		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("32", "JAWA BARAT", -6.920, 107.604)
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi WHERE kode = \\$1$").
			WithArgs("32").
			WillReturnRows(rows)

		req := httptest.NewRequest(http.MethodGet, "/v2/provinsi/32", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var res map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}

		data, ok := res["data"].(map[string]interface{})
		if !ok || data["nama"] != "JAWA BARAT" {
			t.Errorf("unexpected response body: %v", res)
		}
	})

	t.Run("GET /v2/provinsi/{kode} 404 when not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to open sqlmock: %v", err)
		}
		defer db.Close()

		repo := repository.NewDaerahRepository(db)
		svc := service.NewDaerahService(repo)
		h := New(svc)

		mux := http.NewServeMux()
		h.RegisterRoutes(mux)

		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi WHERE kode = \\$1$").
			WithArgs("99").
			WillReturnRows(sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}))

		req := httptest.NewRequest(http.MethodGet, "/v2/provinsi/99", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})
}

func TestHandleSearchRoute(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewDaerahRepository(db)
	svc := service.NewDaerahService(repo)
	h := New(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	t.Run("GET /v2/search with short query returns 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v2/search?q=a", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("GET /v2/search with valid query returns results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"kode", "nama", "tipe", "lat", "lng", "kode_pos", "hierarki"}).
			AddRow("32", "JAWA BARAT", "provinsi", -6.920, 107.604, "", "JAWA BARAT")

		mock.ExpectQuery("SELECT kode, nama, 'provinsi' AS tipe").
			WithArgs("%jawa%", "jawa").
			WillReturnRows(rows)

		req := httptest.NewRequest(http.MethodGet, "/v2/search?q=jawa&tipe=provinsi&pagination=false", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d: %s", w.Code, w.Body.String())
		}
	})
}

func TestHandleReverseGeocodeRoute(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewDaerahRepository(db)
	svc := service.NewDaerahService(repo)
	h := New(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	t.Run("GET /v2/reverse with missing lat/lng returns 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v2/reverse?lat=12", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("GET /v2/reverse outside Indonesia returns 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v2/reverse?lat=51.5074&lng=-0.1278", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("GET /v2/reverse valid coordinates returns 200", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"p_kode", "p_nama",
			"kk_kode", "kk_nama",
			"kc_kode", "kc_nama",
			"d_kode", "d_nama", "d_kode_pos",
			"jarak_meter",
		}).AddRow("32", "JAWA BARAT", "32.73", "KOTA BANDUNG", "32.73.01", "SUKASARI", "32.73.01.1001", "SARIJADI", "40151", 100.5)

		mock.ExpectQuery("FROM nk_desa_kelurahan").
			WithArgs(-6.874, 107.589).
			WillReturnRows(rows)

		req := httptest.NewRequest(http.MethodGet, "/v2/reverse?lat=-6.874&lng=107.589", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d: %s", w.Code, w.Body.String())
		}

		var res map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}

		data, ok := res["data"].(map[string]interface{})
		if !ok {
			t.Fatalf("expected data object: %v", res)
		}
		if data["jarak_meter"].(float64) != 100.5 {
			t.Errorf("expected 100.5, got %v", data["jarak_meter"])
		}
	})
}
