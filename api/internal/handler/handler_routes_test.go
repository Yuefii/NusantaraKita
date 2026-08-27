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
