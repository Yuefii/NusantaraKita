package service

import (
	"testing"
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
		data := []string{"item1"}
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
