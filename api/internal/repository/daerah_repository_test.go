package repository

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"nusantarakita/internal/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetProvinsi(t *testing.T) {
	// Create mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	repo := NewDaerahRepository(db)

	t.Run("Fetch without pagination", func(t *testing.T) {
		// Expect the query
		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("11", "ACEH", 4.695135, 96.749399).
			AddRow("12", "SUMATERA UTARA", 2.115355, 99.545097)

		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi ORDER BY kode ASC$").WillReturnRows(rows)

		provs, total, err := repo.GetProvinsi(context.Background(), 10, 0, false)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if total != 0 {
			t.Errorf("expected total 0 for non-paginated, got %d", total)
		}
		if len(provs) != 2 {
			t.Errorf("expected 2 items, got %d", len(provs))
		}
		if provs[0].Nama != "ACEH" {
			t.Errorf("expected ACEH, got %s", provs[0].Nama)
		}
	})

	t.Run("Fetch with pagination", func(t *testing.T) {
		// For paginated data, fetchPaginatedData executes COUNT query and SELECT query concurrently
		// Due to concurrency in fetchPaginatedData, order of mock matching can be tricky.
		// go-sqlmock can MatchExpectationsInOrder(false)
		mock.MatchExpectationsInOrder(false)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(38)
		mock.ExpectQuery("^SELECT COUNT\\(\\*\\) FROM nk_provinsi$").WillReturnRows(countRows)

		dataRows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("31", "DKI JAKARTA", -6.175110, 106.827153)
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi ORDER BY kode ASC LIMIT \\$1 OFFSET \\$2$").
			WithArgs(10, 0).
			WillReturnRows(dataRows)

		provs, total, err := repo.GetProvinsi(context.Background(), 10, 0, true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if total != 38 {
			t.Errorf("expected total 38, got %d", total)
		}
		if len(provs) != 1 {
			t.Errorf("expected 1 item, got %d", len(provs))
		}
		if provs[0].Nama != "DKI JAKARTA" {
			t.Errorf("expected DKI JAKARTA, got %s", provs[0].Nama)
		}

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestGetByKode(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	repo := NewDaerahRepository(db)

	t.Run("GetProvinsiByKode found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("32", "JAWA BARAT", -6.920, 107.604)
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi WHERE kode = \\$1$").
			WithArgs("32").
			WillReturnRows(rows)

		prov, err := repo.GetProvinsiByKode(context.Background(), "32")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if prov == nil || prov.Nama != "JAWA BARAT" {
			t.Errorf("unexpected result: %+v", prov)
		}
	})

	t.Run("GetProvinsiByKode not found", func(t *testing.T) {
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi WHERE kode = \\$1$").
			WithArgs("99").
			WillReturnRows(sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}))

		prov, err := repo.GetProvinsiByKode(context.Background(), "99")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if prov != nil {
			t.Errorf("expected nil, got %+v", prov)
		}
	})

	t.Run("GetKabKotaByKode found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng", "kode_provinsi"}).
			AddRow("32.73", "KOTA BANDUNG", -6.917, 107.619, "32")
		mock.ExpectQuery("^SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode = \\$1$").
			WithArgs("32.73").
			WillReturnRows(rows)

		kab, err := repo.GetKabKotaByKode(context.Background(), "32.73")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if kab == nil || kab.Nama != "KOTA BANDUNG" {
			t.Errorf("unexpected result: %+v", kab)
		}
	})

	t.Run("GetKecamatanByKode found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng", "kode_kabupaten_kota"}).
			AddRow("32.73.01", "SUKASARI", -6.874, 107.589, "32.73")
		mock.ExpectQuery("^SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan WHERE kode = \\$1$").
			WithArgs("32.73.01").
			WillReturnRows(rows)

		kec, err := repo.GetKecamatanByKode(context.Background(), "32.73.01")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if kec == nil || kec.Nama != "SUKASARI" {
			t.Errorf("unexpected result: %+v", kec)
		}
	})

	t.Run("GetDesaKelurahanByKode found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng", "kode_kecamatan", "kode_pos"}).
			AddRow("32.73.01.1001", "SARIJADI", -6.874, 107.589, "32.73.01", "40151")
		mock.ExpectQuery("^SELECT kode, nama, lat, lng, kode_kecamatan, COALESCE\\(kode_pos, ''\\) FROM nk_desa_kelurahan WHERE kode = \\$1$").
			WithArgs("32.73.01.1001").
			WillReturnRows(rows)

		desa, err := repo.GetDesaKelurahanByKode(context.Background(), "32.73.01.1001")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if desa == nil || desa.Nama != "SARIJADI" || desa.KodePos != "40151" {
			t.Errorf("unexpected result: %+v", desa)
		}
	})
}

func TestSearchWilayah(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	repo := NewDaerahRepository(db)

	t.Run("Search with tipe provinsi", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"kode", "nama", "tipe", "lat", "lng", "kode_pos", "hierarki"}).
			AddRow("32", "JAWA BARAT", "provinsi", -6.920, 107.604, "", "JAWA BARAT")

		mock.ExpectQuery("SELECT kode, nama, 'provinsi' AS tipe").
			WithArgs("%jawa%", "jawa").
			WillReturnRows(rows)

		items, total, err := repo.SearchWilayah(context.Background(), "jawa", "provinsi", 10, 0, false)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if total != 0 {
			t.Errorf("expected 0 total for non-paginated, got %d", total)
		}
		if len(items) != 1 || items[0].Nama != "JAWA BARAT" {
			t.Errorf("unexpected items: %+v", items)
		}
	})
}

func TestReverseGeocode(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	repo := NewDaerahRepository(db)

	t.Run("Found within bounding box", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"p_kode", "p_nama",
			"kk_kode", "kk_nama",
			"kc_kode", "kc_nama",
			"d_kode", "d_nama", "d_kode_pos",
			"jarak_meter",
		}).AddRow("32", "JAWA BARAT", "32.73", "KOTA BANDUNG", "32.73.01", "SUKASARI", "32.73.01.1001", "SARIJADI", "40151", 123.456)

		mock.ExpectQuery("FROM nk_desa_kelurahan").
			WithArgs(-6.874, 107.589).
			WillReturnRows(rows)

		res, err := repo.ReverseGeocode(context.Background(), -6.874, 107.589)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if res == nil {
			t.Fatal("expected result, got nil")
		}
		if res.DesaKelurahan.Nama != "SARIJADI" {
			t.Errorf("expected SARIJADI, got %s", res.DesaKelurahan.Nama)
		}
		if res.Provinsi.Nama != "JAWA BARAT" {
			t.Errorf("expected JAWA BARAT, got %s", res.Provinsi.Nama)
		}
		if res.JarakMeter != 123.46 {
			t.Errorf("expected 123.46, got %f", res.JarakMeter)
		}
	})
}

func TestListQueriesOrderByKode(t *testing.T) {
	type listFunc func(repo *DaerahRepository) error

	cases := []struct {
		name    string
		query   string
		args    []driver.Value
		columns []string
		call    listFunc
	}{
		{"GetProvinsi", "FROM nk_provinsi", nil, []string{"kode", "nama", "lat", "lng"},
			func(r *DaerahRepository) error {
				_, _, err := r.GetProvinsi(context.Background(), 10, 20, true)
				return err
			}},
		{"GetKabKota", "FROM nk_kabupaten_kota", nil, []string{"kode", "nama", "lat", "lng", "kode_provinsi"},
			func(r *DaerahRepository) error {
				_, _, err := r.GetKabKota(context.Background(), 10, 20, true)
				return err
			}},
		{"GetKabKotaByProvinsi", "FROM nk_kabupaten_kota WHERE kode_provinsi = \\$1", []driver.Value{"32"}, []string{"kode", "nama", "lat", "lng", "kode_provinsi"},
			func(r *DaerahRepository) error {
				_, _, err := r.GetKabKotaByProvinsi(context.Background(), "32", 10, 20, true)
				return err
			}},
		{"GetKecamatan", "FROM nk_kecamatan", nil, []string{"kode", "nama", "lat", "lng", "kode_kabupaten_kota"},
			func(r *DaerahRepository) error {
				_, _, err := r.GetKecamatan(context.Background(), 10, 20, true)
				return err
			}},
		{"GetKecamatanByKabKota", "FROM nk_kecamatan WHERE kode_kabupaten_kota = \\$1", []driver.Value{"32.01"}, []string{"kode", "nama", "lat", "lng", "kode_kabupaten_kota"},
			func(r *DaerahRepository) error {
				_, _, err := r.GetKecamatanByKabKota(context.Background(), "32.01", 10, 20, true)
				return err
			}},
		{"GetDesaKelurahan", "FROM nk_desa_kelurahan", nil, []string{"kode", "nama", "lat", "lng", "kode_kecamatan", "kode_pos"},
			func(r *DaerahRepository) error {
				_, _, err := r.GetDesaKelurahan(context.Background(), 10, 20, true)
				return err
			}},
		{"GetDesaKelurahanByKecamatan", "FROM nk_desa_kelurahan WHERE kode_kecamatan = \\$1", []driver.Value{"32.01.01"}, []string{"kode", "nama", "lat", "lng", "kode_kecamatan", "kode_pos"},
			func(r *DaerahRepository) error {
				_, _, err := r.GetDesaKelurahanByKecamatan(context.Background(), "32.01.01", 10, 20, true)
				return err
			}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to open sqlmock database: %s", err)
			}
			defer db.Close()
			mock.MatchExpectationsInOrder(false)

			repo := NewDaerahRepository(db)
			n := len(tc.args)
			mock.ExpectQuery("^SELECT COUNT\\(\\*\\) " + tc.query + "$").
				WithArgs(tc.args...).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			// Pages must come from a fixed order, so ORDER BY comes before LIMIT and OFFSET.
			dataQuery := fmt.Sprintf("^SELECT .* %s ORDER BY kode ASC LIMIT \\$%d OFFSET \\$%d$", tc.query, n+1, n+2)
			mock.ExpectQuery(dataQuery).
				WithArgs(append(append([]driver.Value{}, tc.args...), 10, 20)...).
				WillReturnRows(sqlmock.NewRows(tc.columns))

			if err := tc.call(repo); err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unmet expectations: %v", err)
			}
		})
	}
}

func TestFetchPaginatedDataRowsErr(t *testing.T) {
	errIter := errors.New("connection reset during iteration")

	t.Run("without pagination", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to open sqlmock database: %s", err)
		}
		defer db.Close()

		repo := NewDaerahRepository(db)
		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("11", "ACEH", 4.695135, 96.749399).
			AddRow("12", "SUMATERA UTARA", 2.115355, 99.545097).
			RowError(1, errIter)
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi ORDER BY kode ASC$").WillReturnRows(rows)

		provs, total, err := repo.GetProvinsi(context.Background(), 10, 0, false)
		if !errors.Is(err, errIter) {
			t.Fatalf("expected iteration error %v, got %v", errIter, err)
		}
		if provs != nil || total != 0 {
			t.Errorf("expected no data on error, got %d items, total %d", len(provs), total)
		}
	})

	t.Run("with pagination", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to open sqlmock database: %s", err)
		}
		defer db.Close()
		mock.MatchExpectationsInOrder(false)

		repo := NewDaerahRepository(db)
		dataQuery := "^SELECT kode, nama, lat, lng FROM nk_provinsi ORDER BY kode ASC LIMIT \\$1 OFFSET \\$2$"

		// A first successful call caches the COUNT result, so the failing call
		// below only runs the data query and its error is reported deterministically.
		mock.ExpectQuery("^SELECT COUNT\\(\\*\\) FROM nk_provinsi$").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(38))
		mock.ExpectQuery(dataQuery).
			WithArgs(10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
				AddRow("11", "ACEH", 4.695135, 96.749399))
		if _, _, err := repo.GetProvinsi(context.Background(), 10, 0, true); err != nil {
			t.Fatalf("expected no error on first call, got %v", err)
		}

		rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng"}).
			AddRow("11", "ACEH", 4.695135, 96.749399).
			AddRow("12", "SUMATERA UTARA", 2.115355, 99.545097).
			RowError(1, errIter)
		mock.ExpectQuery(dataQuery).
			WithArgs(10, 0).
			WillReturnRows(rows)

		provs, total, err := repo.GetProvinsi(context.Background(), 10, 0, true)
		if !errors.Is(err, errIter) {
			t.Fatalf("expected iteration error %v, got %v", errIter, err)
		}
		if provs != nil || total != 0 {
			t.Errorf("expected no data on error, got %d items, total %d", len(provs), total)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unmet expectations: %v", err)
		}
	})
}

func TestDesaListsCoalescePostalCodes(t *testing.T) {
	for _, filtered := range []bool{false, true} {
		for _, paginated := range []bool{false, true} {
			t.Run(fmt.Sprintf("filtered=%t/paginated=%t", filtered, paginated), func(t *testing.T) {
				db, mock, err := sqlmock.New()
				if err != nil {
					t.Fatal(err)
				}
				defer db.Close()
				mock.MatchExpectationsInOrder(false)
				repo := NewDaerahRepository(db)
				query := "SELECT kode, nama, lat, lng, kode_kecamatan, COALESCE(kode_pos, '') FROM nk_desa_kelurahan"
				count := "SELECT COUNT(*) FROM nk_desa_kelurahan"
				var args []driver.Value
				if filtered {
					query += " WHERE kode_kecamatan = $1"
					count += " WHERE kode_kecamatan = $1"
					args = append(args, "32.01.01")
				}
				if paginated {
					mock.ExpectQuery(regexp.QuoteMeta(count)).WithArgs(args...).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
				}
				query += " ORDER BY kode ASC"
				if paginated {
					query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
					args = append(args, 10, 20)
				}
				// sqlmock does not evaluate SQL: the empty string represents PostgreSQL's
				// COALESCE result for a NULL postal code. Require that expression above.
				rows := sqlmock.NewRows([]string{"kode", "nama", "lat", "lng", "kode_kecamatan", "kode_pos"}).
					AddRow("32.01.01.1001", "Example A", -6.0, 106.0, "32.01.01", "").
					AddRow("32.01.01.1002", "Example B", -6.1, 106.1, "32.01.01", "12345")
				mock.ExpectQuery("^" + regexp.QuoteMeta(query) + "$").WithArgs(args...).WillReturnRows(rows)
				var result []model.DesaKelurahan
				var total int
				if filtered {
					result, total, err = repo.GetDesaKelurahanByKecamatan(context.Background(), "32.01.01", 10, 20, paginated)
				} else {
					result, total, err = repo.GetDesaKelurahan(context.Background(), 10, 20, paginated)
				}
				if err != nil {
					t.Fatal(err)
				}
				if len(result) != 2 || result[0].KodePos != "" || result[1].KodePos != "12345" {
					t.Fatalf("unexpected postal codes: %+v", result)
				}
				wantTotal := 0
				if paginated {
					wantTotal = 2
				}
				if total != wantTotal {
					t.Errorf("total = %d, want %d", total, wantTotal)
				}
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Fatal(err)
				}
			})
		}
	}
}
