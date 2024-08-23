package response

type Pokemon struct {
	Name string `json:"name"`
	Type string `json:"type"`
	HP   int    `json:"hp"`
}

type PokemonResponse struct {
	MyPokemon []Pokemon `json:"my_pokemon"`
}
