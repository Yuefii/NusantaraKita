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
