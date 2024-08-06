package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://example.com/", nil)
	w := httptest.NewRecorder()
	HomeHandler(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read response body: %s", err)
	}

	expected := `{"Message":"Hello World - Welcome to my first Go homepage"}`
	if expected != string(body) {
		t.Errorf("expected body %s but got %s", expected, string(body))
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type header to be application/json but got %s", resp.Header.Get("Content-Type"))
	}
}
