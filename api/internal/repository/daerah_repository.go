package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"nusantarakita/internal/model"

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
		// rows.Next returns false on iteration errors too, so check rows.Err
		// to avoid returning partial data as if it were complete.
		if err := rows.Err(); err != nil {
			return nil, 0, err
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
		if err := rows.Err(); err != nil {
			cancel()
			dataCh <- dataResult{err: err}
			return
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
	query := "SELECT kode, nama, lat, lng FROM nk_provinsi ORDER BY kode ASC"
	countQuery := "SELECT COUNT(*) FROM nk_provinsi"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.Provinsi, error) {
		var p model.Provinsi
		err := rows.Scan(&p.Kode, &p.Nama, &p.Lat, &p.Lng)
		return p, err
	})
}

// ---------------- Kabupaten / Kota ----------------
func (r *DaerahRepository) GetKabKota(ctx context.Context, limit, offset int, pagination bool) ([]model.KabupatenKota, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota ORDER BY kode ASC"
	countQuery := "SELECT COUNT(*) FROM nk_kabupaten_kota"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.KabupatenKota, error) {
		var k model.KabupatenKota
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi)
		return k, err
	})
}

func (r *DaerahRepository) GetKabKotaByProvinsi(ctx context.Context, kodeProv string, limit, offset int, pagination bool) ([]model.KabupatenKota, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode_provinsi = $1 ORDER BY kode ASC"
	countQuery := "SELECT COUNT(*) FROM nk_kabupaten_kota WHERE kode_provinsi = $1"
	return fetchPaginatedData(ctx, r, query, countQuery, []interface{}{kodeProv}, limit, offset, pagination, func(rows *sql.Rows) (model.KabupatenKota, error) {
		var k model.KabupatenKota
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi)
		return k, err
	})
}

// ---------------- Kecamatan ----------------
func (r *DaerahRepository) GetKecamatan(ctx context.Context, limit, offset int, pagination bool) ([]model.Kecamatan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan ORDER BY kode ASC"
	countQuery := "SELECT COUNT(*) FROM nk_kecamatan"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.Kecamatan, error) {
		var k model.Kecamatan
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota)
		return k, err
	})
}

func (r *DaerahRepository) GetKecamatanByKabKota(ctx context.Context, kodeKab string, limit, offset int, pagination bool) ([]model.Kecamatan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan WHERE kode_kabupaten_kota = $1 ORDER BY kode ASC"
	countQuery := "SELECT COUNT(*) FROM nk_kecamatan WHERE kode_kabupaten_kota = $1"
	return fetchPaginatedData(ctx, r, query, countQuery, []interface{}{kodeKab}, limit, offset, pagination, func(rows *sql.Rows) (model.Kecamatan, error) {
		var k model.Kecamatan
		err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota)
		return k, err
	})
}

// ---------------- Desa / Kelurahan ----------------
func (r *DaerahRepository) GetDesaKelurahan(ctx context.Context, limit, offset int, pagination bool) ([]model.DesaKelurahan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan ORDER BY kode ASC"
	countQuery := "SELECT COUNT(*) FROM nk_desa_kelurahan"
	return fetchPaginatedData(ctx, r, query, countQuery, nil, limit, offset, pagination, func(rows *sql.Rows) (model.DesaKelurahan, error) {
		var d model.DesaKelurahan
		err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos)
		return d, err
	})
}

func (r *DaerahRepository) GetDesaKelurahanByKecamatan(ctx context.Context, kodeKec string, limit, offset int, pagination bool) ([]model.DesaKelurahan, int, error) {
	query := "SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan WHERE kode_kecamatan = $1 ORDER BY kode ASC"
	countQuery := "SELECT COUNT(*) FROM nk_desa_kelurahan WHERE kode_kecamatan = $1"
	return fetchPaginatedData(ctx, r, query, countQuery, []interface{}{kodeKec}, limit, offset, pagination, func(rows *sql.Rows) (model.DesaKelurahan, error) {
		var d model.DesaKelurahan
		err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos)
		return d, err
	})
}

// ---------------- Detail By Kode ----------------

func (r *DaerahRepository) GetProvinsiByKode(ctx context.Context, kode string) (*model.Provinsi, error) {
	query := "SELECT kode, nama, lat, lng FROM nk_provinsi WHERE kode = $1"
	var p model.Provinsi
	err := r.db.QueryRowContext(ctx, query, kode).Scan(&p.Kode, &p.Nama, &p.Lat, &p.Lng)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *DaerahRepository) GetKabKotaByKode(ctx context.Context, kode string) (*model.KabupatenKota, error) {
	query := "SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode = $1"
	var k model.KabupatenKota
	err := r.db.QueryRowContext(ctx, query, kode).Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &k, nil
}

func (r *DaerahRepository) GetKecamatanByKode(ctx context.Context, kode string) (*model.Kecamatan, error) {
	query := "SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan WHERE kode = $1"
	var k model.Kecamatan
	err := r.db.QueryRowContext(ctx, query, kode).Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &k, nil
}

func (r *DaerahRepository) GetDesaKelurahanByKode(ctx context.Context, kode string) (*model.DesaKelurahan, error) {
	query := "SELECT kode, nama, lat, lng, kode_kecamatan, COALESCE(kode_pos, '') FROM nk_desa_kelurahan WHERE kode = $1"
	var d model.DesaKelurahan
	err := r.db.QueryRowContext(ctx, query, kode).Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

// ---------------- Search / Autocomplete ----------------

func (r *DaerahRepository) SearchWilayah(ctx context.Context, q, tipe string, limit, offset int, pagination bool) ([]model.SearchItem, int, error) {
	pattern := "%" + q + "%"
	exact := q

	var query string
	var countQuery string
	var args []interface{}

	switch tipe {
	case "provinsi":
		query = `SELECT kode, nama, 'provinsi' AS tipe, lat, lng, '' AS kode_pos, nama AS hierarki
			FROM nk_provinsi
			WHERE nama ILIKE $1
			ORDER BY CASE WHEN LOWER(nama) = LOWER($2) THEN 1 WHEN LOWER(nama) LIKE LOWER($2) || '%' THEN 2 ELSE 3 END, nama ASC`
		countQuery = "SELECT COUNT(*) FROM nk_provinsi WHERE nama ILIKE $1"
		args = []interface{}{pattern, exact}

	case "kab-kota", "kabupaten_kota":
		query = `SELECT k.kode, k.nama, 'kabupaten_kota' AS tipe, k.lat, k.lng, '' AS kode_pos,
				k.nama || ', ' || p.nama AS hierarki
			FROM nk_kabupaten_kota k
			JOIN nk_provinsi p ON k.kode_provinsi = p.kode
			WHERE k.nama ILIKE $1
			ORDER BY CASE WHEN LOWER(k.nama) = LOWER($2) THEN 1 WHEN LOWER(k.nama) LIKE LOWER($2) || '%' THEN 2 ELSE 3 END, k.nama ASC`
		countQuery = "SELECT COUNT(*) FROM nk_kabupaten_kota WHERE nama ILIKE $1"
		args = []interface{}{pattern, exact}

	case "kecamatan":
		query = `SELECT kc.kode, kc.nama, 'kecamatan' AS tipe, kc.lat, kc.lng, '' AS kode_pos,
				kc.nama || ', ' || kk.nama || ', ' || p.nama AS hierarki
			FROM nk_kecamatan kc
			JOIN nk_kabupaten_kota kk ON kc.kode_kabupaten_kota = kk.kode
			JOIN nk_provinsi p ON kk.kode_provinsi = p.kode
			WHERE kc.nama ILIKE $1
			ORDER BY CASE WHEN LOWER(kc.nama) = LOWER($2) THEN 1 WHEN LOWER(kc.nama) LIKE LOWER($2) || '%' THEN 2 ELSE 3 END, kc.nama ASC`
		countQuery = "SELECT COUNT(*) FROM nk_kecamatan WHERE nama ILIKE $1"
		args = []interface{}{pattern, exact}

	case "desa-kel", "desa_kelurahan":
		query = `SELECT d.kode, d.nama, 'desa_kelurahan' AS tipe, d.lat, d.lng, COALESCE(d.kode_pos, '') AS kode_pos,
				d.nama || ', ' || kc.nama || ', ' || kk.nama || ', ' || p.nama AS hierarki
			FROM nk_desa_kelurahan d
			JOIN nk_kecamatan kc ON d.kode_kecamatan = kc.kode
			JOIN nk_kabupaten_kota kk ON kc.kode_kabupaten_kota = kk.kode
			JOIN nk_provinsi p ON kk.kode_provinsi = p.kode
			WHERE d.nama ILIKE $1 OR d.kode_pos = $2
			ORDER BY CASE WHEN LOWER(d.nama) = LOWER($3) THEN 1 WHEN LOWER(d.nama) LIKE LOWER($3) || '%' THEN 2 ELSE 3 END, d.nama ASC`
		countQuery = "SELECT COUNT(*) FROM nk_desa_kelurahan WHERE nama ILIKE $1 OR kode_pos = $2"
		args = []interface{}{pattern, exact, exact}

	default: // All types
		query = `SELECT kode, nama, tipe, lat, lng, kode_pos, hierarki FROM (
				SELECT kode, nama, 'provinsi' AS tipe, lat, lng, '' AS kode_pos, nama AS hierarki FROM nk_provinsi WHERE nama ILIKE $1
				UNION ALL
				SELECT k.kode, k.nama, 'kabupaten_kota' AS tipe, k.lat, k.lng, '' AS kode_pos, k.nama || ', ' || p.nama AS hierarki FROM nk_kabupaten_kota k JOIN nk_provinsi p ON k.kode_provinsi = p.kode WHERE k.nama ILIKE $1
				UNION ALL
				SELECT kc.kode, kc.nama, 'kecamatan' AS tipe, kc.lat, kc.lng, '' AS kode_pos, kc.nama || ', ' || kk.nama || ', ' || p.nama AS hierarki FROM nk_kecamatan kc JOIN nk_kabupaten_kota kk ON kc.kode_kabupaten_kota = kk.kode JOIN nk_provinsi p ON kk.kode_provinsi = p.kode WHERE kc.nama ILIKE $1
				UNION ALL
				SELECT d.kode, d.nama, 'desa_kelurahan' AS tipe, d.lat, d.lng, COALESCE(d.kode_pos, '') AS kode_pos, d.nama || ', ' || kc.nama || ', ' || kk.nama || ', ' || p.nama AS hierarki FROM nk_desa_kelurahan d JOIN nk_kecamatan kc ON d.kode_kecamatan = kc.kode JOIN nk_kabupaten_kota kk ON kc.kode_kabupaten_kota = kk.kode JOIN nk_provinsi p ON kk.kode_provinsi = p.kode WHERE d.nama ILIKE $1 OR d.kode_pos = $2
			) AS combined
			ORDER BY
				CASE WHEN LOWER(nama) = LOWER($3) THEN 1 WHEN LOWER(nama) LIKE LOWER($3) || '%' THEN 2 ELSE 3 END,
				CASE tipe WHEN 'provinsi' THEN 1 WHEN 'kabupaten_kota' THEN 2 WHEN 'kecamatan' THEN 3 ELSE 4 END,
				nama ASC`
		countQuery = `SELECT (
				(SELECT COUNT(*) FROM nk_provinsi WHERE nama ILIKE $1) +
				(SELECT COUNT(*) FROM nk_kabupaten_kota WHERE nama ILIKE $1) +
				(SELECT COUNT(*) FROM nk_kecamatan WHERE nama ILIKE $1) +
				(SELECT COUNT(*) FROM nk_desa_kelurahan WHERE nama ILIKE $1 OR kode_pos = $2)
			)`
		args = []interface{}{pattern, exact, exact}
	}

	return fetchPaginatedData(ctx, r, query, countQuery, args, limit, offset, pagination, func(rows *sql.Rows) (model.SearchItem, error) {
		var item model.SearchItem
		err := rows.Scan(&item.Kode, &item.Nama, &item.Tipe, &item.Lat, &item.Lng, &item.KodePos, &item.Hierarki)
		return item, err
	})
}

// ---------------- Reverse Geocoding ----------------

func (r *DaerahRepository) ReverseGeocode(ctx context.Context, lat, lng float64) (*model.ReverseGeocodeResult, error) {
	// Query with bounding box (+-0.5 degrees, approx 55 km)
	boxQuery := `SELECT 
			p.kode, p.nama,
			kk.kode, kk.nama,
			kc.kode, kc.nama,
			d.kode, d.nama, COALESCE(d.kode_pos, ''),
			d.jarak_meter
		FROM (
			SELECT kode, nama, kode_kecamatan, kode_pos,
				   6371000 * 2 * ASIN(
					   SQRT(
						   POWER(SIN(RADIANS(lat - $1) / 2), 2) +
						   COS(RADIANS($1)) * COS(RADIANS(lat)) *
						   POWER(SIN(RADIANS(lng - $2) / 2), 2)
					   )
				   ) AS jarak_meter
			FROM nk_desa_kelurahan
			WHERE lat BETWEEN $1 - 0.5 AND $1 + 0.5
			  AND lng BETWEEN $2 - 0.5 AND $2 + 0.5
			ORDER BY jarak_meter ASC
			LIMIT 1
		) d
		JOIN nk_kecamatan kc ON d.kode_kecamatan = kc.kode
		JOIN nk_kabupaten_kota kk ON kc.kode_kabupaten_kota = kk.kode
		JOIN nk_provinsi p ON kk.kode_provinsi = p.kode`

	var res model.ReverseGeocodeResult
	err := r.db.QueryRowContext(ctx, boxQuery, lat, lng).Scan(
		&res.Provinsi.Kode, &res.Provinsi.Nama,
		&res.KabupatenKota.Kode, &res.KabupatenKota.Nama,
		&res.Kecamatan.Kode, &res.Kecamatan.Nama,
		&res.DesaKelurahan.Kode, &res.DesaKelurahan.Nama, &res.DesaKelurahan.KodePos,
		&res.JarakMeter,
	)

	// If no rows found within bounding box, fallback to global nearest search
	if errors.Is(err, sql.ErrNoRows) {
		fallbackQuery := `SELECT 
				p.kode, p.nama,
				kk.kode, kk.nama,
				kc.kode, kc.nama,
				d.kode, d.nama, COALESCE(d.kode_pos, ''),
				d.jarak_meter
			FROM (
				SELECT kode, nama, kode_kecamatan, kode_pos,
					   6371000 * 2 * ASIN(
						   SQRT(
							   POWER(SIN(RADIANS(lat - $1) / 2), 2) +
							   COS(RADIANS($1)) * COS(RADIANS(lat)) *
							   POWER(SIN(RADIANS(lng - $2) / 2), 2)
						   )
					   ) AS jarak_meter
				FROM nk_desa_kelurahan
				ORDER BY (POWER(lat - $1, 2) + POWER(lng - $2, 2)) ASC
				LIMIT 1
			) d
			JOIN nk_kecamatan kc ON d.kode_kecamatan = kc.kode
			JOIN nk_kabupaten_kota kk ON kc.kode_kabupaten_kota = kk.kode
			JOIN nk_provinsi p ON kk.kode_provinsi = p.kode`

		err = r.db.QueryRowContext(ctx, fallbackQuery, lat, lng).Scan(
			&res.Provinsi.Kode, &res.Provinsi.Nama,
			&res.KabupatenKota.Kode, &res.KabupatenKota.Nama,
			&res.Kecamatan.Kode, &res.Kecamatan.Nama,
			&res.DesaKelurahan.Kode, &res.DesaKelurahan.Nama, &res.DesaKelurahan.KodePos,
			&res.JarakMeter,
		)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	res.Koordinat = model.CoordinateInfo{Lat: lat, Lng: lng}
	res.JarakMeter = math.Round(res.JarakMeter*100) / 100
	return &res, nil
}
