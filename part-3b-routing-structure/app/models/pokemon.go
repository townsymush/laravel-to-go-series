package models

type MyPokemon struct {
	Pokemon []Pokemon `json:"My Pokemon"`
}
type Pokemon struct {
	Name string `json:"name"`
	Type string `json:"type"`
	HP   int    `json:"hp"`
}
