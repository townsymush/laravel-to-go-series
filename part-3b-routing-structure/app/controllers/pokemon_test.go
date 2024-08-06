package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPokemonRoot(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://example.com/pokemon", nil)
	w := httptest.NewRecorder()
	PokemonRoot(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read response body: %s", err)
	}

	expected := `{"My Pokemon":[{"name":"Pikachu","type":"Electric","hp":52},{"name":"Geodude","type":"Ground/Fighting","hp":70},{"name":"Bulbasaur","type":"Grass","hp":31}]}`
	if expected != string(body) {
		t.Errorf("expected body %s but got %s", expected, string(body))
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type header to be application/json but got %s", resp.Header.Get("Content-Type"))
	}
}
