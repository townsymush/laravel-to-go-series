package controllers

import (
	"encoding/json"
	"github.com/townsymush/laravel-to-go-series/part-4-request-validation/app/models"
	"net/http"
)

func PokemonRoot(w http.ResponseWriter, _ *http.Request) {
	myPokemon := models.MyPokemon{
		Pokemon: []models.Pokemon{
			{
				Name: "Pikachu",
				Type: "Electric",
				HP:   52,
			},

			{
				Name: "Geodude",
				Type: "Ground/Fighting",
				HP:   70,
			},
			{
				Name: "Bulbasaur",
				Type: "Grass",
				HP:   31,
			},
		},
	}

	payload, err := json.Marshal(myPokemon)
	if err != nil {
		// todo log
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(payload)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
