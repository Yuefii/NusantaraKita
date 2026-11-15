package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"nusantarakita/internal/model"
	"nusantarakita/internal/repository"

	"github.com/patrickmn/go-cache"
)

var (
	ErrNotFound          = errors.New("tidak ditemukan data")
	ErrBadRequest        = errors.New("parameter tidak valid")
	ErrQueryTooShort     = errors.New("parameter 'q' minimal 2 karakter")
	ErrOutsideIndonesia = errors.New("koordinat berada di luar wilayah Indonesia")
)

type ErrPageExceeded struct {
	MaxPage int
}

func (e ErrPageExceeded) Error() string {
	return fmt.Sprintf("nomor halaman melebihi total halaman. Halaman maksimum adalah %d", e.MaxPage)
}

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

// pageOffset returns (halaman-1)*limit, capped at math.MaxInt instead of
// overflowing into a negative OFFSET for very large page numbers.
func pageOffset(limit, halaman int) int {
	if halaman-1 > math.MaxInt/limit {
		return math.MaxInt
	}
	return (halaman - 1) * limit
}

// buildResponse wraps the pagination calculation, empty data checks, and response formatting generically
func buildResponse[T any](data []T, limit, halaman, totalItem int, pagination bool) (any, error) {
	// Check the page bound first: a page past the end returns no rows.
	// Written this way so a huge limit cannot overflow.
	totalHalaman := 0
	if totalItem > 0 {
		totalHalaman = (totalItem-1)/limit + 1
	}
	if pagination && halaman > totalHalaman && totalHalaman > 0 {
		return nil, ErrPageExceeded{MaxPage: totalHalaman}
	}

	if len(data) == 0 {
		return nil, ErrNotFound
	}

	if !pagination {
		return map[string]interface{}{"data": data}, nil
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
		offset := pageOffset(limit, halaman)
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
		offset := pageOffset(limit, halaman)
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
		offset := pageOffset(limit, halaman)
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
		offset := pageOffset(limit, halaman)
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
		offset := pageOffset(limit, halaman)
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
		offset := pageOffset(limit, halaman)
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
		offset := pageOffset(limit, halaman)
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

// ---------------- Detail By Kode ----------------

func (s *DaerahService) GetProvinsiByKode(ctx context.Context, kode string) (any, error) {
	cacheKey := fmt.Sprintf("provinsi_detail_%s", kode)
	return s.withCache(ctx, cacheKey, func() (any, error) {
		data, err := s.repo.GetProvinsiByKode(ctx, kode)
		if err != nil {
			return nil, err
		}
		if data == nil {
			return nil, ErrNotFound
		}
		data.GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", s.cdnBaseURL, data.Kode)
		return model.SingleResponse[*model.Provinsi]{Data: data}, nil
	})
}

func (s *DaerahService) GetKabKotaByKode(ctx context.Context, kode string) (any, error) {
	cacheKey := fmt.Sprintf("kabkota_detail_%s", kode)
	return s.withCache(ctx, cacheKey, func() (any, error) {
		data, err := s.repo.GetKabKotaByKode(ctx, kode)
		if err != nil {
			return nil, err
		}
		if data == nil {
			return nil, ErrNotFound
		}
		data.GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data.Kode)
		return model.SingleResponse[*model.KabupatenKota]{Data: data}, nil
	})
}

func (s *DaerahService) GetKecamatanByKode(ctx context.Context, kode string) (any, error) {
	cacheKey := fmt.Sprintf("kecamatan_detail_%s", kode)
	return s.withCache(ctx, cacheKey, func() (any, error) {
		data, err := s.repo.GetKecamatanByKode(ctx, kode)
		if err != nil {
			return nil, err
		}
		if data == nil {
			return nil, ErrNotFound
		}
		data.GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data.Kode)
		return model.SingleResponse[*model.Kecamatan]{Data: data}, nil
	})
}

func (s *DaerahService) GetDesaKelurahanByKode(ctx context.Context, kode string) (any, error) {
	cacheKey := fmt.Sprintf("desa_detail_%s", kode)
	return s.withCache(ctx, cacheKey, func() (any, error) {
		data, err := s.repo.GetDesaKelurahanByKode(ctx, kode)
		if err != nil {
			return nil, err
		}
		if data == nil {
			return nil, ErrNotFound
		}
		data.GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data.Kode)
		return model.SingleResponse[*model.DesaKelurahan]{Data: data}, nil
	})
}

// ---------------- Search / Autocomplete ----------------

func (s *DaerahService) SearchWilayah(ctx context.Context, q, tipe string, limit, halaman int, pagination bool) (any, error) {
	if len(q) < 2 {
		return nil, ErrQueryTooShort
	}

	cacheKey := fmt.Sprintf("search_%s_%s_%d_%d_%v", q, tipe, limit, halaman, pagination)
	return s.withCache(ctx, cacheKey, func() (any, error) {
		offset := pageOffset(limit, halaman)
		data, totalItem, err := s.repo.SearchWilayah(ctx, q, tipe, limit, offset, pagination)
		if err != nil {
			return nil, err
		}

		for i := range data {
			switch data[i].Tipe {
			case "provinsi":
				data[i].GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", s.cdnBaseURL, data[i].Kode)
			case "kabupaten_kota":
				data[i].GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, data[i].Kode)
			case "kecamatan":
				data[i].GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, data[i].Kode)
			case "desa_kelurahan":
				data[i].GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, data[i].Kode)
			}
		}

		return buildResponse(data, limit, halaman, totalItem, pagination)
	})
}

// ---------------- Reverse Geocoding ----------------

func (s *DaerahService) ReverseGeocode(ctx context.Context, lat, lng float64) (any, error) {
	// Quick boundary check for Indonesian coordinates
	if lat < -15.0 || lat > 10.0 || lng < 90.0 || lng > 145.0 {
		return nil, ErrOutsideIndonesia
	}

	cacheKey := fmt.Sprintf("reverse_%.4f_%.4f", lat, lng)
	return s.withCache(ctx, cacheKey, func() (any, error) {
		res, err := s.repo.ReverseGeocode(ctx, lat, lng)
		if err != nil {
			return nil, err
		}
		if res == nil || res.JarakMeter > 300000 {
			return nil, ErrOutsideIndonesia
		}

		res.Provinsi.GeojsonURL = fmt.Sprintf("%s/provinsi/%s.geojson", s.cdnBaseURL, res.Provinsi.Kode)
		res.KabupatenKota.GeojsonURL = fmt.Sprintf("%s/kabupaten_kota/%s.geojson", s.cdnBaseURL, res.KabupatenKota.Kode)
		res.Kecamatan.GeojsonURL = fmt.Sprintf("%s/kecamatan/%s.geojson", s.cdnBaseURL, res.Kecamatan.Kode)
		res.DesaKelurahan.GeojsonURL = fmt.Sprintf("%s/desa_kelurahan/%s.geojson", s.cdnBaseURL, res.DesaKelurahan.Kode)

		return model.SingleResponse[*model.ReverseGeocodeResult]{Data: res}, nil
	})
}
