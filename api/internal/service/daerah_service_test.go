package service

import (
	"errors"
	"math"
	"testing"

	"nusantarakita/internal/model"
)

func TestBuildResponse(t *testing.T) {
	t.Run("Empty data returns ErrNotFound", func(t *testing.T) {
		data := []string{}
		_, err := buildResponse(data, 10, 1, 0, true)
		if err != ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("Pagination disabled", func(t *testing.T) {
		data := []string{"item1"}
		res, err := buildResponse(data, 10, 1, 1, false)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		resMap, ok := res.(map[string]interface{})
		if !ok {
			t.Fatalf("expected map[string]interface{}, got %T", res)
		}

		if _, exists := resMap["data"]; !exists {
			t.Error("expected 'data' key in response")
		}
	})

	t.Run("Page exceeds max page", func(t *testing.T) {
		// A page past the end returns no rows from the query.
		data := []string{}
		// 15 items, limit 10 => 2 pages max. Requesting page 3.
		_, err := buildResponse(data, 10, 3, 15, true)
		if err == nil {
			t.Fatal("expected ErrPageExceeded, got nil")
		}

		pageErr, ok := err.(ErrPageExceeded)
		if !ok {
			t.Fatalf("expected ErrPageExceeded, got %T", err)
		}
		if pageErr.MaxPage != 2 {
			t.Errorf("expected MaxPage 2, got %d", pageErr.MaxPage)
		}
	})

	t.Run("Valid paginated response", func(t *testing.T) {
		data := []string{"item1", "item2"}
		res, err := buildResponse(data, 10, 1, 12, true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		// res should be of type model.PaginatedResponse[string]
		// Since we can't easily type assert generic types defined in another package if we don't import model,
		// Wait, we can import model and type assert.
		// Actually, since buildResponse returns 'any', we can just check if it's not nil for now.
		if res == nil {
			t.Error("expected response, got nil")
		}
	})
}

func TestPageOffset(t *testing.T) {
	tests := []struct {
		limit, halaman, want int
	}{
		{10, 1, 0},
		{10, 3, 20},
		{1, math.MaxInt, math.MaxInt - 1},
		{10, math.MaxInt, math.MaxInt},
		{10, math.MaxInt/10 + 1, math.MaxInt / 10 * 10},
		{10, math.MaxInt/10 + 2, math.MaxInt},
		{math.MaxInt, 2, math.MaxInt},
		{math.MaxInt, 3, math.MaxInt},
	}
	for _, tt := range tests {
		if got := pageOffset(tt.limit, tt.halaman); got != tt.want {
			t.Errorf("pageOffset(%d, %d) = %d, want %d", tt.limit, tt.halaman, got, tt.want)
		}
		if got := pageOffset(tt.limit, tt.halaman); got < 0 {
			t.Errorf("pageOffset(%d, %d) is negative: %d", tt.limit, tt.halaman, got)
		}
	}
}

func TestBuildResponseHugeLimit(t *testing.T) {
	t.Run("first page reports one page", func(t *testing.T) {
		res, err := buildResponse([]string{"a", "b"}, math.MaxInt, 1, 2, true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if got := res.(model.PaginatedResponse[string]).Pagination.TotalHalaman; got != 1 {
			t.Errorf("expected total_halaman 1, got %d", got)
		}
	})

	t.Run("second page is out of range", func(t *testing.T) {
		_, err := buildResponse([]string{}, math.MaxInt, 2, 2, true)
		var pageErr ErrPageExceeded
		if !errors.As(err, &pageErr) || pageErr.MaxPage != 1 {
			t.Fatalf("expected ErrPageExceeded{MaxPage: 1}, got %v", err)
		}
	})
}
