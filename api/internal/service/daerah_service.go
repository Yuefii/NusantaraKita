package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"nusantarakita/internal/model"
	"nusantarakita/internal/repository"

	"github.com/patrickmn/go-cache"
)

type DaerahService struct {
	repo         *repository.DaerahRepository
	cdnBaseURL   string
	upstashURL   string
	upstashToken string
	httpClient   *http.Client
	localCache   *cache.Cache
}

func NewDaerahService(repo *repository.DaerahRepository) *DaerahService {
	cdnBaseURL := os.Getenv("CDN_BASE_URL")
	if cdnBaseURL == "" {
		cdnBaseURL = "https://cdn.jsdelivr.net/gh/yuefii/NusantaraKita@main/geojson"
	}

	return &DaerahService{
		repo:         repo,
		cdnBaseURL:   cdnBaseURL,
		upstashURL:   os.Getenv("UPSTASH_REDIS_REST_URL"),
		upstashToken: os.Getenv("UPSTASH_REDIS_REST_TOKEN"),
		httpClient:   &http.Client{Timeout: 5 * time.Second},
		localCache:   cache.New(24*time.Hour, 1*time.Hour),
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

// withCache is a generic helper to check Upstash cache first, fallback to in-memory, then call the fetcher if missing
func (s *DaerahService) withCache(ctx context.Context, cacheKey string, fetcher func() (any, error)) (any, error) {
	useRedis := s.upstashURL != "" && s.upstashToken != ""

	// 1. Try to get from Upstash Redis
	if useRedis {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/get/%s", s.upstashURL, cacheKey), nil)
		if err == nil {
			req.Header.Set("Authorization", "Bearer "+s.upstashToken)
			resp, err := s.httpClient.Do(req)
			if err == nil {
				defer resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					var upstashResp struct {
						Result string `json:"result"`
					}
					if err := json.NewDecoder(resp.Body).Decode(&upstashResp); err == nil && upstashResp.Result != "" {
						return json.RawMessage(upstashResp.Result), nil
					}
				}
			}
		}
	} else {
		// 1b. Try to get from Local Cache if Redis is not configured
		if cached, found := s.localCache.Get(cacheKey); found {
			return cached, nil
		}
	}

	// 2. Fetch fresh data
	result, err := fetcher()
	if err != nil {
		return nil, err
	}

	// 3. Store Data
	if useRedis {
		// Store in Upstash
		resultBytes, err := json.Marshal(result)
		if err == nil {
			// Upstash array format for POST: ["SET", "key", "value", "EX", seconds]
			payload := []interface{}{"SET", cacheKey, string(resultBytes), "EX", 86400} // 24 hours
			payloadBytes, _ := json.Marshal(payload)

			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, s.upstashURL, bytes.NewBuffer(payloadBytes))
			if err == nil {
				req.Header.Set("Authorization", "Bearer "+s.upstashToken)
				req.Header.Set("Content-Type", "application/json")
				// Fire and forget
				go func() {
					postResp, postErr := s.httpClient.Do(req)
					if postErr == nil {
						postResp.Body.Close()
					}
				}()
			}
		}
	} else {
		// Store in Local Cache
		s.localCache.Set(cacheKey, result, cache.DefaultExpiration)
	}

	return result, nil
}

func (s *DaerahService) GetProvinsi(ctx context.Context, limit, halaman int, pagination bool) (any, error) {
	cacheKey := fmt.Sprintf("provinsi_%d_%d_%v", limit, halaman, pagination)
	return s.withCache(ctx, cacheKey, func() (any, error) {
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
	return s.withCache(ctx, cacheKey, func() (any, error) {
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
	return s.withCache(ctx, cacheKey, func() (any, error) {
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
	return s.withCache(ctx, cacheKey, func() (any, error) {
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
	return s.withCache(ctx, cacheKey, func() (any, error) {
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
	return s.withCache(ctx, cacheKey, func() (any, error) {
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
	return s.withCache(ctx, cacheKey, func() (any, error) {
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
