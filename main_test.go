package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth2(t *testing.T) {
	mux := newMux()

	req := httptest.NewRequest(http.MethodGet, "/health2", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	body, err := io.ReadAll(rec.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if string(body) != "ok" {
		t.Fatalf("expected body %q, got %q", "ok", string(body))
	}

	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "text/plain") {
		t.Fatalf("expected Content-Type to contain %q, got %q", "text/plain", ct)
	}
}

func TestHealth2MethodNotAllowed(t *testing.T) {
	mux := newMux()

	req := httptest.NewRequest(http.MethodPost, "/health2", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatalf("expected non-200 status for POST /health2, got %d", rec.Code)
	}
}
