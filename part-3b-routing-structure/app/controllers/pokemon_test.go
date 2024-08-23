package controllers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPokemonRoot(t *testing.T) {
	tests := []struct {
		Name               string
		URL                string
		ExpectedPayload    string
		ExpectedStatusCode int
		Handler            http.HandlerFunc
		Err                error
	}{
		{
			Name:               "Pokemon Home",
			URL:                "/",
			ExpectedPayload:    `{"my_pokemon":[{"name":"Pikachu","type":"Electric","hp":52},{"name":"Geodude","type":"Ground/Fighting","hp":70},{"name":"Bulbasaur","type":"Grass","hp":31}]}`,
			ExpectedStatusCode: http.StatusOK,
			Handler:            PokemonRoot,
			Err:                nil,
		},
	}

	for _, test := range tests {
		r := httptest.NewRequest(http.MethodGet, test.URL, nil)
		w := httptest.NewRecorder()

		test.Handler(w, r)
		resp := w.Result()

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type header to be application/json but got %s", resp.Header.Get("Content-Type"))
		}

		if resp.StatusCode != test.ExpectedStatusCode {
			t.Errorf("expected status code %d but got %d", test.ExpectedStatusCode, resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("failed to read response body: %s", err)
		}

		if reflect.DeepEqual(test.ExpectedPayload, string(body)) {
			t.Errorf("expected body %s but got %s", test.ExpectedPayload, string(body))
		}
	}
}

func TestPokemonGet(t *testing.T) {
	tests := []struct {
		Name               string
		PokemonName        string
		ExpectedPayload    string
		ExpectedStatusCode int
	}{
		{
			Name:               "Pokemon Get - Success",
			PokemonName:        "Bulbasaur",
			ExpectedPayload:    `{"name":"Bulbasaur","type":"Grass","hp":31}`,
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "Pokemon Get - No Pokemon",
			PokemonName:        "Agumon",
			ExpectedPayload:    `{"message":"Pokemon not found"}`,
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Name:               "Pokemon Get - No supplied pokemon name",
			PokemonName:        "",
			ExpectedPayload:    `{"message":"Pokemon name is required"}`,
			ExpectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/{pokemonName}", nil)
			w := httptest.NewRecorder()
			r = addParams("pokemonName", test.PokemonName, r)

			PokemonGet(w, r)
			resp := w.Result()

			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf(
					"expected Content-Type header to be application/json but got %s",
					resp.Header.Get("Content-Type"),
				)
			}

			if resp.StatusCode != test.ExpectedStatusCode {
				t.Errorf("expected status code %d but got %d", test.ExpectedStatusCode, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("failed to read response body: %s", err)
			}

			if !reflect.DeepEqual(test.ExpectedPayload, strings.Trim(string(body), "\n")) {
				t.Errorf("expected body %s but got %s", test.ExpectedPayload, string(body))
			}
		})
	}
}

func addParams(key string, value string, r *http.Request) *http.Request {
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add(key, value)

	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
}
