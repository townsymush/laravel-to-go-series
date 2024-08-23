package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/controllers"
)

func New() chi.Router {
	r := chi.NewRouter()

	// set up middlewares
	r.Use(middleware.Logger)

	// add routes like Laravel
	r.Get("/", controllers.HomeHandler)

	r.Route("/pokemon", func(r chi.Router) {
		r.Get("/", controllers.PokemonRoot)
		r.Get("/{pokemonName}", controllers.PokemonGet)
	})

	return r
}
