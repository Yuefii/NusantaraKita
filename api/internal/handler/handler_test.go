package handler

import (
	"net/http"
	"net/url"
	"testing"
)

func TestParseQueryParams(t *testing.T) {
	tests := []struct {
		name           string
		queryURL       string
		expectedLimit  int
		expectedHal    int
		expectedPag    bool
	}{
		{
			name:          "Default values",
			queryURL:      "http://example.com/api",
			expectedLimit: 10,
			expectedHal:   1,
			expectedPag:   true,
		},
		{
			name:          "Custom valid values",
			queryURL:      "http://example.com/api?limit=50&halaman=3&pagination=false",
			expectedLimit: 50,
			expectedHal:   3,
			expectedPag:   false,
		},
		{
			name:          "Invalid types fallback to defaults",
			queryURL:      "http://example.com/api?limit=abc&halaman=xyz&pagination=maybe",
			expectedLimit: 10,
			expectedHal:   1,
			expectedPag:   true,
		},
		{
			name:          "Negative limit and halaman fallback to defaults",
			queryURL:      "http://example.com/api?limit=-5&halaman=-1",
			expectedLimit: 10,
			expectedHal:   1,
			expectedPag:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				URL: &url.URL{
					RawQuery: parseRawQuery(tt.queryURL),
				},
			}
			result := parseQueryParams(req)

			if result.limit != tt.expectedLimit {
				t.Errorf("expected limit %d, got %d", tt.expectedLimit, result.limit)
			}
			if result.halaman != tt.expectedHal {
				t.Errorf("expected halaman %d, got %d", tt.expectedHal, result.halaman)
			}
			if result.pagination != tt.expectedPag {
				t.Errorf("expected pagination %v, got %v", tt.expectedPag, result.pagination)
			}
		})
	}
}

func parseRawQuery(u string) string {
	parsed, _ := url.Parse(u)
	return parsed.RawQuery
}
