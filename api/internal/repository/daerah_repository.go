package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"vercel-go-starter/internal/model"
)

type DaerahRepository struct {
	db *sql.DB
}

func NewDaerahRepository(db *sql.DB) *DaerahRepository {
	return &DaerahRepository{db: db}
}

// replacePlaceholders replaces ? with $1, $2, etc for PostgreSQL
func replacePlaceholders(query string) string {
	n := 1
	for {
		idx := strings.Index(query, "?")
		if idx == -1 {
			break
		}
		query = query[:idx] + fmt.Sprintf("$%d", n) + query[idx+1:]
		n++
	}
	return query
}

// fetchPaginatedData is a generic function to fetch paginated or non-paginated data from the database
func fetchPaginatedData[T any](ctx context.Context, r *DaerahRepository, query string, countQuery string, args []interface{}, limit, offset int, pagination bool, scanFunc func(*sql.Rows) (T, error)) ([]T, int, error) {
	var data []T
	var totalItem int
	var err error

	if !pagination {
		// Use QueryContext for request cancellation
		rows, err := r.db.QueryContext(ctx, replacePlaceholders(query), args...)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		for rows.Next() {
			item, err := scanFunc(rows)
			if err != nil {
				return nil, 0, err
			}
			data = append(data, item)
		}
		return data, 0, nil
	}

	// Use QueryRowContext for request cancellation
	err = r.db.QueryRowContext(ctx, replacePlaceholders(countQuery), args...).Scan(&totalItem)
	if err != nil {
		return nil, 0, err
	}

	finalArgs := append(args, limit, offset)
	// Use QueryContext for request cancellation
	rows, err := r.db.QueryContext(ctx, replacePlaceholders(query+" LIMIT ? OFFSET ?"), finalArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		item, err := scanFunc(rows)
		if err != nil {
			return nil, 0, err
		}
		data = append(data, item)
	}

	return data, totalItem, nil
}

// ---------------- Provinsi ----------------
func (r *DaerahRepository) GetProvinsi(ctx context.Context, limit, offset int, pagination bool) ([]model.Provinsi, int, error) {
	query := "SELECT kode, nama, lat, lng FROM nk_provinsi"
	countQuery := "SELECT COUNT(*) FROM nk_provinsi"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.Provinsi, error) {
		var p model.Provinsi
		err := rows.Scan(&p.Kode, &p.Nama, &p.Lat, &p.Lng)
		return p, err
	})
}

// ---------------- Kabupaten / Kota ----------------
func (r *DaerahRepository) GetKabKota(ctx context.Context, limit, offset int, pagination bool) ([]model.KabupatenKota, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota"
	countQuery := "SELECT COUNT(*) FROM nk_kabupaten_kota"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.KabupatenKota, error) {
		var k model.KabupatenKota
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi)
		return k, err
	})
}

func (r *DaerahRepository) GetKabKotaByProvinsi(ctx context.Context, kodeProv string, limit, offset int, pagination bool) ([]model.KabupatenKota, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode_provinsi = ?"
	countQuery := "SELECT COUNT(*) FROM nk_kabupaten_kota WHERE kode_provinsi = ?"
	return fetchPaginatedData(ctx, r, query, countQuery, []interface{}{kodeProv}, limit, offset, pagination, func(rows *sql.Rows) (model.KabupatenKota, error) {
		var k model.KabupatenKota
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi)
		return k, err
	})
}

// ---------------- Kecamatan ----------------
func (r *DaerahRepository) GetKecamatan(ctx context.Context, limit, offset int, pagination bool) ([]model.Kecamatan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan"
	countQuery := "SELECT COUNT(*) FROM nk_kecamatan"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.Kecamatan, error) {
		var k model.Kecamatan
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota)
		return k, err
	})
}

func (r *DaerahRepository) GetKecamatanByKabKota(ctx context.Context, kodeKab string, limit, offset int, pagination bool) ([]model.Kecamatan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan WHERE kode_kabupaten_kota = ?"
	countQuery := "SELECT COUNT(*) FROM nk_kecamatan WHERE kode_kabupaten_kota = ?"
	return fetchPaginatedData(ctx, r, query, countQuery, []interface{}{kodeKab}, limit, offset, pagination, func(rows *sql.Rows) (model.Kecamatan, error) {
		var k model.Kecamatan
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota)
		return k, err
	})
}

// ---------------- Desa / Kelurahan ----------------
func (r *DaerahRepository) GetDesaKelurahan(ctx context.Context, limit, offset int, pagination bool) ([]model.DesaKelurahan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan"
	countQuery := "SELECT COUNT(*) FROM nk_desa_kelurahan"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.DesaKelurahan, error) {
		var d model.DesaKelurahan
		err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos)
		return d, err
	})
}

func (r *DaerahRepository) GetDesaKelurahanByKecamatan(ctx context.Context, kodeKec string, limit, offset int, pagination bool) ([]model.DesaKelurahan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan WHERE kode_kecamatan = ?"
	countQuery := "SELECT COUNT(*) FROM nk_desa_kelurahan WHERE kode_kecamatan = ?"
	return fetchPaginatedData(ctx, r, query, countQuery, []interface{}{kodeKec}, limit, offset, pagination, func(rows *sql.Rows) (model.DesaKelurahan, error) {
		var d model.DesaKelurahan
		err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos)
		return d, err
	})
}
