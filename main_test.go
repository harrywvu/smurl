package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestShortenAndRedirect(t *testing.T) {
	router := newRouter()
	const original = "https://example.com/path?q=1"
	req := httptest.NewRequest(http.MethodPost, "/api/v1/urls",
		strings.NewReader(`{"original_url":"  `+original+`  "}`))
	req.Header.Set("Content-Type", "application/json")
	created := httptest.NewRecorder()
	router.ServeHTTP(created, req)

	if created.Code != http.StatusCreated {
		t.Fatalf("POST status = %d, want 201, body: %s", created.Code, created.Body)
	}

	var link struct {
		Code        string `json:"code"`
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
		CreatedAt   string `json:"created_at"`
	}
	if err := json.Unmarshal(created.Body.Bytes(), &link); err != nil {
		t.Fatalf("decode creation response: %v", err)
	}
	if len(link.Code) != 7 {
		t.Errorf("code = %q, want 7 characters", link.Code)
	}
	if want := "http://localhost:8080/" + link.Code; link.ShortURL != want {
		t.Errorf("short_url = %q, want %q", link.ShortURL, want)
	}
	if link.OriginalURL != original {
		t.Errorf("original_url = %q, want %q", link.OriginalURL, original)
	}
	if _, err := time.Parse(time.RFC3339, link.CreatedAt); err != nil {
		t.Errorf("created_at must be RFC3339: %v", err)
	}

	redirect := httptest.NewRecorder()
	router.ServeHTTP(redirect, httptest.NewRequest(http.MethodGet, "/"+link.Code, nil))
	if redirect.Code != http.StatusFound {
		t.Errorf("GET status = %d, want 302", redirect.Code)
	}
	if got := redirect.Header().Get("Location"); got != original {
		t.Errorf("Location = %q, want %q", got, original)
	}
}

func TestShortenBadRequest(t *testing.T) {
	router := newRouter()
	tests := []struct {
		name string
		body string
		code string
	}{
		{"malformed JSON", `{"original_url":`, "invalid_json"},
		{"missing URL", `{}`, "invalid_json"},
		{"blank URL", `{"original_url":"   "}`, "invalid_json"},
		{"overlong URL", `{"original_url":"https://example.com/` + strings.Repeat("x", 2048) + `"}`, "invalid_url"},
		{"unparseable URL", `{"original_url":"https://example.com/%zz"}`, "invalid_url"},
		{"unsupported scheme", `{"original_url":"ftp://example.com"}`, "invalid_url"},
		{"missing host", `{"original_url":"https:///path"}`, "invalid_url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/urls", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			checkAPIError(t, res, http.StatusBadRequest, tt.code)
		})
	}
}

func TestRedirectNotFound(t *testing.T) {
	res := httptest.NewRecorder()
	newRouter().ServeHTTP(res, httptest.NewRequest(http.MethodGet, "/missing", nil))
	checkAPIError(t, res, http.StatusNotFound, "not_found")
}

func checkAPIError(t *testing.T, res *httptest.ResponseRecorder, status int, code string) {
	t.Helper()
	if res.Code != status {
		t.Fatalf("status = %d, want %d, body: %s", res.Code, status, res.Body)
	}
	var body struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode error response: %v", err)
	}
	if body.Error.Code != code || body.Error.Message == "" {
		t.Errorf("error response = %s, want code %q and a message", res.Body, code)
	}
}
