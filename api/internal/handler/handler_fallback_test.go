package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Requests no registered route answers still get the JSON error body.
func TestUnmatchedRoutesReturnJSON(t *testing.T) {
	mux, _ := newTestMux(t)
	// A route registered after RegisterRoutes, as cmd/server does for /health.
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {})

	t.Run("unknown one-segment path is a JSON 404", func(t *testing.T) {
		assertError(t, serve(mux, "/v2/tidak-ada"), http.StatusNotFound, codeNotFound, "Endpoint tidak ditemukan")
	})

	t.Run("unknown three-segment path is a JSON 404", func(t *testing.T) {
		assertError(t, serve(mux, "/v2/a/b/c"), http.StatusNotFound, codeNotFound, "Endpoint tidak ditemukan")
	})

	t.Run("path outside /v2 is a JSON 404", func(t *testing.T) {
		assertError(t, serve(mux, "/tidak-ada"), http.StatusNotFound, codeNotFound, "Endpoint tidak ditemukan")
	})

	for _, target := range []string{"/v2/provinsi", "/v2/provinsi/32", "/v2/provinsi/32/kab-kota", "/health"} {
		t.Run("POST "+target+" is a JSON 405", func(t *testing.T) {
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, target, nil))
			assertError(t, w, http.StatusMethodNotAllowed, codeMethodNotAllowed, "Metode tidak diizinkan")
			if allow := w.Header().Get("Allow"); allow != "GET, HEAD" {
				t.Errorf("expected Allow \"GET, HEAD\", got %q", allow)
			}
		})
	}

	t.Run("POST to an unknown path is a JSON 404", func(t *testing.T) {
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/v2/a/b/c", nil))
		assertError(t, w, http.StatusNotFound, codeNotFound, "Endpoint tidak ditemukan")
	})
}
