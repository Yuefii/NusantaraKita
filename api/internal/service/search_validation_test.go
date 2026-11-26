package service

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"nusantarakita/internal/model"
	"nusantarakita/internal/repository"
)

func newSearchService(t *testing.T) (*DaerahService, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
		db.Close()
	})
	svc := NewDaerahService(repository.NewDaerahRepository(db))
	svc.upstashURL = ""
	return svc, mock
}

func TestSearchRejectsInvalidInputBeforeFetching(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		json.NewEncoder(w).Encode(map[string]string{"result": ""})
	}))
	defer server.Close()

	for _, tt := range []struct {
		name, q, tipe string
		want          error
	}{
		{"empty", "", "", ErrQueryTooShort},
		{"one ASCII character", "a", "", ErrQueryTooShort},
		{"whitespace", " \t\u2003", "", ErrQueryTooShort},
		{"padded character", " a ", "", ErrQueryTooShort},
		{"accented character", "é", "", ErrQueryTooShort},
		{"CJK character", "中", "", ErrQueryTooShort},
		{"emoji", "🙂", "", ErrQueryTooShort},
		{"unknown type", "bandung", "unknown", ErrBadRequest},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newSearchService(t)
			svc.upstashURL, svc.upstashToken = server.URL, "test-token"
			svc.httpClient = server.Client()
			before := requests.Load()
			_, err := svc.SearchWilayah(context.Background(), tt.q, tt.tipe, 10, 1, true)
			if !errors.Is(err, tt.want) {
				t.Errorf("expected %v, got %v", tt.want, err)
			}
			if requests.Load() != before {
				t.Error("invalid search input reached the cache")
			}
		})
	}
}

func TestSearchTrimsQueryAndPreservesTypes(t *testing.T) {
	for _, tt := range []struct {
		name, tipe, resultType, queryPart, q string
		padded                               bool
	}{
		{"all types", "", "provinsi", "FROM (", "Banda Aceh", true},
		{"province", "provinsi", "provinsi", "FROM nk_provinsi", "Banda Aceh", false},
		{"city alias", "kab-kota", "kabupaten_kota", "FROM nk_kabupaten_kota", "Banda Aceh", false},
		{"city", "kabupaten_kota", "kabupaten_kota", "FROM nk_kabupaten_kota", "Banda Aceh", false},
		{"district", "kecamatan", "kecamatan", "FROM nk_kecamatan", "Banda Aceh", false},
		{"village alias", "desa-kel", "desa_kelurahan", "FROM nk_desa_kelurahan", "Banda Aceh", false},
		{"village", "desa_kelurahan", "desa_kelurahan", "FROM nk_desa_kelurahan", "Banda Aceh", false},
		{"two Unicode characters", "provinsi", "provinsi", "FROM nk_provinsi", "é中", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc, mock := newSearchService(t)
			args := []driver.Value{"%" + tt.q + "%", tt.q}
			if tt.tipe == "" || tt.resultType == "desa_kelurahan" {
				args = append(args, tt.q)
			}
			rows := sqlmock.NewRows([]string{"kode", "nama", "tipe", "lat", "lng", "kode_pos", "hierarki"}).
				AddRow("32", tt.q, tt.resultType, -6.9, 107.6, "", tt.q)
			mock.ExpectQuery(regexp.QuoteMeta(tt.queryPart)).WithArgs(args...).WillReturnRows(rows)

			input := tt.q
			if tt.padded {
				input = " \t" + tt.q + "\u2003"
			}
			result, err := svc.SearchWilayah(context.Background(), input, tt.tipe, 10, 1, false)
			if err != nil {
				t.Fatal(err)
			}
			response := result.(model.PaginatedResponse[model.SearchItem])
			if response.Pagination != nil || len(response.Data) != 1 {
				t.Fatalf("unexpected response: %+v", response)
			}
			item := response.Data[0]
			if item.Nama != tt.q || item.Tipe != tt.resultType || !strings.HasSuffix(item.GeojsonURL, "/"+tt.resultType+"/32.geojson") {
				t.Errorf("unexpected search item: %+v", item)
			}

			// The unpadded query must reuse the same local cache entry.
			cached, err := svc.SearchWilayah(context.Background(), tt.q, tt.tipe, 10, 1, false)
			if err != nil || !reflect.DeepEqual(cached, result) {
				t.Errorf("expected the cached result, got %+v, %v", cached, err)
			}
		})
	}
}

func TestUpstashCacheKeyRoundTrip(t *testing.T) {
	for _, tt := range []struct{ name, key, path string }{
		{"plain", "search_jawa_provinsi_10_1_true", "/get/search_jawa_provinsi_10_1_true"},
		{"reserved characters", "search_Banda Aceh/é?x#part_provinsi_10_1_true", "/get/search_Banda%20Aceh%2F%C3%A9%3Fx%23part_provinsi_10_1_true"},
		{"percent", "search_100%_provinsi_10_1_true", "/get/search_100%25_provinsi_10_1_true"},
	} {
		for _, hit := range []bool{true, false} {
			name := tt.name + "/miss"
			if hit {
				name = tt.name + "/hit"
			}
			t.Run(name, func(t *testing.T) {
				const data = `{"data":["Banda Aceh"]}`
				var gets atomic.Int32
				posts := make(chan []any, 1)
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.Header.Get("Authorization") != "Bearer test-token" {
						t.Error("missing cache authorization")
					}
					switch r.Method {
					case http.MethodGet:
						gets.Add(1)
						if r.URL.EscapedPath() != tt.path || r.URL.RawQuery != "" {
							t.Errorf("GET target = %q, want %q", r.RequestURI, tt.path)
						}
						value := ""
						if hit {
							value = data
						}
						json.NewEncoder(w).Encode(map[string]string{"result": value})
					case http.MethodPost:
						if r.URL.RequestURI() != "/" || r.Header.Get("Content-Type") != "application/json" {
							t.Errorf("unexpected POST target or content type: %s, %s", r.RequestURI, r.Header.Get("Content-Type"))
						}
						var payload []any
						if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
							t.Error(err)
						}
						posts <- payload
					default:
						t.Errorf("unexpected cache method: %s", r.Method)
					}
				}))
				defer server.Close()
				svc := NewDaerahService(nil)
				svc.upstashURL, svc.upstashToken = server.URL, "test-token"
				svc.httpClient = server.Client()
				fetches := 0
				result, err := svc.withCache(context.Background(), tt.key, func() (any, error) {
					fetches++
					if hit {
						return nil, errors.New("unexpected fetch on cache hit")
					}
					return json.RawMessage(data), nil
				})
				if err != nil {
					t.Fatal(err)
				}
				if string(result.(json.RawMessage)) != data || gets.Load() != 1 {
					t.Errorf("unexpected cache result or GET count: %s, %d", result, gets.Load())
				}
				if hit {
					if fetches != 0 {
						t.Errorf("cache hit fetched %d times", fetches)
					}
					return
				}
				if fetches != 1 {
					t.Errorf("cache miss fetched %d times", fetches)
				}
				select {
				case payload := <-posts:
					want := []any{"SET", tt.key, data, "EX", float64(86400)}
					if !reflect.DeepEqual(payload, want) {
						t.Errorf("POST payload = %#v, want %#v", payload, want)
					}
				case <-time.After(5 * time.Second):
					t.Fatal("cache miss did not POST the result")
				}
			})
		}
	}
}
