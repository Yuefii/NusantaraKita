package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"nusantarakita/internal/model"
	"nusantarakita/internal/service"
)

func TestSearchInputErrors(t *testing.T) {
	for _, tt := range []struct{ name, q, tipe, message string }{
		{"short ASCII", "a", "", service.ErrQueryTooShort.Error()},
		{"whitespace", " \t\u2003", "", service.ErrQueryTooShort.Error()},
		{"padded character", " a ", "", service.ErrQueryTooShort.Error()},
		{"one Unicode character", "é", "", service.ErrQueryTooShort.Error()},
		{"unknown type", "bandung", "unknown", service.ErrBadRequest.Error()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mux, mock := newTestMux(t)
			query := url.Values{"q": {tt.q}, "tipe": {tt.tipe}}
			assertError(t, serve(mux, "/v2/search?"+query.Encode()), http.StatusBadRequest, codeBadRequest, tt.message)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestSearchAcceptsTrimmedUnicodeQuery(t *testing.T) {
	mux, mock := newTestMux(t)
	rows := sqlmock.NewRows([]string{"kode", "nama", "tipe", "lat", "lng", "kode_pos", "hierarki"}).
		AddRow("32", "é中", "provinsi", -6.9, 107.6, "", "é中")
	mock.ExpectQuery("SELECT kode, nama, 'provinsi' AS tipe").
		WithArgs("%é中%", "é中").WillReturnRows(rows)
	query := url.Values{"q": {"\té中 \u2003"}, "tipe": {"provinsi"}, "pagination": {"false"}}
	w := serve(mux, "/v2/search?"+query.Encode())
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	var result model.PaginatedResponse[model.SearchItem]
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Pagination != nil || len(result.Data) != 1 || result.Data[0].Nama != "é中" {
		t.Errorf("unexpected response: %+v", result)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
