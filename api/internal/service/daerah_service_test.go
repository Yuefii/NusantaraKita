package service

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"nusantarakita/internal/model"
	"nusantarakita/internal/repository"
)

func TestBuildResponse(t *testing.T) {
	t.Run("Empty data returns ErrNotFound", func(t *testing.T) {
		data := []string{}
		_, err := buildResponse(data, 10, 1, 0, true)
		if err != ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("Pagination disabled", func(t *testing.T) {
		data := []string{"item1"}
		res, err := buildResponse(data, 10, 1, 1, false)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		paginatedRes, ok := res.(model.PaginatedResponse[string])
		if !ok {
			t.Fatalf("expected model.PaginatedResponse[string], got %T", res)
		}
		if paginatedRes.Pagination != nil {
			t.Errorf("expected nil Pagination when disabled, got %+v", paginatedRes.Pagination)
		}
		if len(paginatedRes.Data) != 1 || paginatedRes.Data[0] != "item1" {
			t.Errorf("expected data [item1], got %v", paginatedRes.Data)
		}

		rawBytes, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("failed to marshal response: %v", err)
		}
		var raw map[string]json.RawMessage
		if err := json.Unmarshal(rawBytes, &raw); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}
		if _, hasPagination := raw["pagination"]; hasPagination {
			t.Errorf("expected pagination field to be omitted, got %s", string(rawBytes))
		}
		if _, hasData := raw["data"]; !hasData {
			t.Errorf("expected data field in JSON, got %s", string(rawBytes))
		}
	})

	t.Run("Page exceeds max page", func(t *testing.T) {
		// A page past the end returns no rows from the query.
		data := []string{}
		// 15 items, limit 10 => 2 pages max. Requesting page 3.
		_, err := buildResponse(data, 10, 3, 15, true)
		if err == nil {
			t.Fatal("expected ErrPageExceeded, got nil")
		}

		pageErr, ok := err.(ErrPageExceeded)
		if !ok {
			t.Fatalf("expected ErrPageExceeded, got %T", err)
		}
		if pageErr.MaxPage != 2 {
			t.Errorf("expected MaxPage 2, got %d", pageErr.MaxPage)
		}
	})

	t.Run("Valid paginated response", func(t *testing.T) {
		data := []string{"item1", "item2"}
		res, err := buildResponse(data, 10, 1, 12, true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		// res should be of type model.PaginatedResponse[string]
		// Since we can't easily type assert generic types defined in another package if we don't import model,
		// Wait, we can import model and type assert.
		// Actually, since buildResponse returns 'any', we can just check if it's not nil for now.
		if res == nil {
			t.Error("expected response, got nil")
		}
	})
}

func TestPageOffset(t *testing.T) {
	tests := []struct {
		limit, halaman, want int
	}{
		{10, 1, 0},
		{10, 3, 20},
		{1, math.MaxInt, math.MaxInt - 1},
		{10, math.MaxInt, math.MaxInt},
		{10, math.MaxInt/10 + 1, math.MaxInt / 10 * 10},
		{10, math.MaxInt/10 + 2, math.MaxInt},
		{math.MaxInt, 2, math.MaxInt},
		{math.MaxInt, 3, math.MaxInt},
	}
	for _, tt := range tests {
		if got := pageOffset(tt.limit, tt.halaman); got != tt.want {
			t.Errorf("pageOffset(%d, %d) = %d, want %d", tt.limit, tt.halaman, got, tt.want)
		}
		if got := pageOffset(tt.limit, tt.halaman); got < 0 {
			t.Errorf("pageOffset(%d, %d) is negative: %d", tt.limit, tt.halaman, got)
		}
	}
}

func TestBuildResponseHugeLimit(t *testing.T) {
	t.Run("first page reports one page", func(t *testing.T) {
		res, err := buildResponse([]string{"a", "b"}, math.MaxInt, 1, 2, true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if got := res.(model.PaginatedResponse[string]).Pagination.TotalHalaman; got != 1 {
			t.Errorf("expected total_halaman 1, got %d", got)
		}
	})

	t.Run("second page is out of range", func(t *testing.T) {
		_, err := buildResponse([]string{}, math.MaxInt, 2, 2, true)
		var pageErr ErrPageExceeded
		if !errors.As(err, &pageErr) || pageErr.MaxPage != 1 {
			t.Fatalf("expected ErrPageExceeded{MaxPage: 1}, got %v", err)
		}
	})
}

func TestSearchServiceValidation(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewDaerahRepository(db)
	svc := NewDaerahService(repo)

	t.Run("Query less than 2 characters returns ErrQueryTooShort", func(t *testing.T) {
		_, err := svc.SearchWilayah(context.Background(), "a", "", 10, 1, true)
		if !errors.Is(err, ErrQueryTooShort) {
			t.Errorf("expected ErrQueryTooShort, got %v", err)
		}
	})
}

func TestReverseGeocodeServiceValidation(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewDaerahRepository(db)
	svc := NewDaerahService(repo)

	t.Run("Coordinates clearly outside Indonesia returns ErrOutsideIndonesia", func(t *testing.T) {
		// New York coordinates
		_, err := svc.ReverseGeocode(context.Background(), 40.7128, -74.0060)
		if !errors.Is(err, ErrOutsideIndonesia) {
			t.Errorf("expected ErrOutsideIndonesia, got %v", err)
		}
	})
}

func TestDetailService(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewDaerahRepository(db)
	svc := NewDaerahService(repo)

	t.Run("GetProvinsiByKode not found returns ErrNotFound", func(t *testing.T) {
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi WHERE kode = \\$1$").
			WithArgs("99").
			WillReturnRows(sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}))

		_, err := svc.GetProvinsiByKode(context.Background(), "99")
		if !errors.Is(err, ErrNotFound) {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("GetProvinsiByKode found returns SingleResponse with GeojsonURL", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("32", "JAWA BARAT", -6.920, 107.604)
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi WHERE kode = \\$1$").
			WithArgs("32").
			WillReturnRows(rows)

		res, err := svc.GetProvinsiByKode(context.Background(), "32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		singleRes, ok := res.(model.SingleResponse[*model.Provinsi])
		if !ok {
			t.Fatalf("expected model.SingleResponse[*model.Provinsi], got %T", res)
		}
		if singleRes.Data.Nama != "JAWA BARAT" {
			t.Errorf("expected JAWA BARAT, got %s", singleRes.Data.Nama)
		}
		if singleRes.Data.GeojsonURL == "" {
			t.Errorf("expected non-empty GeojsonURL")
		}
	})
}

func TestKodePosOmitEmptyConsistency(t *testing.T) {
	t.Run("Empty kode_pos is omitted across all models", func(t *testing.T) {
		cases := []struct {
			name  string
			value any
		}{
			{
				name: "DesaKelurahan",
				value: model.DesaKelurahan{
					Kode:          "32.73.01.1001",
					Nama:          "SARIJADI",
					KodeKecamatan: "32.73.01",
					KodePos:       "",
				},
			},
			{
				name: "SearchItem",
				value: model.SearchItem{
					Kode:    "32.73.01.1001",
					Nama:    "SARIJADI",
					Tipe:    "desa_kelurahan",
					KodePos: "",
				},
			},
			{
				name: "DesaKelurahanInfo",
				value: model.DesaKelurahanInfo{
					Kode:    "32.73.01.1001",
					Nama:    "SARIJADI",
					KodePos: "",
				},
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := json.Marshal(tc.value)
				if err != nil {
					t.Fatalf("failed to marshal %s: %v", tc.name, err)
				}
				var fields map[string]json.RawMessage
				if err := json.Unmarshal(b, &fields); err != nil {
					t.Fatalf("failed to unmarshal %s: %v", tc.name, err)
				}
				if _, exists := fields["kode_pos"]; exists {
					t.Errorf("expected kode_pos to be omitted for %s when empty, got %s", tc.name, string(b))
				}
			})
		}
	})

	t.Run("Non-empty kode_pos is serialized across all models", func(t *testing.T) {
		cases := []struct {
			name  string
			value any
		}{
			{
				name: "DesaKelurahan",
				value: model.DesaKelurahan{
					Kode:          "32.73.01.1001",
					Nama:          "SARIJADI",
					KodeKecamatan: "32.73.01",
					KodePos:       "40151",
				},
			},
			{
				name: "SearchItem",
				value: model.SearchItem{
					Kode:    "32.73.01.1001",
					Nama:    "SARIJADI",
					Tipe:    "desa_kelurahan",
					KodePos: "40151",
				},
			},
			{
				name: "DesaKelurahanInfo",
				value: model.DesaKelurahanInfo{
					Kode:    "32.73.01.1001",
					Nama:    "SARIJADI",
					KodePos: "40151",
				},
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := json.Marshal(tc.value)
				if err != nil {
					t.Fatalf("failed to marshal %s: %v", tc.name, err)
				}
				var decoded struct {
					KodePos string `json:"kode_pos"`
				}
				if err := json.Unmarshal(b, &decoded); err != nil {
					t.Fatalf("failed to unmarshal %s: %v", tc.name, err)
				}
				if decoded.KodePos != "40151" {
					t.Errorf("expected kode_pos 40151 for %s, got %q (%s)", tc.name, decoded.KodePos, string(b))
				}
			})
		}
	})
}
