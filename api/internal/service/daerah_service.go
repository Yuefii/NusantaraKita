package service

import (
	"errors"
	"fmt"
	"os"

	"vercel-go-starter/internal/model"
	"vercel-go-starter/internal/repository"
)

type DaerahService struct {
	repo       *repository.DaerahRepository
	cdnBaseURL string
}

func NewDaerahService(repo *repository.DaerahRepository) *DaerahService {
	// Fallback to default if not set in .env
	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	if cdnBaseURL == "" {
		cdnBaseURL = "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson"
	}

	return &DaerahService{
		repo:       repo,
		cdnBaseURL: cdnBaseURL,
	}
}

func (s *DaerahService) processPagination(halaman, limit, totalItem int, pagination bool) (*model.PaginationMeta, error) {
	if !pagination {
		return nil, nil
	}

	totalHalaman := (totalItem + limit - 1) / limit
	if halaman > totalHalaman && totalHalaman > 0 {
		return nil, fmt.Errorf("nomor halaman melebihi total halaman. Halaman maksimum adalah %d", totalHalaman)
	}

	return &model.PaginationMeta{
		TotalItem:      totalItem,
		TotalHalaman:   totalHalaman,
		HalamanSaatIni: halaman,
		UkuranHalaman:  limit,
	}, nil
}

func (s *DaerahService) GetProvinsi(limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetProvinsi(limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	paginationMeta, err := s.processPagination(halaman, limit, totalItem, pagination)
	if err != nil {
		return nil, err
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	return model.PaginatedProvinsiResponse{
		Pagination: paginationMeta,
		Data:       data,
	}, nil
}

func (s *DaerahService) GetKabKota(limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKabKota(limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	paginationMeta, err := s.processPagination(halaman, limit, totalItem, pagination)
	if err != nil {
		return nil, err
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	return model.PaginatedKabupatenKotaResponse{
		Pagination: paginationMeta,
		Data:       data,
	}, nil
}

func (s *DaerahService) GetKabKotaByProvinsi(kodeProv string, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKabKotaByProvinsi(kodeProv, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	paginationMeta, err := s.processPagination(halaman, limit, totalItem, pagination)
	if err != nil {
		return nil, err
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	return model.PaginatedKabupatenKotaResponse{
		Pagination: paginationMeta,
		Data:       data,
	}, nil
}

func (s *DaerahService) GetKecamatan(limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKecamatan(limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	paginationMeta, err := s.processPagination(halaman, limit, totalItem, pagination)
	if err != nil {
		return nil, err
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	return model.PaginatedKecamatanResponse{
		Pagination: paginationMeta,
		Data:       data,
	}, nil
}

func (s *DaerahService) GetKecamatanByKabKota(kodeKab string, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKecamatanByKabKota(kodeKab, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	paginationMeta, err := s.processPagination(halaman, limit, totalItem, pagination)
	if err != nil {
		return nil, err
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	return model.PaginatedKecamatanResponse{
		Pagination: paginationMeta,
		Data:       data,
	}, nil
}

func (s *DaerahService) GetDesaKelurahan(limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetDesaKelurahan(limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	paginationMeta, err := s.processPagination(halaman, limit, totalItem, pagination)
	if err != nil {
		return nil, err
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	return model.PaginatedDesaKelurahanResponse{
		Pagination: paginationMeta,
		Data:       data,
	}, nil
}

func (s *DaerahService) GetDesaKelurahanByKecamatan(kodeKec string, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetDesaKelurahanByKecamatan(kodeKec, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	paginationMeta, err := s.processPagination(halaman, limit, totalItem, pagination)
	if err != nil {
		return nil, err
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	return model.PaginatedDesaKelurahanResponse{
		Pagination: paginationMeta,
		Data:       data,
	}, nil
}
