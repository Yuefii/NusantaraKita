package repository

import (
	"context"
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

		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi$").WillReturnRows(rows)

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
		mock.ExpectQuery("^SELECT kode, nama, lat, lng FROM nk_provinsi LIMIT \\$1 OFFSET \\$2$").
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
