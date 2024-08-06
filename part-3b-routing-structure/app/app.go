package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/router"
)

type App struct {
	router chi.Router
}

func New() *App {
	return &App{router: router.New()}
}

func (app *App) Router() chi.Router {
	return app.router
}
