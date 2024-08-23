package storage

import (
	"errors"
	"strings"
)

type Pokemon struct {
	Name string
	Type string
	HP   int
}

type PokemonStore struct {
	data []Pokemon
}

// NewPokemonStore returns an instance of the data store to get Pokemon Data
// note: this initialises data in the store. We will expand on this and use DB connection in future articles
func NewPokemonStore() *PokemonStore {
	return &PokemonStore{
		data: []Pokemon{
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
}

// Get returns a Pokemon by their name
// note: we wouldn't usually use something like a name as a unique identifier but for this article is suits the use case.
func (p *PokemonStore) Get(name string) (Pokemon, error) {
	if p == nil {
		return Pokemon{}, errors.New("pokemon storage is nil")
	}
	for _, pkmn := range p.data {
		if strings.ToLower(pkmn.Name) == strings.ToLower(name) {
			return pkmn, nil
		}
	}
	return Pokemon{}, NewNotFoundError("pokemon")
}

// All returns all the Pokemon in our store
func (p *PokemonStore) All() []Pokemon {
	return p.data
}

// Insert accepts a Pokemon and stores it. If the Pokemon already exists we return an error
// note we would usually have this logic in our programme. This will be expanded in a future article
func (p *PokemonStore) Insert(pokemon Pokemon) error {
	if p == nil {
		return errors.New("pokemon storage is nil")
	}
	for _, pkmn := range p.data {
		if pkmn.Name == pokemon.Name {
			return errors.New("pokemon already exists")
		}
	}
	p.data = append(p.data, pokemon)
	return nil
}
