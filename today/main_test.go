package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	mux := mainHandler()
	t.Run("root", func(t *testing.T) {
		r, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)
		res := w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected %d got %d", http.StatusOK, res.StatusCode)
		}
		if w.Body.String() != "Hello!" {
			t.Errorf("expected body to be %q got %q", "Hello!", w.Body.String())
		}
	})
	t.Run("Status NotFound", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/notfound", nil)
		if err != nil {
			t.Fatal(err)
		}
		mux.ServeHTTP(w, r)
		resp := w.Result()
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected %d got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}
