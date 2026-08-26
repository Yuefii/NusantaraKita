package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"vercel-go-starter/internal/model"

	"github.com/patrickmn/go-cache"
)

type DaerahRepository struct {
	db    *sql.DB
	cache *cache.Cache
}

func NewDaerahRepository(db *sql.DB) *DaerahRepository {
	// Cache for COUNT queries with 24 hour expiration
	c := cache.New(24*time.Hour, 1*time.Hour)
	return &DaerahRepository{
		db:    db,
		cache: c,
	}
}

// fetchPaginatedData is a generic function to fetch paginated or non-paginated data from the database
func fetchPaginatedData[T any](ctx context.Context, r *DaerahRepository, query string, countQuery string, args []interface{}, limit, offset int, pagination bool, scanFunc func(*sql.Rows) (T, error)) ([]T, int, error) {
	if !pagination {
		rows, err := r.db.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		var data []T
		for rows.Next() {
			item, err := scanFunc(rows)
			if err != nil {
				return nil, 0, err
			}
			data = append(data, item)
		}
		return data, 0, nil
	}

	// Create a sub-context to cancel the other query early if one fails
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a cache key for the count query
	countCacheKey := fmt.Sprintf("count_%s_%v", countQuery, args)

	type countResult struct {
		total int
		err   error
	}

	type dataResult struct {
		data []T
		err  error
	}

	countCh := make(chan countResult, 1)
	dataCh := make(chan dataResult, 1)

	// Goroutine 1: Execute COUNT(*) query or get from cache
	go func() {
		// Check cache first
		if cachedTotal, found := r.cache.Get(countCacheKey); found {
			countCh <- countResult{total: cachedTotal.(int), err: nil}
			return
		}

		var total int
		err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
		if err != nil {
			cancel() // Abort the data query immediately
		} else {
			r.cache.Set(countCacheKey, total, cache.DefaultExpiration)
		}
		countCh <- countResult{total: total, err: err}
	}()

	// Goroutine 2: Execute SELECT DATA query
	go func() {
		var result []T
		finalArgs := append(args, limit, offset)
		paginatedQuery := fmt.Sprintf("%s LIMIT $%d OFFSET $%d", query, len(args)+1, len(args)+2)
		rows, err := r.db.QueryContext(ctx, paginatedQuery, finalArgs...)
		if err != nil {
			cancel() // Abort the count query immediately
			dataCh <- dataResult{err: err}
			return
		}
		defer rows.Close()

		for rows.Next() {
			item, err := scanFunc(rows)
			if err != nil {
				cancel()
				dataCh <- dataResult{err: err}
				return
			}
			result = append(result, item)
		}
		dataCh <- dataResult{data: result, err: nil}
	}()

	// Wait and collect results from both Goroutines
	cRes := <-countCh
	if cRes.err != nil {
		return nil, 0, cRes.err
	}

	dRes := <-dataCh
	if dRes.err != nil {
		return nil, 0, dRes.err
	}

	return dRes.data, cRes.total, nil
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
	query := "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode_provinsi = $1"
	countQuery := "SELECT COUNT(*) FROM nk_kabupaten_kota WHERE kode_provinsi = $1"
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
	query := "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan WHERE kode_kabupaten_kota = $1"
	countQuery := "SELECT COUNT(*) FROM nk_kecamatan WHERE kode_kabupaten_kota = $1"
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
	query := "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan WHERE kode_kecamatan = $1"
	countQuery := "SELECT COUNT(*) FROM nk_desa_kelurahan WHERE kode_kecamatan = $1"
	return fetchPaginatedData(ctx, r, query, countQuery, []interface{}{kodeKec}, limit, offset, pagination, func(rows *sql.Rows) (model.DesaKelurahan, error) {
		var d model.DesaKelurahan
		err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos)
		return d, err
	})
}
