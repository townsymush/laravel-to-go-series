package mapping

import (
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/response"
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/storage"
)

// ToPokemonResponse accepts a storage.Pokemon and converts it to a response.Pokemon type
func ToPokemonResponse(in storage.Pokemon) response.Pokemon {
	return response.Pokemon{
		Name: in.Name,
		HP:   in.HP,
		Type: in.Type,
	}
}
