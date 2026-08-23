package repository

import (
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

func (r *DaerahRepository) GetProvinsi(limit, offset int, pagination bool) ([]model.Provinsi, int, error) {
	var data []model.Provinsi
	var totalItem int
	var err error

	if !pagination {
		rows, err := r.db.Query("SELECT kode, nama, lat, lng FROM nk_provinsi")
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		for rows.Next() {
			var p model.Provinsi
			if err := rows.Scan(&p.Kode, &p.Nama, &p.Lat, &p.Lng); err != nil {
				return nil, 0, err
			}
			data = append(data, p)
		}
		return data, 0, nil
	}

	err = r.db.QueryRow("SELECT COUNT(*) FROM nk_provinsi").Scan(&totalItem)
	if err != nil {
		return nil, 0, err
	}

	query := replacePlaceholders("SELECT kode, nama, lat, lng FROM nk_provinsi LIMIT ? OFFSET ?")
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Provinsi
		if err := rows.Scan(&p.Kode, &p.Nama, &p.Lat, &p.Lng); err != nil {
			return nil, 0, err
		}
		data = append(data, p)
	}

	return data, totalItem, nil
}

func (r *DaerahRepository) fetchKabKotaRaw(query string, countQuery string, args []interface{}, limit, offset int, pagination bool) ([]model.KabupatenKota, int, error) {
	var data []model.KabupatenKota
	var totalItem int
	var err error

	if !pagination {
		rows, err := r.db.Query(replacePlaceholders(query), args...)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		for rows.Next() {
			var k model.KabupatenKota
			if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi); err != nil {
				return nil, 0, err
			}
			data = append(data, k)
		}
		return data, 0, nil
	}

	err = r.db.QueryRow(replacePlaceholders(countQuery), args...).Scan(&totalItem)
	if err != nil {
		return nil, 0, err
	}

	finalArgs := append(args, limit, offset)
	rows, err := r.db.Query(replacePlaceholders(query+" LIMIT ? OFFSET ?"), finalArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var k model.KabupatenKota
		if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeProvinsi); err != nil {
			return nil, 0, err
		}
		data = append(data, k)
	}

	return data, totalItem, nil
}

func (r *DaerahRepository) GetKabKota(limit, offset int, pagination bool) ([]model.KabupatenKota, int, error) {
	return r.fetchKabKotaRaw("SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota", "SELECT COUNT(*) FROM nk_kabupaten_kota", nil, limit, offset, pagination)
}

func (r *DaerahRepository) GetKabKotaByProvinsi(kodeProv string, limit, offset int, pagination bool) ([]model.KabupatenKota, int, error) {
	return r.fetchKabKotaRaw("SELECT kode, nama, lat, lng, kode_provinsi FROM nk_kabupaten_kota WHERE kode_provinsi = ?", "SELECT COUNT(*) FROM nk_kabupaten_kota WHERE kode_provinsi = ?", []interface{}{kodeProv}, limit, offset, pagination)
}


func (r *DaerahRepository) fetchKecamatanRaw(query string, countQuery string, args []interface{}, limit, offset int, pagination bool) ([]model.Kecamatan, int, error) {
	var data []model.Kecamatan
	var totalItem int
	var err error

	if !pagination {
		rows, err := r.db.Query(replacePlaceholders(query), args...)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		for rows.Next() {
			var k model.Kecamatan
			if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota); err != nil {
				return nil, 0, err
			}
			data = append(data, k)
		}
		return data, 0, nil
	}

	err = r.db.QueryRow(replacePlaceholders(countQuery), args...).Scan(&totalItem)
	if err != nil {
		return nil, 0, err
	}

	finalArgs := append(args, limit, offset)
	rows, err := r.db.Query(replacePlaceholders(query+" LIMIT ? OFFSET ?"), finalArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var k model.Kecamatan
		if err := rows.Scan(&k.Kode, &k.Nama, &k.Lat, &k.Lng, &k.KodeKabupatenKota); err != nil {
			return nil, 0, err
		}
		data = append(data, k)
	}

	return data, totalItem, nil
}

func (r *DaerahRepository) GetKecamatan(limit, offset int, pagination bool) ([]model.Kecamatan, int, error) {
	return r.fetchKecamatanRaw("SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan", "SELECT COUNT(*) FROM nk_kecamatan", nil, limit, offset, pagination)
}

func (r *DaerahRepository) GetKecamatanByKabKota(kodeKab string, limit, offset int, pagination bool) ([]model.Kecamatan, int, error) {
	return r.fetchKecamatanRaw("SELECT kode, nama, lat, lng, kode_kabupaten_kota FROM nk_kecamatan WHERE kode_kabupaten_kota = ?", "SELECT COUNT(*) FROM nk_kecamatan WHERE kode_kabupaten_kota = ?", []interface{}{kodeKab}, limit, offset, pagination)
}


func (r *DaerahRepository) fetchDesaKelRaw(query string, countQuery string, args []interface{}, limit, offset int, pagination bool) ([]model.DesaKelurahan, int, error) {
	var data []model.DesaKelurahan
	var totalItem int
	var err error

	if !pagination {
		rows, err := r.db.Query(replacePlaceholders(query), args...)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()

		for rows.Next() {
			var d model.DesaKelurahan
			if err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos); err != nil {
				return nil, 0, err
			}
			data = append(data, d)
		}
		return data, 0, nil
	}

	err = r.db.QueryRow(replacePlaceholders(countQuery), args...).Scan(&totalItem)
	if err != nil {
		return nil, 0, err
	}

	finalArgs := append(args, limit, offset)
	rows, err := r.db.Query(replacePlaceholders(query+" LIMIT ? OFFSET ?"), finalArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var d model.DesaKelurahan
		if err := rows.Scan(&d.Kode, &d.Nama, &d.Lat, &d.Lng, &d.KodeKecamatan, &d.KodePos); err != nil {
			return nil, 0, err
		}
		data = append(data, d)
	}

	return data, totalItem, nil
}

func (r *DaerahRepository) GetDesaKelurahan(limit, offset int, pagination bool) ([]model.DesaKelurahan, int, error) {
	return r.fetchDesaKelRaw("SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan", "SELECT COUNT(*) FROM nk_desa_kelurahan", nil, limit, offset, pagination)
}

func (r *DaerahRepository) GetDesaKelurahanByKecamatan(kodeKec string, limit, offset int, pagination bool) ([]model.DesaKelurahan, int, error) {
	return r.fetchDesaKelRaw("SELECT kode, nama, lat, lng, kode_kecamatan, kode_pos FROM nk_desa_kelurahan WHERE kode_kecamatan = ?", "SELECT COUNT(*) FROM nk_desa_kelurahan WHERE kode_kecamatan = ?", []interface{}{kodeKec}, limit, offset, pagination)
}
