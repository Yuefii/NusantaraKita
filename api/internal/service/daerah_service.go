package service

import (
	"context"
	"errors"
	"fmt"
	"os"

	"vercel-go-starter/internal/model"
	"vercel-go-starter/internal/repository"

	"github.com/patrickmn/go-cache"
	"time"
)

type DaerahService struct {
	repo       *repository.DaerahRepository
	cdnBaseURL string
	cache      *cache.Cache
}

func NewDaerahService(repo *repository.DaerahRepository) *DaerahService {
	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	if cdnBaseURL == "" {
		cdnBaseURL = "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson"
	}

	// Create a cache with a default expiration time of 24 hours, and which
	// purges expired items every 1 hour (though data is mostly static)
	c := cache.New(24*time.Hour, 1*time.Hour)

	return &DaerahService{
		repo:       repo,
		cdnBaseURL: cdnBaseURL,
		cache:      c,
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

// withCache is a generic helper to check cache first, then call the fetcher if missing
func (s *DaerahService) withCache(cacheKey string, fetcher func() (any, error)) (any, error) {
	if cached, found := s.cache.Get(cacheKey); found {
		return cached, nil
	}

	result, err := fetcher()
	if err != nil {
		return nil, err
	}

	// Store in cache with default expiration
	s.cache.Set(cacheKey, result, cache.DefaultExpiration)
	return result, nil
}

func (s *DaerahService) GetProvinsi(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("provinsi_%d_%d_%v", limit, halaman, pagination)
	return s.withCache(cacheKey, func() (any, error) {
		offset := (halaman - 1) * limit
		data, totalItem, err := s.repo.GetProvinsi(ctx, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			data[i].GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", s.cdnBaseURL, data[i].Kode)
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}

func (s *DaerahService) GetKabKota(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("kabkota_%d_%d_%v", limit, halaman, pagination)
	return s.withCache(cacheKey, func() (any, error) {
		offset := (halaman - 1) * limit
		data, totalItem, err := s.repo.GetKabKota(ctx, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			data[i].GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data[i].Kode)
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}

func (s *DaerahService) GetKabKotaByProvinsi(ctx context.Context, kodeProv string, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("kabkota_prov_%s_%d_%d_%v", kodeProv, limit, halaman, pagination)
	return s.withCache(cacheKey, func() (any, error) {
		offset := (halaman - 1) * limit
		data, totalItem, err := s.repo.GetKabKotaByProvinsi(ctx, kodeProv, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			data[i].GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data[i].Kode)
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}

func (s *DaerahService) GetKecamatan(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("kecamatan_%d_%d_%v", limit, halaman, pagination)
	return s.withCache(cacheKey, func() (any, error) {
		offset := (halaman - 1) * limit
		data, totalItem, err := s.repo.GetKecamatan(ctx, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			data[i].GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data[i].Kode)
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}

func (s *DaerahService) GetKecamatanByKabKota(ctx context.Context, kodeKab string, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("kecamatan_kab_%s_%d_%d_%v", kodeKab, limit, halaman, pagination)
	return s.withCache(cacheKey, func() (any, error) {
		offset := (halaman - 1) * limit
		data, totalItem, err := s.repo.GetKecamatanByKabKota(ctx, kodeKab, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			data[i].GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data[i].Kode)
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}

func (s *DaerahService) GetDesaKelurahan(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("desa_%d_%d_%v", limit, halaman, pagination)
	return s.withCache(cacheKey, func() (any, error) {
		offset := (halaman - 1) * limit
		data, totalItem, err := s.repo.GetDesaKelurahan(ctx, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			data[i].GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data[i].Kode)
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}

func (s *DaerahService) GetDesaKelurahanByKecamatan(ctx context.Context, kodeKec string, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("desa_kec_%s_%d_%d_%v", kodeKec, limit, halaman, pagination)
	return s.withCache(cacheKey, func() (any, error) {
		offset := (halaman - 1) * limit
		data, totalItem, err := s.repo.GetDesaKelurahanByKecamatan(ctx, kodeKec, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			data[i].GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data[i].Kode)
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}
