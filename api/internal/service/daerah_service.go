package service

import (
	"context"
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
	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	if cdnBaseURL == "" {
		cdnBaseURL = "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson"
	}

	return &DaerahService{
		repo:       repo,
		cdnBaseURL: cdnBaseURL,
	}
}

// buildResponse wraps the pagination calculation, empty data checks, and response formatting generically
func buildResponse[T any](data []T, limit, halaman, totalItem int, pagination bool) (any, error) {
	if len(data) == 0 {
		return nil, errors.New("tidak ditemukan data")
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
	}

	totalHalaman := (totalItem + limit - 1) / limit
	if halaman > totalHalaman && totalHalaman > 0 {
		return nil, fmt.Errorf("nomor halaman melebihi total halaman. Halaman maksimum adalah %d", totalHalaman)
	}

	meta := &model.PaginationMeta{
		TotalItem:      totalItem,
		TotalHalaman:   totalHalaman,
		HalamanSaatIni: halaman,
		UkuranHalaman:  limit,
	}

	return model.PaginatedResponse[T]{
		Pagination: meta,
		Data:       data,
	}, nil
}

func (s *DaerahService) GetProvinsi(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetProvinsi(ctx, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	return buildResponse(data, limit, halaman, totalItem, pagination)
}

func (s *DaerahService) GetKabKota(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKabKota(ctx, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	return buildResponse(data, limit, halaman, totalItem, pagination)
}

func (s *DaerahService) GetKabKotaByProvinsi(ctx context.Context, kodeProv string, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKabKotaByProvinsi(ctx, kodeProv, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	return buildResponse(data, limit, halaman, totalItem, pagination)
}

func (s *DaerahService) GetKecamatan(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKecamatan(ctx, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	return buildResponse(data, limit, halaman, totalItem, pagination)
}

func (s *DaerahService) GetKecamatanByKabKota(ctx context.Context, kodeKab string, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetKecamatanByKabKota(ctx, kodeKab, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	return buildResponse(data, limit, halaman, totalItem, pagination)
}

func (s *DaerahService) GetDesaKelurahan(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetDesaKelurahan(ctx, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	return buildResponse(data, limit, halaman, totalItem, pagination)
}

func (s *DaerahService) GetDesaKelurahanByKecamatan(ctx context.Context, kodeKec string, limit, halaman int, pagination bool) (any, error) {
	offset := (halaman - 1) * limit
	data, totalItem, err := s.repo.GetDesaKelurahanByKecamatan(ctx, kodeKec, limit, offset, pagination)
	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data[i].Kode)
	}

	return buildResponse(data, limit, halaman, totalItem, pagination)
}
