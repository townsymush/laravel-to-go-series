package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/mapping"
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/response"
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/storage"
	"net/http"
)

// todo add context
type PokemonReader interface {
	Get(name string) (storage.Pokemon, error)
	All() ([]storage.Pokemon, error)
}

// todo add logger
func PokemonRoot(w http.ResponseWriter, _ *http.Request) {
	store := storage.NewPokemonStore()

	p := response.PokemonResponse{}
	for _, c := range store.All() {
		p.MyPokemon = append(p.MyPokemon, mapping.ToPokemonResponse(c))
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func PokemonGet(w http.ResponseWriter, r *http.Request) {
	store := storage.NewPokemonStore()                 // todo: Let's inject this later
	w.Header().Set("Content-Type", "application/json") //todo we should set this better

	// get param (pokemon name)
	name := chi.URLParam(r, "pokemonName")
	if name == "" {
		resp := response.NewJSON(http.StatusBadRequest, "Pokemon name is required")
		if err := resp.WriteResponse(w); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
		return
	}

	pokemon, err := store.Get(name)
	if err != nil {
		if errors.As(err, &storage.NotFoundErr{}) {
			resp := response.NewJSON(http.StatusNotFound, "Pokemon not found")
			if err = resp.WriteResponse(w); err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	payload := mapping.ToPokemonResponse(pokemon)

	err = json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	return
}
