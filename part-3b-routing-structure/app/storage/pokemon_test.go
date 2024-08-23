package storage

import (
	"reflect"
	"testing"
)

func TestNewPokemonStore(t *testing.T) {
	p := NewPokemonStore()

	expected := []Pokemon{
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
	}

	actual := p.All()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v\nActual: %v", expected, actual)
	}
}

func TestPokemonStore_Insert(t *testing.T) {
	p := NewPokemonStore()

	expected := []Pokemon{
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
		{
			Name: "Charmander",
			Type: "Fire",
			HP:   45,
		},
	}

	err := p.Insert(Pokemon{
		Name: "Charmander",
		Type: "Fire",
		HP:   45,
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, p.All()) {
		t.Errorf("Expected: %v\nActual: %v", expected, p.All())
	}
}

func TestPokemonStore_Get(t *testing.T) {
	p := NewPokemonStore()
	expected := Pokemon{
		Name: "Bulbasaur",
		Type: "Grass",
		HP:   31,
	}

	actual, err := p.Get("Bulbasaur")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v\nActual: %v", expected, actual)
	}
}
